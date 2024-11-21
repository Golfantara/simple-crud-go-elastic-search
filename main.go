package main

import (
	"elasticsearch/config"
	"elasticsearch/feature/user"
	uh "elasticsearch/feature/user/handler"
	ur "elasticsearch/feature/user/repository"
	ure "elasticsearch/feature/user/repository/elasticsearch"
	uu "elasticsearch/feature/user/usecase"
	"elasticsearch/routes"
	"elasticsearch/utils"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	routes.User(mux, UserHandler())

	log.Println("Server is running at :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func UserHandler() user.Handler {
	cfg := config.LoadDBConfig()
	client, err := utils.NewElasticClient()
	if err != nil {
		log.Printf("Elastic configuration: URL=%s, User=%s", cfg.ELASTIC_URL, cfg.ELASTIC_USER)
		log.Fatalf("Error creating ElasticSearch client: %v", err)
	}

	db := utils.InitDB()

	repoElastic := ure.NewUserRepository(client)
	repo := ur.New(db)
	uc := uu.New(repoElastic, repo)
	return uh.New(uc)
}
