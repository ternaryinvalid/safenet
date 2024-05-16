package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/config"
)

func New() (config config.Config) {
	err := cleanenv.ReadConfig("/bin/config.yml", &config)
	if err != nil {
		err = fmt.Errorf(strings.ReplaceAll(err.Error(), ", ", ",\n"))
		log.Fatal(err)
	}

	return
}
