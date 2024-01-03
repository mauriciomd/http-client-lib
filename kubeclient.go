package kubeclient

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mauriciomd/http-client-lib/services"
)

type KubeClient struct {
	baseURL    string
	port       uint
	timeout    time.Duration
	httpClient *http.Client

	Deployment *services.Deployment
}

func New(opt ...func(*KubeClient) error) (*KubeClient, error) {
	kubeClient := &KubeClient{}
	for _, o := range opt {
		if err := o(kubeClient); err != nil {
			return nil, err
		}
	}

	if kubeClient.port > 0 {
		kubeClient.baseURL = fmt.Sprintf("%s:%d", kubeClient.baseURL, kubeClient.port)
	}

	if kubeClient.timeout > 0 {
		kubeClient.httpClient.Timeout = kubeClient.timeout
	}

	kubeClient.Deployment = services.NewDeploytment(kubeClient.httpClient, kubeClient.baseURL)

	return kubeClient, nil
}

func WithURL(u string) func(*KubeClient) error {
	return func(kc *KubeClient) error {
		_, err := url.ParseRequestURI(u)
		if err != nil {
			return err
		}

		kc.baseURL = u
		return nil
	}
}

func WithHttpClient(hc *http.Client) func(*KubeClient) error {
	return func(kc *KubeClient) error {
		if hc == nil {
			return errors.New("HTTP Client should not be nil")
		}

		kc.httpClient = hc
		return nil
	}
}

func WithPort(p uint) func(*KubeClient) error {
	return func(kc *KubeClient) error {
		kc.port = p
		return nil
	}
}

func WithTimeout(t uint) func(*KubeClient) error {
	return func(kc *KubeClient) error {
		kc.timeout = time.Duration(t * uint(time.Millisecond))
		return nil
	}
}
