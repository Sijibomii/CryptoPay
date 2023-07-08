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
var Overlay = /** @class */ (function (_super) {
    __extends(Overlay, _super);
    function Overlay() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.mountPoint = null;
        return _this;
    }
    Overlay.prototype.set = function (mountPoint) {
        this.mountPoint = mountPoint;
        mountPoint.appendChild(this.element);
        this.render();
    };
    Overlay.prototype.open = function () {
        this.element.querySelector("#checkout-overlay").classList.add("active");
    };
    Overlay.prototype.close = function () {
        var _this = this;
        this.element.querySelector("#checkout-overlay").classList.remove("active");
        setTimeout(function () {
            _this.mountPoint.removeChild(_this.element);
        }, 300);
    };
    Overlay.prototype.render = function () {
        var template = "<div id=\"checkout-overlay\"><div class=\"checkout-overlay-content\"></div></div>";
        this.element.innerHTML = template.trim();
    };
    return Overlay;
}(Component));
export { Overlay };
