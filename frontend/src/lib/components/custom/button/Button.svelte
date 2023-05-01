<script lang="ts">
  //export let variant: "text" | "raised" | "unelevated" | "outlined" =
  export let variant: "unelevated" | "outlined" = "unelevated";
  export let color: "primary" | "secondary" = "primary";

  // TODO: add ripple effect, just like smui buttons
</script>

<button
  class={`custom-button ${variant} ${color}`}
  style={`` + $$props.style || ""}
  {...$$restProps}
>
  <slot />
</button>

<style>
  .custom-button {
    margin-left: 0;
    margin-right: 0;
  }
</style>
