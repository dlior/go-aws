package app

import (
	"lambda/cronJob"
	"lambda/database"
)

type App struct {
	CronJobHandler cronJob.CronJobHandler
}

func NewApp() App {
	db := database.NewDynamoDBClient()
	cronJobHandler := cronJob.NewCronJobHandler(db)

	return App{
		CronJobHandler: cronJobHandler,
	}
}
