package main

import (
	"context"
	"fmt"
	apikey "frgt/apiKey"
	"frgt/bash"
	"frgt/errormessage"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/urfave/cli"
	"google.golang.org/api/option"
)

type err struct {
	errors string
}

func generateText() (string, error) {

	ctx := context.Background()

	//Get API key
	apiKey := apikey.GetAPIKey()

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	errormessage.Error(err)

	defer client.Close()

	sliceOfStrings := os.Args[1:]
	joinedString := strings.Join(sliceOfStrings, " ")
	overrideAction := "generate a linux command line for  " + joinedString + " and extract the command line only"

	model := client.GenerativeModel("gemini-pro")
	config := model.GenerationConfig
	resp, err := model.GenerateContent(ctx, genai.Text(overrideAction))
	config.SetTemperature(0)
	config.SetTopK(1)
	extractResponse := resp.Candidates[0].Content.Parts[0]
	resToString := fmt.Sprint(extractResponse)

	return resToString, err
}

func main() {

	app := &cli.App{
		EnableBashCompletion: true,
		Copyright:            "Amr Taha (https://github.com/amroshworld)",
		Email:                "amr.taha.fue@gmail.com",
		Commands: []cli.Command{

			{

				Name:  "frgt",
				Usage: "generate  command for your specific task using ai from just text you enter",
				Action: func(cCtx *cli.Context) error {
					response, err := generateText()

					if response != "" {

						fmt.Println(response)
						bash.StoreText(response)
						bash.ReloadHistory()

					} else {
						fmt.Println("no response received")
						return err

					}
					return err
				},
				Subcommands: []cli.Command{

					{
						Name:  "rmvkey",
						Usage: "remove api key of frgt to allow you to add your own",
						Action: func(cCtx *cli.Context) error {
							apikey.ClearFile()
							return nil
						},
					},
				}},
		}}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
