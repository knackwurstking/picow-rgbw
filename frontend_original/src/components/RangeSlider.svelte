<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let min = 0;
    export let max = 100;
    export let value = 50;
    $: value >= min &&  dispatch("input");
    export let orient: "horizontal" | "vertical" = "horizontal";

    interface Position {
        x: number;
        y: number
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
                value = 100
            } else {
                value = v
            }
        }
    };
</script>

<div bind:this={container} class="container" class:vertical={orient === "vertical"}
    on:pointerdown={(ev) => {
        const rect = container.getBoundingClientRect();
        pointer = {
            x: ev.clientX - rect.left,
            y: ev.clientY - rect.top,
        };
    }}
    on:pointermove={(ev) => {
        if (!pointer || ev.buttons === 0) return;
        const rect = container.getBoundingClientRect();
        pointer = {
            x: ev.clientX - rect.left,
            y: rect.height - (ev.clientY - rect.top),
        }
    }}
    on:pointerout={(ev) => {
        if (!pointer || ev.buttons === 0) return;
        const rect = container.getBoundingClientRect();
        if (ev.clientY < rect.top) {
            pointer = {
                x: ev.clientX - rect.left,
                y: rect.height,
            }
        } else if (ev.clientY > rect.bottom) {
            pointer = {
                x: ev.clientX - rect.left,
                y: 0,
            }
        } else {
            pointer = {
                x: ev.clientX - rect.left,
                y: rect.height - (ev.clientY - rect.top),
            }
        }
    }}
    on:pointerup={() => {
        if (!pointer) {
            pointer = null;
        }
    }}
    {...$$restProps}
>
    <div class="track"/>
    {#if orient === "vertical"}
        <div class="progress"
            style={`
                height: ${value}%;
            `}
        />
        <div class="thumb"
            style={`
                top: calc(${value}% - 8px);
            `}
        />
    {:else}
        <div class="progress"
            style={`
                width: ${value}%;
            `}
        />
        <div class="thumb"
            style={`
                left: calc(${value}% - 8px);
            `}
        />
    {/if}
</div>

<style>
    .container {
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

    .container.vertical {
        width: 1.75rem;
        height: 100%;
        transform: rotate(180deg);
    }

    .container:not(.vertical) .track {
        position: absolute;
        background: var(--theme-secondary);
        width: 100%;
        height: 4px;
        left: 0;
    }

    .container.vertical .track {
        position: absolute;
        background: var(--theme-secondary);
        width: 4px;
        height: 100%;
        top: 0;
    }

    .container:not(.vertical) .progress {
        position: absolute;
        background: var(--theme-primary);
        height: 4px;
        left: 0;
    }

    .container.vertical .progress {
        position: absolute;
        background: var(--theme-primary);
        height: 100%;
        width: 4px;
        top: 0;
    }

    .container .thumb {
        position: absolute;
        height: 100%;
        width: 16px;
        background: var(--theme-primary);
    }

    .container.vertical .thumb {
        height: 16px;
        width: 100%;
    }
</style>
