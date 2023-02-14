import ctypes
import json

# PRINT HELLO WORLD
lib = ctypes.cdll.LoadLibrary('./lib.so')
hello = lib.hello()

# SEND JSON
document = {
    "nome": "Tom Oliveira",
    "idade": "20"
}
lib.fromJSON(json.dumps(document).encode('utf-8'))

# PRINT MY NAME WITH PARAMETER
lib.sayMyName("PRINT MY NAME WITH PARAMETER: Tom Oliveira".encode("utf-8"))

# PRINT MY NAME WITH C POINTER
nameFromCGoString = lib.returnCStringMyName
nameFromCGoString.restype = ctypes.c_void_p
out = nameFromCGoString("Tom Oliveira".encode("utf-8"))
pointer = ctypes.string_at(out)
outVal = pointer.decode("utf-8")
print("PRINT MY NAME WITH C POINTER ", outVal)


# GET PERSON OBJ WITH POINTER
obj = lib.getPessoaObject
obj.restype = ctypes.c_void_p
out = obj(
    "Tom Oliveira".encode("utf-8"),
    "21".encode("utf-8")
)
pointer = ctypes.string_at(out)
outVal = pointer.decode("utf-8")
print("PRINT GET PERSON OBJ WITH POINTER ", outVal)
