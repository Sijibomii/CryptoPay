var Component = /** @class */ (function () {
    function Component(options) {
        this.config = options.config;
        this.dataset = options.dataset || {};
        this.navigation = options.navigation;
        // this.apiClient = options.apiClient;
        this.modal = options.modal;
        this.onSuccess = options.onSuccess;
        this.element = document.createElement("div");
        this.overlay = options.overlay;
        this.apiClient = options.apiClient;
    }
    Component.prototype.willMount = function () {
        if (this.navigatorParams && this.navigatorParams.dataset) {
            this.dataset = Object.assign(this.dataset, this.navigatorParams.dataset);
        }
    };
    Component.prototype.willUnmount = function () { };
    Component.prototype.onMounted = function () { };
    Component.prototype.onUnmounted = function () { };
    Component.prototype.render = function () { };
    return Component;
}());
export { Component };
