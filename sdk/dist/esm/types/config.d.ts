export type Config = Options & {
    modal?: HTMLElement | null;
};
export type Options = {
    price: number;
    fiat: string;
    identifier: string;
    apiUrl: string;
    apiKey: string;
    currencies: string[];
    onSuccess: (ticket: any) => null;
    button?: HTMLElement;
};
