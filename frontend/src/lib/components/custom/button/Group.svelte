<script lang="ts">
  export let variant: "text" | "raised" | "unelevated" | "outlined" =
    "unelevated";
</script>

<div class="custom-button__group" {...$$restProps}>
  <slot />
</div>

<style>
  .custom-button__group {
    display: inline-flex;
  }
</style>
