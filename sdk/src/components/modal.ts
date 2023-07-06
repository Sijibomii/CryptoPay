import { Component } from "./components";

export class Modal extends Component {

    public mountPoint: HTMLElement | null = null;

    set(mountPoint: HTMLElement) {
        this.mountPoint = mountPoint;
        mountPoint.appendChild(this.element);
        this.render(); 
    }

    open() {
        setImmediate(() => {
        this.element.querySelector("#modal-container")!.classList.add("active");
        });
        this.element
        .querySelector(".btn-close")!
        .addEventListener("click", this.close.bind(this), { once: true });
    }

    close() {
        this.element.querySelector("#modal-container")!.classList.remove("active");
        this.overlay!.close();

        setTimeout(() => {
        this.mountPoint!.removeChild(this.element);
        }, 300);

        this.navigation!.clear();
    }

    override render() {
        const template = `
                <div id="modal-container">
                    <div class="sk-fading-circle">
                        <div class="sk-circle1 sk-circle"></div>
                        <div class="sk-circle2 sk-circle"></div>
                        <div class="sk-circle3 sk-circle"></div>
                        <div class="sk-circle4 sk-circle"></div>
                        <div class="sk-circle5 sk-circle"></div>
                        <div class="sk-circle6 sk-circle"></div>
                        <div class="sk-circle7 sk-circle"></div>
                        <div class="sk-circle8 sk-circle"></div>
                        <div class="sk-circle9 sk-circle"></div>
                        <div class="sk-circle10 sk-circle"></div>
                        <div class="sk-circle11 sk-circle"></div>
                        <div class="sk-circle12 sk-circle"></div>
                    </div>
                    
                    <div class="modal">
                        <div class="content-root"></div>
                        <button class="btn-close" />
                    </div>
                </div>`;

        this.element.innerHTML = template.trim();
    }
}
