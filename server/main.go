package main

import (
	userRepository "github.com/ternaryinvalid/safenet/server/internal/app/adapters/secondary/repositories/userrepository"
	"github.com/ternaryinvalid/safenet/server/internal/app/service"
	"github.com/ternaryinvalid/safenet/server/internal/pkg/config"
)

func main() {
	cfg := config.New()

	userRepository := userRepository.New(cfg.Databases.UserRepository)
	_ = userRepository

	apiService := service.New(userRepository)
	_ = apiService
}
