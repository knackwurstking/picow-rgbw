<script lang="ts">
    import { createEventDispatcher } from "svelte";

    import Button from "./Button.svelte";

    const dispatch = createEventDispatcher();
</script>

<div class="container"
    {...$$restProps}
>
    <Button
        style="
            width: 40%;
            height: 100%;
            max-width: 100px;
        "
        class="off"
        on:click={() => dispatch("change", { state: "off" })}
    ><span class="off">OFF</span></Button>
    <Button
        style="
            width: 40%;
            height: 100%;
            max-width: 100px;
        "
        class="set"
        on:click={() => dispatch("change", { state: "set" })}
    ><span class="set">SET<span></Button>
</div>

<style>
    div.container {
        display: flex;
        justify-content: space-evenly;
        align-items: center;
    }

    div.container span {
        font-size: 1.5rem;
        font-weight: bolder;
    }

    div.container span.off {
        text-shadow: 0 0 0.25rem red;
    }

    div.container span.set {
        text-shadow: 0 0 0.35rem green;
    }
</style>
