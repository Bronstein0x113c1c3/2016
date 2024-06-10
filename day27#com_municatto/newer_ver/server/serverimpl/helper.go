package serverimpl

import (
	pb "serv/protobuf"
)

func receive(conn pb.Calling_VoIPServer, data_chan chan Chunk, id int, signal chan struct{}) {
	defer close(signal)
	for {
		data, err := conn.Recv()
		if err != nil {
			return
		}
		// fmt.Printf("%v\n%v\n%v\n")
		c := CreateChunk(data.GetChunk(), data.GetName(), id)

		data_chan <- c
	}
}
func send(conn pb.Calling_VoIPServer, data_chan chan Chunk, id int, signal chan struct{}) {
	defer close(signal)
	for {
		data, ok := <-data_chan
		if !ok {
			return
		}
		conn.Send(
			&pb.ServerRES{
				Msg: &pb.ClientMSG{
					Chunk: data.data,
					Name:  data.name,
				},
				Id: int32(id),
			},
		)
	}
}
