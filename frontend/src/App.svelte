<script lang="ts">
    import { onMount } from "svelte";

    onMount(() => {
        // ...
    });
</script>

<svelte:head>
    <!-- SMUI Light Theme -->
    <link rel="stylesheet" href="/theme.css" />
    <title>Pico Web | Home</title>
</svelte:head>

<div class="container debug">
    <div class="devices-container debug" />
    <div class="spacer" />
    <div class="control-container debug" />
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

        overflow: hidden;
        overflow-x: scroll;
        scroll-snap-type: x mandatory;
    }

    .container > div:not(.spacer) {
        min-width: calc(100vw - 32px);
        width: calc(100vw - 32px);
        max-width: calc(100vw - 32px);
        height: calc(100% - 32px);
        margin: 16px 0;

        scroll-snap-align: center;
    }

    .container > div:not(.spacer):last-child {
        margin-right: 16px;
    }

    .container > div:not(.spacer):first-child {
        margin-left: 16px;
    }

    /* TODO: adding a line which connects the divs connected to it (use ::before and ::after) */
    .container > div.spacer {
        min-width: 32px;
        width: 32px;
        max-width: 32px;
        height: 100%;
    }
</style>
