package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nurmanhabib/go-rest-skeleton/interfaces/routers"

	cmd2 "github.com/nurmanhabib/go-rest-skeleton/interfaces/cmd"

	"github.com/nurmanhabib/go-rest-skeleton/config"
	"github.com/urfave/cli/v2"
)

func main() {
	// Check .env file
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()

	timeLoc, _ := time.LoadLocation(conf.App.Timezone)
	time.Local = timeLoc

	app := cmd2.NewCli()
	app.Action = func(c *cli.Context) error {
		// Init Router
		router := routers.New(conf).Init()

		// Run app at defined port
		appPort := conf.App.Port

		log.Println(router.Run(":" + appPort))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(app)
	}
}
