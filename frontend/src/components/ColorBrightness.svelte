<script lang="ts">
    export let min: number = 0;
    export let max: number = 100;
    export let value: number = 100;
</script>

<div class="container debug" {...$$restProps}>
    <input
        type="range"
        {min}
        {max}
        bind:value
        orient="vertical"
    />
</div>

<style>
    .container {
        height: 100%;
        width: 100%;
        padding: 8px 4px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    input {
        -webkit-appearance: none;
        appearance: none;
        background-color: var(--theme-secondary);
        height: calc(100% - 16px);
        margin: 8px 0;
        width: 16px;
        cursor: pointer;
    }

    input::-webkit-slider-thumb {
    }

    input::-moz-range-thumb {
        border: none;
        background: var(--theme-primary);
        width: 100%;
        height: 16px;
    }
</style>
