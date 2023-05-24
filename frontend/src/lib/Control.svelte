<script lang="ts">
  import ColorPicker from "svelteui/src/misc/ColorPicker.svelte";

  import ColorStorage from "./ControlColorStorage.svelte";

  export let r = 100;
  export let g = 100;
  export let b = 100;
  export let w = 100;

  // TODO: load stored items from the `localStorage`
  let items: Color[] = [
    // NOTE: placeholder items...
    [100, 100, 100, 100],
    [90, 0, 0, 0],
    [100, 100, 50, 0],
    [0, 100, 0, 0],
    [0, 50, 80, 0],
    [100, 100, 100, 100],
    [90, 0, 0, 0],
    [100, 100, 50, 0],
    [0, 100, 0, 0],
    [0, 50, 80, 0],
  ];

  function itemadd() {
    // TODO: add current rgbw to items
  }

  function itemdelete() {
    // TODO: select items from list and delete (needs some
    //       select prop in ColorStorage component)
  }
</script>

<fieldset class="control">
  <legend>Control</legend>

  <div class="spacer" style="height: 100%;" />

  <div class="scene">
    <!-- TODO: manage scenes here (Quick and Easy) -->
  </div>

  <div class="store">
    <ColorStorage bind:items on:itemadd={itemadd} on:itemdelete={itemdelete} />
  </div>
  <div class="picker">
    <ColorPicker bind:r bind:g bind:b bind:w />
  </div>
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
