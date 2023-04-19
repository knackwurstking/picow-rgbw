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
            // TODO: Moving the ".thumb" here...
            // ...
        }
    }
</script>

<div bind:this={container} class="container"
    on:pointerdown={(ev) => {
        const rect = container.getBoundingClientRect();
        pointer = {
            x: ev.clientX - rect.left,
            y: ev.clientY - rect.top,
        };
        console.log("pointerdown", pointer);
    }}
    on:pointermove={(ev) => {
        if (!pointer) return;
        const rect = container.getBoundingClientRect();
        pointer = {
            x: ev.clientX - rect.left,
            y: ev.clientY - rect.top,
        }
    }}
    on:pointerup={(ev) => {
        pointer = null;
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
            left: calc(${value}% - 16px);
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
