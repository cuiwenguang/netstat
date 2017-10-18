# encoding: utf-8
NET_URL="127.0.0.1"
NET_PORT=3563
WEB_URL="http://10.20.22.180:8080/api/v1"

MASTER_KEY="c13a6abc012b6ba79a8f4064d5173673"

GAME={
    "NAME": "com.baiyigame.fourking"
}

# gonet 测试接口
TEST_MENU={
    1: {
        "name": "登陆",
        "roter": 'Fourking/ReqLogin',
        "data": {
            "SessionToken": "xxxx",
            "BundleIdentifier":  GAME["NAME"]
        },
        'pb': ['pb.login_msg_pb2', 'LoginRequest'],
        "msg_type": 0
    }
}
