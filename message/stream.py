from __future__ import absolute_import, division, print_function, with_statement

import struct

class Stream(object) :

    def __init__(self, data = "") :
        self.index = 0
        self.data = data
        self.size = len(self.data)


    def tell(self) :
        return self.index


    def seek(self, seek) :
        seek = max(0, seek)
        self.index = min(self.size, seek)


    def read(self, size = None) :
        if self.size <= self.index :
            return ''
        if size is None :
            size = self.size - self.index
        start = self.index
        end = min(self.size, self.index + size)
        self.index = end
        return (self.data[start : end])


    def write(self, data) :
        self.data += data
        self.size = len(self.data)


    def getvalue(self) :
        return self.data