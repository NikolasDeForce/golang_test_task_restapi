package main

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("not invalid apiToken and idInstance", func(t *testing.T) {
		statusCode := GetSettingsHandler("http://localhost", ":8010", "123", "A123PI", GetSettings)

		if statusCode != http.StatusNotFound {
			t.Errorf("got %v want %v", statusCode, http.StatusNotFound)
		}
	})

	t.Run("check GET method to valid apiToken and idInstance", func(t *testing.T) {
		statusCode := GetSettingsHandler("http://localhost", ":8010", "1234567890", "A123p456I789", GetSettings)

		if statusCode != http.StatusOK {
			t.Errorf("got %v want %v", statusCode, http.StatusOK)
		}
	})

	t.Run("check POST method to valid apiToken and idInstance", func(t *testing.T) {
		statusCode := SendMessageHandler("http://localhost", ":8010", "1234567890", "A123p456I789", SendMessage)

		if statusCode != http.StatusOK {
			t.Errorf("got %v want %v", statusCode, http.StatusOK)
		}
	})
}
