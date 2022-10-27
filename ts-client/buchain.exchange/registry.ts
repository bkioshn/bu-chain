import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteExchangeRate } from "./types/buchain/exchange/tx";
import { MsgCreateExchangeRate } from "./types/buchain/exchange/tx";
import { MsgExchangeToken } from "./types/buchain/exchange/tx";
import { MsgUpdateExchangeRate } from "./types/buchain/exchange/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/buchain.exchange.MsgDeleteExchangeRate", MsgDeleteExchangeRate],
    ["/buchain.exchange.MsgCreateExchangeRate", MsgCreateExchangeRate],
    ["/buchain.exchange.MsgExchangeToken", MsgExchangeToken],
    ["/buchain.exchange.MsgUpdateExchangeRate", MsgUpdateExchangeRate],
    
];

export { msgTypes }