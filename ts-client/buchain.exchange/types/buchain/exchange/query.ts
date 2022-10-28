/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { ExchangeRate, Params } from "./exchange";

export const protobufPackage = "buchain.exchange";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetExchangeRateRequest {
  index: string;
}

export interface QueryGetExchangeRateResponse {
  exchangeRate: ExchangeRate | undefined;
}

export interface QueryAllExchangeRateRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllExchangeRateResponse {
  exchangeRate: ExchangeRate[];
  pagination: PageResponse | undefined;
}

export interface QueryExchangeAmountRequest {
  /** cosmos.base.v1beta1.Coin denom = 1 [(gogoproto.nullable) = false]; */
  denom: string;
  amount: string;
  exchangeToken: string;
}

export interface QueryExchangeAmountResponse {
  amount: number;
}

export interface QueryExchangePairsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryExchangePairsResponse {
  exchangePair: string[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetExchangeRateRequest(): QueryGetExchangeRateRequest {
  return { index: "" };
}

export const QueryGetExchangeRateRequest = {
  encode(message: QueryGetExchangeRateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetExchangeRateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetExchangeRateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetExchangeRateRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetExchangeRateRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetExchangeRateRequest>, I>>(object: I): QueryGetExchangeRateRequest {
    const message = createBaseQueryGetExchangeRateRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetExchangeRateResponse(): QueryGetExchangeRateResponse {
  return { exchangeRate: undefined };
}

export const QueryGetExchangeRateResponse = {
  encode(message: QueryGetExchangeRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.exchangeRate !== undefined) {
      ExchangeRate.encode(message.exchangeRate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetExchangeRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetExchangeRateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.exchangeRate = ExchangeRate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetExchangeRateResponse {
    return { exchangeRate: isSet(object.exchangeRate) ? ExchangeRate.fromJSON(object.exchangeRate) : undefined };
  },

  toJSON(message: QueryGetExchangeRateResponse): unknown {
    const obj: any = {};
    message.exchangeRate !== undefined
      && (obj.exchangeRate = message.exchangeRate ? ExchangeRate.toJSON(message.exchangeRate) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetExchangeRateResponse>, I>>(object: I): QueryGetExchangeRateResponse {
    const message = createBaseQueryGetExchangeRateResponse();
    message.exchangeRate = (object.exchangeRate !== undefined && object.exchangeRate !== null)
      ? ExchangeRate.fromPartial(object.exchangeRate)
      : undefined;
    return message;
  },
};

function createBaseQueryAllExchangeRateRequest(): QueryAllExchangeRateRequest {
  return { pagination: undefined };
}

export const QueryAllExchangeRateRequest = {
  encode(message: QueryAllExchangeRateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllExchangeRateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllExchangeRateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllExchangeRateRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllExchangeRateRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllExchangeRateRequest>, I>>(object: I): QueryAllExchangeRateRequest {
    const message = createBaseQueryAllExchangeRateRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllExchangeRateResponse(): QueryAllExchangeRateResponse {
  return { exchangeRate: [], pagination: undefined };
}

export const QueryAllExchangeRateResponse = {
  encode(message: QueryAllExchangeRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.exchangeRate) {
      ExchangeRate.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllExchangeRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllExchangeRateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.exchangeRate.push(ExchangeRate.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllExchangeRateResponse {
    return {
      exchangeRate: Array.isArray(object?.exchangeRate)
        ? object.exchangeRate.map((e: any) => ExchangeRate.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllExchangeRateResponse): unknown {
    const obj: any = {};
    if (message.exchangeRate) {
      obj.exchangeRate = message.exchangeRate.map((e) => e ? ExchangeRate.toJSON(e) : undefined);
    } else {
      obj.exchangeRate = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllExchangeRateResponse>, I>>(object: I): QueryAllExchangeRateResponse {
    const message = createBaseQueryAllExchangeRateResponse();
    message.exchangeRate = object.exchangeRate?.map((e) => ExchangeRate.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryExchangeAmountRequest(): QueryExchangeAmountRequest {
  return { denom: "", amount: "", exchangeToken: "" };
}

export const QueryExchangeAmountRequest = {
  encode(message: QueryExchangeAmountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.exchangeToken !== "") {
      writer.uint32(26).string(message.exchangeToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryExchangeAmountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryExchangeAmountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.exchangeToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryExchangeAmountRequest {
    return {
      denom: isSet(object.denom) ? String(object.denom) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      exchangeToken: isSet(object.exchangeToken) ? String(object.exchangeToken) : "",
    };
  },

  toJSON(message: QueryExchangeAmountRequest): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    message.amount !== undefined && (obj.amount = message.amount);
    message.exchangeToken !== undefined && (obj.exchangeToken = message.exchangeToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryExchangeAmountRequest>, I>>(object: I): QueryExchangeAmountRequest {
    const message = createBaseQueryExchangeAmountRequest();
    message.denom = object.denom ?? "";
    message.amount = object.amount ?? "";
    message.exchangeToken = object.exchangeToken ?? "";
    return message;
  },
};

function createBaseQueryExchangeAmountResponse(): QueryExchangeAmountResponse {
  return { amount: 0 };
}

export const QueryExchangeAmountResponse = {
  encode(message: QueryExchangeAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== 0) {
      writer.uint32(8).uint64(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryExchangeAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryExchangeAmountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryExchangeAmountResponse {
    return { amount: isSet(object.amount) ? Number(object.amount) : 0 };
  },

  toJSON(message: QueryExchangeAmountResponse): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryExchangeAmountResponse>, I>>(object: I): QueryExchangeAmountResponse {
    const message = createBaseQueryExchangeAmountResponse();
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseQueryExchangePairsRequest(): QueryExchangePairsRequest {
  return { pagination: undefined };
}

export const QueryExchangePairsRequest = {
  encode(message: QueryExchangePairsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryExchangePairsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryExchangePairsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryExchangePairsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryExchangePairsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryExchangePairsRequest>, I>>(object: I): QueryExchangePairsRequest {
    const message = createBaseQueryExchangePairsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryExchangePairsResponse(): QueryExchangePairsResponse {
  return { exchangePair: [], pagination: undefined };
}

export const QueryExchangePairsResponse = {
  encode(message: QueryExchangePairsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.exchangePair) {
      writer.uint32(10).string(v!);
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryExchangePairsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryExchangePairsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.exchangePair.push(reader.string());
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryExchangePairsResponse {
    return {
      exchangePair: Array.isArray(object?.exchangePair) ? object.exchangePair.map((e: any) => String(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryExchangePairsResponse): unknown {
    const obj: any = {};
    if (message.exchangePair) {
      obj.exchangePair = message.exchangePair.map((e) => e);
    } else {
      obj.exchangePair = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryExchangePairsResponse>, I>>(object: I): QueryExchangePairsResponse {
    const message = createBaseQueryExchangePairsResponse();
    message.exchangePair = object.exchangePair?.map((e) => e) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a ExchangeRate by index. */
  ExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse>;
  /** Queries a list of ExchangeRate items. */
  ExchangeRateAll(request: QueryAllExchangeRateRequest): Promise<QueryAllExchangeRateResponse>;
  /** Queries exchange-amount */
  ExchangeAmount(request: QueryExchangeAmountRequest): Promise<QueryExchangeAmountResponse>;
  /** Queries list of exchange pair */
  ExchangePairs(request: QueryExchangePairsRequest): Promise<QueryExchangePairsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.ExchangeRate = this.ExchangeRate.bind(this);
    this.ExchangeRateAll = this.ExchangeRateAll.bind(this);
    this.ExchangeAmount = this.ExchangeAmount.bind(this);
    this.ExchangePairs = this.ExchangePairs.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  ExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse> {
    const data = QueryGetExchangeRateRequest.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Query", "ExchangeRate", data);
    return promise.then((data) => QueryGetExchangeRateResponse.decode(new _m0.Reader(data)));
  }

  ExchangeRateAll(request: QueryAllExchangeRateRequest): Promise<QueryAllExchangeRateResponse> {
    const data = QueryAllExchangeRateRequest.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Query", "ExchangeRateAll", data);
    return promise.then((data) => QueryAllExchangeRateResponse.decode(new _m0.Reader(data)));
  }

  ExchangeAmount(request: QueryExchangeAmountRequest): Promise<QueryExchangeAmountResponse> {
    const data = QueryExchangeAmountRequest.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Query", "ExchangeAmount", data);
    return promise.then((data) => QueryExchangeAmountResponse.decode(new _m0.Reader(data)));
  }

  ExchangePairs(request: QueryExchangePairsRequest): Promise<QueryExchangePairsResponse> {
    const data = QueryExchangePairsRequest.encode(request).finish();
    const promise = this.rpc.request("buchain.exchange.Query", "ExchangePairs", data);
    return promise.then((data) => QueryExchangePairsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
