package controllers

import (
	"encoding/json"
	"httpReq/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find((&products))

	c.JSON(http.StatusOK, gin.H{"prodcuts": products})
}
func Show(c *gin.Context) {
	var products []models.Product
	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Meassage": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Meassage": "Data not found"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"products": products})

}
func Create(c *gin.Context) {
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid Payload Data Product"})
		return
	}
	models.DB.Create(&products)
	c.JSON(http.StatusOK, gin.H{"Message": "Data successsfully added",
		"payload": products})
}
func Update(c *gin.Context) {
	var products models.Product

	id := c.Param("id")
	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&products).Where("id = ?", id).Updates(&products).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to Update data"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"Messaage": "Data Product successfully updated",
		"Update": products})
}
func Delete(c *gin.Context) {
	var products models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&products, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Succesfully delete",
		"result": products})
}
