var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
import { Component } from "./components";
import errorImg from "../images/btc.png";
var PaymentError = /** @class */ (function (_super) {
    __extends(PaymentError, _super);
    function PaymentError() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    PaymentError.prototype.onMounted = function () {
        var closeButton = this.modal.element.querySelector(".btn-close");
        if (closeButton.classList.contains("disabled")) {
            closeButton.classList.remove("disabled");
        }
    };
    PaymentError.prototype.render = function () {
        var message = this.navigatorParams.dataset.message;
        var template = "\n            <div class=\"modal-content simple\">\n                <div class=\"icon-error\">\n                  <img src=\"".concat(errorImg, "\" width=\"9\" height=\"34\" />\n                </div>\n                <div class=\"title\">Error</div>\n                <p class=\"message\">\n                    ").concat(message, "\n                </p>\n            </div>\n        ");
        this.element.innerHTML = template.trim();
    };
    return PaymentError;
}(Component));
export { PaymentError };
