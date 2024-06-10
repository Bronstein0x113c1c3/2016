package serverimpl

type Chunk struct {
	data []byte
	name string
	id   int
}

// chunk for recognizing the packet, also for channel

func (c Chunk) Get() ([]byte, string, int) {
	return c.data, c.name, c.id
}
func CreateChunk(data []byte, name string, id int) Chunk {
	return Chunk{
		data: data,
		name: name,
		id:   id,
	}
}
