import { Navigation } from "../lib/navigation";
import { Config } from "../types";
import { Overlay } from "./overlay";
import { Modal } from "./modal";
import { ApiClient } from "../lib";
export type navigatorParam = {
    dataset: dataset;
};
type dataset = any;
export type ComponentInputs = {
    config?: Config;
    navigation?: Navigation;
    modal?: Modal;
    onSuccess: (ticket: any) => void;
    overlay?: Overlay;
    dataset?: any;
    currencies: string[];
    apiClient: ApiClient;
};
export declare class Component {
    config: Config | undefined;
    navigation: Navigation | undefined;
    modal: Modal | undefined;
    onSuccess: ((ticket: any) => void);
    element: HTMLElement;
    overlay?: Overlay | undefined;
    navigatorParams?: navigatorParam;
    dataset?: dataset;
    apiClient: ApiClient;
    constructor(options: ComponentInputs);
    willMount(): void;
    willUnmount(): void;
    onMounted(): void;
    onUnmounted(): void;
    render(): void;
}
export {};
