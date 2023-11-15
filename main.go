package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// create a channel and loop over it, for every event print data
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	// set environment with variables for the bot and app tokens
	// not necessary, you can directly pass the tokens in the slacker.NewClient underneath, but good practice
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6220711908880-6220764286672-KPSrhiW05CuO53t8NfLwSeQq")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A065TMZ0U1G-6194824494869-999f0366a3e48eeacdd25441ad28886c67be8bd0dc6a9a884004eae40cd6ab72")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// prints command events -> when a command is passed to the bot it will print the results... must write func ^
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition {
		Description: "yob calculator",
		Examples: []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			// convert the string in order to calculate age
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	// stop the bot
	ctx, cancel := context.WithCancel(context.Background())
	// defer the cancel until after the context (the function) has run its course
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}