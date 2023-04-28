<script lang="ts">
  import { onMount } from "svelte";

  import CheckLabel from "./CheckLabel.svelte";
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
  <div class="content">
    {#each devices as device}
      <CheckLabel
        label={device.addr}
        currentColor={device.rgbw.map((gp) => gp.duty)}
        offline={device.offline}
        on:change={(ev) => {
          if (!ev.detail.checked) {
            // remove device from selected
            selected = selected.filter((d) => d != device);
          } else {
            selected = [...selected, device];
          }

          if (!selected.length && Math.max(r, g, b, w) === 0) {
            [r, g, b, w] = [100, 100, 100, 100];
          } else {
            if (
              selected.find((device) => !!device.rgbw.find((gp) => gp.duty > 0))
            ) {
              [r, g, b, w] = selected[selected.length - 1].rgbw.map(
                (d) => d.duty
              );
            }
          }
        }}
      />
    {/each}
  </div>
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

  fieldset > .content {
    width: 100%;
    height: 100%;
  }
</style>
