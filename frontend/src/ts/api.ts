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

export class Api {
    protocol: "http:" | "https:";
    host: string;
    version: "v1";
    paths: ApiPaths;

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
}

export default new Api();
