package dll

import (
	"log"
	"syscall"
)
func Dll_Call(){
	// Define qual DLL será chamada
	dll := syscall.NewLazyDLL("./dll/hello_Dll.dll")

	// Procura a função na DLL pelo nome
	hello := dll.NewProc("helloWorld")

	// Chama a função da DLL
	_, _, err := hello.Call()
	if err == nil{
		log.Fatalln(err)
	}

}