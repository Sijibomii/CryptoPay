import { Options } from "./types";
import "./main.css";
export default class CryptoPay {
    private config;
    private navigation;
    private apiClient;
    private overlay;
    private modal;
    constructor(options: Options);
    init(): void;
    resize(): void;
    onSuccess(ticket: any): void;
    closeModal(): void;
    activate(): void;
    mount(): void;
}
