import { Navigation, ApiClient } from "./lib";
import { Config, Options } from "./types";
import {
    Overlay,
    Modal,
    Select as Selector,
    Payment,
    Confirmation,
    Success,
    PaymentError
} from "./components";

import "./main.css";

// take redirect link. redirect to payment page after payment redirect back

export default class CryptoPay{

    private config: Config;
    private navigation: Navigation
    private apiClient: ApiClient
    private overlay: Overlay
    private modal: Modal
    constructor(options: Options) {
        this.config = Object.assign({
            modal: null
        }, options)


        this.navigation = new Navigation()

        this.apiClient = new ApiClient({
            config: this.config
        });

        this.overlay = new Overlay({
            navigation: this.navigation,
            apiClient: this.apiClient,
            config: this.config,
            currencies: this.config.currencies,
            onSuccess: this.onSuccess.bind(this),
          });

        this.modal = new Modal({
            overlay: this.overlay,
            navigation: this.navigation,
            apiClient: this.apiClient,
            config: this.config,
            currencies: this.config.currencies,
            onSuccess: this.onSuccess.bind(this),
        });

        this.navigation.init({
            routes: {
              selector: new Selector({
                navigation: this.navigation,
                apiClient: this.apiClient,
                config: this.config,
                modal: this.modal,
                overlay: this.overlay,
                currencies: this.config.currencies,
                onSuccess: this.onSuccess.bind(this),
              }),
              payment: new Payment({
                navigation: this.navigation,
                apiClient: this.apiClient,
                config: this.config,
                modal: this.modal,
                overlay: this.overlay,
                currencies: this.config.currencies,
                onSuccess: this.onSuccess.bind(this),
              }), 
              confirmation: new Confirmation({
                navigation: this.navigation,
                apiClient: this.apiClient,
                config: this.config,
                modal: this.modal,
                overlay: this.overlay,
                currencies: this.config.currencies,
                onSuccess: this.onSuccess.bind(this),
              }),
              success: new Success({
                navigation: this.navigation,
                apiClient: this.apiClient,
                onSuccess: this.onSuccess.bind(this),
                modal: this.modal,
                overlay: this.overlay,
                currencies: this.config.currencies,
              }),
              error: new PaymentError({
                navigation: this.navigation,
                apiClient: this.apiClient,
                onSuccess: this.onSuccess.bind(this),
                modal: this.modal,
                overlay: this.overlay,
                currencies: this.config.currencies,
              })
            }
          });

    }

    init() {
      this.mount();
    }
    
    resize() {}
    
    onSuccess(ticket: any) {
        this.config.onSuccess(ticket);
    }
    
    closeModal() {
        this.modal.close();
        this.navigation.reset();
    }
    
      activate() {

        // redirect page to new page
        // redirect to payment page
        
        this.overlay.set(document.body);
    
        this.navigation.target = this.overlay.element.querySelector(
          ".checkout-overlay-content"
        ) as HTMLElement;
    
        this.navigation.navigate("selector", {
            dataset: ''
        });
    
        setTimeout(() => {
          this.overlay.open();
        }, 0);
      }
    
      mount() {
        this.config.button!.addEventListener("click", this.activate.bind(this));
      }
}