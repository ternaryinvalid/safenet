package http_adapter

import (
	"context"
	"fmt"
	"log"
)

func (a *HttpAdapter) Start() {
	startMsg := fmt.Sprint("Сервер запущен на", a.config.Server.Port)

	log.Println(startMsg)

	a.notify <- a.server.ListenAndServe()
	close(a.notify)
}

func (a *HttpAdapter) Notify() <-chan error {
	return a.notify
}

func (a *HttpAdapter) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
