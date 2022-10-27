/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "buchain.exchange";

export interface MsgExchangeToken {
  creator: string;
  receiver: string;
  denom: Coin | undefined;
  exchangeDenom: string;
}

export interface MsgExchangeTokenResponse {
}

function createBaseMsgExchangeToken(): MsgExchangeToken {
  return { creator: "", receiver: "", denom: undefined, exchangeDenom: "" };
}

export const MsgExchangeToken = {
  encode(message: MsgExchangeToken, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.denom !== undefined) {
      Coin.encode(message.denom, writer.uint32(26).fork()).ldelim();
    }
    if (message.exchangeDenom !== "") {
      writer.uint32(34).string(message.exchangeDenom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgExchangeToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgExchangeToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.denom = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.exchangeDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExchangeToken {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      denom: isSet(object.denom) ? Coin.fromJSON(object.denom) : undefined,
      exchangeDenom: isSet(object.exchangeDenom) ? String(object.exchangeDenom) : "",
    };
  },

  toJSON(message: MsgExchangeToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.denom !== undefined && (obj.denom = message.denom ? Coin.toJSON(message.denom) : undefined);
    message.exchangeDenom !== undefined && (obj.exchangeDenom = message.exchangeDenom);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgExchangeToken>, I>>(object: I): MsgExchangeToken {
    const message = createBaseMsgExchangeToken();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.denom = (object.denom !== undefined && object.denom !== null) ? Coin.fromPartial(object.denom) : undefined;
    message.exchangeDenom = object.exchangeDenom ?? "";
    return message;
  },
};

function createBaseMsgExchangeTokenResponse(): MsgExchangeTokenResponse {
  return {};
}

export const MsgExchangeTokenResponse = {
  encode(_: MsgExchangeTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgExchangeTokenResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgExchangeTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgExchangeTokenResponse {
    return {};
  },

  toJSON(_: MsgExchangeTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgExchangeTokenResponse>, I>>(_: I): MsgExchangeTokenResponse {
    const message = createBaseMsgExchangeTokenResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ExchangeToken(request: MsgExchangeToken): Promise<MsgExchangeTokenResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ExchangeToken = this.ExchangeToken.bind(this);
  }
  ExchangeToken(request: MsgExchangeToken): Promise<MsgExchangeTokenResponse> {
    const data = MsgExchangeToken.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Msg", "ExchangeToken", data);
    return promise.then((data) => MsgExchangeTokenResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
