package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/wrlou/delinkcious/pkg/db_util"
	"github.com/wrlou/delinkcious/pkg/social_graph_client"
	. "github.com/wrlou/delinkcious/pkg/test_util"
	"log"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	db, err := db_util.RunLocalDB("social_graph_manager")
	check(err)

	// Ignore if table doesn't exist (will be created by service)
	err = db_util.DeleteFromTableIfExist(db, "social_graph")
	check(err)
}

func main() {
	initDB()

	ctx := context.Background()
	defer StopService(ctx)
	RunService(ctx, ".", "social_graph_service")

	// Run some tests with the client
	cli, err := social_graph_client.NewClient("localhost:9090")
	check(err)

	following, err := cli.GetFollowing("wrlou")
	check(err)
	log.Print("wrlou is following:", following)
	followers, err := cli.GetFollowers("wrlou")
	check(err)
	log.Print("wrlou is followed by:", followers)

	err = cli.Follow("wrlou", "liat")
	check(err)
	err = cli.Follow("wrlou", "guy")
	check(err)
	err = cli.Follow("guy", "wrlou")
	check(err)
	err = cli.Follow("saar", "wrlou")
	check(err)
	err = cli.Follow("saar", "ophir")
	check(err)

	following, err = cli.GetFollowing("wrlou")
	check(err)
	log.Print("wrlou is following:", following)
	followers, err = cli.GetFollowers("wrlou")
	check(err)
	log.Print("wrlou is followed by:", followers)
}
