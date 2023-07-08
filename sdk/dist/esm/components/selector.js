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
import { Component } from "./components";
import bitcoinImg from "../images/btc.png";
var Select = /** @class */ (function (_super) {
    __extends(Select, _super);
    function Select(options) {
        var _this = _super.call(this, options) || this;
        // could not get a better initializer
        _this.buttons = _this.element.querySelectorAll("");
        return _this;
    }
    Select.prototype.onMounted = function () {
        this.activateCurrencySelectButton();
        this.element.querySelector(".btn-close").addEventListener("click", this.overlay.close.bind(this.overlay));
    };
    Select.prototype.activateCurrencySelectButton = function () {
        var _this = this;
        this.buttons = this.element.querySelectorAll(".btn-currency");
        this.buttons.forEach(function (button) {
            var el = button;
            button.addEventListener("click", function (event) { return _this.navigateToPayment(el.dataset['type'], event); });
        });
    };
    Select.prototype.navigateToPayment = function (currency, e) {
        var _a;
        return __awaiter(this, void 0, void 0, function () {
            var data;
            return __generator(this, function (_b) {
                switch (_b.label) {
                    case 0:
                        e.stopPropagation();
                        this.navigation.target.removeChild(this.element);
                        this.modal.set(this.overlay.element.querySelector("checkout-overlay-content"));
                        this.navigation.target = (_a = this.modal) === null || _a === void 0 ? void 0 : _a.element.querySelector(".content-root");
                        return [4 /*yield*/, this.apiClient.createPayment(currency)];
                    case 1:
                        data = _b.sent();
                        this.navigation.navigate.call(this.navigation, "payment", {
                            dataset: Object.assign(this.dataset, {
                                currency: currency,
                                payment: data.payment,
                                store: data.store,
                                sessionToken: data.token
                            })
                        });
                        this.modal.open();
                        return [2 /*return*/];
                }
            });
        });
    };
    Select.prototype.renderButton = function (currency) {
        switch (currency) {
            case "btc":
                return "<button class=\"btn-currency\" data-type=\"btc\">\n                        <img src=\"".concat(bitcoinImg, "\" width=\"98\" height=\"109\" />\n                        <span class=\"label\">Bitcoin</span>\n                        </button>");
            default:
                return "";
        }
    };
    Select.prototype.render = function () {
        var _this = this;
        var currencies = this.config.currencies;
        var template = "\n        <div class=\"currency-selector\">\n            <button class=\"btn-close\"></button>\n            <p class=\"message\">\n            Select the cryptocurrency you want to use.\n            </p>\n            <div class=\"body\">\n            ".concat(currencies.map(function (c) { return _this.renderButton(c); }).join(""), "\n            </div>\n        </div>");
        this.element.innerHTML = template.trim();
    };
    return Select;
}(Component));
export { Select };
