

export type BrewSettings = BrewSetting[]

export interface BrewSetting {
    brewtype: BrewTypes,
    coffee: string,
    grinder: string,
    grindlevel: number,
    watertemperature: number,
}

export enum BrewTypes {
    Siebträger = "Siebträger",
    Mokkakocher = "Mokkakocher",
    Filter = "Filter",
    Frenchpress = "French press"
}