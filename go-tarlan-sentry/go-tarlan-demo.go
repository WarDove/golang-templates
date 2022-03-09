package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://bb7e4775d8a44c7cacb9370306917ea1@sentry.rabita.az/6",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
		fmt.Println("we generated an issuee braaah!")
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Eziz xanımlar, Bayramınız mübarek! Size beraberlik, azadlıq ve xoşbextliklerle dolu bir heyat arzulayırıq")
	fmt.Println("how ya doin? :P")

}
