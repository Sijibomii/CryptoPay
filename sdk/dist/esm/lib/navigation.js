var Navigation = /** @class */ (function () {
    function Navigation() {
        this.target = document.createElement("div");
        this.current = '';
        this.routes = {};
    }
    Navigation.prototype.reset = function () {
        if (this.current) {
            this.routes[this.current].onUnmounted();
            this.current = '';
        }
    };
    Navigation.prototype.init = function (options) {
        this.routes = options.routes;
    };
    Navigation.prototype.clear = function () {
        if (this.current) {
            this.routes[this.current].onUnmounted();
        }
        this.target;
        this.current = '';
    };
    Navigation.prototype.navigate = function (to, params) {
        if (params) {
            this.routes[to].navigatorParams = params;
        }
        if (this.target.hasChildNodes() && this.routes[to]) {
            var node = this.routes[to].element;
            this.target.replaceChild(node, this.routes[this.current].element);
        }
        else {
            this.target.appendChild(this.routes[to].element);
        }
        if (this.current) {
            this.routes[this.current].onUnmounted();
        }
        this.routes[to].willMount();
        this.routes[to].render();
        this.routes[to].onMounted();
        this.current = to;
    };
    return Navigation;
}());
export { Navigation };
