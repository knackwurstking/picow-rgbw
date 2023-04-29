<script lang="ts">
    import Devices from "./lib/Devices.svelte";
    import Control from "./lib/Control.svelte";
    import type { Device } from "./lib/api";

    let r: number = 100;
    let g: number = 100;
    let b: number = 100;
    let w: number = 100;

    let selected: Device[] = [];
    $: console.debug("[app]", { selected });
</script>

<svelte:head>
    <link
        rel="stylesheet"
        href="node_modules/svelte-material-ui/themes/material-dark.css"
    />

    <title>Pico Web | Home</title>
</svelte:head>

<div class="container">
    <div class="x-scroll">
        <Devices
            style="margin-left: 16px;"
            bind:r
            bind:g
            bind:b
            bind:w
            bind:selected
        />
        <div class="spacer" />
        <Control style="margin-right: 16px;" />
    </div>
    <fieldset class="action">
        <legend>Action</legend>
        <!-- TODO: Add ON/OFF buttons, change ON to SET if .control if visible? -->
    </fieldset>
</div>

<style>
    :global(html, body) {
        overflow: hidden;
    }

    :global(:root) {
        --theme-border: rgba(255, 255, 255, 0.12);
        --theme-border-radius: var(--mdc-shape-medium, 4px);
    }

    :global(fieldset) {
        border-radius: var(--theme-border-radius);
        border-color: var(--theme-border);
    }

    .container,
    .x-scroll {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;
    }

    .x-scroll {
        height: calc(100% - 16px - 90px);
        padding: 8px 0;
        padding-bottom: 16px;
        z-index: 1;

        overflow: hidden;
        overflow-x: scroll;
        scroll-snap-type: x mandatory;

        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;
    }

    div.spacer {
        position: relative;
        min-width: 32px;
        width: 32px;
        max-width: 32px;
        height: 100%;
        margin: 16px 0;
    }

    div.spacer::before {
        content: "";
        position: absolute;
        top: calc(50% - 1px);
        left: 0;
        width: 100%;
        height: 2px;
        background: var(--theme-border);
    }

    .action {
        position: absolute;
        bottom: 0;
        left: 0;
        width: calc(100% - 32px);
        height: 95px;
        margin: 8px 16px;
    }
</style>
