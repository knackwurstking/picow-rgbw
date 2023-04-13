<script lang="ts">
    import { onMount, onDestroy } from "svelte";

    import CheckLabel from "../components/CheckLabel.svelte";
    import ColorDefaults from "../components/ColorDefaults.svelte";
    import ColorBrightness from "../components/ColorBrightness.svelte";
    import ColorPicker from "../components/ColorPicker.svelte";
    import PowerToggle from "../components/PowerToggle.svelte";

    import Api, { type Device, type Events } from "../ts/api";

    let selected: Device[] = [];

    let devices: Device[] = [];
    $: {
        const newSelected = [];
        for (const s of selected) {
            if (!!devices.find((d) => d.addr === s.addr)) {
                newSelected.push(s);
            }
        }
        selected = newSelected;
    }

    let brightness = 0;
    $: console.log("brightness:", brightness);

    // TODO: need to fix the max value
    let r: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    };
    let g: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    };
    let b: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    };
    let w: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    };

    /*
        rgbw:  90 70 90 80
        0-100: 5 - (100 - 30 + 10) (min:5 - (100% - min:rgbw + diff:max:rgbw:to:100%))
        range: 5 - 80
        current: 70
    */

    let bMin = 5;
    let bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));

    const forDestroy: Events = {
        devices: [],
        device: [],
    };

    onMount(() => {
        // sse: "devices"
        let devicesHandler = async (data: Device[]) => {
            devices = data;
        };
        Api.addEventListener("devices", devicesHandler);
        forDestroy.devices.push(devicesHandler);

        Api.addEventListener("devices", devicesHandler);
        forDestroy.devices.push(devicesHandler);

        // sse: "device"
        const deviceHandler = async (data: Device) => {
            const device = devices.find((d) => d.addr == data.addr);
            device.rgbw = data.rgbw;
            device.offline = data.offline;
            devices = devices;
        };
        Api.addEventListener("device", deviceHandler);
        forDestroy.device.push(deviceHandler);
    });

    onDestroy(() => {
        for (const t in forDestroy) {
            for (const l of forDestroy[t]) {
                Api.removeEventListener(t, l);
            }
        }
    });
</script>

<svelte:head>
    <title>Pico Web | Devices</title>
</svelte:head>

<div class="devices container">
    <section class="list">
        <fieldset>
            <legend>Devices</legend>
            <div class="content list">
                {#each devices as device}
                    <CheckLabel
                        checked={!device.offline &&
                            !!selected.find((sd) => sd.addr === device.addr)}
                        label={device.addr}
                        currentColor={device.rgbw.map((gp) => gp.duty)}
                        offline={device.offline}
                        on:change={() => {
                            if (!!selected.find((d) => d === device)) {
                                // remove device from selected
                                selected = selected.filter((d) => d != device);
                            } else {
                                selected = [...selected, device];
                            }

                            if (!selected.length) {
                                [r, g, b, w] = [100, 100, 100, 100];
                            } else {
                                if (
                                    selected.find(
                                        (device) =>
                                            !!device.rgbw.find(
                                                (gp) => gp.duty > 0
                                            )
                                    )
                                ) {
                                    [r, g, b, w] = selected[
                                        selected.length - 1
                                    ].rgbw.map((d) => d.duty);
                                }
                            }
                        }}
                    />
                {/each}
            </div>
        </fieldset>
    </section>

    <section class="ctrl">
        <fieldset>
            <legend>Control</legend>
            <div class="content">
                <div
                    style="
                        display: flex;
                        flex-direction: column;
                        width: 100%;
                        height: 100%;
                    "
                >
                    <div style="height:100%;" />

                    <!-- TODO: Add some horiz. color default values picker -->
                    <ColorDefaults />

                    <ColorPicker
                        style="height: 180px;"
                        bind:r bind:g bind:b bind:w
                    />
                </div>

                <!-- TODO: height should be the same like the color picker height (?) -->
                <div
                    style="
                        width: 42px;
                        display: flex;
                        flex-direction: column;
                    "
                >
                    <div style="height:100%;" />
                    <ColorBrightness
                        style="
                            height: 168px;
                            padding: 0 8px;
                            padding-bottom: 8px;
                        "
                        value={brightness}
                        on:change={(ev) => {
                            const currentMin = Math.min(...[r, g, b, w]);
                            const diff = currentMin - ev.detail.value;
                            console.log(diff);
                            r -= diff;
                            g -= diff;
                            b -= diff;
                            w -= diff;
                        }}
                    />
                </div>
            </div>
            <div class="bottom">
                <PowerToggle
                    style="
                        width: 100%;
                        height: 100%;
                    "
                    on:change={async (ev) => {
                        switch (ev.detail.state) {
                            case "set":
                                await Api.postDevices(
                                    ...selected.map((d) => ({
                                        addr: d.addr,
                                        rgbw: [
                                            Math.round((r / 100) * brightness),
                                            Math.round((g / 100) * brightness),
                                            Math.round((b / 100) * brightness),
                                            Math.round((w / 100) * brightness),
                                        ],
                                    }))
                                );
                                break;
                            case "off":
                                await Api.postDevices(
                                    ...selected.map((d) => ({
                                        addr: d.addr,
                                        rgbw: [0, 0, 0, 0],
                                    }))
                                );
                                break;
                        }
                    }}
                />
            </div>
        </fieldset>
    </section>
</div>

<style>
    div.devices.container {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
    }

    div.devices.container > section {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

    div.devices.container > section.list {
        height: 60%;
    }

    div.devices.container > section.ctrl {
        height: 40%;
    }

    /* default: mobile (portrait) */

    @media (min-width: 769px) {
        div.devices.container {
            flex-direction: row;
        }

        div.devices.container > section.list {
            width: 50%;
            height: 100%;
        }

        div.devices.container > section.ctrl {
            width: 50%;
            height: 100%;
        }
    }

    div.devices.container > section.list fieldset {
        margin: 16px;
        border-color: var(--theme-border);
        overflow: hidden;
        height: calc(100% - 32px);
    }

    div.devices.container > section.list > fieldset > div.content.list {
        display: flex;
        flex-direction: column;
        scroll-behavior: smooth;
        width: 100%;
        height: 100%;
    }

    div.devices.container > section.ctrl > fieldset {
        margin: 16px;
        padding-bottom: 0;
        height: calc(100% - 32px);
        border-color: var(--theme-border);
    }

    div.devices.container > section.ctrl > fieldset > div {
        margin: 8px;
    }

    div.devices.container > section.ctrl fieldset div.content {
        width: calc(100% - 16px);
        height: calc(100% - 64px - 20px);
        margin-bottom: 0;
        overflow: hidden;
        overflow-y: auto;
        display: flex;
    }

    div.devices.container > section.ctrl fieldset div.bottom {
        display: flex;
        justify-content: space-evenly;
        align-items: center;
        width: calc(100% - 16px);
        height: 64px;
        bottom: 0;
        left: 0;
        padding: 8px;
        margin-top: 0;
        border-top: 1px solid var(--theme-border);
    }
</style>
