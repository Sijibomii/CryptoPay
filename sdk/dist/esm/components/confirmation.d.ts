/// <reference types="node" />
import { Component, ComponentInputs } from "./components";
export declare class Confirmation extends Component {
    polling: NodeJS.Timeout;
    constructor(options: ComponentInputs);
    onMounted(): Promise<void>;
    onUnmounted(): void;
    pollPaymentStatus(): Promise<any>;
    render(): void;
}
