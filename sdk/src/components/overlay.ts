import { Component }from "./components";

export class Overlay extends Component {
    public mountPoint: HTMLElement | null = null;

    set(mountPoint: HTMLElement) {
        this.mountPoint = mountPoint; 
        mountPoint.appendChild(this.element);
        this.render();
    }

    open() {
        this.element.querySelector("#checkout-overlay")!.classList.add("active");
    }

    close() {
        this.element.querySelector("#checkout-overlay")!.classList.remove("active");

        setTimeout(() => {
        this.mountPoint!.removeChild(this.element);
        }, 300);
    }

    override render() {
        const template = `<div id="checkout-overlay"><div class="checkout-overlay-content"></div></div>`;

        this.element.innerHTML = template.trim();
    }
}
