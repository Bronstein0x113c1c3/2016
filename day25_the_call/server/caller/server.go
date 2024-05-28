package caller

import (
	"fmt"
	"log"
	pb "server/protobuf"
	"sync"
	"sync/atomic"
)

// var ListOfClient sync.Map

type Caller struct {
	host string
	port int
	// input *in.Input
	pb.UnimplementedTheCallServer
	ListOfClient map[uint32]chan []byte
	// ListOfClient *sync.Map
	mutex        *sync.RWMutex
	ChangeSignal chan struct{}
	counter      uint32
	// receiver     chan []byte
	// Done           chan struct{}
}

func New(host string, port int /* input *in.Input*/) *Caller {
	return &Caller{
		host: host,
		port: port,
		// input:          input,
		ListOfClient: make(map[uint32]chan []byte),
		mutex:        &sync.RWMutex{},
		ChangeSignal: make(chan struct{}),
		// Done:           make(chan struct{}),
		counter: 0,
	}
}

// func (s *Caller) LoadAChan(i uint) chan []byte {
// 	c, _ := s.ListOfClient.Load(i)
// 	channel := c.(chan []byte)
// 	return channel
// }

func (s *Caller) GetAmountOfChannel() uint32 {
	return s.counter
}
func (s *Caller) add() (uint32, chan []byte) {
	// s.mutex.Lock()
	s.ChangeSignal <- struct{}{}
	// s.ListOfClient = append(s.ListOfClient, make(chan []byte))
	index := atomic.AddUint32(&s.counter, 1)
	channel := make(chan []byte)
	s.mutex.Lock()
	s.ListOfClient[index] = channel
	s.mutex.Unlock()
	// i := len(s.ListOfClient) - 1
	// s.mutex.Unlock()
	log.Printf("Channel %v is created!!!", index)
	return index, channel
}
func (s *Caller) delete_chan(i uint32, closed bool) {
	// s.mutex.Lock()
	if !closed {
		s.ChangeSignal <- struct{}{}
		s.mutex.RLock()
		channel, _ := s.ListOfClient[i]
		s.mutex.RUnlock()

		close(channel)
	}
	s.mutex.Lock()
	delete(s.ListOfClient, i)
	s.mutex.Unlock()
	// s.ListOfClient = slices.Delete(s.ListOfClient, i, i+1)
	// s.mutex.Unlock()
	// s.Done <- struct{}{}
	// x := s.counter
	_ = atomic.AddUint32(&s.counter, ^uint32(0))
	log.Printf("Channel %v is deleted from the reservation!!!", i)
}

func (s *Caller) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}
func (s *Caller) Calling(caller pb.TheCall_CallingServer) error {
	// reader := bufio.NewReader(s.input.GetStream())
	index, channel := s.add()
	// defer log.Printf("%v is completely closed \n", index)
	// channel := s.LoadAChan(index)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer s.delete_chan(index, true)
		defer log.Printf("%v is closed \n", index)
		defer wg.Done()

		for {
			data, ok := <-channel
			if !ok {
				log.Printf("%v is closed \n", index)
				return
			}
			caller.Send(&pb.Server_MSGSound{Sound: data})
		}
	}()
	wg.Wait()
	log.Printf("%v is completely closed \n", index)
	return nil
}
