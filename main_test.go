package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ashishsingh4u/bookmicroservice/controllers"
	"github.com/ashishsingh4u/bookmicroservice/models"
	"github.com/gin-gonic/gin"
)

// Private Method
func initConfiguration() (r *gin.Engine) {
	models.ConnectDatabase()

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r = gin.Default()
	// Grouping
	v1 := r.Group("/v1")
	{
		v1.GET("/books", controllers.GetBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}

	return
}

func TestCreateBookMethod(t *testing.T) {
	r := initConfiguration()

	input := "{ \"title\": \"Start with Why\", \"author\": \"Simon Sinek\"}"
	req, err := http.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(input))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusCreated {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusCreated, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
	}
}

func TestGetBooksMethod(t *testing.T) {
	r := initConfiguration()

	req, err := http.NewRequest(http.MethodGet, "/v1/books", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestFindBookMethod(t *testing.T) {
	r := initConfiguration()

	req, err := http.NewRequest(http.MethodGet, "/v1/books/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestUpdateBookMethod(t *testing.T) {
	r := initConfiguration()

	input := "{ \"title\": \"Start with Why\", \"author\": \"Simon Sinek\"}"
	req, err := http.NewRequest(http.MethodPatch, "/v1/books/1", strings.NewReader(input))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestDeleteBookMethod(t *testing.T) {
	r := initConfiguration()

	req, err := http.NewRequest(http.MethodDelete, "/v1/books/107", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else if w.Code == http.StatusInternalServerError {
		if w.Body.String() == "{\"error\":\"database error: record not found\"}" {
			t.Logf("Expected to get status %d is same ast %d\n", http.StatusInternalServerError, w.Code)
		}
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
