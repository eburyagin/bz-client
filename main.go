package main

import (
	"bz-client/internal/serv"
	"bz-lib/cfg"
	"flag"
	"log"
	"runtime"
)

var cf = flag.String("c", "./config.json", "конфигурационный файл")

func main() {

	flag.Parse()

	config, err := cfg.Load(*cf)
	if err != nil {
		log.Fatalln("Ошибка загрузки конфигурации!")
	}

	serv.Start_Services(config)

	runtime.Goexit()
	log.Println("Завершил работу")

}
