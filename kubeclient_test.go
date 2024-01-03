package kubeclient

import (
	"net/http"
	"testing"
)

func TestDefaultKubeClient(t *testing.T) {
	kc, err := New()

	if kc == nil {
		t.Errorf("Client should be created")
	}

	if err != nil {
		t.Errorf("Error must be nil. Got: %s", err.Error())
	}
}

func TestValidtKubeClient(t *testing.T) {
	kc, err := New(
		WithURL("http://localhost"),
		WithPort(3000),
		WithTimeout(30),
		WithHttpClient(&http.Client{}),
	)

	if kc == nil {
		t.Errorf("Client should be created")
	}

	if err != nil {
		t.Errorf("Error must be nil. Got: %s", err.Error())
	}
}

func TestClientWithInvalidUrl(t *testing.T) {
	kc, err := New(WithURL("http//localhost"))
	if kc != nil {
		t.Errorf("Client should not be created.")
	}

	if err == nil {
		t.Errorf("Error should not be nil")
	}
}

func TestClientWithInvalidHTTPClient(t *testing.T) {
	kc, err := New(WithHttpClient(nil))
	if kc != nil {
		t.Errorf("Client should not be created.")
	}

	if err == nil {
		t.Errorf("Error should not be nil")
	}
}
