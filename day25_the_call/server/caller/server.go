package caller

import (
	"fmt"
	"log"
	pb "server/protobuf"
	"sync"
	"sync/atomic"

	"github.com/gordonklaus/portaudio"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Caller struct {
	host string
	port int
	pb.UnimplementedTheCallServer
	ListOfClient map[uint32]chan []byte
	Mutex        *sync.RWMutex
	ChangeSignal chan uint32
	counter      uint32
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

func (s *Caller) GetAmountOfChannel() uint32 {
	return s.counter
}
func (s *Caller) add() (uint32, chan []byte) {
	// s.Mutex.Lock()
	s.ChangeSignal <- 0
	// s.ListOfClient = append(s.ListOfClient, make(chan []byte))
	// index := atomic.AddUint32(&s.counter, 1)
	var index uint32
	channel := make(chan []byte)
	s.Mutex.Lock()
	for i := 1; i <= (1<<32)-1; i++ {
		if _, found := s.ListOfClient[uint32(i)]; !found {
			index = uint32(i)
			s.ListOfClient[index] = channel
			break
		}
	}
	// s.ListOfClient[index] = channel
	s.Mutex.Unlock()
	// i := len(s.ListOfClient) - 1
	// s.Mutex.Unlock()
	if index != 0 {
		_ = atomic.AddUint32(&s.counter, 1)
		log.Printf("Channel %v is created!!!", index)
		return index, channel
	} else {
		return 0, nil
	}
}
func (s *Caller) delete_chan(i uint32, closed bool) {
	// s.Mutex.Lock()
	// defer log.Printf("There is %v\n", s.counter)
	if !closed {
		s.ChangeSignal <- i

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

	_ = atomic.AddUint32(&s.counter, ^uint32(0))
	log.Printf("Channel %v is deleted from the reservation!!!", i)
	return

}

func (s *Caller) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}

// func processing_sound()
func (s *Caller) Calling(caller pb.TheCall_CallingServer) error {
	// reader := bufio.NewReader(s.input.GetStream())

	index, channel := s.add()

	if index == 0 {
		return status.Error(codes.Canceled, "The server is full of connections")
	}

	portaudio.Initialize()
	defer portaudio.Terminate()

	defer log.Printf("%v is completely closed \n", index)

	signal_1 := make(chan struct{})
	signal_2 := make(chan struct{})

	go send_to_client(index, signal_1, channel, caller)

	go hear_from_client(caller, signal_2)
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

}
