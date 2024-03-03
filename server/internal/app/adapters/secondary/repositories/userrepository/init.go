package userRepository

import (
	"github.com/ternaryinvalid/safenet/server/internal/pkg/config"
	"github.com/ternaryinvalid/safenet/server/internal/pkg/repohelpers"

	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type UserRepository struct {
	config config.Database
	DB     *sql.DB
}

func New(cfg config.Database) *UserRepository {
	hostString := fmt.Sprintf("DB [%s] on address %s:%s ", cfg.Type, cfg.Host, cfg.Port)

	log.Println("Попытка подключения к", hostString)

	connectionString := repohelpers.GetConnectionString(cfg.Type, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Type, connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	log.Println(hostString, "поделючена успешно!")

	return &UserRepository{
		config: cfg,
		DB:     db,
	}
}
