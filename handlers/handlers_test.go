package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"github.com/carlosm27/apiexercise/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsRouteOk(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/product", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductsRouteNotFound(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostProductsRouteBadRequest(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	body := `{
		"sku": "FAL-881898502",
		"name":"Camisa Manga Corta Hombre",
		"brand": "Basement",
		"price": 45.0,
		"size": "",
	  "image":
	  "https://www.coldcutsmerch.com/cdn/shop/products/BASEMENT_BURNING_TIGER_SHIRT_MOCK_WHITE.jpg?v=1619622274"
	}`

	b := bytes.NewBufferString(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/product", b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestPostProductsRouteCreated(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	body := `{
		"sku": "FAL-881898502",
		"name":"Camisa Manga Corta Hombre",
		"brand": "Basement",
		"price": 24990.00,
		"size": "M",
	  "image":
	  "https://www.coldcutsmerch.com/cdn/shop/products/BASEMENT_BURNING_TIGER_SHIRT_MOCK_WHITE.jpg?v=1619622274"
	}`

	b := bytes.NewBufferString(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/product", b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"message":"Product created"}`, w.Body.String())
}

func TestGetProductRouteNotFound(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/product/FAL-881898", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

}

func TestGetProductRouteOk(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/product/FAL-881898502", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestPutProductsRouteOK(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	body := `{
		"sku": "FAL-881898502",
		"name":"Camisa Manga Corta Hombre",
		"brand": "Basement",
		"price": 24990.00,
		"size": "S",
	  "image":
	  "https://www.coldcutsmerch.com/cdn/shop/products/BASEMENT_BURNING_TIGER_SHIRT_MOCK_WHITE.jpg?v=1619622274"
	}`

	b := bytes.NewBufferString(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/product/FAL-881898502", b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"Product updated"}`, w.Body.String())
}

func TestPutProductsRouteNotFound(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	body := `{
		"sku": "FAL-881898502",
		"name":"Camisa Manga Corta Hombre",
		"brand": "Basement",
		"price": 24990.00,
		"size": "S",
	  "image":
	  "https://www.coldcutsmerch.com/cdn/shop/products/BASEMENT_BURNING_TIGER_SHIRT_MOCK_WHITE.jpg?v=1619622274"
	}`

	b := bytes.NewBufferString(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/product/FAL-881898", b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"Product not found!"}`, w.Body.String())
}

func TestDeleteProductsRouteOK(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/product/FAL-881898502", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"Product deleted"}`, w.Body.String())
}

func TestDeleteProductsRouteNotFound(t *testing.T) {
	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := NewServer(db)
	server.RegisterRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/product/FAL-881898", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"Product not found!"}`, w.Body.String())
}
