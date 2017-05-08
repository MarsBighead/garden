package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("5 seconds cron running:", i)
	})
	c.AddFunc("@every 1m", func() {
		i++
		log.Println("1 minutes cron running:", i)
	})
	c.Start()
	select {}
}
