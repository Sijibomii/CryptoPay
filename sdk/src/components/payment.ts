import { Component, ComponentInputs } from "./components";
import QRCode from "qrcode";

export class Payment extends Component {

    public timeLeft: string;
    public copyButton:  Element;
    public timer: NodeJS.Timeout
    public polling: NodeJS.Timeout
    constructor(options: ComponentInputs) {
        super(options)
        // could not get a better initializer
        this.timeLeft = "";
        this.copyButton = this.element.querySelector("") as Element;
        this.timer =  setTimeout(() => {}, 100);
        this.polling =  setTimeout(() => {}, 100);
    }

    override async onMounted() {

        this.timeLeft = "15:00";

        try {
            this.copyButton = this.element.querySelector(".copy-address") as Element;

            this.copyButton.addEventListener("click", this.copyAddress.bind(this));

            const qrcodeElement = this.element.querySelector(".qrcode");
            const qrcodeOptions = {
            text: `${this.dataset.payment.address}`,
            width: 100,
            height: 100,
            colorDark: "#4a4a4a",
            colorLight: "#ffffff"
            };

            QRCode.toCanvas(qrcodeElement, JSON.stringify(qrcodeOptions), (error) => {
                if (error) {
                  console.error("Failed to generate QR code:", error);
                } else {
                  console.log("QR code generated successfully");
                }
              });
        

            this.startCountdown();
            await this.pollPaymentStatus();
        } catch (error) {
          this.navigation!.navigate.call(this.navigation, "error", {
            dataset: Object.assign(this.dataset, {
              message: "Error detected during payment session."
            })
          });
          return;
        }
    }
    
    override onUnmounted() {
        clearTimeout(this.polling);
        clearTimeout(this.timer);
    
        if (this.copyButton)
          this.copyButton.removeEventListener("click", this.copyAddress);
    }
    
    copyAddress() {
        var el = document.createElement("textarea");
        el.value = this.dataset.payment.address;
        el.setAttribute("readonly", "");
        el.style.position = "absolute"
        el.style.left= "-9999px"

        document.body.appendChild(el);
        el.select();
        this.copyToClipboard(el.value)
        document.body.removeChild(el);
      }

    copyToClipboard(text: string): void {
        navigator.clipboard.writeText(text)
          .then(() => {
            console.log("Text copied to clipboard");
          })
          .catch((error) => {
            console.error("Failed to copy text to clipboard:", error);
          });
    }
    
    startCountdown() {
        let presentTime = this.timeLeft;
        let timeArray = presentTime.split(/[:]+/);
        let m: any = timeArray[0];
    
        let s: any= parseInt(timeArray[1]!) - 1;
    
        if (s < 10 && s >= 0) {
          s = "0" + s;
        }
    
        if (s < 0) {
          s = "59";
        }
    
        if (s == 59) {
          m = parseInt(m!) - 1;
        }
    
        this.timeLeft = `${m}:${s}`;
    
        if (this.timeLeft == "0:00") {
          this.navigation!.navigate.call(this.navigation, "error", {
            dataset: Object.assign(this.dataset, {
              message: "Payment session has expired."
            })
          });
          return;
    }
    
        this.element.querySelector(".timer")!.innerHTML = this.timeLeft;
    
        this.timer = setTimeout(this.startCountdown.bind(this), 1000);
      }

    
    async pollPaymentStatus(): Promise<any>{
        const resp = await this.apiClient.getPaymentStatus(
            this.dataset.payment.id,
            this.dataset.sessionToken
        );

        if (resp.status == "pending") {
            await new Promise(resolve => (this.polling = setTimeout(resolve, 4000)));
            return await this.pollPaymentStatus();
        }

        this.navigation!.navigate.call(this.navigation, "confirmation", {
            dataset: this.dataset
        });
    }
    
    override render() {
        let expiration = new Date();
        expiration.setMinutes(expiration.getMinutes() + 15);

        let expiresAt = `${expiration.getHours()}:${(
            "0" + expiration.getMinutes()
        ).substring(-2)}`;

        let { currency, payment, store } = this.dataset;

        let network;
        switch (payment.crypto) {
            case "btc":
            network = payment.btc_network;
            break;
        }

        const template = `
                <div class="modal-content payment ${currency}">
                    <div class="header">
                        <div class="title">${store.name}</div>
                        <p class="message">
                            Please send crypto to the address below.
                        </p>
                        <div class="amount">${
                            payment.charge
                        } ${currency.toUpperCase()}</div>
                        <div class="converted-amount">(${
                            payment.price
                        } ${payment.fiat.toUpperCase()})</div>
                        <div class="payment-qrcode-wrapper">
                            <div class="qrcode">

                            </div>
                        </div>
                    </div>
                    <div class="body">
                        <div class="address">${payment.address}</div>
                        <button class="copy-address">Copy the Address</button>
                        <div class="separator"><span class="or">or<span></div>
                        ${
                            currency === "eth"
                            ? `<img class="pay-with-metamask" src="" width="180" height="48" alt="Pay with metamask" /><div class="metamask-error"></div>`
                            : `<a href="bitcoin:${payment.address}?amount=${
                                payment.charge
                                }" class="open-in-wallet">Open in Wallet</a>`
                        }
                    </div>
                    <div class="network">Network: ${network.toUpperCase()}</div>
                    <div class="footer">
                        <div class="timer footer-left">15:00</div>
                        <div class="footer-right">
                            <div class="rate"><span class="bold">1${currency.toUpperCase()} = ${(
            payment.price / payment.charge
        ).toFixed(2)} ${payment.fiat.toUpperCase()}</span></div>
                            <div class="footer-message">This payment modal is valid until: <span class="bold">${expiresAt}</span></div>
                        </div>
                    </div>
                </div>`;

        this.element.innerHTML = template.trim();
    }

}