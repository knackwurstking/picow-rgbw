<script lang="ts">
    export let name = "";
    export let min = 0;
    export let max = 100;
    export let value = 50;

    interface Position {
        x: number;
        y: number
    }

    let container: HTMLDivElement;
    let pointer: Position = null;
    $: {
        if (pointer) {
            const rect = container.getBoundingClientRect();
            const width = rect.right - rect.left; // 100% == max value
            const v = Math.round(pointer.x / (width / max));

            if (v < 0) {
                value = 0;
            } else if (v > 100) {
                value = 100
            } else {
                value = v
            }
        }
    };
</script>

<div bind:this={container} class="container"
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
            y: ev.clientY - rect.top,
        }
    }}
    on:pointerout={(ev) => {
        if (!pointer || ev.buttons === 0) return;
        const rect = container.getBoundingClientRect();
        if (ev.clientX < rect.left) {
            pointer = {
                x: 0,
                y: ev.clientY - rect.top,
            }
        } else if (ev.clientX > rect.right) {
            pointer = {
                x: rect.width,
                y: ev.clientY - rect.top,
            }
        } else {
            pointer = {
                x: ev.clientX - rect.left,
                y: ev.clientY - rect.top,
            }
        }
    }}
    on:pointerup={() => {
        if (!pointer) {
            pointer = null;
        }
    }}
>
    <div class="track"/>
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
</div>

<style>
    .container {
        position: relative;
        width: 100%;
        font-size: 1.5rem;
        display: flex;
        justify-content: center;
        place-items: center;
        user-select: none;
        touch-action: none;
    }

    .container .track {
        position: absolute;
        background: var(--theme-secondary);
        width: 100%;
        height: 4px;
        left: 0;
    }

    .container .progress {
        position: absolute;
        background: var(--theme-primary);
        height: 4px;
        left: 0;
    }

    .container .thumb {
        position: absolute;
        height: 100%;
        width: 16px;
        background: var(--theme-primary);
    }
</style>
