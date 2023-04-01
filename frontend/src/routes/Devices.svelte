<script lang="ts">
    import { onMount } from "svelte";

    import CheckLabel from "../components/CheckLabel.svelte";

    import Devices, { type Device } from "./devices";

    let devices: Device[] = [];
    let selected: Device[] = [];

    onMount(() => {
        devices = Devices.get();
    });
</script>

<svelte:head>
    <title>Pico Web | Devices</title>
</svelte:head>

<!-- TODO: ...
    ...split view (landscape or horizontal orientation)
    ...checkbox like group of devices for selecting device(s) to control
    ...button group for on/off
    ...input group (range slider) for rgbw (0-100)
    ...color preview somewhere
    ...color history (stored on client side)
-->

<div class="devices container">
    <section class="list">
        <fieldset>
            <legend>Devices</legend>

            <div class="content list">
                {#each devices as device }
                    <CheckLabel
                        checked={!!selected.find(d => d === device)}
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
