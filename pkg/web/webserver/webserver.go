package webserver

import "github.com/fahmifan/commurz/pkg/service"

type Config struct {
	CartService *service.Service
}

type Webserver struct {
	*Config
}

func (server *Webserver) Run() error {

	return nil
}
