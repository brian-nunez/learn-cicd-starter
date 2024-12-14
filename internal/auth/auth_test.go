package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"testing"
)

func Test_GetAPIKey(t *testing.T) {
	var header map[string][]string = make(map[string][]string)

	result, err := auth.GetAPIKey(header)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
