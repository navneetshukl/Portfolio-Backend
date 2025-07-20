package main

import (
	"log"
	"os"
	email "portfolio/internals/adapter/external"
	"portfolio/internals/config"
	routes "portfolio/internals/interface"
	"portfolio/internals/interface/api/rest/handlers"
	"portfolio/internals/usecase"
	"time"
)

func main() {
	// Load env
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("error in loading .env ",err)
	// 	return
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	conf, err := config.LoadConfig("./.config")
	if err != nil {
		log.Println("error in loading the config file ", err)
		return
	}

	emailSvc := email.NewMailSvc(conf)
	emailUseCase := usecase.NewEmailUseCase(conf, emailSvc)
	emailhandler := handlers.NewHandler(emailUseCase)

	// Start cron job in background
	go func() {
		log.Println("I am fine") // Print immediately on start

		ticker := time.NewTicker(14*time.Minute + 30*time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C
			log.Println("I am fine")
		}
	}()

	router := routes.SetUpRoutes(&emailhandler)
	err = router.Run(":" + port)
	if err != nil {
		log.Println("error in starting the server")
		return
	}

}
