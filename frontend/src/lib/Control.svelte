<script lang="ts">
  import ColorPicker from "svelteui/src/misc/ColorPicker.svelte";

  import ColorStorage from "./ControlColorStorage.svelte";

  export let r = 100;
  export let g = 100;
  export let b = 100;
  export let w = 100;
</script>

<fieldset class="control">
  <legend>Control</legend>

  <div class="spacer" style="height: 100%;" />

  <div class="scene">
    <!-- TODO: manage scenes here (Quick and Easy) -->
  </div>

  <ColorStorage />
  <ColorPicker bind:r bind:g bind:b bind:w />
</fieldset>

<style>
  fieldset.control {
    position: relative;

    min-width: calc(100vw - 32px);
    width: calc(100vw - 32px);
    max-width: calc(100vw - 32px);
    height: 100%;

    margin-right: 8px;

    display: flex;
    flex-direction: column;

    overflow: hidden;
    overflow-y: auto;
    scroll-behavior: smooth;
    scroll-snap-align: center;
  }
</style>
