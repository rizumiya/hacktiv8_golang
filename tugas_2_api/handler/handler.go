package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "tugas_2_api/database"
    "tugas_2_api/models"
)

func CreateOrder(c *gin.Context) {
    db := database.GetDb()
    var order models.Order
    if err := c.BindJSON(&order); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if err := db.Create(&order).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusCreated, gin.H{"pesan": "Data berhasil ditambahkan", "order": order})
}

func GetOrders(c *gin.Context) {
    db := database.GetDb()
    var orders []models.Order
    if err := db.Preload("Items").Find(&orders).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func UpdateOrder(c *gin.Context) {
    db := database.GetDb()
    ID, err := strconv.Atoi(c.Param("orderId"))
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID tidak sesuai"})
      return
    }

    var updatedOrder models.Order
    if err := c.BindJSON(&updatedOrder); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    for _, item := range updatedOrder.Items {
      db.Model(&models.Item{}).Where("id = ?", item.ID).Updates(item)
    }

    if err := db.Model(&models.Order{ID: uint(ID)}).Updates(updatedOrder).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"pesan": "Data berhasil diubah", "order": updatedOrder})
}

func DeleteOrder(c *gin.Context) {
    db := database.GetDb()
    ID, err := strconv.Atoi(c.Param("orderId"))
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID tidak sesuai"})
      return
    }

    if err := db.Delete(&models.Order{ID: uint(ID)}).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"pesan": "Data berhasil dihapus"})
}
