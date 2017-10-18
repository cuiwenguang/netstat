import os, sys
import sqlite3

if os.name == 'nt':
    conn_str=sys.path[0] + "\\test.db"
else:
    conn_str = sys.path[0] + "/test.db"

def create_user(account_id, token, code):
    if update_token(code, token)>0: return

    sql="insert into users(mobile_code, account_id, token) VALUES (?, ?, ?)"
    con = sqlite3.connect(conn_str)
    con.execute(sql, (code, account_id, token))
    con.commit()
    con.close()

def update_token(code, token):
    sql="update users set token=? where mobile_code=?"
    con = sqlite3.connect(conn_str)
    rows = con.execute(sql,(token, code))
    return rows.rowcount

def create_history(addr, params):
    pass