package main

import (
    "net"
    "net/rpc"
    "net/http"
    "librpc"
    "log"
)

type Service int

func (t * Service) Div(arg * librpc.Args, res * librpc.Res) error {
    if 0 == arg.B {
        log.Fatal("divied by zero")
    }
    res.A = arg.A / arg.B
    res.B = arg.A % arg.B
    return nil
}

func (t * Service) Add(arg * librpc.Args, res * int) error {
    *res = arg.A + arg.B
    return nil
}


func main() {
    service := new(Service)
    rpc.Register(service)
    rpc.HandleHTTP()

    server,error := net.Listen("tcp", ":9090")
    if error != nil {
        log.Fatal("net.Listen error:", error)
    }
    http.Serve(server, nil)
}
