package timer

import (
	"github.com/robfig/cron/v3"
	"log"
)

func TimerFunc() {

	c := cron.New()

	_, _ = c.AddFunc("@every 4s", func() {
		log.Println("tick every 4 second")
	})

	c.Start()

}
