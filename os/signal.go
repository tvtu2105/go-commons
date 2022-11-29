package os

func ListenOsSignal() chan bool {
	signal := make(chan bool)
	go func() {

	}()
	return signal
}
