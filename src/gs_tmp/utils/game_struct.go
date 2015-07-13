package utils

import (
	"net"
)

type Buffer struct {
	Cur_p int
	Data  []byte
}

type Client struct {
	Client  *net.TCPConn
	Handler chan *Msg
}

type Msg struct {
	PlayerId int
	Category int
	Buff     *Buffer
}
