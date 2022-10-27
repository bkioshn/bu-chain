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

export interface MsgCreateExchangeRate {
  creator: string;
  index: string;
  rate: string;
}

export interface MsgCreateExchangeRateResponse {
}

export interface MsgUpdateExchangeRate {
  creator: string;
  index: string;
  rate: string;
}

export interface MsgUpdateExchangeRateResponse {
}

export interface MsgDeleteExchangeRate {
  creator: string;
  index: string;
}

export interface MsgDeleteExchangeRateResponse {
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

function createBaseMsgCreateExchangeRate(): MsgCreateExchangeRate {
  return { creator: "", index: "", rate: "" };
}

export const MsgCreateExchangeRate = {
  encode(message: MsgCreateExchangeRate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.rate !== "") {
      writer.uint32(26).string(message.rate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateExchangeRate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateExchangeRate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.rate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateExchangeRate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      rate: isSet(object.rate) ? String(object.rate) : "",
    };
  },

  toJSON(message: MsgCreateExchangeRate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.rate !== undefined && (obj.rate = message.rate);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateExchangeRate>, I>>(object: I): MsgCreateExchangeRate {
    const message = createBaseMsgCreateExchangeRate();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.rate = object.rate ?? "";
    return message;
  },
};

function createBaseMsgCreateExchangeRateResponse(): MsgCreateExchangeRateResponse {
  return {};
}

export const MsgCreateExchangeRateResponse = {
  encode(_: MsgCreateExchangeRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateExchangeRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateExchangeRateResponse();
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

  fromJSON(_: any): MsgCreateExchangeRateResponse {
    return {};
  },

  toJSON(_: MsgCreateExchangeRateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateExchangeRateResponse>, I>>(_: I): MsgCreateExchangeRateResponse {
    const message = createBaseMsgCreateExchangeRateResponse();
    return message;
  },
};

function createBaseMsgUpdateExchangeRate(): MsgUpdateExchangeRate {
  return { creator: "", index: "", rate: "" };
}

export const MsgUpdateExchangeRate = {
  encode(message: MsgUpdateExchangeRate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.rate !== "") {
      writer.uint32(26).string(message.rate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateExchangeRate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateExchangeRate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.rate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateExchangeRate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      rate: isSet(object.rate) ? String(object.rate) : "",
    };
  },

  toJSON(message: MsgUpdateExchangeRate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.rate !== undefined && (obj.rate = message.rate);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateExchangeRate>, I>>(object: I): MsgUpdateExchangeRate {
    const message = createBaseMsgUpdateExchangeRate();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.rate = object.rate ?? "";
    return message;
  },
};

function createBaseMsgUpdateExchangeRateResponse(): MsgUpdateExchangeRateResponse {
  return {};
}

export const MsgUpdateExchangeRateResponse = {
  encode(_: MsgUpdateExchangeRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateExchangeRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateExchangeRateResponse();
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

  fromJSON(_: any): MsgUpdateExchangeRateResponse {
    return {};
  },

  toJSON(_: MsgUpdateExchangeRateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateExchangeRateResponse>, I>>(_: I): MsgUpdateExchangeRateResponse {
    const message = createBaseMsgUpdateExchangeRateResponse();
    return message;
  },
};

function createBaseMsgDeleteExchangeRate(): MsgDeleteExchangeRate {
  return { creator: "", index: "" };
}

export const MsgDeleteExchangeRate = {
  encode(message: MsgDeleteExchangeRate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteExchangeRate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteExchangeRate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteExchangeRate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
    };
  },

  toJSON(message: MsgDeleteExchangeRate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteExchangeRate>, I>>(object: I): MsgDeleteExchangeRate {
    const message = createBaseMsgDeleteExchangeRate();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseMsgDeleteExchangeRateResponse(): MsgDeleteExchangeRateResponse {
  return {};
}

export const MsgDeleteExchangeRateResponse = {
  encode(_: MsgDeleteExchangeRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteExchangeRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteExchangeRateResponse();
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

  fromJSON(_: any): MsgDeleteExchangeRateResponse {
    return {};
  },

  toJSON(_: MsgDeleteExchangeRateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteExchangeRateResponse>, I>>(_: I): MsgDeleteExchangeRateResponse {
    const message = createBaseMsgDeleteExchangeRateResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ExchangeToken(request: MsgExchangeToken): Promise<MsgExchangeTokenResponse>;
  CreateExchangeRate(request: MsgCreateExchangeRate): Promise<MsgCreateExchangeRateResponse>;
  UpdateExchangeRate(request: MsgUpdateExchangeRate): Promise<MsgUpdateExchangeRateResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteExchangeRate(request: MsgDeleteExchangeRate): Promise<MsgDeleteExchangeRateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ExchangeToken = this.ExchangeToken.bind(this);
    this.CreateExchangeRate = this.CreateExchangeRate.bind(this);
    this.UpdateExchangeRate = this.UpdateExchangeRate.bind(this);
    this.DeleteExchangeRate = this.DeleteExchangeRate.bind(this);
  }
  ExchangeToken(request: MsgExchangeToken): Promise<MsgExchangeTokenResponse> {
    const data = MsgExchangeToken.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Msg", "ExchangeToken", data);
    return promise.then((data) => MsgExchangeTokenResponse.decode(new _m0.Reader(data)));
  }

  CreateExchangeRate(request: MsgCreateExchangeRate): Promise<MsgCreateExchangeRateResponse> {
    const data = MsgCreateExchangeRate.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Msg", "CreateExchangeRate", data);
    return promise.then((data) => MsgCreateExchangeRateResponse.decode(new _m0.Reader(data)));
  }

  UpdateExchangeRate(request: MsgUpdateExchangeRate): Promise<MsgUpdateExchangeRateResponse> {
    const data = MsgUpdateExchangeRate.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Msg", "UpdateExchangeRate", data);
    return promise.then((data) => MsgUpdateExchangeRateResponse.decode(new _m0.Reader(data)));
  }

  DeleteExchangeRate(request: MsgDeleteExchangeRate): Promise<MsgDeleteExchangeRateResponse> {
    const data = MsgDeleteExchangeRate.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Msg", "DeleteExchangeRate", data);
    return promise.then((data) => MsgDeleteExchangeRateResponse.decode(new _m0.Reader(data)));
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
