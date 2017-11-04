package main

import (
	"log"
	"net"

	"github.com/dustin/go-coap"
	"fmt"
	logger "github.com/shengkehua/xlog4go"
)

func handleA(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	//logger.Info("start handlerA")
	logger.Info("Got message in handleA: path=%q: %#v from %v", m.Path(), m, a)
	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("hello world"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		logger.Info("Transmitting from A %#v", res)
		return res
	}
	return nil
}

func handleB(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	/*timeNow := time.Now().UnixNano() / 1000000
	spendTime := timeNow - startTime
	fmt.Println("speed time :", spendTime)
	*/
	logger.Info("Got message in handleB: path=%q: %#v from %v", m.Path(), m, a)
	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("good bye!"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)


		logger.Info("Transmitting from B %#v", res)
		return res
	}
	return nil
}

func main() {
	fmt.Println("start server")
	//startTime = time.Now().UnixNano() / 1000000
	mux := coap.NewServeMux()
	mux.Handle("/a", coap.FuncHandler(handleA))
	mux.Handle("/b", coap.FuncHandler(handleB))

	log.Fatal(coap.ListenAndServe("udp", ":5684", mux))
}
