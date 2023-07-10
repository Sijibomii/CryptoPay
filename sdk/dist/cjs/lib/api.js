"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ApiClient = void 0;
var ApiClient = /** @class */ (function () {
    function ApiClient(options) {
        this.apiUrl = options.config.apiUrl;
        this.apiKey = options.config.apiKey;
        this.price = options.config.price;
        this.fiat = options.config.fiat;
        this.identifier = options.config.identifier;
    }
    ApiClient.prototype.createPayment = function (crypto) {
        return fetch("".concat(this.apiUrl, "/payments"), {
            mode: "cors",
            headers: {
                Authorization: "Bearer ".concat(this.apiKey),
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
        }).then(function (response) {
            if (!response.ok)
                throw Error(response.statusText);
            return response.json();
        });
    };
    ApiClient.prototype.getPaymentStatus = function (id, sessionToken) {
        return fetch("".concat(this.apiUrl, "/payments/").concat(id, "/status"), {
            mode: "cors",
            headers: {
                Authorization: "Bearer ".concat(sessionToken),
                accept: "application/json"
            }
        }).then(function (response) {
            if (!response.ok)
                throw Error(response.statusText);
            return response.json();
        });
    };
    ApiClient.prototype.createVoucher = function (paymentId, sessionToken) {
        return fetch("".concat(this.apiUrl, "/vouchers"), {
            mode: "cors",
            headers: {
                Authorization: "Bearer ".concat(sessionToken),
                accept: "application/json",
                "content-Type": "application/json"
            },
            method: "POST",
            body: JSON.stringify({
                payment_id: paymentId
            })
        }).then(function (response) {
            if (!response.ok)
                throw Error(response.statusText);
            return response.json();
        });
    };
    return ApiClient;
}());
exports.ApiClient = ApiClient;
