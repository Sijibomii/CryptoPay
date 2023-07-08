"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var lib_1 = require("./lib");
var components_1 = require("./components");
require("./main.css");
// take redirect link. redirect to payment page after payment redirect back
var CryptoPay = /** @class */ (function () {
    function CryptoPay(options) {
        this.config = Object.assign({
            modal: null
        }, options);
        this.navigation = new lib_1.Navigation();
        this.apiClient = new lib_1.ApiClient({
            config: this.config
        });
        this.overlay = new components_1.Overlay({
            navigation: this.navigation,
            apiClient: this.apiClient,
            config: this.config,
            currencies: this.config.currencies,
            onSuccess: this.onSuccess.bind(this),
        });
        this.modal = new components_1.Modal({
            overlay: this.overlay,
            navigation: this.navigation,
            apiClient: this.apiClient,
            config: this.config,
            currencies: this.config.currencies,
            onSuccess: this.onSuccess.bind(this),
        });
        this.navigation.init({
            routes: {
                selector: new components_1.Select({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    config: this.config,
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                    onSuccess: this.onSuccess.bind(this),
                }),
                payment: new components_1.Payment({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    config: this.config,
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                    onSuccess: this.onSuccess.bind(this),
                }),
                confirmation: new components_1.Confirmation({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    config: this.config,
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                    onSuccess: this.onSuccess.bind(this),
                }),
                success: new components_1.Success({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    onSuccess: this.onSuccess.bind(this),
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                }),
                error: new components_1.PaymentError({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    onSuccess: this.onSuccess.bind(this),
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                })
            }
        });
    }
    CryptoPay.prototype.init = function () {
        this.mount();
    };
    CryptoPay.prototype.resize = function () { };
    CryptoPay.prototype.onSuccess = function (ticket) {
        this.config.onSuccess(ticket);
    };
    CryptoPay.prototype.closeModal = function () {
        this.modal.close();
        this.navigation.reset();
    };
    CryptoPay.prototype.activate = function () {
        // redirect page to new page
        // redirect to payment page
        var _this = this;
        this.overlay.set(document.body);
        this.navigation.target = this.overlay.element.querySelector(".checkout-overlay-content");
        this.navigation.navigate("selector", {
            dataset: ''
        });
        setTimeout(function () {
            _this.overlay.open();
        }, 0);
    };
    CryptoPay.prototype.mount = function () {
        this.config.button.addEventListener("click", this.activate.bind(this));
    };
    return CryptoPay;
}());
exports.default = CryptoPay;
