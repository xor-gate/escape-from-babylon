import ctypes

lib = ctypes.cdll.LoadLibrary('library.dll')
main = library.executeMain
main()
