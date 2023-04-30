<script lang="ts">
  export let color: Color = {
    r: 255,
    g: 255,
    b: 255,
    w: 255,
  };

  let r: number = 255;
  let g: number = 255;
  let b: number = 255;
  let w: number = 255;

  $: {
    // NOTE: does this work?
    console.log("[control] color:", r, g, b, w);
    color = { r, g, b, w };
  }
</script>

<fieldset class="control" {...$$restProps}>
  <legend>Control</legend>
</fieldset>

<style>
  fieldset.control {
    min-width: calc(100vw - 32px);
    width: calc(100vw - 32px);
    max-width: calc(100vw - 32px);
    height: 100%;

    scroll-snap-align: center;
  }
</style>
