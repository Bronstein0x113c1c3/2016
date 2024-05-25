package caller

import (
	"fmt"
	"log"
	pb "server/protobuf"
	"slices"
	"sync"

	"google.golang.org/grpc/status"
)

type Caller struct {
	host string
	port int
	// input *in.Input
	pb.UnimplementedTheCallServer
	list_of_client []chan []byte
	mutex          *sync.RWMutex
	DeleteSignal   chan string
	WaitSignal     chan struct{}
}

func New(host string, port int /*, ctx context.Context  input *in.Input*/) *Caller {
	return &Caller{
		host: host,
		port: port,
		// input:          input,
		list_of_client: make([]chan []byte, 0),
		mutex:          &sync.RWMutex{},
		DeleteSignal:   make(chan string),
		WaitSignal:     make(chan struct{}),
		// Done:           make(chan struct{}),
	}
}

//	func (s *Caller) Close() {
//		for _, j := range s.list_of_client {
//			close(j)
//		}
//		close(s.DeleteSignal)
//	}
func (s *Caller) add() int {
	s.mutex.Lock()
	s.list_of_client = append(s.list_of_client, make(chan []byte))
	i := len(s.list_of_client) - 1
	s.mutex.Unlock()
	log.Printf("Channel %v is created!!!", i)
	return i
}
func (s *Caller) delete_chan(i int) {
	s.DeleteSignal <- fmt.Sprintf("%v is leaving", i)
	s.mutex.Lock()
	close(s.list_of_client[i])
	s.list_of_client = slices.Delete(s.list_of_client, i, i+1)
	s.mutex.Unlock()
	// s.Done <- struct{}{}
	log.Printf("Channel %v is deleted!!!", i)
	s.WaitSignal <- struct{}{}
}
func (s *Caller) Close() {
	// s.DeleteSignal <- "Shutting down all!!!"
	log.Println("Shutting down all!!!")
	// s.mutex.Lock()
	for i := range s.list_of_client {
		// s.delete_chan(i)
		close(s.list_of_client[i])
	}
	clear(s.list_of_client)
	// s.mutex.Unlock()
	close(s.WaitSignal)
	close(s.DeleteSignal)
}
func (s *Caller) Update(data []byte) {
	for _, client := range s.list_of_client {
		client <- data
	}
}
func (s *Caller) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}
func (s *Caller) Calling(caller pb.TheCall_CallingServer) error {
	index := s.add()
	defer s.delete_chan(index)
	go func() error {
		for {
			data, ok := <-s.list_of_client[index]
			// n, err := reader.Read(data)

			if !ok {
				log.Println("Reading done!!!!")
				return nil
			}
			// if _, err := caller.RecvMsg(; err != io.EOF {
			// 	log.Println("EOF is received!!!!")
			// 	return nil
			// }
			err := caller.Send(&pb.Server_MSGSound{
				Sound: data,
			})
			if err != nil {
				return status.Error(status.Code(err), status.Code(err).String())
			}

		}
	}()
	for {
		if data, _ := caller.Recv(); string(data.Sound) == "Goodbye!!" {
			return nil
		}
	}
}
