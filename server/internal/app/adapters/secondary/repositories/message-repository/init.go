package message_repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/config"
	"github.com/ternaryinvalid/safenet/server/internal/pkg/repohelpers"
	"log"
	"os"
)

type MessageRepository struct {
	cfg config.Database
	DB  *sql.DB
}

func New(cfg config.Database) *MessageRepository {
	currentHostString := fmt.Sprintf("DB host [%s:%s]. ", cfg.Host, cfg.Port)

	log.Println(currentHostString, "CONNECTING...")

	connectionString := repohelpers.GetConnectionString(cfg.Type, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Type, connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())

		os.Exit(1)
	}

	log.Println(currentHostString, "CONNECTED!")

	return &MessageRepository{
		cfg: cfg,
		DB:  db,
	}
}
