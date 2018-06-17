// Package kongo comes with an a simple API for communicating with Kong Admin.
//
// Example Usage
//
// The following is a simple example using the api:
//    import (
//      "fmt"
//      "github.com/fabiorphp/kongo"
//      "log"
//    )
//
//    func main() {
//      k, err := kongo.New(nil, "http://127.0.0.1:8001")
//
//      if err != nil {
//        log.Fatal(err)
//      }
//
//      status, _, err := k.Node.Status()
//
//      if err != nil {
//        log.Fatal(err)
//      }
//
//      fmt.Printf("Requests: %d\n", status.Server.TotalRequests)
//    }
//
// Please look at the [Kong](https://getkong.org/docs/) docs for more information about the Rest-API.
package kongo
