export interface GpPWM {
    nr: number;
    duty: number;
}

export interface Device {
    addr: string;
    rgbw: GpPWM[];
}

export class Api {
    protocol: ("http:" | "https:")

    constructor() {
        this.protocol = "http:"
    }

    async get(host: string | null = null): Promise<Device[]> {
        // TODO: do a "/api/v1/devices" call here parse and return the result

        return [];
    }
}

export default new Api();
