package caller

import (
	"fmt"
	"log"
	pb "server/protobuf"
	"slices"
	"sync"
)

type Caller struct {
	host string
	port int
	// input *in.Input
	pb.UnimplementedTheCallServer
	ListOfClient []chan []byte
	mutex        *sync.RWMutex
	ChangeSignal chan struct{}
	// receiver     chan []byte
	// Done           chan struct{}
}

func New(host string, port int /* input *in.Input*/) *Caller {
	return &Caller{
		host: host,
		port: port,
		// input:          input,
		ListOfClient: make([]chan []byte, 0),
		mutex:        &sync.RWMutex{},
		ChangeSignal: make(chan struct{}),
		// Done:           make(chan struct{}),
	}
}

//	func (s *Caller) Close() {
//		for _, j := range s.ListOfClient {
//			close(j)
//		}
//		close(s.DeleteSignal)
//	}
func (s *Caller) add() int {
	// s.mutex.Lock()
	s.ChangeSignal <- struct{}{}
	s.ListOfClient = append(s.ListOfClient, make(chan []byte))
	i := len(s.ListOfClient) - 1
	// s.mutex.Unlock()
	log.Printf("Channel %v is created!!!", i)
	return i
}
func (s *Caller) delete_chan(i int, closed bool) {

	// s.mutex.Lock()
	if !closed {
		s.ChangeSignal <- struct{}{}
		close(s.ListOfClient[i])
	}
	s.ListOfClient = slices.Delete(s.ListOfClient, i, i+1)
	// s.mutex.Unlock()
	// s.Done <- struct{}{}
	log.Printf("Channel %v is deleted!!!", i)
}

// func (s *Caller) UpdateOrClose(receiver chan []byte, signal chan struct{}) {
// 	// for _, client := range s.ListOfClient {
// 	// 	client <- data
// 	// }
// 	for data := range receiver {
// 		select {
// 		case <-s.DeleteSignal:
// 			log.Println("Change received....")
// 			return
// 		default:
// 			for _, client := range s.ListOfClient {
// 				client <- data
// 			}
// 		}
// 	}
// 	log.Println("Starting to close channel")
// 	for _, client := range s.ListOfClient {
// 		close(client)
// 	}
// 	log.Println("Closed these channel....")
// 	close(s.DeleteSignal)
// 	log.Println("No more addition or deletion from the channel, closing all....")
// 	return

// }

func (s *Caller) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}
func (s *Caller) Calling(caller pb.TheCall_CallingServer) error {
	// reader := bufio.NewReader(s.input.GetStream())
	index := s.add()
	// defer log.Printf("%v is completely closed \n", index)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// defer s.delete_chan(index, true)
		defer log.Printf("%v is closed \n", index)
		defer wg.Done()
		for {
			data, ok := <-s.ListOfClient[index]
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
