package util

func Merge(channels ... <-chan bool) <-chan bool {
	signal := make(chan bool)

	go func() {
		for _, channel := range channels {
			<-channel
		}

		signal <- true
	}()

	return signal
}
