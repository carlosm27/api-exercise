package handlers

import (
	//"log"
	"net/http"

	"github.com/carlosm27/apiexercise/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewProduct struct {
	Sku            string  `json:"sku" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Brand          string  `json:"brand" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
	Size           string  `json:"Size" binding:"required"`
	PrincipalImage string  `json:"image" binding:"required"`
}

type ProductUpdate struct {
	Sku            string  `json:"sku"`
	Name           string  `json:"name"`
	Brand          string  `json:"brand"`
	Price          float64 `json:"price"`
	Size           string  `json:"Size"`
	PrincipalImage string  `json:"image"`
}

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) RegisterRouter(router *gin.Engine) {
	router.GET("/product", s.getProducts)
	router.GET("/product/:sku", s.getProduct)
	router.POST("/product", s.postProduct)
	router.PUT("/product/:sku", s.updateProduct)
	router.DELETE("/product/:sku", s.deleteProduct)

}

func (s *Server) getProducts(c *gin.Context) {

	var products []model.Product

	if err := s.db.Find(&products).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)

}

func (s *Server) getProduct(c *gin.Context) {

	var product model.Product

	if err := s.db.Where("sku= ?", c.Param("sku")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (s *Server) postProduct(c *gin.Context) {

	var product NewProduct

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct := model.Product{Sku: product.Sku, Name: product.Name, Brand: product.Brand, Price: product.Price, Size: product.Size, PrincipalImage: product.PrincipalImage}

	if err := s.db.Create(&newProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func (s *Server) updateProduct(c *gin.Context) {

	var product model.Product

	if err := s.db.Where("sku = ?", c.Param("sku")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	var updateGrocery ProductUpdate

	if err := c.ShouldBindJSON(&updateGrocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Model(&product).Updates(model.Product{Name: updateGrocery.Name, Brand: updateGrocery.Brand, Price: updateGrocery.Price, Size: updateGrocery.Size}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})

}

func (s *Server) deleteProduct(c *gin.Context) {

	var product model.Product

	if err := s.db.Where("sku = ?", c.Param("sku")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	if err := s.db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})

}
