/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "parrots.parrots";

export interface MsgSetProfile {
  creator: string;
  username: string;
  displayName: string;
  description: string;
}

export interface MsgSetProfileResponse {
  id: number;
}

export interface MsgUploadBeak {
  creator: string;
  fileIndex: string;
  name: string;
  creatorUsername: string;
  description: string;
  license: string;
  linkedBeaks: number[];
  tags: string[];
}

export interface MsgUploadBeakResponse {
  id: number;
}

const baseMsgSetProfile: object = {
  creator: "",
  username: "",
  displayName: "",
  description: "",
};

export const MsgSetProfile = {
  encode(message: MsgSetProfile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.displayName !== "") {
      writer.uint32(26).string(message.displayName);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetProfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetProfile } as MsgSetProfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.username = reader.string();
          break;
        case 3:
          message.displayName = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetProfile {
    const message = { ...baseMsgSetProfile } as MsgSetProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.username !== undefined && object.username !== null) {
      message.username = String(object.username);
    } else {
      message.username = "";
    }
    if (object.displayName !== undefined && object.displayName !== null) {
      message.displayName = String(object.displayName);
    } else {
      message.displayName = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    return message;
  },

  toJSON(message: MsgSetProfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.username !== undefined && (obj.username = message.username);
    message.displayName !== undefined &&
      (obj.displayName = message.displayName);
    message.description !== undefined &&
      (obj.description = message.description);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetProfile>): MsgSetProfile {
    const message = { ...baseMsgSetProfile } as MsgSetProfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.username !== undefined && object.username !== null) {
      message.username = object.username;
    } else {
      message.username = "";
    }
    if (object.displayName !== undefined && object.displayName !== null) {
      message.displayName = object.displayName;
    } else {
      message.displayName = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    return message;
  },
};

const baseMsgSetProfileResponse: object = { id: 0 };

export const MsgSetProfileResponse = {
  encode(
    message: MsgSetProfileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetProfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetProfileResponse } as MsgSetProfileResponse;
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

  fromJSON(object: any): MsgSetProfileResponse {
    const message = { ...baseMsgSetProfileResponse } as MsgSetProfileResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgSetProfileResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetProfileResponse>
  ): MsgSetProfileResponse {
    const message = { ...baseMsgSetProfileResponse } as MsgSetProfileResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUploadBeak: object = {
  creator: "",
  fileIndex: "",
  name: "",
  creatorUsername: "",
  description: "",
  license: "",
  linkedBeaks: 0,
  tags: "",
};

export const MsgUploadBeak = {
  encode(message: MsgUploadBeak, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fileIndex !== "") {
      writer.uint32(18).string(message.fileIndex);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.creatorUsername !== "") {
      writer.uint32(34).string(message.creatorUsername);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.license !== "") {
      writer.uint32(50).string(message.license);
    }
    writer.uint32(58).fork();
    for (const v of message.linkedBeaks) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.tags) {
      writer.uint32(66).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUploadBeak {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUploadBeak } as MsgUploadBeak;
    message.linkedBeaks = [];
    message.tags = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fileIndex = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.creatorUsername = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.license = reader.string();
          break;
        case 7:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.linkedBeaks.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.linkedBeaks.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 8:
          message.tags.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUploadBeak {
    const message = { ...baseMsgUploadBeak } as MsgUploadBeak;
    message.linkedBeaks = [];
    message.tags = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fileIndex !== undefined && object.fileIndex !== null) {
      message.fileIndex = String(object.fileIndex);
    } else {
      message.fileIndex = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (
      object.creatorUsername !== undefined &&
      object.creatorUsername !== null
    ) {
      message.creatorUsername = String(object.creatorUsername);
    } else {
      message.creatorUsername = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.license !== undefined && object.license !== null) {
      message.license = String(object.license);
    } else {
      message.license = "";
    }
    if (object.linkedBeaks !== undefined && object.linkedBeaks !== null) {
      for (const e of object.linkedBeaks) {
        message.linkedBeaks.push(Number(e));
      }
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgUploadBeak): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fileIndex !== undefined && (obj.fileIndex = message.fileIndex);
    message.name !== undefined && (obj.name = message.name);
    message.creatorUsername !== undefined &&
      (obj.creatorUsername = message.creatorUsername);
    message.description !== undefined &&
      (obj.description = message.description);
    message.license !== undefined && (obj.license = message.license);
    if (message.linkedBeaks) {
      obj.linkedBeaks = message.linkedBeaks.map((e) => e);
    } else {
      obj.linkedBeaks = [];
    }
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUploadBeak>): MsgUploadBeak {
    const message = { ...baseMsgUploadBeak } as MsgUploadBeak;
    message.linkedBeaks = [];
    message.tags = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fileIndex !== undefined && object.fileIndex !== null) {
      message.fileIndex = object.fileIndex;
    } else {
      message.fileIndex = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (
      object.creatorUsername !== undefined &&
      object.creatorUsername !== null
    ) {
      message.creatorUsername = object.creatorUsername;
    } else {
      message.creatorUsername = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.license !== undefined && object.license !== null) {
      message.license = object.license;
    } else {
      message.license = "";
    }
    if (object.linkedBeaks !== undefined && object.linkedBeaks !== null) {
      for (const e of object.linkedBeaks) {
        message.linkedBeaks.push(e);
      }
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(e);
      }
    }
    return message;
  },
};

const baseMsgUploadBeakResponse: object = { id: 0 };

export const MsgUploadBeakResponse = {
  encode(
    message: MsgUploadBeakResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUploadBeakResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUploadBeakResponse } as MsgUploadBeakResponse;
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

  fromJSON(object: any): MsgUploadBeakResponse {
    const message = { ...baseMsgUploadBeakResponse } as MsgUploadBeakResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgUploadBeakResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUploadBeakResponse>
  ): MsgUploadBeakResponse {
    const message = { ...baseMsgUploadBeakResponse } as MsgUploadBeakResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SetProfile(request: MsgSetProfile): Promise<MsgSetProfileResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UploadBeak(request: MsgUploadBeak): Promise<MsgUploadBeakResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SetProfile(request: MsgSetProfile): Promise<MsgSetProfileResponse> {
    const data = MsgSetProfile.encode(request).finish();
    const promise = this.rpc.request("parrots.parrots.Msg", "SetProfile", data);
    return promise.then((data) =>
      MsgSetProfileResponse.decode(new Reader(data))
    );
  }

  UploadBeak(request: MsgUploadBeak): Promise<MsgUploadBeakResponse> {
    const data = MsgUploadBeak.encode(request).finish();
    const promise = this.rpc.request("parrots.parrots.Msg", "UploadBeak", data);
    return promise.then((data) =>
      MsgUploadBeakResponse.decode(new Reader(data))
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
