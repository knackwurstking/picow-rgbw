<script lang="ts"<>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
</script>

<div class="container">
    <button
        class="off"
        on:click={() => dispatch("change", { state: "off" })}
    >OFF</button>
    <button
        class="set"
        on:click={() => dispatch("change", { state: "set" })}
    >SET</button>
</div>

<style>
    div.container {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: space-evenly;
        align-items: center;
    }

    div.container button {
        height: 90%;
        width: 40%;
        max-width: 100px;
        background-color: var(--theme-surface);
        font-size: 1.5rem;
        font-weight: bolder;
        border-color: var(--theme-border);
        color: var(--theme-on-surface);
        transition: transform 0.1s ease, background-color 0.1 ease;
    }

    div.container button:hover {
        background-color: var(--theme-primary--hover);
    }

    div.container button:active {
        background-color: var(--theme-primary--focus);
    }

    div.container button.off {
        text-shadow: 0 0 0.25rem red;
    }

    div.container button.set{
        text-shadow: 0 0 0.35rem green;
    }

    div.container button:active {
        transform: scale(0.99);
    }
</style>
