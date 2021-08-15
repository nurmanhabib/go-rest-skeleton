package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rollbar/rollbar-go"

	"github.com/joho/godotenv"
	"github.com/nurmanhabib/go-rest-skeleton/interfaces/routers"

	"github.com/nurmanhabib/go-rest-skeleton/interfaces/cmd"

	"github.com/nurmanhabib/go-rest-skeleton/config"
	"github.com/urfave/cli/v2"
)

func catchError() {
	if r := recover(); r != nil {
		rollbar.Critical(r)
		rollbar.Wait()
	}
}

func main() {
	defer catchError()

	// Check .env file
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided")
	}

	conf := config.New()

	timeLoc, _ := time.LoadLocation(conf.App.Timezone)
	time.Local = timeLoc

	if conf.Rollbar.IsEnabled() {
		rollbar.SetToken(conf.Rollbar.Token)
		rollbar.SetEnvironment(conf.Env)
	}

	app := cmd.NewCli()
	app.Action = func(c *cli.Context) error {
		// Init Router
		router := routers.New(conf).Init()

		// Run app at defined port
		appPort := conf.App.Port
		srv := &http.Server{
			Addr:    ":" + appPort,
			Handler: router,
		}

		errSrv := make(chan error, 1)

		go func() {
			errSrv <- srv.ListenAndServe()
		}()

		shutdownChannel := make(chan os.Signal, 1)
		signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-shutdownChannel:
			log.Println("signal:", sig)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatalln("Server forced to shutdown:", err)
			}

			log.Println("Server shutdown")

		case err := <-errSrv:
			if err != nil {
				log.Fatalln("Server error:", err)
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(app)
	}
}
