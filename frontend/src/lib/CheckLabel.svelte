<script lang="ts">
    import { createEventDispatcher } from "svelte";

    import StatusLED from "./StatusLED.svelte";

    const dispatch = createEventDispatcher();

    export let label: string = "";
    export let currentColor: number[] = [0, 0, 0, 0];
    export let offline: boolean = false;

    let checked: boolean = false;
    $: dispatch("change", { checked: checked });
</script>

<label class:checked {...$$restProps}>
    <StatusLED
        style="
            float: right;
        "
        active={!offline}
    />
    <input
        type="checkbox"
        checked={!offline && checked}
        on:input={() => (checked = !checked)}
    />
    {label}
    <br />
    <code style="font-size:0.85rem;">[{currentColor.join(",")}]</code>
</label>

<style>
    label {
        display: block;
        padding: 8px;
        margin: 8px;
        border: 1px solid var(--theme-border);
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
