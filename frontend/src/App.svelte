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

<div class="container">
    <fieldset class="devices-container" />
    <div class="spacer" />
    <fieldset class="control-container" />
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

    .container > fieldset {
        min-width: calc(100vw - 32px);
        width: calc(100vw - 32px);
        max-width: calc(100vw - 32px);
        height: calc(100% - 32px);
        margin: 16px 0;

        scroll-snap-align: center;
    }

    .container > fieldset:last-child {
        margin-right: 16px;
    }

    .container > fieldset:first-child {
        margin-left: 16px;
    }

    .container > div.spacer {
        position: relative;
        min-width: 32px;
        width: 32px;
        max-width: 32px;
        height: 100%;
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
