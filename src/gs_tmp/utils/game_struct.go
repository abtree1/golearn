package utils

type Buffer struct {
	Cur_p int
	Data  []byte
}

type Msg struct {
	Category int32
	Buff     *Buffer
}

type ObMsg struct {
	PlayerId int
	Buff     *Buffer
	Handler  chan<- *ObMsg
}
