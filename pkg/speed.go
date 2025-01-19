package pkg

import "syscall"

type BufferSize struct {
}

func NewBufferSize() *BufferSize {
	return &BufferSize{}
}
func (s *BufferSize) getMaxBufferSize() int {
	return syscall.Getpagesize() * 64
}
