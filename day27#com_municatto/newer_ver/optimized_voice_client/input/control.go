package input

func (i *Input) Close() {
	close(i.signal)
}
