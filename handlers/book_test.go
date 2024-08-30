package handlers

import (
	"bytes"
	"encoding/json"
	"example/bookstore/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/books", GetBooks)
	r.POST("/books", PostBooks)
	r.GET("/books/:id", GetBookByID)
	return r
}

func TestGetBooks(t *testing.T) {
	router := SetupRouter()
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostBooks(t *testing.T) {
	router := SetupRouter()
	book := models.Book{Title: "Test Book", Author: "Test Author", Price: 9.99}
	jsonValue, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Couldn't marshal book: %v", err)
	}
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBookByID(t *testing.T) {
	router := SetupRouter()
	req, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
