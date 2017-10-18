# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: msg.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='msg.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\tmsg.proto\x1a\x19google/protobuf/any.proto\"I\n\x08Response\x12\x0c\n\x04\x43ode\x18\x01 \x01(\x11\x12\x0b\n\x03Msg\x18\x02 \x01(\t\x12\"\n\x04\x44\x61ta\x18\x03 \x01(\x0b\x32\x14.google.protobuf.Any\"1\n\x0c\x45nterRoomReq\x12\x0f\n\x07RoomNum\x18\x01 \x01(\x05\x12\x10\n\x08Password\x18\x02 \x01(\t\"$\n\x11QuickEnterRoomReq\x12\x0f\n\x07RoomCId\x18\x01 \x01(\x05\x42\x32\n\x10\x63om.baiyigame.pbZ\x02pb\xaa\x02\x19\x43om.Baiyigame.Godnet.Commb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_any__pb2.DESCRIPTOR,])




_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Code', full_name='Response.Code', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Msg', full_name='Response.Msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Data', full_name='Response.Data', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=40,
  serialized_end=113,
)


_ENTERROOMREQ = _descriptor.Descriptor(
  name='EnterRoomReq',
  full_name='EnterRoomReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='RoomNum', full_name='EnterRoomReq.RoomNum', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Password', full_name='EnterRoomReq.Password', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=115,
  serialized_end=164,
)


_QUICKENTERROOMREQ = _descriptor.Descriptor(
  name='QuickEnterRoomReq',
  full_name='QuickEnterRoomReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='RoomCId', full_name='QuickEnterRoomReq.RoomCId', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=166,
  serialized_end=202,
)

_RESPONSE.fields_by_name['Data'].message_type = google_dot_protobuf_dot_any__pb2._ANY
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
DESCRIPTOR.message_types_by_name['EnterRoomReq'] = _ENTERROOMREQ
DESCRIPTOR.message_types_by_name['QuickEnterRoomReq'] = _QUICKENTERROOMREQ
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), dict(
  DESCRIPTOR = _RESPONSE,
  __module__ = 'msg_pb2'
  # @@protoc_insertion_point(class_scope:Response)
  ))
_sym_db.RegisterMessage(Response)

EnterRoomReq = _reflection.GeneratedProtocolMessageType('EnterRoomReq', (_message.Message,), dict(
  DESCRIPTOR = _ENTERROOMREQ,
  __module__ = 'msg_pb2'
  # @@protoc_insertion_point(class_scope:EnterRoomReq)
  ))
_sym_db.RegisterMessage(EnterRoomReq)

QuickEnterRoomReq = _reflection.GeneratedProtocolMessageType('QuickEnterRoomReq', (_message.Message,), dict(
  DESCRIPTOR = _QUICKENTERROOMREQ,
  __module__ = 'msg_pb2'
  # @@protoc_insertion_point(class_scope:QuickEnterRoomReq)
  ))
_sym_db.RegisterMessage(QuickEnterRoomReq)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\020com.baiyigame.pbZ\002pb\252\002\031Com.Baiyigame.Godnet.Comm'))
# @@protoc_insertion_point(module_scope)