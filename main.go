package main

import "C"
import (
	"encoding/json"
	"fmt"
	"main/pessoa"
)

//export hello
func hello() {
	fmt.Println("PRINT HELLO WORLD")
}

//export returnCStringMyName
func returnCStringMyName(namePtr *C.char) *C.char {
	name := C.GoString(namePtr)

	return C.CString(name)
}

//export sayMyName
func sayMyName(namePtr *C.char) {
	name := C.GoString(namePtr)
	fmt.Println(name)
}

//export getPessoaObject
func getPessoaObject(nome *C.char, idade *C.char) *C.char {

	goNome := C.GoString(nome)
	goIdade := C.GoString(idade)

	pessoa := pessoa.Pessoa{
		Nome:  goNome,
		Idade: goIdade,
	}

	v, _ := json.Marshal(pessoa)

	return C.CString(string(v))
}

//export fromJSON
func fromJSON(documentPtr *C.char) {

	documentString := C.GoString(documentPtr)

	jsonDocument := map[string]string{}

	json.Unmarshal([]byte(documentString), &jsonDocument)

	pessoa := pessoa.Pessoa{
		Nome:  jsonDocument["nome"],
		Idade: jsonDocument["idade"],
	}

	fmt.Println("PRINT FROM JSON", pessoa)
}

func main() {}
