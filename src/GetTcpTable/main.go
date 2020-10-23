package main

/*
#cgo windows LDFLAGS:-lWS2_32 -liphlpapi
void cRun();
 */
import "C"

func main(){
	C.cRun()
}
