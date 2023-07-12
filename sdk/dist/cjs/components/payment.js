"use strict";
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
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Payment = void 0;
var components_1 = require("./components");
// import QRCode from "qrcode";
var Payment = /** @class */ (function (_super) {
    __extends(Payment, _super);
    function Payment(options) {
        var _this = _super.call(this, options) || this;
        // could not get a better initializer
        _this.timeLeft = "";
        _this.copyButton = _this.element.querySelector(".copy-btn");
        _this.timer = setTimeout(function () { }, 100);
        _this.polling = setTimeout(function () { }, 100);
        return _this;
    }
    Payment.prototype.onMounted = function () {
        return __awaiter(this, void 0, void 0, function () {
            var error_1;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        this.timeLeft = "15:00";
                        _a.label = 1;
                    case 1:
                        _a.trys.push([1, 3, , 4]);
                        this.copyButton = this.element.querySelector(".copy-address");
                        this.copyButton.addEventListener("click", this.copyAddress.bind(this));
                        // const qrcodeElement = this.element.querySelector(".qrcode");
                        // const qrcodeOptions = {
                        //     text: `${this.dataset.payment.address}`,
                        //     width: 100,
                        //     height: 100,
                        //     colorDark: "#4a4a4a",
                        //     colorLight: "#ffffff"
                        // };
                        // const ss =  JSON.stringify(qrcodeOptions)
                        // QRCode.toCanvas(qrcodeElement, JSON.stringify(qrcodeOptions), (error : any) => {
                        //     if (error) {
                        //       console.error("Failed to generate QR code:", error);
                        //     } else {
                        //       console.log("QR code generated successfully");
                        //     }
                        //   });
                        this.startCountdown();
                        return [4 /*yield*/, this.pollPaymentStatus()];
                    case 2:
                        _a.sent();
                        return [3 /*break*/, 4];
                    case 3:
                        error_1 = _a.sent();
                        this.navigation.navigate.call(this.navigation, "error", {
                            dataset: Object.assign(this.dataset, {
                                message: "Error detected during payment session."
                            })
                        });
                        return [2 /*return*/];
                    case 4: return [2 /*return*/];
                }
            });
        });
    };
    Payment.prototype.onUnmounted = function () {
        clearTimeout(this.polling);
        clearTimeout(this.timer);
        if (this.copyButton)
            this.copyButton.removeEventListener("click", this.copyAddress);
    };
    Payment.prototype.copyAddress = function () {
        var el = document.createElement("textarea");
        el.value = this.dataset.payment.address;
        el.setAttribute("readonly", "");
        el.style.position = "absolute";
        el.style.left = "-9999px";
        document.body.appendChild(el);
        el.select();
        this.copyToClipboard(el.value);
        document.body.removeChild(el);
    };
    Payment.prototype.copyToClipboard = function (text) {
        navigator.clipboard.writeText(text)
            .then(function () {
            console.log("Text copied to clipboard");
        })
            .catch(function (error) {
            console.error("Failed to copy text to clipboard:", error);
        });
    };
    Payment.prototype.startCountdown = function () {
        var presentTime = this.timeLeft;
        var timeArray = presentTime.split(/[:]+/);
        var m = timeArray[0];
        var s = parseInt(timeArray[1]) - 1;
        if (s < 10 && s >= 0) {
            s = "0" + s;
        }
        if (s < 0) {
            s = "59";
        }
        if (s == 59) {
            m = parseInt(m) - 1;
        }
        this.timeLeft = "".concat(m, ":").concat(s);
        if (this.timeLeft == "0:00") {
            this.navigation.navigate.call(this.navigation, "error", {
                dataset: Object.assign(this.dataset, {
                    message: "Payment session has expired."
                })
            });
            return;
        }
        this.element.querySelector(".timer").innerHTML = this.timeLeft;
        this.timer = setTimeout(this.startCountdown.bind(this), 1000);
    };
    Payment.prototype.pollPaymentStatus = function () {
        return __awaiter(this, void 0, void 0, function () {
            var resp;
            var _this = this;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.apiClient.getPaymentStatus(this.dataset.payment.id, this.dataset.sessionToken)];
                    case 1:
                        resp = _a.sent();
                        if (!(resp.status == "pending")) return [3 /*break*/, 4];
                        return [4 /*yield*/, new Promise(function (resolve) { return (_this.polling = setTimeout(resolve, 4000)); })];
                    case 2:
                        _a.sent();
                        return [4 /*yield*/, this.pollPaymentStatus()];
                    case 3: return [2 /*return*/, _a.sent()];
                    case 4:
                        this.navigation.navigate.call(this.navigation, "confirmation", {
                            dataset: this.dataset
                        });
                        return [2 /*return*/];
                }
            });
        });
    };
    Payment.prototype.render = function () {
        var expiration = new Date();
        expiration.setMinutes(expiration.getMinutes() + 15);
        var expiresAt = "".concat(expiration.getHours(), ":").concat(("0" + expiration.getMinutes()).substring(-2));
        var _a = this.dataset, currency = _a.currency, payment = _a.payment, store = _a.store;
        var network;
        switch (payment.crypto) {
            case "btc":
                network = payment.btc_network;
                break;
        }
        var template = "\n                <div class=\"modal-content payment ".concat(currency, "\">\n                    <div class=\"header\">\n                        <div class=\"title\">").concat(store.name, "</div>\n                        <p class=\"message\">\n                            Please send crypto to the address below.\n                        </p>\n                        <div class=\"amount\">").concat(payment.charge, " ").concat(currency.toUpperCase(), "</div>\n                        <div class=\"converted-amount\">(").concat(payment.price, " ").concat(payment.fiat.toUpperCase(), ")</div>\n                        <div class=\"payment-qrcode-wrapper\">\n                            <div class=\"qrcode\">\n\n                            </div>\n                        </div>\n                    </div>\n                    <div class=\"body\">\n                        <div class=\"address\">").concat(payment.address, "</div>\n                        <button class=\"copy-address\">Copy the Address</button>\n                        <div class=\"separator\"><span class=\"or\">or<span></div>\n                        ").concat(currency === "eth"
            ? "<img class=\"pay-with-metamask\" src=\"\" width=\"180\" height=\"48\" alt=\"Pay with metamask\" /><div class=\"metamask-error\"></div>"
            : "<a href=\"bitcoin:".concat(payment.address, "?amount=").concat(payment.charge, "\" class=\"open-in-wallet\">Open in Wallet</a>"), "\n                    </div>\n                    <div class=\"network\">Network: ").concat(network.toUpperCase(), "</div>\n                    <div class=\"footer\">\n                        <div class=\"timer footer-left\">15:00</div>\n                        <div class=\"footer-right\">\n                            <div class=\"rate\"><span class=\"bold\">1").concat(currency.toUpperCase(), " = ").concat((payment.price / payment.charge).toFixed(2), " ").concat(payment.fiat.toUpperCase(), "</span></div>\n                            <div class=\"footer-message\">This payment modal is valid until: <span class=\"bold\">").concat(expiresAt, "</span></div>\n                        </div>\n                    </div>\n                </div>");
        this.element.innerHTML = template.trim();
    };
    return Payment;
}(components_1.Component));
exports.Payment = Payment;
