import {Navigation} from "../lib/navigation";
import { Config } from "../types";
import { Overlay } from "./overlay";
import { Modal } from "./modal";
import { ApiClient } from "../lib";

export type navigatorParam = {
    dataset: dataset
}
type dataset = any
export type ComponentInputs = {
    config?: Config
    navigation?: Navigation
    modal?: Modal
    onSuccess: (ticket: any) => void
    overlay?: Overlay
    dataset?: any
    currencies: string[]
    apiClient: ApiClient
}

export class Component {
    public config: Config | undefined;
    public navigation: Navigation  | undefined;
    public modal: Modal  | undefined
    public onSuccess: ((ticket: any) => void);
    public element: HTMLElement
    public overlay?: Overlay | undefined
    public navigatorParams?: navigatorParam
    public dataset?: dataset
    public apiClient: ApiClient
    
    constructor(options: ComponentInputs) {
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