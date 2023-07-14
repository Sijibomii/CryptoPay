import { Component } from "./components";

export class Success extends Component {

    override onMounted() {
        this.modal!.element.querySelector(".btn-close")!.classList.remove("disabled");

        // const dataset = this.dataset;
        // this.apiClient
        // .createVoucher(dataset.payment.id, dataset.sessionToken)
        // .then(data => {
        //     this.onSuccess(data.voucher);
        // });
    }

    override render() {
        const { currency } = this.dataset;

        const template = `
                <div class="modal-content simple ${currency}">
                    <div class="icon-check"></div>
                    <div class="title">Sucessfully Completed</div>
                    <p class="message">
                        We have confirmed your payment
                    </p>
                </div>
            `;

        this.element.innerHTML = template.trim();
    }
}
