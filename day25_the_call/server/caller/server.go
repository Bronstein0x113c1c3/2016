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
	Mutex        *sync.RWMutex
	ChangeSignal chan uint32
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
		Mutex:        &sync.RWMutex{},
		ChangeSignal: make(chan uint32),
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
	// s.Mutex.Lock()
	s.ChangeSignal <- 0
	// s.ListOfClient = append(s.ListOfClient, make(chan []byte))
	index := atomic.AddUint32(&s.counter, 1)
	channel := make(chan []byte)
	s.Mutex.Lock()
	s.ListOfClient[index] = channel
	s.Mutex.Unlock()
	// i := len(s.ListOfClient) - 1
	// s.Mutex.Unlock()
	log.Printf("Channel %v is created!!!", index)
	return index, channel
}
func (s *Caller) delete_chan(i uint32, closed bool) {
	// s.Mutex.Lock()
	// defer log.Printf("There is %v\n", s.counter)
	if !closed {
		s.ChangeSignal <- i
		// s.Mutex.RLock()
		// channel := s.ListOfClient[i]
		// s.Mutex.RUnlock()
		// s.Mutex.Lock()
		// delete(s.ListOfClient, i)
		// s.Mutex.Unlock()
		// close(channel)
		_ = atomic.AddUint32(&s.counter, ^uint32(0))
		log.Printf("Channel %v is deleted from the reservation!!!", i)
		return
	}
	s.Mutex.Lock()
	delete(s.ListOfClient, i)
	s.Mutex.Unlock()
	// s.ListOfClient = slices.Delete(s.ListOfClient, i, i+1)
	// s.Mutex.Unlock()
	// s.Done <- struct{}{}
	// x := s.counter
	_ = atomic.AddUint32(&s.counter, ^uint32(0))
	log.Printf("Channel %v is deleted from the reservation!!!", i)
	return

}

func (s *Caller) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}
func (s *Caller) Calling(caller pb.TheCall_CallingServer) error {
	// reader := bufio.NewReader(s.input.GetStream())
	index, channel := s.add()
	defer log.Printf("%v is completely closed \n", index)
	// defer log.Printf("%v is completely closed \n", index)
	// channel := s.LoadAChan(index)
	// wg := &sync.WaitGroup{}
	// wg.Add(1)
	signal_1 := make(chan struct{})
	signal_2 := make(chan struct{})
	// defer close(signal)
	go func() {
		defer func(signal chan struct{}) {
			close(signal)
		}(signal_1)
		// defer s.delete_chan(index, true)
		// defer log.Printf("%v is closed \n", index)
		// defer wg.Done()

		for {
			data, ok := <-channel
			if !ok {
				log.Printf("%v is forcing closed \n", index)
				return
			}
			caller.Send(&pb.Server_MSGSound{Sound: data})
		}
	}()
	go func() {
		defer close(signal_2)
		for {
			data, err := caller.Recv()
			if err != nil {
				return
			}
			x := data.Sound
			if string(x) == "Goodbye!!" {
				return
			}
		}

	}()
	for {
		select {
		case <-signal_1:
			log.Println("Forcing closed....")
			s.delete_chan(index, true)
			return nil
		case <-signal_2:
			log.Println("Disconnected....")
			s.delete_chan(index, false)
			return nil
		}

	}

	// wg.Wait()

}
