package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/revacholiere-moralist/blogbackend/database"
	"github.com/revacholiere-moralist/blogbackend/models"
	"github.com/revacholiere-moralist/blogbackend/util"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost models.Blog

	if err := c.BodyParser(&blogPost); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post added",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit

	var total int64
	var blogs []models.Blog

	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&blogs)
	database.DB.Model(&models.Blog{}).Count(&total)

	lastPage := int(total) / limit
	if int(total)%limit > 0 {
		lastPage += 1
	}

	return c.JSON(fiber.Map{
		"data": blogs,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogPost models.Blog

	database.DB.Where("id=?", id).Preload("User").First(&blogPost)
	return c.JSON(fiber.Map{
		"data": blogPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)
	return c.JSON(blog)
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)

	var blogs []models.Blog
	database.DB.Model(&blogs).Where("user_id=?", id).Preload("User").Find(&blogs)

	return c.JSON(blogs)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog {
		Id:uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		return c.JSON(fiber.Map {
			"message": "Record not found",
		})
	}

	return c.JSON(fiber.Map {
		"message": "post deleted successfully",
	})
}
