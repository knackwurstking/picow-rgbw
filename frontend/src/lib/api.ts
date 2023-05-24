export type GpPin = number;
export type Duty = number;

export interface Gp {
  nr: GpPin;
  duty: Duty;
}

export interface Device {
  addr: string;
  offline: boolean;
  rgbw: Gp[];
}

export interface ApiPathsV1 {
  devices: () => string;
  postDevices: () => string;
  device: (id: string) => string;
}

export interface ApiPaths {
  v1: ApiPathsV1;
}

export interface Events {
  devices: ((data: Device[]) => Promise<void> | void)[];
  device: ((data: Device) => Promise<void> | void)[];
  offline: (() => Promise<void> | void)[];
}

export interface Sources {
  devices: EventSource | null;
  device: EventSource | null;
}

export interface PutDeviceRequestData {
  addr: string;
  rgbw: Duty[];
}

export class Api {
  protocol: "http:" | "https:";
  host: string;
  version: "v1";
  paths: ApiPaths;
  events: Events;
  sources: Sources;

  constructor() {
    this.protocol = "http:";
    this.host = "";
    this.version = "v1";

    this.paths = {
      v1: {
        devices: () => "/api/v1/devices",
        postDevices: () => "/api/v1/devices",
        device: (id) => `/api/v1/devices/${id}`,
      },
    };

    this.events = {
      devices: [],
      device: [],
      offline: [],
    };

    this.sources = {
      devices: null,
      device: null,
    }

    this.sse();
  }

  url(key: string, ...param: any): string {
    const origin = !!this.host ? `${this.protocol}//${this.host}` : "";
    return `${origin}${this.paths[this.version][key](...param)}`;
  }

  async devices(): Promise<Device[]> {
    const url = this.url("devices");
    const resp = await fetch(url);

    if (!resp.ok) {
      throw `response error: ${resp.statusText} (${url})`;
    }

    return (await resp.json()) as Device[];
  }

  async postDevices(...data: PutDeviceRequestData[]): Promise<void> {
    const url = this.url("postDevices");
    const resp = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    if (!resp.ok) {
      throw `response error: ${resp.statusText} (${url})`;
    }
  }

  addEventListener(type: "offline", listener: () => Promise<void> | void): void
  addEventListener(type: "devices", listener: (data: Device[]) => Promise<void> | void): void
  addEventListener(type: "device", listener: (data: Device) => Promise<void> | void): void
  addEventListener(
    type: ("offline" | "devices" | "device"),
    listener: any
  ): void {
    if (!(type in this.events)) {
      return;
    }

    this.events[type].push(listener);
  }

  removeEventListener(type: "offline", listener: () => Promise<void> | void): void
  removeEventListener(type: "devices", listener: (data: Device[]) => Promise<void> | void): void
  removeEventListener(type: "device", listener: (data: Device) => Promise<void> | void): void
  removeEventListener(
    type: ("offline" | "devices" | "device"),
    listener: any
  ): void {
    if (!(type in this.events)) {
      return;
    }

    let i = 0;
    const listeners = this.events[type];
    for (const l of listeners) {
      if (l == listener) {
        // @ts-expect-error
        this.events[type] = [
          ...listeners.slice(0, i),
          ...listeners.slice(i + 1),
        ];
      }
      i++;
    }
  }

  sse() {
    for (const p of ["devices", "device"]) {
      console.log(`[api] Connect to sse EventSource: "${p}".`);

      const path = "/api/v1/events/" + p;
      if (this.sources[p]) {
        this.sources[p].close();
      }

      const source = (this.sources[p] = new EventSource(path));

      if (p === "devices") { // Only "devices" will handle the reconnect
        source.onerror = () => {
          console.error(`Oops, sse EventSource for "${p}" failed. Try re-connect...`);

          setTimeout(() => {
            console.log("[api] ...reconnecting to sse event source!");
            this.sse();
          }, 3000);

          for (const l of this.events.offline) {
            l();
          }
        };
      }

      source.addEventListener("update", (ev) => {
        for (const l of this.events[p]) {
          l(JSON.parse(ev.data));
        }
      });
    }
  }
}

export default new Api();
