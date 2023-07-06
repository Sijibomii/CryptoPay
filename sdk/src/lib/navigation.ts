import { Component, navigatorParam } from "../components";
import { Current, Target } from "../types";

type navigationInitInput = {
    routes: Record<string, Component>
}

export class Navigation {
    public target?: Target;
    private current: Current;
    private routes: Record<string, Component>;

    constructor() {
        this.target = document.createElement("div");
        this.current = ''; 
        this.routes = {};       
    }

    reset() {
        if (this.current){
            this.routes[this.current]!.onUnmounted();
            this.current = '';
        } 
    }

    init(options: navigationInitInput) {
        this.routes = options.routes;
    }
    
    clear() {
        if (this.current) {
            this.routes[this.current]!.onUnmounted();
        }
        this.target;
        this.current = '';
    }

    navigate(to: string, params: navigatorParam) {

        if (params) {
          this.routes[to]!.navigatorParams = params;
        }
    
        if (this.target!.hasChildNodes() && this.routes[to]) {
            const node: Node =  this.routes[to]!.element
            this.target!.replaceChild(node, this.routes[this.current]!.element);
        } else {
            this.target!.appendChild(this.routes[to]!.element);
        }
    
        if (this.current) {
          this.routes[this.current]!.onUnmounted();
        }
    
        this.routes[to]!.willMount();
        this.routes[to]!.render();
        this.routes[to]!.onMounted();
        this.current = to;
    }
}