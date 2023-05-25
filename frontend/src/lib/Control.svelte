<script lang="ts">
  import ColorPicker from "svelteui/misc/ColorPicker.svelte";

  import ColorStorage from "./ControlColorStorage.svelte";

  export let r = 100;
  export let g = 100;
  export let b = 100;
  export let w = 100;

  let items: Color[] = [];

  function itemadd() {
    for (let i = 0; i < items.length; i++) {
      const item = items[i];

      if (item[0] === r && item[1] === g && item[2] === b && item[3] === w) {
        items = [items[i], ...items.slice(0, i), ...items.slice(i + 1)];
        return;
      }
    }

    items.unshift([r, g, b, w]);
    items = items;
  }

  function itemdelete() {
    for (let i = 0; i < items.length; i++) {
      const item = items[i];

      if (item[0] === r && item[1] === g && item[2] === b && item[3] === w) {
        items = [...items.slice(0, i), ...items.slice(i + 1)];
        return;
      }
    }
  }

  function itemselected(ev: CustomEvent<Color>) {
    const color = ev.detail;
    r = color[0];
    g = color[1];
    b = color[2];
    w = color[3];
  }
</script>

<fieldset class="control">
  <legend>Control</legend>

  <div class="spacer" style="height: 100%;" />

  <div class="scene">
    <!-- TODO: manage scenes here (Quick and Easy) -->
  </div>

  <div class="store">
    <ColorStorage
      bind:items
      on:itemadd={itemadd}
      on:itemdelete={itemdelete}
      on:itemselected={itemselected}
    />
  </div>
  <div class="picker">
    <ColorPicker style="user-select: none;" bind:r bind:g bind:b bind:w />
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
