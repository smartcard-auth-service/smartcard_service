package transfer

var BatchCh chan []byte

func InitTransferChan() {
	BatchCh = make(chan []byte)
}

func CloseTransferChan() {
	close(BatchCh)
}
