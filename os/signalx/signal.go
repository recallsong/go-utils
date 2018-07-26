package signalx

import (
	"os"
	"os/signal"
)

func Notify(sig ...os.Signal) <-chan os.Signal {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, sig...)
	return sc
}
