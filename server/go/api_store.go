/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
	"github.com/rwxrob/opsapi/internal/api"
)

// DeleteOrder - Delete purchase order by ID
func DeleteOrder(c *gin.Context) {
	api.Handle(c, "DeleteOrder")
}

// GetInventory - Returns pet inventories by status
func GetInventory(c *gin.Context) {
	api.Handle(c, "GetInventory")
}

// GetOrderById - Find purchase order by ID
func GetOrderById(c *gin.Context) {
	api.Handle(c, "GetOrderById")
}

// PlaceOrder - Place an order for a pet
func PlaceOrder(c *gin.Context) {
	api.Handle(c, "PlaceOrder")
}
