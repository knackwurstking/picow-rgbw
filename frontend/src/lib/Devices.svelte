<script lang="ts">
  import { onMount } from "svelte";

  import Checkbox from "@smui/checkbox";

  import List, {
    Item,
    Text,
    PrimaryText,
    SecondaryText,
    Separator,
    Meta,
  } from "@smui/list";

  import StatusLED from "./StatusLED.svelte";
  import api, { type Device, type Events } from "./api";

  const forDestroy: Events = {
    devices: [],
    device: [],
    offline: [],
  };

  export let r: number = 100;
  export let g: number = 100;
  export let b: number = 100;
  export let w: number = 100;

  export let selected: Device[] = [];

  let devices: Device[] = [];
  $: {
    console.debug("[devices]", { devices });
    const newSelected = [];
    for (const s of selected) {
      if (!!devices.find((d) => d.addr === s.addr)) {
        newSelected.push(s);
      }
    }
    selected = newSelected;
  }

  onMount(() => {
    // sse: "offline"
    let offlineHandler = async () => {
      console.debug(`[devices] event: "offline"`);
      for (const d of devices) {
        d.offline = true;
      }
      devices = devices;
    };
    api.addEventListener("offline", offlineHandler);
    forDestroy.offline.push(offlineHandler);

    // sse: "devices"
    let devicesHandler = async (data: Device[]) => {
      console.debug(`[devices] event: "devices"`);
      devices = data;
    };
    api.addEventListener("devices", devicesHandler);
    forDestroy.devices.push(devicesHandler);

    // sse: "device"
    const deviceHandler = async (data: Device) => {
      console.debug(`[devices] event: "device"`);
      const device = devices.find((d) => d.addr == data.addr);
      device.rgbw = data.rgbw;
      device.offline = data.offline;
      devices = devices;
    };
    api.addEventListener("device", deviceHandler);
    forDestroy.device.push(deviceHandler);
  });
</script>

<fieldset {...$$restProps}>
  <legend>Devices</legend>
  <List checkList>
    {#each devices as device}
      <Item style="height: 65px;">
        <Text>
          <PrimaryText>{device.addr}</PrimaryText>
          <SecondaryText
            >[{device.rgbw.map((gp) => gp.duty).join(",")}]</SecondaryText
          >
        </Text>
        <Meta>
          <Checkbox
            style="margin-right: 8px;"
            bind:group={selected}
            value={device}
          />
          <StatusLED
            style="
              position: absolute;
              top: 4px;
              right: 4px;
            "
            active={!device.offline}
          />
        </Meta>
      </Item>
      <Separator />
    {/each}
  </List>
  <!-- TODO: Add some kind of a action bar to bottom -->
</fieldset>

<style>
  fieldset {
    min-width: calc(100vw - 32px);
    width: calc(100vw - 32px);
    max-width: calc(100vw - 32px);
    height: calc(100% - 32px);
    margin: 16px 0;

    scroll-snap-align: center;
  }
</style>
