package serverimpl

import (
	"errors"
	"fmt"
	"log"
	pb "serv/protobuf"
	"sync"
)

type Server struct {
	host   string
	port   int
	Input  chan Chunk
	Output [10]chan Chunk

	lock         *sync.Mutex
	ChangeSignal chan int

	pb.UnimplementedCallingServer
}

func New(host string, port int, input chan Chunk) *Server {
	return &Server{
		host:         host,
		port:         port,
		Input:        input,
		lock:         &sync.Mutex{},
		ChangeSignal: make(chan int),
		// InputChangeSignal: make(chan int),
	}
}
func (s Server) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}
func (s *Server) Add() (int, error) {
	s.ChangeSignal <- 0

	index := -1
	s.lock.Lock()
	for i := range 10 {
		if s.Output[i] == nil {
			s.Output[i] = make(chan Chunk)

			index = i
			break
		}
	}
	s.lock.Unlock()
	err := errors.New("the channel list is full, cannot add")
	if index == -1 {
		return -1, err
	} else {
		return index, nil
	}
}
func (s *Server) delete_chan(i int, closed bool) {
	// s.InputChangeSignal <- -(i + 1)
	log.Printf("Delete for %v is called \n", i)
	defer log.Printf("%v is deleted!!! \n", i)
	if !closed {
		s.ChangeSignal <- -(i + 1)
		log.Printf("Deletion requested!!!")
		return
	}
	log.Printf("Released %v!!!", i)
	return
}

// func
