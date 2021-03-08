package utils

import (
	"fmt"
)

const (
	KeepClear     = 32
	DefaultLength = 128
	ExpandLength  = 128
)

var (
	NewLine = []byte("\n")
)

type LocalLog struct {
	buffer []byte
	index  int
}

func NewLocalLog() *LocalLog {
	return &LocalLog{make([]byte, DefaultLength), 0}
}

func (l *LocalLog) Printf(f string, i ...interface{}) {
	l.write([]byte(fmt.Sprintf(f, i...)))
}

func (l *LocalLog) Print(i ...interface{}) {
	l.write([]byte(fmt.Sprint(i...)))
}

func (l *LocalLog) Println(i ...interface{}) {
	if len(i) == 0 {
		l.write(NewLine)
		return
	}

	for _, v := range i {
		l.Print(v)
		l.write(NewLine)
	}
}

func (l *LocalLog) String() string {
	return string(l.buffer)
}

func (l *LocalLog) write(buffer []byte) {
	dataLen := len(buffer)
	if len(l.buffer)-l.index <= dataLen+KeepClear {
		l.buffer = append(l.buffer, make([]byte, dataLen+ExpandLength)...)
	}

	for i, v := range buffer {
		l.buffer[l.index+i] = v
	}

	l.index += dataLen
}
