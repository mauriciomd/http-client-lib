package services

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mauriciomd/http-client-lib/http_errors"
	"github.com/mauriciomd/http-client-lib/models"
)

func TestCreateDeployment(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	deployment := models.Deployment{
		Id:       uuid.New(),
		Replicas: 1,
		Image:    "golang",
		Labels: models.Label{
			MonitorLabel:  "service-y",
			DeploymentTag: "tag-xpto",
		},
		Ports: []models.Port{
			{Name: "http", Port: 65},
		},
	}

	createdDeployment, err := service.Create(context.Background(), deployment)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if createdDeployment.Id != deployment.Id {
		t.Errorf("Both Ids should be the same")
	}

	defer service.Delete(context.Background(), deployment.Id)
}

func TestCreateDeploymentWithShortTimeout(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	deployment := models.Deployment{
		Id:       uuid.New(),
		Replicas: 1,
		Image:    "golang",
		Labels: models.Label{
			MonitorLabel:  "service-y",
			DeploymentTag: "tag-xpto",
		},
		Ports: []models.Port{
			{Name: "http", Port: 65},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	defer cancel()

	_, err := service.Create(ctx, deployment)
	if err == nil {
		t.Errorf("It should fail due to a timeout. %s", err.Error())
		return
	}

	service.Delete(context.Background(), deployment.Id)
}

func TestCreateNonValidDeployment(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	deployment := models.Deployment{
		Replicas: 0,
		Labels: models.Label{
			MonitorLabel:  "service-y",
			DeploymentTag: "tag-xpto",
		},
		Ports: []models.Port{
			{Name: "http", Port: 65},
		},
	}

	_, err := service.Create(context.Background(), deployment)
	if err == nil {
		t.Errorf("Error should be nil")
	}
}

func TestCreateDuplicatedDeployment(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	deployment := models.Deployment{
		Id:       uuid.New(),
		Replicas: 1,
		Image:    "golang",
		Labels: models.Label{
			MonitorLabel:  "service-y",
			DeploymentTag: "tag-xpto",
		},
		Ports: []models.Port{
			{Name: "http", Port: 65},
		},
	}

	service.Create(context.Background(), deployment)
	defer service.Delete(context.Background(), deployment.Id)

	_, err := service.Create(context.Background(), deployment)
	if err == nil {
		t.Errorf("Error should be nil")
	}
}

func TestGetDeploymentById(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	deployment := models.Deployment{
		Id:       uuid.New(),
		Replicas: 1,
		Image:    "golang",
		Labels: models.Label{
			MonitorLabel:  "service-y",
			DeploymentTag: "tag-xpto",
		},
		Ports: []models.Port{
			{Name: "http", Port: 65},
		},
	}

	service.Create(context.Background(), deployment)
	defer service.Delete(context.Background(), deployment.Id)

	retrievedDeploy, err := service.GetById(context.Background(), deployment.Id)
	if err != nil {
		t.Errorf("Error should be nil")
	}

	if retrievedDeploy.Id != deployment.Id {
		t.Errorf("The retrieved deploy is not the same as the created one.")
	}
}

func TestGetDeploymentWithInvalidId(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	id := uuid.New()

	deployment, err := service.GetById(context.Background(), id)
	if deployment != nil {
		t.Errorf("Deployment should be nil")
	}

	if _, ok := err.(http_errors.NotFoundResource); !ok {
		t.Errorf("Error should be NotFoundError")
	}
}

func TestDeleteDeploymentById(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	deployment := models.Deployment{
		Id:       uuid.New(),
		Replicas: 1,
		Image:    "golang",
		Labels: models.Label{
			MonitorLabel:  "service-y",
			DeploymentTag: "tag-xpto",
		},
		Ports: []models.Port{
			{Name: "http", Port: 65},
		},
	}

	service.Create(context.Background(), deployment)
	if err := service.Delete(context.Background(), deployment.Id); err != nil {
		t.Errorf("Error should be nil")
	}

	retrievedDeploy, err := service.GetById(context.Background(), deployment.Id)
	if err == nil {
		t.Errorf("It should return NotFoundError")
	}

	if retrievedDeploy != nil {
		t.Errorf("Retrieved deploy should be nil")
	}
}

func TestDeleteDeploymentWithInvalidId(t *testing.T) {
	service := NewDeploytment(&http.Client{}, "http://localhost:3000")
	id := uuid.New()

	err := service.Delete(context.Background(), id)

	if _, ok := err.(http_errors.NotFoundResource); !ok {
		t.Errorf("Error should be NotFoundError")
	}
}
