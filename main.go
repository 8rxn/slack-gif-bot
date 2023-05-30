package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/joho/godotenv"
	libgiphy "github.com/sanzaru/go-giphy"
	"github.com/shomali11/slacker"
)

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

func main() {

	//Loading environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env %v", err)
	}

	//Bot definition
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//Giphy init
	giphy := libgiphy.NewGiphy(os.Getenv("GIPHY_API_KEY"))

	//Keep a log of events
	go printCommandEvents(bot.CommandEvents())

	// Bot Command Definitions

	helloBot := &slacker.CommandDefinition{
		Description: "Hello back with a random gif",
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			name := request.Param("name")
			listOfHello := [3]string{"hello", "hola", "namaste"}
			index := rand.Intn(3)
			dataRandom, err := giphy.GetRandom("hello")
			if err != nil {
				fmt.Println("error in giphy:", err)
			}

			r := fmt.Sprintf("%s %s. I'm Slack-gif-Bot \n Link to a random gif for you: %+v", listOfHello[index], name, dataRandom.Data.Images.Fixed_height.Webp)

			response.Reply(r, slacker.WithThreadReply(true))
		},
	}

	gifTextBot := &slacker.CommandDefinition{
		Description: "Get gif on their prompt",
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			prompt := request.Param("prompt")
			dataRandom, err := giphy.GetRandom(prompt)
			if err != nil {
				fmt.Println("error in giphy:", err)
			}

			r := fmt.Sprintf("Here's a %s Gif for you: %+v", prompt, dataRandom.Data.Images.Fixed_height.Webp)

			response.Reply(r, slacker.WithThreadReply(true))
		},
	}

	gifBot := &slacker.CommandDefinition{
		Description: "Get gif on their prompt",
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			prompt := request.Param("prompt")
			dataRandom, err := giphy.GetRandom(prompt)
			if err != nil {
				fmt.Println("error in giphy:", err)
			}

			r := fmt.Sprintf("Here's a %s Gif for you: %+v", prompt, dataRandom.Data.Images.Fixed_height.Webp)

			response.Reply(r, slacker.WithThreadReply(true))
		},
	}

	//Bot commands listed

	bot.Command("Hello, I am <name>", helloBot)

	bot.Command("Give me a gif of <prompt>", gifTextBot)

	bot.Command(" <prompt>", gifBot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
