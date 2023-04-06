<script lang="ts">
    import { createEventDispatcher } from "svelte";

    import StatusLED from "./StatusLED.svelte";

    const dispatch = createEventDispatcher();

    export let checked: boolean = false;
    export let label: string = "";
    export let offline: boolean = false;
</script>

<label class:checked>
    <input
        disabled={offline}
        type="checkbox"
        {checked}
        on:change={() => dispatch("change", { checked: checked })}
    />
    {label}
    <StatusLED
        style="
            float: right;
        "
        active={!offline}
    />
</label>

<style>
    label {
        margin: 8px;
        padding: 16px;
        border: 1px solid var(--theme-border);
        padding-left: 16px;
        background-color: transparent;
        transition: background-color 0.35s ease;
        user-select: none;
    }

    label:hover {
        background-color: var(--theme-primary--hover);
    }

    label.checked {
        background-color: var(--theme-primary--focus);
    }

    label input {
        display: none;
    }
</style>
