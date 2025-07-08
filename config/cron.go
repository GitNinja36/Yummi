package config

import (
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

func StartCronJob() {
	c := cron.New()
	c.AddFunc("*/14 * * * *", func() {
		resp, err := http.Get(os.Getenv("API_URL"))
		if err != nil {
			log.Printf("Cron failed to hit API:", err)
			return
		}
		defer resp.Body.Close()
		log.Fatal("Cron hit", os.Getenv("API_URL"), "Status:", resp.StatusCode)
	})
	c.Start()
}
