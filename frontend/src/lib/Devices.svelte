<script lang="ts">
  import List, { Item, Meta, Separator } from "svelteui/list";
  import Checkbox from "svelteui/checkbox";
  import StatusLED from "svelteui/misc/StatusLED.svelte";

  import type { Device } from "./api";

  export let devices: Device[];
  export let selected: string[];
</script>

<!-- "Devices" view to the left -->
<fieldset class="devices">
  <legend>Devices</legend>

  <List checkable multiple checklist bind:group={selected}>
    {#each devices as device}
      <Item
        primaryText={device.addr}
        secondaryText={`[${device.rgbw.map((gp) => gp.duty).join(",")}]`}
        value={device.addr}
        checked={!!selected.find((a) => a === device.addr)}
      >
        <Meta slot="right">
          <Checkbox
            disableUserActions
            style="margin-right: 28px; float: right;"
            group={selected}
            value={device.addr}
          />

          <StatusLED
            style="
              position: absolute;
              top: 4px;
              right: 4px;
              font-size: 0.8rem;
            "
            active={!device.offline}
          />
        </Meta>
      </Item>
      <Separator />
    {/each}
  </List>
</fieldset>

<style>
  fieldset.devices {
    position: relative;

    min-width: calc(100vw - 32px);
    width: calc(100vw - 32px);
    max-width: calc(100vw - 32px);
    height: 100%;

    margin-left: 8px;

    overflow: hidden;

    scroll-snap-align: center;
  }
</style>
