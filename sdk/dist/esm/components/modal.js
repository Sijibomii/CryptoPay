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
var Modal = /** @class */ (function (_super) {
    __extends(Modal, _super);
    function Modal() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.mountPoint = null;
        return _this;
    }
    Modal.prototype.set = function (mountPoint) {
        this.mountPoint = mountPoint;
        mountPoint.appendChild(this.element);
        this.render();
    };
    Modal.prototype.open = function () {
        var _this = this;
        setTimeout(function () {
            _this.element.querySelector("#modal-container").classList.add("active");
        }, 0);
        this.element
            .querySelector(".btn-close")
            .addEventListener("click", this.close.bind(this), { once: true });
    };
    Modal.prototype.close = function () {
        var _this = this;
        this.element.querySelector("#modal-container").classList.remove("active");
        this.overlay.close();
        setTimeout(function () {
            _this.mountPoint.removeChild(_this.element);
        }, 300);
        this.navigation.clear();
    };
    Modal.prototype.render = function () {
        var template = "\n                <div id=\"modal-container\">\n                    <div class=\"sk-fading-circle\">\n                        <div class=\"sk-circle1 sk-circle\"></div>\n                        <div class=\"sk-circle2 sk-circle\"></div>\n                        <div class=\"sk-circle3 sk-circle\"></div>\n                        <div class=\"sk-circle4 sk-circle\"></div>\n                        <div class=\"sk-circle5 sk-circle\"></div>\n                        <div class=\"sk-circle6 sk-circle\"></div>\n                        <div class=\"sk-circle7 sk-circle\"></div>\n                        <div class=\"sk-circle8 sk-circle\"></div>\n                        <div class=\"sk-circle9 sk-circle\"></div>\n                        <div class=\"sk-circle10 sk-circle\"></div>\n                        <div class=\"sk-circle11 sk-circle\"></div>\n                        <div class=\"sk-circle12 sk-circle\"></div>\n                    </div>\n                    \n                    <div class=\"modal\">\n                        <div class=\"content-root\"></div>\n                        <button class=\"btn-close\" />\n                    </div>\n                </div>";
        this.element.innerHTML = template.trim();
    };
    return Modal;
}(Component));
export { Modal };
