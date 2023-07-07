import { Component, ComponentInputs }from "./components";

export class Confirmation extends Component {
    public polling: NodeJS.Timeout

    constructor(options: ComponentInputs) {
        super(options)

        this.polling =  setTimeout(() => {}, 100);
    }
    override async onMounted() {
        this.modal?.element.querySelector(".btn-close")!.classList.add("disabled");
        try {
        await this.pollPaymentStatus();
        } catch (error: any) {
        this.navigation!.navigate.call(this.navigation, "error", {
            dataset: Object.assign(this.dataset, {
            message: error.message
            })
        });
        }
    }

    override onUnmounted() {
        clearTimeout(this.polling);
    }

    async pollPaymentStatus(): Promise<any> {
        const dataset = this.dataset;
        const resp = await this.apiClient.getPaymentStatus(
        dataset.payment.id,
        dataset.sessionToken
        );

        this.dataset.confirmations_required = parseInt(resp.confirmations_required);
        this.dataset.remaining_confirmations = parseInt(
        resp.remaining_confirmations
        );

        if (resp.status == "insufficient_amount") {
            throw new Error(
                "Insufficient amount paid. Please contact store for the refundation."
            );
        }

        this.render();

        if (this.dataset.remaining_confirmations == 0) {
            await new Promise(resolve => setTimeout(resolve, 500));
            this.navigation!.navigate.call(this.navigation, "success", {
                dataset: {
                payment: dataset.payment,
                sessionToken: dataset.sessionToken,
                currency: dataset.currency
                }
            });
            return;
        }

        await new Promise(resolve => (this.polling = setTimeout(resolve, 5000)));
        return await this.pollPaymentStatus();
    }

    override render() {
        let {
        payment,
        remaining_confirmations,
        confirmations_required,
        currency,
        transactionHash
        } = this.dataset;
        let progress = `${((confirmations_required - remaining_confirmations) /
        confirmations_required) *
        100}%`;

        let etherscan_url;
        if (transactionHash && currency === "eth") {
        etherscan_url = `https://${
            payment.eth_network !== "main" ? `${payment.eth_network}.` : ""
        }etherscan.io/tx/${transactionHash}`;
        }

        let template;

        if (!this.dataset.confirmations_required) {
        template = `
                    <div class="modal-content confirmation ${currency}">
                        <div class="header">
                            <div class="title">Waiting for Confirmation</div>
                            <p class="message">
                                We have detected your payment. Please wait for - more confirmations.
                            </p>
                        </div>
                        <div class="body">
                            <div class="confirmations">- / -</div>
                            <div class="progress-bar">

                            </div>
                            <div class="message">Serveral confirmations are required to ensure that the transaction is included to the blockchain.</div>
                        </div>
                    </div>`;
        } else {
        template = `
                    <div class="modal-content confirmation ${currency}">
                        <div class="header">
                            <div class="title">Waiting for Confirmation</div>
                            <p class="message">
                                We have detected your payment. Please wait for ${remaining_confirmations} more confirmations.
                            </p>
                        </div>
                        <div class="body">
                            <div class="confirmations">${confirmations_required -
                            remaining_confirmations} / ${confirmations_required}</div>
                            <div class="progress-bar">
                                <div class="progress" style="width: ${progress};"></div>
                            </div>
                            <div class="message">
                                Serveral confirmations are required to ensure that the transaction is included to the blockchain.
                                ${
                                etherscan_url
                                    ? `You can also check confirmation status here. <a href="${etherscan_url}" target="_blank" rel="noopener noreferrer">${etherscan_url}</a>`
                                    : ""
                                }
                            </div>
                        </div>
                    </div>`;
        }

        this.element.innerHTML = template.trim();
    }
}
