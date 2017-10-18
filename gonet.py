import socket
import config
from message.message import Message
import pb.login_msg_pb2
from message.protocol import Protocol


def conn_gonet():
    tcpclient = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    ser_address = (config.NET_URL, config.NET_PORT)
    tcpclient.connect(ser_address)
    return tcpclient


def login(tcpclient,token):
    model = pb.login_msg_pb2.LoginRequest()
    model.SessionToken=token
    model.BundleIdentifier = config.GAME["NAME"]
    req_data = model.SerializeToString()
    msg = Message(Message.MSG_TYPE_REQUEST, route='Login/ReqLogin', body=req_data)
    protocol_pack = Protocol(Protocol.PROTO_TYPE_DATA, msg.encode())
    return tcpclient.send(protocol_pack.pack())

def request(tcpclient,type, router, data, pb):
    module = __import__(pb[0], globals(), locals(), [pb[1]])
    c = getattr(module, pb[1])()
    for k, v in data.items():
        setattr(c, k, v)

    req_data = c.SerializeToString()
    msg = Message(type, route=router, body=req_data)
    protocol_pack = Protocol(Protocol.PROTO_TYPE_DATA, msg.encode())
    return tcpclient.send(protocol_pack.pack())

