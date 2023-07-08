import { Navigation, ApiClient } from "./lib";
import { Overlay, Modal, Select as Selector, Payment, Confirmation, Success, PaymentError } from "./components";
import "./main.css";
// take redirect link. redirect to payment page after payment redirect back
var CryptoPay = /** @class */ (function () {
    function CryptoPay(options) {
        this.config = Object.assign({
            modal: null
        }, options);
        this.navigation = new Navigation();
        this.apiClient = new ApiClient({
            config: this.config
        });
        this.overlay = new Overlay({
            navigation: this.navigation,
            apiClient: this.apiClient,
            config: this.config,
            currencies: this.config.currencies,
            onSuccess: this.onSuccess.bind(this),
        });
        this.modal = new Modal({
            overlay: this.overlay,
            navigation: this.navigation,
            apiClient: this.apiClient,
            config: this.config,
            currencies: this.config.currencies,
            onSuccess: this.onSuccess.bind(this),
        });
        this.navigation.init({
            routes: {
                selector: new Selector({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    config: this.config,
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                    onSuccess: this.onSuccess.bind(this),
                }),
                payment: new Payment({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    config: this.config,
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                    onSuccess: this.onSuccess.bind(this),
                }),
                confirmation: new Confirmation({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    config: this.config,
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                    onSuccess: this.onSuccess.bind(this),
                }),
                success: new Success({
                    navigation: this.navigation,
                    apiClient: this.apiClient,
                    onSuccess: this.onSuccess.bind(this),
                    modal: this.modal,
                    overlay: this.overlay,
                    currencies: this.config.currencies,
                }),
                error: new PaymentError({
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
export default CryptoPay;
