/// <reference types="svelte" />
/// <reference types="vite/client" />

declare type GpPin = number;
declare type Duty = number;

declare interface Gp {
  nr: GpPin;
  duty: Duty;
}

declare interface Device {
  addr: string;
  offline: boolean;
  rgbw: Gp[];
}

declare interface ReqPutDevice {
  addr: string;
  rgbw: Duty[];
}

declare interface ApiPathsV1 {
  devices: () => string;
  postDevices: () => string;
  device: (id: string) => string;
}

declare interface ApiPaths {
  v1: ApiPathsV1;
}

declare interface Events {
  devices: ((data: Device[]) => Promise<void> | void)[];
  device: ((data: Device) => Promise<void> | void)[];
  offline: (() => Promise<void> | void)[];
}

declare interface Sources {
  devices: EventSource | null;
  device: EventSource | null;
}
