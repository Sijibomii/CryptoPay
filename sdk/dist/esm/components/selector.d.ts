import { Component, ComponentInputs } from "./components";
export declare class Select extends Component {
    buttons: NodeListOf<Element>;
    constructor(options: ComponentInputs);
    onMounted(): void;
    activateCurrencySelectButton(): void;
    navigateToPayment(currency: string, e: Event): Promise<void>;
    renderButton(currency: string): string;
    render(): void;
}
