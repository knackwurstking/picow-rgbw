<script lang="ts">
  import { createEventDispatcher } from "svelte";

  //export let variant: "text" | "raised" | "unelevated" | "outlined" =
  export let variant: "unelevated" | "outlined" = "unelevated";
  export let color: "primary" | "secondary" = "primary";

  const dispatch = createEventDispatcher();

  let button: HTMLButtonElement;

  // TODO: add ripple effect, just like smui buttons
  function _click(
    ev: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement }
  ) {
    const r = button.getBoundingClientRect();

    const cursorX = Math.round(ev.pageX - r.x);
    const cursorY = Math.round(ev.pageY - r.y);

    // TODO: set ripple top and left positions
    // TODO: set width and height to 0 first
    // TODO: set display to block
    // TODO: start animation on width and height
    // NOTE: animation time 2.5s?
  }
</script>

<button
  bind:this={button}
  class={`custom-button ${variant} ${color}`}
  style={`` + $$props.style || ""}
  {...$$restProps}
  on:click={_click}
>
  <slot />
  <div class="ripple" />
</button>

<style>
  .custom-button {
    margin-left: 0;
    margin-right: 0;

    position: relative;
  }

  .custom-button > .ripple {
    display: none;
    position: absolute;
    width: 0;
    height: 0;
  }
</style>
