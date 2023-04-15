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

    let r: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    }
    let g: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    }
    let b: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    }
    let w: number = 100;
    $: {
        brightness = Math.min(...[r, g, b, w]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b, w]));
    }

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

<div class="container">
    <section class="devices-list">
        <fieldset>
            <legend>Devices</legend>
            <div class="content">
                {#each devices as device}
                    <CheckLabel
                        label={device.addr}
                        currentColor={device.rgbw.map((gp) => gp.duty)}
                        offline={device.offline}
                        on:change={(ev) => {
                            if (!ev.detail.checked) {
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

    <section class="devices-ctrl">
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
                        bind:r
                        bind:g
                        bind:b
                        bind:w
                    />
                </div>

                <!-- TODO: height should be the same like the color picker height (?) -->
                <div
                    style="
                        display: flex;
                        flex-direction: column;
                    "
                >
                    <div style="height:100%;" />
                    <ColorBrightness
                        style="
                            height: 168px;
                            padding: 0 8px 8px 8px;
                        "
                        value={brightness}
                        on:change={(ev) => {
                            const currentMin = Math.min(...[r, g, b, w]);
                            const diff = currentMin - ev.detail.value;

                            if (
                                r - diff > 100 ||
                                g - diff > 100 ||
                                b - diff > 100 ||
                                w - diff > 100
                            ) {
                                const rest = 100 - Math.max(...[r, g, b, w]);
                                r += rest;
                                g += rest;
                                b += rest;
                                w += rest;
                                return;
                            }

                            r -= diff;
                            g -= diff;
                            b -= diff;
                            w -= diff;
                        }}
                    />
                </div>
            </div>
            <div class="button-group">
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
                                            Math.round(r),
                                            Math.round(g),
                                            Math.round(b),
                                            Math.round(w),
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
    fieldset {
        margin-top: 0;
        margin-bottom: 0;
    }

    .container {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        padding: 8px;
    }

    .container > * {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

    .container > .devices-list {
        height: 100%;
    }

    .container > .devices-ctrl {
        min-height: 300px;
        height: 300px;
        max-height: 300px;
    }

    /* default: mobile (portrait) */

    @media (orientation: landscape) {
        .container {
            flex-direction: row;
        }

        .container > .devices-list {
            width: 50%;
            height: 100%;
        }

        .container > .devices-ctrl {
            width: 50%;
            min-height: 100%;
            height: 100%;
            max-height: 100%;
        }
    }

    .container > * > fieldset {
        margin: 4px 8px;
        border-color: var(--theme-border);
    }

    .container > .devices-list fieldset {
        overflow: hidden;
        height: calc(100% - 8px);
    }

    .container > .devices-list > fieldset > .content {
        display: flex;
        flex-direction: column;
        overflow: hidden;
        overflow-y: auto;
        scroll-behavior: smooth;
        width: 100%;
        height: 100%;
    }

    .container > .devices-ctrl > fieldset {
        padding-bottom: 0;
        height: calc(100% - 8px);
    }

    .container > .devices-ctrl > fieldset > div {
        margin: 8px;
    }

    .container > .devices-ctrl fieldset div.content {
        width: calc(100% - 16px);
        height: calc(100% - 56px - 20px);
        overflow: hidden;
        overflow-y: auto;
        display: flex;
    }

    .container > .devices-ctrl fieldset div.button-group {
        display: flex;
        justify-content: space-evenly;
        align-items: center;
        width: calc(100% - 16px);
        height: 56px;
        bottom: 0;
        left: 0;
        padding: 8px;
        padding-bottom: 0;
        margin-top: 0;
        border-top: 1px solid var(--theme-border);
    }
</style>
