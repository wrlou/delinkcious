package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/wrlou/delinkcious/pkg/db_util"
	om "github.com/wrlou/delinkcious/pkg/object_model"
	. "github.com/wrlou/delinkcious/pkg/test_util"
	"github.com/wrlou/delinkcious/pkg/user_client"
	"log"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	db, err := db_util.RunLocalDB("user_manager")
	if err != nil {
		return
	}

	tables := []string{"sessions", "users"}
	for _, table := range tables {
		err = db_util.DeleteFromTableIfExist(db, table)
		check(err)
	}
}

func main() {
	initDB()

	ctx := context.Background()
	defer StopService(ctx)
	RunService(ctx, ".", "user_service")

	// Run some tests with the client
	cli, err := user_client.NewClient("localhost:7070")
	check(err)

	err = cli.Register(om.User{"gg@gg.com", "wrlou"})
	check(err)
	log.Print("wrlou has registered successfully")

	session, err := cli.Login("wrlou", "secret")
	check(err)
	log.Print("wrlou has logged in successfully. the session is: ", session)

	err = cli.Logout("wrlou", session)
	check(err)
	log.Print("wrlou has logged out successfully.")

}
