import login_msg_pb2

model = login_msg_pb2.LoginRequest()
model.SessionToken = "xxxxx"
model.BundleIdentifier = "ccsdfasfasdf"


data = model.SerializeToString()
print data