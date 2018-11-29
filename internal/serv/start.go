package serv

import (
	"log"

	"bz-lib/cfg"
)

func Start_Services(config *cfg.Config) error {
	log.Println("Запуск сервисов...")
	Start_ListClients_v1(config)
	log.Println("Сервисы запущены.")
	return nil
}
