package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/config"
	"log"
	"strings"
)

func New() (config config.Config) {
	err := cleanenv.ReadConfig("/bin/config.yml", &config)
	if err != nil {
		err = fmt.Errorf(strings.ReplaceAll(err.Error(), ", ", ",\n"))
		log.Fatal(err)
	}

	return
}
