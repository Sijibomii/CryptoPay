import { Current, Target } from "../types";

export default class Navigation {
    private target: Target;
    private current: Current;

    constructor() {
        this.target = document.createElement("div");
        this.current = '';        
    }

    
}