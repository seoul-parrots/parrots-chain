/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../parrots/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Profile } from "../parrots/models";

export const protobufPackage = "parrots.parrots";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetProfilesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryGetProfilesResponse {
  Profile: Profile[];
  pagination: PageResponse | undefined;
}

export interface QueryProfileCountRequest {}

export interface QueryProfileCountResponse {
  count: number;
}

export interface QueryGetProfileRequest {
  id: number;
}

export interface QueryGetProfileResponse {
  profile: Profile | undefined;
}

export interface QueryGetProfileByUsernameRequest {
  username: string;
}

export interface QueryGetProfileByUsernameResponse {
  profile: Profile | undefined;
}

export interface QueryGetBeaksCountRequest {}

export interface QueryGetBeaksCountResponse {
  count: number;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetProfilesRequest: object = {};

export const QueryGetProfilesRequest = {
  encode(
    message: QueryGetProfilesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProfilesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfilesRequest,
    } as QueryGetProfilesRequest;
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

  fromJSON(object: any): QueryGetProfilesRequest {
    const message = {
      ...baseQueryGetProfilesRequest,
    } as QueryGetProfilesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetProfilesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfilesRequest>
  ): QueryGetProfilesRequest {
    const message = {
      ...baseQueryGetProfilesRequest,
    } as QueryGetProfilesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetProfilesResponse: object = {};

export const QueryGetProfilesResponse = {
  encode(
    message: QueryGetProfilesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Profile) {
      Profile.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetProfilesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfilesResponse,
    } as QueryGetProfilesResponse;
    message.Profile = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Profile.push(Profile.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryGetProfilesResponse {
    const message = {
      ...baseQueryGetProfilesResponse,
    } as QueryGetProfilesResponse;
    message.Profile = [];
    if (object.Profile !== undefined && object.Profile !== null) {
      for (const e of object.Profile) {
        message.Profile.push(Profile.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetProfilesResponse): unknown {
    const obj: any = {};
    if (message.Profile) {
      obj.Profile = message.Profile.map((e) =>
        e ? Profile.toJSON(e) : undefined
      );
    } else {
      obj.Profile = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfilesResponse>
  ): QueryGetProfilesResponse {
    const message = {
      ...baseQueryGetProfilesResponse,
    } as QueryGetProfilesResponse;
    message.Profile = [];
    if (object.Profile !== undefined && object.Profile !== null) {
      for (const e of object.Profile) {
        message.Profile.push(Profile.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryProfileCountRequest: object = {};

export const QueryProfileCountRequest = {
  encode(
    _: QueryProfileCountRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryProfileCountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryProfileCountRequest,
    } as QueryProfileCountRequest;
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

  fromJSON(_: any): QueryProfileCountRequest {
    const message = {
      ...baseQueryProfileCountRequest,
    } as QueryProfileCountRequest;
    return message;
  },

  toJSON(_: QueryProfileCountRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryProfileCountRequest>
  ): QueryProfileCountRequest {
    const message = {
      ...baseQueryProfileCountRequest,
    } as QueryProfileCountRequest;
    return message;
  },
};

const baseQueryProfileCountResponse: object = { count: 0 };

export const QueryProfileCountResponse = {
  encode(
    message: QueryProfileCountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryProfileCountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryProfileCountResponse,
    } as QueryProfileCountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryProfileCountResponse {
    const message = {
      ...baseQueryProfileCountResponse,
    } as QueryProfileCountResponse;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    return message;
  },

  toJSON(message: QueryProfileCountResponse): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryProfileCountResponse>
  ): QueryProfileCountResponse {
    const message = {
      ...baseQueryProfileCountResponse,
    } as QueryProfileCountResponse;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    return message;
  },
};

const baseQueryGetProfileRequest: object = { id: 0 };

export const QueryGetProfileRequest = {
  encode(
    message: QueryGetProfileRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProfileRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetProfileRequest } as QueryGetProfileRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProfileRequest {
    const message = { ...baseQueryGetProfileRequest } as QueryGetProfileRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetProfileRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfileRequest>
  ): QueryGetProfileRequest {
    const message = { ...baseQueryGetProfileRequest } as QueryGetProfileRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetProfileResponse: object = {};

export const QueryGetProfileResponse = {
  encode(
    message: QueryGetProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.profile !== undefined) {
      Profile.encode(message.profile, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfileResponse,
    } as QueryGetProfileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.profile = Profile.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProfileResponse {
    const message = {
      ...baseQueryGetProfileResponse,
    } as QueryGetProfileResponse;
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromJSON(object.profile);
    } else {
      message.profile = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetProfileResponse): unknown {
    const obj: any = {};
    message.profile !== undefined &&
      (obj.profile = message.profile
        ? Profile.toJSON(message.profile)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfileResponse>
  ): QueryGetProfileResponse {
    const message = {
      ...baseQueryGetProfileResponse,
    } as QueryGetProfileResponse;
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromPartial(object.profile);
    } else {
      message.profile = undefined;
    }
    return message;
  },
};

const baseQueryGetProfileByUsernameRequest: object = { username: "" };

export const QueryGetProfileByUsernameRequest = {
  encode(
    message: QueryGetProfileByUsernameRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetProfileByUsernameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfileByUsernameRequest,
    } as QueryGetProfileByUsernameRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProfileByUsernameRequest {
    const message = {
      ...baseQueryGetProfileByUsernameRequest,
    } as QueryGetProfileByUsernameRequest;
    if (object.username !== undefined && object.username !== null) {
      message.username = String(object.username);
    } else {
      message.username = "";
    }
    return message;
  },

  toJSON(message: QueryGetProfileByUsernameRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfileByUsernameRequest>
  ): QueryGetProfileByUsernameRequest {
    const message = {
      ...baseQueryGetProfileByUsernameRequest,
    } as QueryGetProfileByUsernameRequest;
    if (object.username !== undefined && object.username !== null) {
      message.username = object.username;
    } else {
      message.username = "";
    }
    return message;
  },
};

const baseQueryGetProfileByUsernameResponse: object = {};

export const QueryGetProfileByUsernameResponse = {
  encode(
    message: QueryGetProfileByUsernameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.profile !== undefined) {
      Profile.encode(message.profile, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetProfileByUsernameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetProfileByUsernameResponse,
    } as QueryGetProfileByUsernameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.profile = Profile.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProfileByUsernameResponse {
    const message = {
      ...baseQueryGetProfileByUsernameResponse,
    } as QueryGetProfileByUsernameResponse;
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromJSON(object.profile);
    } else {
      message.profile = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetProfileByUsernameResponse): unknown {
    const obj: any = {};
    message.profile !== undefined &&
      (obj.profile = message.profile
        ? Profile.toJSON(message.profile)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProfileByUsernameResponse>
  ): QueryGetProfileByUsernameResponse {
    const message = {
      ...baseQueryGetProfileByUsernameResponse,
    } as QueryGetProfileByUsernameResponse;
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromPartial(object.profile);
    } else {
      message.profile = undefined;
    }
    return message;
  },
};

const baseQueryGetBeaksCountRequest: object = {};

export const QueryGetBeaksCountRequest = {
  encode(
    _: QueryGetBeaksCountRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetBeaksCountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBeaksCountRequest,
    } as QueryGetBeaksCountRequest;
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

  fromJSON(_: any): QueryGetBeaksCountRequest {
    const message = {
      ...baseQueryGetBeaksCountRequest,
    } as QueryGetBeaksCountRequest;
    return message;
  },

  toJSON(_: QueryGetBeaksCountRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetBeaksCountRequest>
  ): QueryGetBeaksCountRequest {
    const message = {
      ...baseQueryGetBeaksCountRequest,
    } as QueryGetBeaksCountRequest;
    return message;
  },
};

const baseQueryGetBeaksCountResponse: object = { count: 0 };

export const QueryGetBeaksCountResponse = {
  encode(
    message: QueryGetBeaksCountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetBeaksCountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBeaksCountResponse,
    } as QueryGetBeaksCountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBeaksCountResponse {
    const message = {
      ...baseQueryGetBeaksCountResponse,
    } as QueryGetBeaksCountResponse;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    return message;
  },

  toJSON(message: QueryGetBeaksCountResponse): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBeaksCountResponse>
  ): QueryGetBeaksCountResponse {
    const message = {
      ...baseQueryGetBeaksCountResponse,
    } as QueryGetBeaksCountResponse;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Profiles items. */
  GetProfiles(
    request: QueryGetProfilesRequest
  ): Promise<QueryGetProfilesResponse>;
  /** Queries a list of ProfileCount items. */
  ProfileCount(
    request: QueryProfileCountRequest
  ): Promise<QueryProfileCountResponse>;
  /** Queries a list of GetProfile items. */
  GetProfile(request: QueryGetProfileRequest): Promise<QueryGetProfileResponse>;
  /** Queries a list of GetProfileByUsername items. */
  GetProfileByUsername(
    request: QueryGetProfileByUsernameRequest
  ): Promise<QueryGetProfileByUsernameResponse>;
  /** Queries a list of GetBeaksCount items. */
  GetBeaksCount(
    request: QueryGetBeaksCountRequest
  ): Promise<QueryGetBeaksCountResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("parrots.parrots.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  GetProfiles(
    request: QueryGetProfilesRequest
  ): Promise<QueryGetProfilesResponse> {
    const data = QueryGetProfilesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "parrots.parrots.Query",
      "GetProfiles",
      data
    );
    return promise.then((data) =>
      QueryGetProfilesResponse.decode(new Reader(data))
    );
  }

  ProfileCount(
    request: QueryProfileCountRequest
  ): Promise<QueryProfileCountResponse> {
    const data = QueryProfileCountRequest.encode(request).finish();
    const promise = this.rpc.request(
      "parrots.parrots.Query",
      "ProfileCount",
      data
    );
    return promise.then((data) =>
      QueryProfileCountResponse.decode(new Reader(data))
    );
  }

  GetProfile(
    request: QueryGetProfileRequest
  ): Promise<QueryGetProfileResponse> {
    const data = QueryGetProfileRequest.encode(request).finish();
    const promise = this.rpc.request(
      "parrots.parrots.Query",
      "GetProfile",
      data
    );
    return promise.then((data) =>
      QueryGetProfileResponse.decode(new Reader(data))
    );
  }

  GetProfileByUsername(
    request: QueryGetProfileByUsernameRequest
  ): Promise<QueryGetProfileByUsernameResponse> {
    const data = QueryGetProfileByUsernameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "parrots.parrots.Query",
      "GetProfileByUsername",
      data
    );
    return promise.then((data) =>
      QueryGetProfileByUsernameResponse.decode(new Reader(data))
    );
  }

  GetBeaksCount(
    request: QueryGetBeaksCountRequest
  ): Promise<QueryGetBeaksCountResponse> {
    const data = QueryGetBeaksCountRequest.encode(request).finish();
    const promise = this.rpc.request(
      "parrots.parrots.Query",
      "GetBeaksCount",
      data
    );
    return promise.then((data) =>
      QueryGetBeaksCountResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
