import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgExchangeToken } from "./types/buchain/exchange/tx";
import { MsgUpdateExchangeRate } from "./types/buchain/exchange/tx";
import { MsgCreateExchangeRate } from "./types/buchain/exchange/tx";
import { MsgDeleteExchangeRate } from "./types/buchain/exchange/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/buchain.exchange.MsgExchangeToken", MsgExchangeToken],
    ["/buchain.exchange.MsgUpdateExchangeRate", MsgUpdateExchangeRate],
    ["/buchain.exchange.MsgCreateExchangeRate", MsgCreateExchangeRate],
    ["/buchain.exchange.MsgDeleteExchangeRate", MsgDeleteExchangeRate],
    
];

export { msgTypes }