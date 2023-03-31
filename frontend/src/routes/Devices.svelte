<script lang="ts">
    import devices from "./devices";
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
        border: 1px solid red; 
        width: 100%;
        height: 100%;
    }

    div.devices.container > section.list {
        height: 50%;
    }

    div.devices.container > section.ctrl {
        height: 50%;
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
</style>
