package controllers

import (
	"shop/src/database"
	"shop/src/models"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func Ambassadors(c fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}

func Ranking(c fiber.Ctx) error {

	ranking, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "ambassadors_ranking", redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		return err
	}
	var result := make(map[string]float64)

	for _, rank := range ranking {
		result[rank.Member.(string)] = rank.Score
	}

	return c.JSON(result)
}
