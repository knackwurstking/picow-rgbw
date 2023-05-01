<script lang="ts">
  import { createEventDispatcher } from "svelte";

  //export let variant: "text" | "raised" | "unelevated" | "outlined" =
  export let variant: "unelevated" | "outlined" = "unelevated";
  export let color: "primary" | "secondary" = "primary";

  const dispatch = createEventDispatcher();

  // TODO: add ripple effect, just like smui buttons
</script>

<button
  class={`custom-button ${variant} ${color}`}
  style={`` + $$props.style || ""}
  {...$$restProps}
  on:click={(ev) => dispatch("click", ev)}
>
  <slot />
</button>

<style>
  .custom-button {
    margin-left: 0;
    margin-right: 0;
  }
</style>
