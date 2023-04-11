<script lang="ts">
    export let min: number = 0;
    export let max: number = 100;
    export let value: number = 100;
</script>

<div class="container" {...$$restProps}>
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
        width: 18px;
        cursor: pointer;
    }

    input::-webkit-slider-thumb {
        -webkit-appearance: none;
        appearance: none;
        background: var(--theme-primary);
        width: 100%;
        height: 18px;
    }

    input::-moz-range-thumb {
        border: none;
        background: var(--theme-primary);
        width: 100%;
        height: 18px;
    }
</style>
