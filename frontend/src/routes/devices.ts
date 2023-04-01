export interface GpPWM {
    nr: number;
    duty: number;
}

export interface Device {
    addr: string;
    rgbw: GpPWM[];
}

export class Devices {
    get(): Device[] {
        return [];
    }
}

export default new Devices();
