<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let min = 0;
    export let max = 100;
    export let value = 50;
    $: value >= min && dispatch("input");
    export let orient: "horizontal" | "vertical" = "horizontal";

    interface Position {
        x: number;
        y: number;
    }

    let container: HTMLDivElement;
    let pointer: Position = null;
    $: {
        if (pointer) {
            const rect = container.getBoundingClientRect();

            let v: number;
            if (orient === "vertical") {
                const height = rect.bottom - rect.top; // 100% == max value
                v = Math.round(pointer.y / (height / max));
            } else {
                const width = rect.right - rect.left; // 100% == max value
                v = Math.round(pointer.x / (width / max));
            }

            if (v < min) {
                value = min;
            } else if (v > 100) {
                value = 100;
            } else {
                value = v;
            }
        }
    }
</script>

<div
    bind:this={container}
    class="custom-slider__container"
    class:vertical={orient === "vertical"}
    on:pointerdown={(ev) => {
        const rect = container.getBoundingClientRect();
        pointer = {
            x: ev.clientX - rect.left,
            y: rect.height - (ev.clientY - rect.top),
        };
    }}
    on:pointermove={(ev) => {
        if (!pointer || ev.buttons === 0) return;
        const rect = container.getBoundingClientRect();
        pointer = {
            x: ev.clientX - rect.left,
            y: rect.height - (ev.clientY - rect.top),
        };
    }}
    on:pointerout={(ev) => {
        if (!pointer || ev.buttons === 0) return;
        const rect = container.getBoundingClientRect();
        if (ev.clientY < rect.top) {
            pointer = {
                x: ev.clientX - rect.left,
                y: rect.height,
            };
        } else if (ev.clientY > rect.bottom) {
            pointer = {
                x: ev.clientX - rect.left,
                y: 0,
            };
        } else {
            pointer = {
                x: ev.clientX - rect.left,
                y: rect.height - (ev.clientY - rect.top),
            };
        }
    }}
    on:pointerup={() => {
        if (!pointer) {
            pointer = null;
        }
    }}
    {...$$restProps}
>
    <div class="custom-slider__track" />

    {#if orient === "vertical"}
        <div
            class="custom-slider__progress"
            style={`
                height: ${value}%;
            `}
        />
        <div
            class="custom-slider__thumb"
            style={`
                top: calc(${value}% - 12px);
            `}
        />
    {:else}
        <div
            class="custom-slider__progress"
            style={`
                width: ${value}%;
            `}
        />
        <div
            class="custom-slider__thumb"
            style={`
                left: calc(${value}% - 12px);
            `}
        />
    {/if}
</div>

<style>
    .custom-slider__container {
        position: relative;
        width: 100%;
        height: 1.75rem;
        font-size: 1.75rem;
        display: flex;
        justify-content: center;
        align-items: center;
        user-select: none;
        touch-action: none;
    }

    .custom-slider__container.vertical {
        width: 1.75rem;
        height: calc(100% - 8px);
        transform: rotate(180deg);
    }

    .custom-slider__container:not(.vertical) .custom-slider__track {
        position: absolute;
        width: 100%;
        height: 6px;
        left: 0;
    }

    .custom-slider__container.vertical .custom-slider__track {
        position: absolute;
        width: 6px;
        height: 100%;
        top: 0;
    }

    .custom-slider__container:not(.vertical) .custom-slider__progress {
        position: absolute;
        height: 6px;
        left: 0;
    }

    .custom-slider__container.vertical .custom-slider__progress {
        position: absolute;
        height: 100%;
        width: 6px;
        top: 0;
    }

    .custom-slider__container .custom-slider__thumb {
        position: absolute;
        height: 24px;
        width: 24px;
    }
</style>
