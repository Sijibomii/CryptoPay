import { Component, ComponentInputs } from "./components";
import bitcoinImg from "../images/btc.png";
import { Config } from "../types";


export class Select extends Component {

    public buttons:  NodeListOf<Element>;


    constructor(options: ComponentInputs) {
        super(options)
        // could not get a better initializer
        this.buttons = this.element.querySelectorAll("");
    }
    override onMounted() {
        this.activateCurrencySelectButton();
        this.element.querySelector(".btn-close")!.addEventListener("click", this.overlay!.close.bind(this.overlay));
    }

    activateCurrencySelectButton() {
        this.buttons = this.element.querySelectorAll(".btn-currency");
        this.buttons.forEach(button => {
            const el = button as HTMLElement
            button.addEventListener(
                "click",
                (event: Event) => this.navigateToPayment(el.dataset['type']!, event)
            );
        });
    }

    async navigateToPayment(currency: string, e: Event) { 
        e.stopPropagation();

        this.navigation!.target!.removeChild(this.element);

        this.modal!.set(this.overlay!.element.querySelector("checkout-overlay-content") as HTMLElement);

        this.navigation!.target = this.modal?.element.querySelector(".content-root") as HTMLElement;

        let data = await this.apiClient.createPayment(currency);

        this.navigation!.navigate.call(this.navigation, "payment", {
            dataset: Object.assign(this.dataset, {
                currency,
                payment: data.payment,
                store: data.store,
                sessionToken: data.token
            })
        });

        this.modal!.open();
    }

    renderButton(currency: string) {
        switch (currency) {
            case "btc":
                return `<button class="btn-currency" data-type="btc">
                        <img src="${bitcoinImg}" width="98" height="109" />
                        <span class="label">Bitcoin</span>
                        </button>`;

            default:
                return ""
            }
    }

    override render() {
        const { currencies } = this.config as Config

        const template = `
        <div class="currency-selector">
            <button class="btn-close"></button>
            <p class="message">
            Select the cryptocurrency you want to use.
            </p>
            <div class="body">
            ${currencies.map(c => this.renderButton(c)).join("")}
            </div>
        </div>`;

        this.element.innerHTML = template.trim();
    }
}
