package goutil
/*
import "reflect"

type ITFC interface{}
type callInfo struct {
	chReq chan ITFC
	chRes chan ITFC
}
type CallFunc func(p ITFC) (r ITFC)
type GenServer struct {
	m map[CallFunc]callInfo
}


func (this *GenServer) RegCall(func(p ITFC) (r ITFC)) {
	tp := reflect.TypeOf(p)
	chp := make(chan tp)
	tr := reflect.TypeOf(r)
	chr := make(chan tr)
	this.m[chp] = f
}

func (this *GenServer) Call(f func(p ITFC) (r ITFC)) {

}

func (this *GenServer) Start(n, loop func()) {
	go this.loop()
}
func (*GenServer) loop() {
	for {

	}
}*/
