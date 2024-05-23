package main

import (
	"lambda/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	myApp := app.NewApp()

	lambda.Start(myApp.CronJobHandler.GetStocksPrice)
}
