<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
</script>

<button {...$$props} on:click={() => dispatch("click")}><slot /></button>

<style>
    button {
        background-color: var(--theme-surface);
        border-color: var(--theme-border);
        color: var(--theme-on-surface);
        transition: transform 0.1s ease, background-color 0.1 ease;
    }

    button:hover {
        background-color: var(--theme-primary--hover);
    }

    button:active {
        background-color: var(--theme-primary--focus);
    }

    button:active {
        transform: scale(0.99);
    }
</style>
