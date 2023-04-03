<script lang="ts">
    import { onMount } from "svelte";

    import CheckLabel from "../components/CheckLabel.svelte";
    import ColorPicker from "../components/ColorPicker.svelte";
    import PowerToggle from "../components/PowerToggle.svelte";

    import Api, { type Device } from "../ts/api";

    let selected: Device[] = [];
    $: console.log("selected:", selected);

    let devices: Device[] = [];
    $: {
        const newSelected = []
        for (const s of selected) {
            if (!!devices.find(d => d.addr === s.addr)) {
                newSelected.push(s)
            }
        }
        selected = newSelected
    }

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
                {#each devices as device }
                    <CheckLabel
                        checked={!!selected.find(sd => sd.addr === device.addr)}
                        label={device.addr}
                        on:change={() => {
                            if (!!selected.find(d => d === device)) {
                                selected = selected.filter(d => d != device);
                            } else {
                                selected = [...selected, device];
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
                <ColorPicker />
            </div>
            <div class="bottom">
                <PowerToggle />
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

    @media (min-width: 1025px) {
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

    /* TODO: remove scroll bar */
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
        width: calc(100% - 16px);
        height: 64px;
        bottom: 0;
        left: 0;
        padding: 8px;
        margin-top: 0;
        border-top: 1px solid var(--theme-border);
        border-bottom: 1px solid red;
    }
</style>
