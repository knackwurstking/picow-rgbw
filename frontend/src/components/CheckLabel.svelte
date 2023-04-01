<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let checked: boolean = false;
    export let label: string = "";
</script>

<label><input
    type="checkbox"
    checked={checked}
    on:change={() => dispatch("change", { checked: checked })}
/>{label}</label>

<style>
    label {
        margin: 8px;
        padding: 16px;
        border: 1px solid var(--theme-border);
    }
</style>
