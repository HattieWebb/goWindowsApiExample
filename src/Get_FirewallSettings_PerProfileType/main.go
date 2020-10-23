package main

/*
#cgo windows LDFLAGS:-lole32 -loleaut32
int cRun();
 */
import "C"

func main(){
	C.cRun()
}
