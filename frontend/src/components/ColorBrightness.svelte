<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let min: number = 5;
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
        on:input={() => dispatch("change", { value })}
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
        height: calc(100% - 16px);
        margin: 8px 0;
    }
</style>
