
export type Config = Options & {
    modal?: HTMLElement | null
    apiUrl: string
    apiKey: string
};

export type Options = {
    price: number
    fiat: string
    identifier: string
}