import Navigation from "../lib/navigation";
import { Config } from "../types";

type navigatorParam = {
    dataset: dataset
}
type dataset = any
type ComponentInputs = {
    config?: Config
    navigation?: Navigation
    modal?: HTMLElement | null
    onSuccess?: (ticket: any) => null
    element?: HTMLElement
    overlay?: HTMLElement | null
    dataset?: any
}

export default class Components {

    private config: Config | undefined;
    private navigation: Navigation  | undefined;
    private modal: HTMLElement | null  | undefined
    private onSuccess?: undefined | ((ticket: any) => null);
    private element?: HTMLElement | undefined
    private overlay?: HTMLElement | null | undefined
    private navigatorParams?: navigatorParam
    private dataset?: dataset
    
    constructor(options: ComponentInputs = {}) {
        this.config = options.config;
        this.dataset = options.dataset || {}; 
        this.navigation = options.navigation;
        // this.apiClient = options.apiClient;
        this.modal = options.modal;
        this.onSuccess = options.onSuccess;
        this.element = document.createElement("div");
        this.overlay = options.overlay;
    }

    willMount() {
        if (this.navigatorParams && this.navigatorParams.dataset) {
            this.dataset = Object.assign(this.dataset, this.navigatorParams.dataset);
        }
    }  

    willUnmount() {}

    onMounted() {}

    onUnmounted() {}
}