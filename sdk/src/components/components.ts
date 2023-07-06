import {Navigation} from "../lib/navigation";
import { Config } from "../types";
import { Overlay } from "./overlay";
export type navigatorParam = {
    dataset: dataset
}
type dataset = any
type ComponentInputs = {
    config?: Config
    navigation?: Navigation
    modal?: HTMLElement | null
    onSuccess?: (ticket: any) => null
    element : HTMLElement
    overlay?: Overlay
    dataset?: any
}

export class Component {
    public config: Config | undefined;
    public navigation: Navigation  | undefined;
    public modal: HTMLElement | null  | undefined
    public onSuccess?: undefined | ((ticket: any) => null);
    public element: HTMLElement
    public overlay?: Overlay | undefined
    public navigatorParams?: navigatorParam
    public dataset?: dataset
    
    constructor(options: ComponentInputs) {
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

    render(){}
}