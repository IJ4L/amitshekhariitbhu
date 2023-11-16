package bootstrap

import (
	"log"
)

type Application struct {
	Env   *Env
	MySQL *Client
}

func App() Application {

	app := &Application{}
	app.Env = NewEnv()
	app.MySQL = NewMySQLDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	err := CloseMySQLConnection(app.MySQL)
	if err != nil {
		log.Fatal(err)
	}
}