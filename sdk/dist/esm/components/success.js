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
var Success = /** @class */ (function (_super) {
    __extends(Success, _super);
    function Success() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Success.prototype.onMounted = function () {
        this.modal.element.querySelector(".btn-close").classList.remove("disabled");
        // const dataset = this.dataset;
        // this.apiClient
        // .createVoucher(dataset.payment.id, dataset.sessionToken)
        // .then(data => {
        //     this.onSuccess(data.voucher);
        // });
    };
    Success.prototype.render = function () {
        var currency = this.dataset.currency;
        var template = "\n                <div class=\"modal-content simple ".concat(currency, "\">\n                    <div class=\"icon-check\"></div>\n                    <div class=\"title\">Sucessfully Completed</div>\n                    <p class=\"message\">\n                        We have confirmed your payment\n                    </p>\n                </div>\n            ");
        this.element.innerHTML = template.trim();
    };
    return Success;
}(Component));
export { Success };
