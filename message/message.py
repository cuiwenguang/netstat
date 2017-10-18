#!/usr/bin/env python
#
# Copyright 2016 leenjewel
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Message encode or decode body of pomelo protocol package

Pomelo protocol package :

+++++++++++++++++++++++++++++++++++++++
+ type +    length    +      body     +
+++++++++++++++++++++++++++++++++++++++
 1 bytes    3 bytes      length bytes

Body of package :

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+  flag  +  message id  +    route    +            data              +
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 byte     0 ~ 5 bytes   0 ~ 256 bytes

Flag of message :

                   1 byte
++++++++++++++++++++++++++++++++++++++++++
+          +  message type  +  is route  +
++++++++++++++++++++++++++++++++++++++++++
   4 bits          3 bits         1 bit

Message Types :
             flag
request :  |----000-| message id | route |
notify  :  |----001-| route |
response:  |----010-| message id |
push    :  |----011-| route |

Route compress :

flag              route
|-------0|        | length (1 byte) | utf8 string |
|-------1|        | route code (2 bytes, big end) |
"""

from __future__ import absolute_import, division, print_function, with_statement

import json
from .stream import Stream
from .protobuf import *

class Message(object) :
    """Encode and decode body of promelo protocol package
    """

    MSG_TYPE_REQUEST  = 0
    MSG_TYPE_NOTIFY   = 1
    MSG_TYPE_RESPONSE = 2
    MSG_TYPE_PUSH     = 3

    MSG_FLAG_BYTES       = 1
    MSG_ROUTE_LEN_BYTES  = 1
    MSG_ROUTE_CODE_BYTES = 2

    def __init__(self, msg_type, msg_id = 0, route = "", body=None) :
        self.msg_type = msg_type
        self.msg_id = msg_id
        self.route = route
        self.body = body

    def has_id(self) :
        return (Message.MSG_TYPE_REQUEST == self.msg_type or \
                Message.MSG_TYPE_RESPONSE == self.msg_type)


    def has_route(self) :
        return (Message.MSG_TYPE_RESPONSE != self.msg_type)


    def id_len(self) :
        if not self.has_id() :
            return 0
        ret = 1
        msg_id = self.msg_id >> 7
        while (msg_id > 0) :
            ret += 1
            msg_id >>= 7
        return ret


    def route_len(self) :
        if not self.has_route() :
            return 0
        return len(self.route)

    def encode_flag(self, stream, compress_route=True):
        if compress_route:
            stream.write(struct.pack('B', (self.msg_type << 1) | 0x1))
        else:
            stream.write(struct.pack('B', (self.msg_type << 1) | 0x0))

    def encode_id(self, stream):
        assert self.has_id(), "Message type %d don't has id" % (self.msg_type)
        msg_id = self.msg_id
        while True:
            tmp = msg_id & 0x7F
            nxt = msg_id >> 7
            if nxt != 0:
                tmp += 128
            stream.write(struct.pack('B', tmp))
            msg_id = nxt
            if msg_id == 0:
                break

    def encode(self) :
        stream = Stream()
        self.encode_flag(stream, False)
        if self.has_id() :
            self.encode_id(stream)
        if self.has_route() :
            route_len = self.route_len()
            stream.write(struct.pack('B', route_len))
            stream.write(self.route)
        stream.write(self.body)
        return stream.getvalue()


    @classmethod
    def decode(cls, stream) :
        route_size = struct.unpack("B", stream.read(1))[0]
        return protobuf_decode_string(stream, route_size)
