# encoding: utf-8
#! /usr/bin/env python
#
# GUI module generated by PAGE version 4.9
# In conjunction with Tcl version 8.6
#    Oct 11, 2017 02:44:09 PM
import sys, json
import threading
import select
from tkinter import *
import tkinter.ttk as ttk

import config
import web
import db
import gonet

class Frame:
    def __init__(self, top=None):
        '''This class configures and populates the toplevel window.
           top is the toplevel containing window.'''
        _bgcolor = '#d9d9d9'  # X11 color: 'gray85'
        _fgcolor = '#000000'  # X11 color: 'black'
        _compcolor = '#d9d9d9' # X11 color: 'gray85'
        _ana1color = '#d9d9d9' # X11 color: 'gray85'
        _ana2color = '#d9d9d9' # X11 color: 'gray85'
        self.style = ttk.Style()
        if sys.platform == "win32":
            self.style.theme_use('winnative')
        self.style.configure('.',background=_bgcolor)
        self.style.configure('.',foreground=_fgcolor)
        self.style.configure('.',font="TkDefaultFont")
        self.style.map('.',background=
            [('selected', _compcolor), ('active',_ana2color)])

        top.geometry("890x605+526+125")
        top.title("netstat")
        top.configure(background="#d9d9d9")
        top.configure(highlightbackground="#d9d9d9")
        top.configure(highlightcolor="black")

        self.init_test_menu(top)

        self.LstHistory = Listbox(top)
        self.LstHistory.place(relx=0.0, rely=0.02, relheight=1.02, relwidth=0.26)

        self.LstHistory.configure(background="white")
        self.LstHistory.configure(disabledforeground="#a3a3a3")
        self.LstHistory.configure(font="TkFixedFont")
        self.LstHistory.configure(foreground="#000000")
        self.LstHistory.configure(highlightbackground="#d9d9d9")
        self.LstHistory.configure(highlightcolor="black")
        self.LstHistory.configure(selectbackground="#c4c4c4")
        self.LstHistory.configure(selectforeground="black")
        self.LstHistory.configure(width=234)

        self.txtResult = Text(top)
        self.txtResult.place(relx=0.27, rely=0.28, relheight=0.75, relwidth=0.72)

        self.txtResult.configure(background="white")
        self.txtResult.configure(cursor="fleur")
        self.txtResult.configure(font="TkTextFont")
        self.txtResult.configure(foreground="black")
        self.txtResult.configure(highlightbackground="#d9d9d9")
        self.txtResult.configure(highlightcolor="black")
        self.txtResult.configure(insertbackground="black")
        self.txtResult.configure(selectbackground="#c4c4c4")
        self.txtResult.configure(selectforeground="black")
        self.txtResult.configure(takefocus="0")
        self.txtResult.configure(undo="1")
        self.txtResult.configure(width=644)
        self.txtResult.configure(wrap=WORD)

        self.TLabel2 = ttk.Label(top)
        self.TLabel2.place(relx=0.28, rely=0.23, height=21, width=52)
        self.TLabel2.configure(background="#d9d9d9")
        self.TLabel2.configure(foreground="#000000")
        self.TLabel2.configure(relief=FLAT)
        self.TLabel2.configure(takefocus="0")
        self.TLabel2.configure(text='''监听返回结果''')

        self.Label2 = Label(top)
        self.Label2.place(relx=0.28, rely=0.12, height=23, width=30)
        self.Label2.configure(activebackground="#f9f9f9")
        self.Label2.configure(activeforeground="black")
        self.Label2.configure(background="#d9d9d9")
        self.Label2.configure(disabledforeground="#a3a3a3")
        self.Label2.configure(foreground="#000000")
        self.Label2.configure(highlightbackground="#d9d9d9")
        self.Label2.configure(highlightcolor="black")
        self.Label2.configure(text='''参数''')

        self.TEntry1 = ttk.Entry(top)
        self.TEntry1.place(relx=0.34, rely=0.12, relheight=0.04, relwidth=0.60)
        self.TEntry1.configure(width=406)
        self.TEntry1.configure(takefocus="")
        self.TEntry1.configure(cursor="ibeam")

        self.Label3 = Label(top)
        self.Label3.place(relx=0.28, rely=0.03, height=23, width=30)
        self.Label3.configure(activebackground="#f9f9f9")
        self.Label3.configure(activeforeground="black")
        self.Label3.configure(background="#d9d9d9")
        self.Label3.configure(disabledforeground="#a3a3a3")
        self.Label3.configure(foreground="#000000")
        self.Label3.configure(highlightbackground="#d9d9d9")
        self.Label3.configure(highlightcolor="black")
        self.Label3.configure(text='''用户''')

        self.TEntry2 = ttk.Entry(top)
        self.TEntry2.place(relx=0.34, rely=0.03, relheight=0.04, relwidth=0.46)
        self.TEntry2.configure(takefocus="")
        self.TEntry2.configure(cursor="ibeam")

        self.btnLogin = ttk.Button(top)
        self.btnLogin.place(relx=0.84, rely=0.03, height=27, width=87)
        self.btnLogin.configure(takefocus="")
        self.btnLogin.configure(text='''登录''')
        self.btnLogin.configure(command=self.on_login)

        self.user = None
        try:
            self.client = gonet.conn_gonet()
        except:
            print('socket 网络链接出错，请检查配置是否正确，服务器是否正常启动')
        self.on_read()

    def init_test_menu(self, top):
        self.menubar = Menu(top, font="TkMenuFont", bg='#d9d9d9', fg='#000000')
        top.configure(menu=self.menubar)

        self.menubar.add_command(
            activebackground="#d8d8d8",
            activeforeground="#000000",
            background="#d9d9d9",
            font="TkMenuFont",
            foreground="#000000",
            label="用户")

        self.test_menu = Menu(self.menubar, tearoff=0)
        for k ,v in config.TEST_MENU.items():
            self.test_menu.add_command(
                label=v['name'],
                command=lambda y=k: self.on_menu_click(y))
        self.menubar.add_cascade(label="测试", menu=self.test_menu)

        self.menubar.add_command(
            activebackground="#d8d8d8",
            activeforeground="#000000",
            background="#d9d9d9",
            font="TkMenuFont",
            foreground="#000000",
            label="配置",
            command=self.on_conf)

    def on_conf(self):
        import tkMessageBox
        tkMessageBox.showinfo('提示', '暂为实现，请打开项目根目录，找到config手工修改')

    def on_menu_click(self, index):
        current = config.TEST_MENU[index]
        try:
            data = json.loads(self.TEntry1.get())
        except:
            data = current['data'] # 默认值
        gonet.request(self.client, current['msg_type'],current['roter'], data, current['pb'])
        self.LstHistory.insert(0, '%s-%s' %(current['roter'], current['name']))

    def on_login(self):
        if self.TEntry2['state'].string == 'readonly':
            self.TEntry2['state']='normal'
            self.btnLogin.configure(text='''登录''')
            config.user = None
            return
        code = self.TEntry2.get()
        if len(code)==0: return
        data = web.create_gust(self.TEntry2.get())
        if data["code"] == 2000:
            ret = gonet.login(self.client, data['data']['Data']['Token'])
            db.create_user(data['data']['AccountID'], data['data']['Data']['Token'], code)
            self.user = data['data']
            self.TEntry2['state'] = 'readonly'
            self.btnLogin.configure(text="注销")


    def on_read(self):
        def read_data():
            while True:
                print ">>> begin accept data:"
                data = self.client.recv(512)
                # todo 获取的数据反序列化，并显示到txtresult窗口
                print data
        threading.Thread(target=read_data).start()


