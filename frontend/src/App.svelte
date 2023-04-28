<script lang="ts">
    import { onMount } from "svelte";

    import Devices from "./lib/Devices.svelte";
    import Control from "./lib/Control.svelte";

    onMount(() => {
        // ...
    });
</script>

<svelte:head>
    <!-- SMUI Light Theme -->
    <link rel="stylesheet" href="/theme.css" />
    <title>Pico Web | Home</title>
</svelte:head>

<div class="container">
    <Devices class="devices-container" style="margin-left: 16px;" />
    <div class="spacer" />
    <Control class="control-container" style="margin-right: 16px;" />
</div>

<style>
    :global(html, body) {
        overflow: hidden;
    }

    .container {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;

        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;

        overflow: hidden;
        overflow-x: scroll;
        scroll-snap-type: x mandatory;
    }

    .container > div.spacer {
        position: relative;
        min-width: 32px;
        width: 32px;
        max-width: 32px;
        height: 100%;
        margin: 16px 0;
    }

    .container > div.spacer::before {
        content: "";
        position: absolute;
        top: calc(50% - 1px);
        left: 0;
        width: 100%;
        height: 2px;
        background: var(--theme-border);
    }
</style>
