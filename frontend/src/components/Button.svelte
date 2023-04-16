<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
</script>

<button {...$$props} on:click={() => dispatch("click")}><slot /></button>

<style>
    button {
        position: relative;
        background-color: transparent;
        border-color: var(--theme-border);
        color: var(--theme-on-surface);
        transition: transform 0.1s ease, background-color 0.1 ease;
    }

    button::before {
        content: "";
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        z-index: -1;
        background-color: var(--theme-surface);
        filter: blur(5px);
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
