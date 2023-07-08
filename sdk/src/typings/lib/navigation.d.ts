import { Component, navigatorParam } from "../components";
import { Target } from "../types";
type navigationInitInput = {
    routes: Record<string, Component>;
};
export declare class Navigation {
    target?: Target;
    private current;
    private routes;
    constructor();
    reset(): void;
    init(options: navigationInitInput): void;
    clear(): void;
    navigate(to: string, params: navigatorParam): void;
}
export {};
