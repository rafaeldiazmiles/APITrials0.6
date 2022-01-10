package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Algo de fill un reply pointer para mandar la data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// Crear un RPC server nuevo
	timeserver := new(TimeServer)
	// Registrar el famoso rpc server que hicimo
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// y ahora vamos a escuchar requests.. en el port 1234..
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error;", err)
	}
	http.Serve(l, nil)

}
