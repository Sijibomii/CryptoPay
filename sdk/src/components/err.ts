import {Component} from "./components";
import errorImg from "../images/btc.png";

export class PaymentError extends Component {

  override onMounted() {
    let closeButton = this.modal!.element.querySelector(".btn-close");

    if (closeButton!.classList.contains("disabled")) {
      closeButton!.classList.remove("disabled");
    }
  }

  override render() {
    const message = this.navigatorParams!.dataset.message;

    const template = `
            <div class="modal-content simple">
                <div class="icon-error">
                  <img src="${errorImg}" width="9" height="34" />
                </div>
                <div class="title">Error</div>
                <p class="message">
                    ${message}
                </p>
            </div>
        `;

    this.element.innerHTML = template.trim();
  }
}