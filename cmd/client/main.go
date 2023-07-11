package main

import (
	"log"

	"GSS/internal/client/config"
	"GSS/internal/metrics"

	"github.com/gofiber/fiber/v2"
)

func HiServer(client *fiber.Client) (err error) {
	agent := client.Get("http://127.0.0.1:4040")
	err = agent.Parse()
	if err != nil {
		log.Fatal(err)
	}
	statusCode, body, errs := agent.String()
	log.Println(statusCode)
	log.Println(body)
	if errs != nil {
		log.Fatal(errs)
	}
	return nil
}

func SendMetrics(client *fiber.Client, metric map[string]uint64) (err error) {
	agent := client.Post("http://127.0.0.1:4040")
	log.Println("Sending metric")
	log.Println(metric)
	err = agent.JSON(metric).Parse()
	if err != nil {
		log.Fatal(err)
	}
	statusCode, _, _ := agent.String()
	log.Println(statusCode)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	cfg, err := config.LoadConfig("../../config/client")
	client := fiber.AcquireClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg.ServerAddr)

	// err = HiServer(client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	metric := metrics.Get()
	err = SendMetrics(client, metric)
	if err != nil {
		log.Fatal(err)
	}
}
