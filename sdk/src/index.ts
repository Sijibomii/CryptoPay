import { Config, Options } from "./types";

export default class CryptoPay{
    private config: Config;
    
    constructor(options: Options={}) {
        this.config = Object.assign({
            modal: null
        }, options)


        
    }
}