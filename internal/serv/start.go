package serv

import (
	"log"

	"bz-client/internal/cfg"
)

func Start_Services(config *cfg.Config) error {
	log.Println("Запуск сервисов...")
	Start_ListClients_v1(config)
	log.Println("Сервисы запущены.")
	return nil
}
