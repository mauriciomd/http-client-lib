package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	kubeclient "github.com/mauriciomd/http-client-lib"
	"github.com/mauriciomd/http-client-lib/models"
)

func main() {
	client, err := kubeclient.New(
		kubeclient.WithURL("http://localhost"),
		kubeclient.WithPort(3000),
		kubeclient.WithTimeout(100),
		kubeclient.WithHttpClient(&http.Client{}),
	)
	if err != nil {
		panic(err)
	}

	d := models.Deployment{
		Id:       uuid.New(),
		Replicas: 1,
		Image:    "golang",
		Labels:   models.Label{MonitorLabel: "service-y", DeploymentTag: "xpto"},
		Ports:    []models.Port{{Name: "http", Port: 65}},
	}

	client.Deployment.Create(context.Background(), d)

	x, _ := client.Deployment.GetById(context.Background(), d.Id)
	fmt.Println(x.Id)

	client.Deployment.Delete(context.Background(), x.Id)
}
