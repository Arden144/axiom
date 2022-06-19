package utility

import (
	"os"
	"os/signal"
)

func OnSignal(sig ...os.Signal) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, sig...)
	<-s
}
