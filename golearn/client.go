package main

import (
    "net/rpc"
    "log"
    "fmt"
    "librpc"
)

func main() {
    serverAddress := "127.0.0.1"
    client,error := rpc.DialHTTP("tcp", serverAddress + ":9090")
    if error != nil {
        log.Fatal("DialHttp error", error)
    }

    args := librpc.Args{5, 4}
    var res librpc.Res
    var addRes int
    error = client.Call("Service.Div", &args, &res)
    if error != nil {
        log.Fatal("client call error", error)
    }
    fmt.Printf("Res.a=%d,Res.b=%d\n", res.A, res.B)

    args = librpc.Args{1,1}
    error = client.Call("Service.Add", &args, &addRes)
    if error != nil {
        log.Fatal("client call error", error)
    }
    fmt.Printf("Res=%d\n", addRes)

}
