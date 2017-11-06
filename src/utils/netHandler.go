package utils

import (
	"net"
	logger "github.com/shengkehua/xlog4go"
)

func BroadCastServer() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Local: <%s>", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			logger.Error("error during read: %s", err)
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func BroadCastClient() {
	ip := net.ParseIP("10.141.127.255")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		logger.Error(err.Error())
	}
	n, err := conn.WriteToUDP([]byte("hello"), dstAddr)
	if err != nil {
		logger.Error(err.Error())
	}
	data := make([]byte, 1024)
	n, _, err = conn.ReadFromUDP(data)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("read %s form <%s>", data[:n], conn.RemoteAddr())
}

func MultiCastServer() {
	addr, err := net.ResolveUDPAddr("udp", "224.0.0.250:9981")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Local: <%s>", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			logger.Error("Error during read: %s", err)
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func MultiCastClient() {
	ip := net.ParseIP("224.0.0.250")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer conn.Close()

	data := make([]byte, 1024)
	for {
		conn.Write([]byte("hello"))
		logger.Info("<%s>", conn.RemoteAddr())
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			logger.Error("Error during read: %s", err)
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
	}
}

func SimpleServer() {
	listener, err:= net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5555})
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Local: <%s>", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			logger.Error(err.Error())
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func SimpleClient() {
	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 5555}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Error(err.Error())
	}
	defer conn.Close()
	conn.Write([]byte("hello"))
	logger.Info("<%s>", conn.RemoteAddr())
}
