/// <reference types="node" />
import { Component, ComponentInputs } from "./components";
export declare class Payment extends Component {
    timeLeft: string;
    copyButton: Element;
    timer: NodeJS.Timeout;
    polling: NodeJS.Timeout;
    constructor(options: ComponentInputs);
    onMounted(): Promise<void>;
    onUnmounted(): void;
    copyAddress(): void;
    copyToClipboard(text: string): void;
    startCountdown(): void;
    pollPaymentStatus(): Promise<any>;
    render(): void;
}
