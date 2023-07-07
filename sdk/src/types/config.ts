
export type Config = Options & {
    modal?: HTMLElement | null
    button?: HTMLElement
};

export type Options = {
    price: number
    fiat: string
    identifier: string
    apiUrl: string
    apiKey: string
    currencies: string[]
    onSuccess: (ticket: any) => null
}