package main

import (
	"log"
	"os"

	"github.com/dustin/go-coap"
	//"time"
	"strconv"
	"time"
	"fmt"
)

func main() {

	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 12345,
		Payload:   []byte("hello, world!"),
	}

	path := "/some/path"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	num, _ := strconv.Atoi(os.Args[2])
	req.SetOption(coap.ETag, "weetag")
	req.SetOption(coap.MaxAge, 3)
	req.SetPathString(path)

	c, err := coap.Dial("udp", "localhost:5684")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	//startTime := time.Now()
	for i := 0; i < num; i++ {
		rv, err := c.Send(req)
		if err != nil {
			log.Printf("%s", err)
		}
		if rv != nil {
			log.Printf("%s", rv.Payload)
		}
	}
	//endTimeCoAP := time.Now()
	//fmt.Println("coap cost: ", endTimeCoAP.Sub(startTime).Nanoseconds()/int64(time.Millisecond), "ms")

}
