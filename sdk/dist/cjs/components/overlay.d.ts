import { Component } from "./components";
export declare class Overlay extends Component {
    mountPoint: HTMLElement | null;
    set(mountPoint: HTMLElement): void;
    open(): void;
    close(): void;
    render(): void;
}
