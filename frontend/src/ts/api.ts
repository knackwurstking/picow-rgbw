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

export interface ReqPutDevice {
    addr: string;
    rgbw: Duty[];
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
}

export class Api {
    protocol: "http:" | "https:";
    host: string;
    version: "v1";
    paths: ApiPaths;
    events: Events;

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
        };

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

    async postDevices(...data: ReqPutDevice[]): Promise<void> {
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

    addEventListener(
        type: string,
        listener: (data: any) => Promise<void> | void
    ) {
        if (!(type in this.events)) {
            return;
        }

        this.events[type].push(listener);
    }

    removeEventListener(
        type: string,
        listener: (data: any) => Promise<void> | void
    ) {
        if (!(type in this.events)) {
            return;
        }

        let i = 0;
        const listeners = this.events[type];
        for (const l of listeners) {
            if (l == listener) {
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
            const source = new EventSource("/api/v1/events/" + p);
            source.onerror = (ev) => {
                console.error("sse: ", ev);
                // TODO: try to reconnect every few seconds
                //       and update devices
            };
            source.onopen = () => {
                console.log("sse: onopen");
            };
            source.addEventListener("update", (ev) => {
                console.log(p, "update");
                for (const l of this.events[p]) {
                    l(JSON.parse(ev.data));
                }
            });
        }
    }
}

export default new Api();
