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

  import StatusLED from "./components/StatusLED.svelte";
  import api, { type Device } from "./ts/api";

  //const forDestroy: Events = {
  //  devices: [],
  //  device: [],
  //  offline: [],
  //};

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
    {
      const handler = async () => {
        console.debug(`[devices, event] "offline"`);
        for (const d of devices) {
          d.offline = true;
        }
        devices = devices;
      };
      api.addEventListener("offline", handler);
      //forDestroy.offline.push(offlineHandler);
    }

    // sse: "devices"
    {
      let handler = async (data: Device[]) => {
        console.debug(`[devices, event] "devices"`);

        // store color for device in localStorage
        data.forEach((d) => {
          if (d.rgbw.find((gp) => gp.duty > 0)) {
            window.localStorage.setItem(
              `color:${d.addr}`,
              JSON.stringify(d.rgbw.map((gp) => gp.duty))
            );
          }
        });

        devices = data;
      };
      api.addEventListener("devices", handler);
      //forDestroy.devices.push(devicesHandler);

      // store color handler
      handler = async (data: Device[]) => {
        console.debug(`[devices, event] "devices" (store color)`);

        // store color for device in localStorage
        data.forEach((d) => {
          if (d.rgbw.find((gp) => gp.duty > 0)) {
            window.localStorage.setItem(
              `color:${d.addr}`,
              JSON.stringify(d.rgbw.map((gp) => gp.duty))
            );
          }
        });
      };
      api.addEventListener("devices", handler);
      //forDestroy.devices.push(devicesHandler);
    }

    // sse: "device", "device" (store color)
    {
      let handler = async (data: Device) => {
        console.debug(`[devices, event] "device"`);
        const device = devices.find((d) => d.addr == data.addr);
        device.rgbw = data.rgbw;
        device.offline = data.offline;
        devices = devices;
      };
      api.addEventListener("device", handler);
      //forDestroy.device.push(deviceHandler);

      handler = async (data: Device) => {
        console.debug(`[devices, event] "device" (store color)`);
        if (data.rgbw.find((gp) => gp.duty > 0)) {
          window.localStorage.setItem(
            `color:${data.addr}`,
            JSON.stringify(data.rgbw.map((gp) => gp.duty))
          );
        }
      };
      api.addEventListener("device", handler);
      //forDestroy.device.push(deviceHandler);
    }
  });
</script>

<fieldset class="devices" {...$$restProps}>
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
</fieldset>

<style>
  fieldset.devices {
    min-width: calc(100vw - 32px);
    width: calc(100vw - 32px);
    max-width: calc(100vw - 32px);
    height: 100%;

    scroll-snap-align: center;
  }
</style>
