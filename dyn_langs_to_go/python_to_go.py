from sys import platform as _platform
from ctypes import cdll

def lib_extension():
    if _platform == 'darwin':
        return 'dylib'
    elif _platform == 'win32':
        return 'dll'
    else:
        return 'so'

lib = cdll.LoadLibrary('go_adder/libadder.' + lib_extension())

print "Python say: about to call Go ..."
total = lib.Add(40, 2)
print "Python says: result is " + str(total)
