
import sys

try:
    from Tkinter import *
except ImportError:
    from tkinter import *

try:
    import ttk
    py3 = 0
except ImportError:
    import tkinter.ttk as ttk
    py3 = 1

import  frame
import config


def destroy_window():
    # Function which closes the window.
    global top_level
    top_level.destroy()
    top_level = None

def vp_start_gui():
    config.root = Tk()
    top = frame.Frame(config.root)
    config.root.mainloop()


if __name__ == '__main__':
    vp_start_gui()

