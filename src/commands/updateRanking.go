package main

import (
	"context"
	"shop/src/database"
	"shop/src/models"

	"github.com/redis/go-redis/v9"
)

func main() {

	database.Connect()
	database.SetupRedis()

	ctx := context.Background()

	var users []models.User
	database.DB.Where("is_ambassador = true").Find(&users)

	for _, user := range users {
		ambassador := models.Ambassador(user)
		ambassador.CalculateRevenue(database.DB)

		database.Cache.ZAdd(ctx, "ambassadors_ranking", redis.Z{
			Score:  *ambassador.Revenue,
			Member: user.Name(),
		})
	}
}
