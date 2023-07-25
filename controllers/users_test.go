package controllers

import (
	"andybrandproject/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------------
//
//	TEST THE GET METHOD - START
//
// ---------------------------------------------------------------------
func TestGetUsers(t *testing.T) {
	app := fiber.New()
	app.Get("/users", Users)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to perform the test: %s", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

// ---------------------------------------------------------------------
//
//	TEST THE GET METHOD - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	TEST THE POST METHOD - START
//
// ---------------------------------------------------------------------
func TestUsersPost(t *testing.T) {
	app := fiber.New()

	app.Post("/users", CreateUsers)

	input := models.Users{
		Name:        "name time",
		Address:     "the des",
		Description: "add ress",
	}

	bodyReq, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal input JSON: %s", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(bodyReq))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to perform the test: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	t.Log(resp.StatusCode)
	t.Log(string(body))

	assert.Equal(t, 200, resp.StatusCode)
}

// ---------------------------------------------------------------------
//
//	TEST THE POST METHOD - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	TEST THE PUT METHOD - START
//
// ---------------------------------------------------------------------
func TestUsersPut(t *testing.T) {
	app := fiber.New()

	app.Put("/users/:id", UpdateUsers)

	input := models.Users{
		Name:        "Update Name",
		Address:     "Updated address",
		Description: "Updated Description",
	}

	bodyReq, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal input JSON: %s", err)
	}

	req, err := http.NewRequest(http.MethodPut, "/users/64c00fbec355ed265e43e152", bytes.NewBuffer(bodyReq))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to perform the test: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	t.Log(resp.StatusCode)
	t.Log(string(body))

	assert.Equal(t, 200, resp.StatusCode)
}

// ---------------------------------------------------------------------
//
//	TEST THE PUT METHOD - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	TEST THE DELETE METHOD - START
//
// ---------------------------------------------------------------------
func TestDeleteUser(t *testing.T) {
	app := fiber.New()

	app.Delete("/users/:id", Delete)

	userID := "64c0120248ea7ab1c8c38c2f"

	req, err := http.NewRequest(http.MethodDelete, "/users/"+userID, nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %s", err)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to perform the test: %s", err)
	}
	defer resp.Body.Close()

	t.Log(resp.StatusCode)

	assert.Equal(t, 200, resp.StatusCode)
}

// ---------------------------------------------------------------------
//
//	TEST THE DELETE METHOD - END
//
// ---------------------------------------------------------------------
