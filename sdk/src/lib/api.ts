import { Config } from "../types";


type ApiClientInput = {
    config: Config
}
export class ApiClient {

    public apiUrl: string;
    public apiKey: string;
    public price: number;
    public fiat: string;
    public identifier: string

    constructor(options: ApiClientInput) { 
        this.apiUrl = options.config.apiUrl;
        this.apiKey = options.config.apiKey;
        this.price = options.config.price;
        this.fiat = options.config.fiat;
        this.identifier = options.config.identifier;
    }

    createPayment(crypto: string) {
        return fetch(`${this.apiUrl}/p/payments`, {
          mode: "cors",
          headers: {
            Authorization: `Bearer ${this.apiKey}`,
            accept: "application/json",
            "content-Type": "application/json"
          },
          method: "POST", 
          body: JSON.stringify({
            fiat: this.fiat,
            crypto: crypto,
            price: this.price,
            identifier: this.identifier
          })
        }).then(response => {
          if (!response.ok) throw Error(response.statusText);
          return response.json();
        });
    }

    getPaymentStatus(id: string, sessionToken: string) {
        return fetch(`${this.apiUrl}/p/payments/${id}/status`, {
          mode: "cors",
          headers: {
            Authorization: `Bearer ${sessionToken}`,
            accept: "application/json"
          }
        }).then(response => {
          if (!response.ok) throw Error(response.statusText);
    
          return response.json();
        });
    }

    createVoucher(paymentId: string, sessionToken: string) {
        return fetch(`${this.apiUrl}/p/vouchers`, {
          mode: "cors",
          headers: {
            Authorization: `Bearer ${sessionToken}`,
            accept: "application/json",
            "content-Type": "application/json"
          },
          method: "POST",
          body: JSON.stringify({
            payment_id: paymentId
          })
        }).then(response => {
          if (!response.ok) throw Error(response.statusText);
    
          return response.json();
        });
      }
}