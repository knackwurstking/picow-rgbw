<script lang="ts">
    import Button, { Group, Label } from "@smui/button";

    import Devices from "./lib/Devices.svelte";
    import Control from "./lib/Control.svelte";
    import type { Device } from "./lib/ts/api";

    let control: HTMLDivElement;
    let actionButtonLabel: "SET" | "ON" = "ON";

    let color: Color = {
        r: 255,
        b: 255,
        g: 255,
        w: 255,
    };

    let selected: Device[] = [];
    $: console.debug("[app]", { selected });

    // TODO: ON/OFF handler (if actionButtonLabel === "ON")
    //       using color stored in localStorage (or the color configured
    //       on the control view)
</script>

<svelte:head>
    <link
        rel="stylesheet"
        href="node_modules/svelte-material-ui/themes/material-dark.css"
    />

    <title>Pico Web | Home</title>
</svelte:head>

<div class="container">
    <div
        class="x-scroll"
        on:scroll={() => {
            if (!control) return;
            const r = control.getBoundingClientRect();
            if (r.left < window.innerWidth / 2) {
                actionButtonLabel = "SET";
            } else {
                actionButtonLabel = "ON";
            }
        }}
    >
        <Devices style="margin-left: 16px;" bind:selected />
        <div class="spacer" />
        <div bind:this={control}>
            <Control style="margin-right: 16px;" bind:color />
        </div>
    </div>
    <fieldset class="action">
        <legend>Action</legend>
        <Group
            variant="unelevated"
            style="display: flex; justify-content: stretch;"
        >
            <Button variant="unelevated" color="primary" style="flex-grow: 1;">
                <Label>OFF</Label>
            </Button>
            <Button variant="outlined" color="primary" style="flex-grow: 1;">
                <Label>{actionButtonLabel}</Label>
            </Button>
        </Group>
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
        height: calc(100% - 16px - 75px);
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
        height: 80px;
        margin: 8px 16px;
    }
</style>
