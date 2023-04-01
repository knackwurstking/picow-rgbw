export interface GpPWM {
    nr: number;
    duty: number;
}

export interface Device {
    addr: string;
    rgbw: GpPWM[];
}

export class Devices {
    async get(): Promise<Device[]> {
        // TODO: do a "/api/v1/devices" call here parse and return the result

        return [];
    }
}

export default new Devices();
