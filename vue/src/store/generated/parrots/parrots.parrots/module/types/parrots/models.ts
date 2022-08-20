/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "parrots.parrots";

export interface Profile {
  id: number;
  creator: string;
  username: string;
  display_name: string;
  description: string;
  respectedBeaks: string[];
}

export interface Beak {
  id: number;
  creator: string;
  file_index: string;
  name: string;
  creator_username: string;
  description: string;
  license: string;
  tags: string[];
  linked_beaks: number[];
}

const baseProfile: object = {
  id: 0,
  creator: "",
  username: "",
  display_name: "",
  description: "",
  respectedBeaks: "",
};

export const Profile = {
  encode(message: Profile, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.username !== "") {
      writer.uint32(26).string(message.username);
    }
    if (message.display_name !== "") {
      writer.uint32(34).string(message.display_name);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    for (const v of message.respectedBeaks) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Profile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProfile } as Profile;
    message.respectedBeaks = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.username = reader.string();
          break;
        case 4:
          message.display_name = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.respectedBeaks.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Profile {
    const message = { ...baseProfile } as Profile;
    message.respectedBeaks = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
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
    if (object.display_name !== undefined && object.display_name !== null) {
      message.display_name = String(object.display_name);
    } else {
      message.display_name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.respectedBeaks !== undefined && object.respectedBeaks !== null) {
      for (const e of object.respectedBeaks) {
        message.respectedBeaks.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Profile): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.username !== undefined && (obj.username = message.username);
    message.display_name !== undefined &&
      (obj.display_name = message.display_name);
    message.description !== undefined &&
      (obj.description = message.description);
    if (message.respectedBeaks) {
      obj.respectedBeaks = message.respectedBeaks.map((e) => e);
    } else {
      obj.respectedBeaks = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Profile>): Profile {
    const message = { ...baseProfile } as Profile;
    message.respectedBeaks = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
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
    if (object.display_name !== undefined && object.display_name !== null) {
      message.display_name = object.display_name;
    } else {
      message.display_name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.respectedBeaks !== undefined && object.respectedBeaks !== null) {
      for (const e of object.respectedBeaks) {
        message.respectedBeaks.push(e);
      }
    }
    return message;
  },
};

const baseBeak: object = {
  id: 0,
  creator: "",
  file_index: "",
  name: "",
  creator_username: "",
  description: "",
  license: "",
  tags: "",
  linked_beaks: 0,
};

export const Beak = {
  encode(message: Beak, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.file_index !== "") {
      writer.uint32(26).string(message.file_index);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.creator_username !== "") {
      writer.uint32(42).string(message.creator_username);
    }
    if (message.description !== "") {
      writer.uint32(50).string(message.description);
    }
    if (message.license !== "") {
      writer.uint32(58).string(message.license);
    }
    for (const v of message.tags) {
      writer.uint32(66).string(v!);
    }
    writer.uint32(74).fork();
    for (const v of message.linked_beaks) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Beak {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBeak } as Beak;
    message.tags = [];
    message.linked_beaks = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.file_index = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.creator_username = reader.string();
          break;
        case 6:
          message.description = reader.string();
          break;
        case 7:
          message.license = reader.string();
          break;
        case 8:
          message.tags.push(reader.string());
          break;
        case 9:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.linked_beaks.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.linked_beaks.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Beak {
    const message = { ...baseBeak } as Beak;
    message.tags = [];
    message.linked_beaks = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.file_index !== undefined && object.file_index !== null) {
      message.file_index = String(object.file_index);
    } else {
      message.file_index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (
      object.creator_username !== undefined &&
      object.creator_username !== null
    ) {
      message.creator_username = String(object.creator_username);
    } else {
      message.creator_username = "";
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
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(String(e));
      }
    }
    if (object.linked_beaks !== undefined && object.linked_beaks !== null) {
      for (const e of object.linked_beaks) {
        message.linked_beaks.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: Beak): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.file_index !== undefined && (obj.file_index = message.file_index);
    message.name !== undefined && (obj.name = message.name);
    message.creator_username !== undefined &&
      (obj.creator_username = message.creator_username);
    message.description !== undefined &&
      (obj.description = message.description);
    message.license !== undefined && (obj.license = message.license);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    if (message.linked_beaks) {
      obj.linked_beaks = message.linked_beaks.map((e) => e);
    } else {
      obj.linked_beaks = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Beak>): Beak {
    const message = { ...baseBeak } as Beak;
    message.tags = [];
    message.linked_beaks = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.file_index !== undefined && object.file_index !== null) {
      message.file_index = object.file_index;
    } else {
      message.file_index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (
      object.creator_username !== undefined &&
      object.creator_username !== null
    ) {
      message.creator_username = object.creator_username;
    } else {
      message.creator_username = "";
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
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(e);
      }
    }
    if (object.linked_beaks !== undefined && object.linked_beaks !== null) {
      for (const e of object.linked_beaks) {
        message.linked_beaks.push(e);
      }
    }
    return message;
  },
};

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