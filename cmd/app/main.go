package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/JonatasMSantos/golang-apis-messaging-hexagonal-architecture/internal/infra/akafka"
	"github.com/JonatasMSantos/golang-apis-messaging-hexagonal-architecture/internal/infra/repository"
	"github.com/JonatasMSantos/golang-apis-messaging-hexagonal-architecture/internal/infra/web"
	"github.com/JonatasMSantos/golang-apis-messaging-hexagonal-architecture/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUserCase(repository)
	listProductsUseCase := usecase.NewListProductUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreatProductHandler)
	r.Get("/products", productHandlers.ListProductsHandter)

	//Goroutines
	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	//Goroutines
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			// salvar log de erro
			// continue
		}
		_, err = createProductUseCase.Execute(dto)
	}
}
