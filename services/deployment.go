package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/mauriciomd/http-client-lib/http_errors"
	"github.com/mauriciomd/http-client-lib/models"
)

type Deployment struct {
	baseUrl    string
	resource   string
	httpClient *http.Client
}

func NewDeploytment(c *http.Client, url string) *Deployment {
	return &Deployment{
		baseUrl:    url,
		resource:   "deployments",
		httpClient: c,
	}
}

func (d *Deployment) Create(ctx context.Context, deployment models.Deployment) (*models.Deployment, error) {
	url := fmt.Sprintf("%s/%s", d.baseUrl, d.resource)
	j, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(j)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	res, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusBadRequest {
		return nil, http_errors.FromBadRequest(res)
	}

	if res.StatusCode != http.StatusCreated {
		return nil, http_errors.FromHTTPResponse(res)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var dep models.Deployment
	err = json.Unmarshal(body, &dep)
	if err != nil {
		return nil, err
	}

	return &dep, nil
}

func (d *Deployment) GetById(ctx context.Context, id uuid.UUID) (*models.Deployment, error) {
	url := fmt.Sprintf("%s/%s/%s", d.baseUrl, d.resource, id)

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, http_errors.NewNotFound(id, d.resource)
	}

	if res.StatusCode != http.StatusOK {
		return nil, http_errors.FromHTTPResponse(res)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var deployment models.Deployment
	err = json.Unmarshal(body, &deployment)
	if err != nil {
		return nil, err
	}

	return &deployment, nil
}

func (d *Deployment) Delete(ctx context.Context, id uuid.UUID) error {
	url := fmt.Sprintf("%s/%s/%s", d.baseUrl, d.resource, id)
	req, err := http.NewRequest(http.MethodDelete, url, http.NoBody)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	res, err := d.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return http_errors.NewNotFound(id, d.resource)
	}

	if res.StatusCode != http.StatusNoContent {
		return http_errors.FromHTTPResponse(res)
	}

	return nil
}
