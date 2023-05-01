<script lang="ts">
    import { onMount } from "svelte";

    import Checkbox from "@smui/checkbox";
    import Button, { Group, Label } from "@smui/button";
    import List, {
        Item,
        Text,
        PrimaryText,
        SecondaryText,
        Separator,
        Meta,
    } from "@smui/list";

    import StatusLED from "./lib/components/StatusLED.svelte";
    import ColorPicker from "./lib/components/ColorPicker.svelte";

    import api, { type Device, type Duty } from "./lib/ts/api";

    // NOTE: Devices
    let devices: Device[] = [];
    let selected: Device[] = [];
    $: {
        const newSelected = [];
        for (const s of selected) {
            if (!!devices.find((d) => d.addr === s.addr)) {
                newSelected.push(s);
            }
        }
        selected = newSelected;
    }

    // NOTE: Controls
    let control: HTMLDivElement;

    let r: number = 100;
    let g: number = 100;
    let b: number = 100;
    let w: number = 100;

    let color: Duty[] = [255, 255, 255, 255];
    $: {
        color = [r, g, b, w];
    }

    // NOTE: Actions
    let actionButtonLabel: "SET" | "ON" = "ON";

    // Turn selected off selected devices. (action button handler)
    function _off() {
        api.postDevices(
            ...selected.map((d) => ({
                addr: d.addr,
                rgbw: [0, 0, 0, 0],
            }))
        );
    }

    // Turn on or set color for selected devices. (action button handler)
    function _set() {
        if (actionButtonLabel === "ON") {
            api.postDevices(
                ...selected.map((d) => {
                    let c = color;

                    const storedColor = window.localStorage.getItem(
                        `color:${d.addr}`
                    );
                    if (storedColor) {
                        c = JSON.parse(storedColor);
                    }

                    return {
                        addr: d.addr,
                        rgbw: c,
                    };
                })
            );

            return;
        }

        if (actionButtonLabel === "SET") {
            api.postDevices(
                ...selected.map((d) => {
                    return {
                        addr: d.addr,
                        rgbw: color,
                    };
                })
            );

            return;
        }
    }

    onMount(() => {
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
    });
</script>

<svelte:head>
    <link
        rel="stylesheet"
        href="node_modules/svelte-material-ui/themes/svelte-dark.css"
    />

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

            <List checkList>
                {#each devices as device}
                    <Item style="height: 65px;">
                        <Text>
                            <PrimaryText>{device.addr}</PrimaryText>

                            <SecondaryText>
                                [{device.rgbw.map((gp) => gp.duty).join(",")}]
                            </SecondaryText>
                        </Text>

                        <Meta>
                            <Checkbox
                                style="margin-right: 8px;"
                                bind:group={selected}
                                value={device}
                            />

                            <StatusLED
                                style="
                                  position: absolute;
                                  top: 4px;
                                  right: 4px;
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

                <div class="scene" />

                <div class="store" />

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

    :global(:root) {
        --mdc-theme-primary: #ff3e00;
        --mdc-theme-on-primary: #fff;
        --mdc-theme-secondary: #018786;
        --mdc-theme-on-secondary: #fff;
        --mdc-theme-surface: #222222;
        --mdc-theme-on-surface: #fff;

        --mdc-shape-medium: 4px;

        --theme-primary: var(--mdc-theme-primary);
        --theme-secondary: var(--mdc-theme-secondary);
        --theme-border: rgba(255, 255, 255, 0.12);
        --theme-border-radius: var(--mdc-shape-medium);
    }

    :global(fieldset) {
        border-radius: var(--theme-border-radius);
        border-color: var(--theme-border);
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

    main .x-scroll fieldset.devices div.spacer {
        position: relative;
        min-width: 32px;
        width: 32px;
        max-width: 32px;
        height: 100%;
        margin: 16px 0;
    }

    main .x-scroll fieldset.devices div.spacer::before {
        content: "";
        position: absolute;
        top: calc(50% - 1px);
        left: 0;
        width: 100%;
        height: 2px;
        background: var(--theme-border);
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
        min-width: calc(100vw - 32px);
        width: calc(100vw - 32px);
        max-width: calc(100vw - 32px);
        height: 100%;

        overflow: hidden;
        overflow-y: auto;
        scroll-behavior: smooth;
        scroll-snap-align: center;
    }

    main .x-scroll fieldset.control {
        min-width: calc(100vw - 32px);
        width: calc(100vw - 32px);
        max-width: calc(100vw - 32px);
        height: 100%;
        margin-right: 16px;

        display: flex;
        flex-direction: column;

        overflow: hidden;
        overflow-y: auto;
        scroll-behavior: smooth;
        scroll-snap-align: center;
    }

    main .x-scroll fieldset.control .scene {
    }

    main .x-scroll fieldset.control .store {
    }
</style>
