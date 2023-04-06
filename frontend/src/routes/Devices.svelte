<script lang="ts">
    import { onMount } from "svelte";

    import Button from "../components/Button.svelte";
    import CheckLabel from "../components/CheckLabel.svelte";
    import ColorPicker from "../components/ColorPicker.svelte";
    import PowerToggle from "../components/PowerToggle.svelte";

    import Api, { type Device } from "../ts/api";

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

    let r: number = 100;
    let g: number = 100;
    let b: number = 100;
    let w: number = 100;

    onMount(async () => {
        devices = await Api.devices();
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
                        offline={device.offline}
                        label={device.addr}
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
                <ColorPicker bind:r bind:g bind:b bind:w />
            </div>
            <div class="bottom">
                <!-- TODO: replace refresh button with an icon type button? -->
                <Button
                    style="
                        width: 30%;
                        max-width: 90px;
                        height: 90%;
                    "
                    on:click={async () => {
                        devices = await Api.devices();
                    }}>Refresh</Button
                >
                <PowerToggle
                    style="
                        width: 60%;
                        height: 100%;
                    "
                    on:change={async (ev) => {
                        switch (ev.detail.state) {
                            case "set":
                                await Api.postDevices(
                                    ...selected.map((d) => ({
                                        addr: d.addr,
                                        rgbw: [r, g, b, w],
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
        overflow: hidden;
        overflow-y: auto;
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
