package repohelpers

import (
	"fmt"
	"log"
)

const postgres = "postgres"

func GetConnectionString(Type, Host, Port, User, Password, Name string) (connectionString string) {
	switch Type {
	case postgres:
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			Host, Port, User, Password, Name, "Europe/Moscow")
	default:
		log.Fatal("Неверный тип БД")
	}

	return
}
