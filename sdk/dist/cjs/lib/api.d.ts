import { Config } from "../types";
type ApiClientInput = {
    config: Config;
};
export declare class ApiClient {
    apiUrl: string;
    apiKey: string;
    price: number;
    fiat: string;
    identifier: string;
    constructor(options: ApiClientInput);
    createPayment(crypto: string): Promise<any>;
    getPaymentStatus(id: string, sessionToken: string): Promise<any>;
    createVoucher(paymentId: string, sessionToken: string): Promise<any>;
}
export {};
