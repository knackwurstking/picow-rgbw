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
</div>

<style>
    div.devices.container {
        display: flex;
        width: 100%;
        height: 100%;
    }

    /* TODO: handle orientation landscape and portrait via available pixel with */
</style>
