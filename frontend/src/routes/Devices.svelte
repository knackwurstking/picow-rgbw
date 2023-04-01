<script lang="ts">
    import { onMount } from "svelte";

    import CheckLabel from "../components/CheckLabel.svelte";

    import Api, { type Device } from "../ts/api";

    let selectedDevices: Device[] = [];
    $: console.log("selected:", selectedDevices);

    let devices: Device[] = [];
    $: {
        const newSelectedDevices = []
        for (const selectedDevice of selectedDevices) {
            if (!!devices.find(d => d.addr === selectedDevice.addr)) {
                newSelectedDevices.push(selectedDevice)
            }
        }
        selectedDevices = newSelectedDevices
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
                        checked={!!selectedDevices.find(sd => sd.addr === device.addr)}
                        label={device.addr}
                        on:change={() => {
                            if (!!selectedDevices.find(d => d === device)) {
                                selectedDevices = selectedDevices.filter(d => d != device);
                            } else {
                                selectedDevices = [...selectedDevices, device];
                            }
                        }}
                    />
                {/each}
            </div>
        </fieldset>
    </section>

    <section class="ctrl">
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
</style>
