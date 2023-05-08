<script lang="ts">
    import { onMount } from "svelte";

    import { Separator } from "@smui/list"; // TODO: replace with my own list stuff

    import Checkbox from "svelteui/src/checkbox";
    import List, { Item, Meta } from "svelteui/src/list";
    import Button, { Group, Label } from "svelteui/src/button";

    import StatusLED from "svelteui/src/misc/StatusLED.svelte";
    import ColorPicker from "svelteui/src/misc/ColorPicker.svelte";

    import api, { type Device, type Duty } from "./lib/api";

    let mounted = false;
    let actionButtonLabel: "SET" | "ON" = "ON";

    // NOTE: Devices
    let devices: Device[] = [];
    let selected: string[] = [];

    // NOTE: Controls
    let control: HTMLDivElement;

    let r: number = 100;
    let g: number = 100;
    let b: number = 100;
    let w: number = 100;

    let color: Duty[];

    // update local storage if selected (group) changed
    $: !!selected.length &&
        window.localStorage.setItem("selected", JSON.stringify(selected));

    // update local storage if color (rgbw) changed
    $: {
        color = [r, g, b, w];
        if (mounted) {
            window.localStorage.setItem("color", JSON.stringify(color));
        }
    }

    // Turn selected off selected devices. (action button handler)
    function _off() {
        api.postDevices(
            ...selected.map((a) => ({
                addr: a,
                rgbw: [0, 0, 0, 0],
            }))
        );
    }

    // Turn on or set color for selected devices. (action button handler)
    function _set() {
        if (actionButtonLabel === "ON") {
            api.postDevices(
                ...selected.map((a) => {
                    let c = color;

                    const storedColor = window.localStorage.getItem(
                        `color:${a}`
                    );
                    if (storedColor) {
                        c = JSON.parse(storedColor);
                    }

                    return { addr: a, rgbw: c };
                })
            );

            return;
        }

        if (actionButtonLabel === "SET") {
            api.postDevices(
                ...selected.map((a) => {
                    return { addr: a, rgbw: color };
                })
            );

            return;
        }
    }

    onMount(() => {
        // get color from storage for the color picker component
        try {
            color = JSON.parse(
                window.localStorage.getItem("color") || `[${r},${g},${b},${w}]`
            );

            r = color[0] || 0;
            g = color[1] || 0;
            b = color[2] || 0;
            w = color[3] || 0;
        } catch (err) {
            console.warn("parse color from local storage failed:", err);
        }

        // sse: "offline"
        api.addEventListener("offline", () => {
            console.debug(`[app, event] "offline"`);

            devices.forEach((d) => (d.offline = true));
            devices = devices;
        });

        // sse: "devices"
        api.addEventListener("devices", (data: Device[]) => {
            console.debug(`[app, event] "devices"`);

            devices = data;

            // load previous selected devices from the localStorage
            const l: string[] = JSON.parse(
                window.localStorage.getItem("selected") || "[]"
            );
            selected = devices
                .filter((d) => l.find((a) => a === d.addr))
                .map((d) => d.addr);
        });

        // sse: "devices" (store color)
        api.addEventListener("devices", async (data: Device[]) => {
            console.debug(`[app, event] "devices" (store color)`);

            data.forEach((d) => {
                if (d.rgbw.find((gp) => gp.duty > 0))
                    window.localStorage.setItem(
                        `color:${d.addr}`,
                        JSON.stringify(d.rgbw.map((gp) => gp.duty))
                    );
            });
        });

        // sse: "device", "device" (store color)
        api.addEventListener("device", (data: Device) => {
            console.debug(`[app, event] "device"`, data);

            const device = devices.find((d) => d.addr === data.addr);
            if (!device) return;

            device.rgbw = data.rgbw;
            device.offline = data.offline;

            devices = devices;
        });

        api.addEventListener("device", async (data: Device) => {
            console.debug(`[app, event] "device" (store color)`);

            if (data.rgbw.find((gp) => gp.duty > 0))
                window.localStorage.setItem(
                    `color:${data.addr}`,
                    JSON.stringify(data.rgbw.map((gp) => gp.duty))
                );
        });

        mounted = true;
    });
</script>

<svelte:head>
    <title>Pico Web | Home</title>
</svelte:head>

<main>
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
        <!-- "Devices" view to the left -->
        <fieldset class="devices">
            <legend>Devices</legend>

            <List checkable multiple checklist bind:group={selected}>
                {#each devices as device}
                    <Item
                        primaryText={device.addr}
                        secondaryText={`[${device.rgbw
                            .map((gp) => gp.duty)
                            .join(",")}]`}
                        value={device.addr}
                        checked={!!selected.find((a) => a === device.addr)}
                    >
                        <Meta slot="right">
                            <Checkbox
                                disableUserActions
                                style="margin-right: 28px; float: right;"
                                group={selected}
                                value={device.addr}
                            />

                            <StatusLED
                                style="
                                  position: absolute;
                                  top: 4px;
                                  right: 4px;
                                  font-size: 0.8rem;
                                "
                                active={!device.offline}
                            />
                        </Meta>
                    </Item>
                    <Separator />
                {/each}
            </List>
        </fieldset>

        <div class="spacer" />

        <!-- "Control" view to the right -->
        <div bind:this={control}>
            <fieldset class="control">
                <legend>Control</legend>

                <div class="spacer" style="height: 100%;" />

                <div class="scene">
                    <!-- TODO: manage scenes here (Quick and Easy) -->
                </div>

                <div class="store">
                    <!-- TODO: store color presets (rgb) -->
                </div>

                <ColorPicker bind:r bind:g bind:b bind:w />
            </fieldset>
        </div>
    </div>

    <!-- Actions "OFF", "ON|SET" -->
    <fieldset class="action">
        <legend>Action</legend>
        <Group
            variant="unelevated"
            style="display: flex; justify-content: stretch;"
        >
            <Button
                variant="unelevated"
                color="primary"
                style="flex-grow: 1;"
                on:click={() => _off()}
            >
                <Label>OFF</Label>
            </Button>
            <Button
                variant="outlined"
                color="primary"
                style="flex-grow: 1;"
                on:click={() => _set()}
            >
                <Label>{actionButtonLabel}</Label>
            </Button>
        </Group>
    </fieldset>
</main>

<style>
    :global(html, body) {
        overflow: hidden;
    }

    :global(fieldset) {
        border-radius: var(--theme-border-radius);
        border-color: var(--theme-border-color);
    }

    main,
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

    main .x-scroll > div.spacer {
        position: relative;
        min-width: 8px;
        width: 8px;
        max-width: 8px;
        height: 100%;
        margin: 16px 0;
    }

    main .x-scroll > div.spacer::before {
        content: "";
        position: absolute;
        top: calc(50% - 1px);
        left: 0;
        width: 100%;
        height: 2px;
        background: var(--theme-border-color);
    }

    main .action {
        position: absolute;
        bottom: 0;
        left: 0;
        width: calc(100% - 32px);
        height: 80px;
        margin: 8px 16px;
    }

    main .x-scroll fieldset.devices {
        position: relative;

        min-width: calc(100vw - 32px);
        width: calc(100vw - 32px);
        max-width: calc(100vw - 32px);
        height: 100%;

        margin-left: 8px;

        overflow: hidden;

        scroll-snap-align: center;
    }

    main .x-scroll fieldset.control {
        position: relative;

        min-width: calc(100vw - 32px);
        width: calc(100vw - 32px);
        max-width: calc(100vw - 32px);
        height: 100%;

        margin-right: 8px;

        display: flex;
        flex-direction: column;

        overflow: hidden;
        overflow-y: auto;
        scroll-behavior: smooth;
        scroll-snap-align: center;
    }

    /*
    main .x-scroll fieldset.control .scene {
    }

    main .x-scroll fieldset.control .store {
    }
    */
</style>
