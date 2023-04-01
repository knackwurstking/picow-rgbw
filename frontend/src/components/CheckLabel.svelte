<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let checked: boolean = false;
    export let label: string = "";
</script>

<label class:checked><input
    type="checkbox"
    checked={checked}
    on:change={() => dispatch("change", { checked: checked })}
/>{label}</label>

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

    label.checked {
        background-color: var(--theme-primary);
    }

    label input {
        display: none;
    }
</style>
