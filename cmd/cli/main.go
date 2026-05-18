package main

import (
	"fmt"
	"os"

	"github.com/dibiler/go-asset-service/internal/assets"
	"github.com/dibiler/go-asset-service/internal/cli"
	"github.com/dibiler/go-asset-service/pkg/utils"
)

func main() {

	data, err := utils.LoadJSON[assets.Asset]("data/assets.json")

	if err != nil {
		fmt.Println("Failed to load assets:", err)
		os.Exit(1)
	}
	params, errs := cli.ParseParams(os.Args[1:])
	if len(errs) > 0 {
		for _, errItem := range errs {
			fmt.Println(errItem.Error())
		}
		os.Exit(1)
	}
	service := assets.NewService(data)
	outputList := cli.Run(params, service)

	for _, line := range outputList {
		fmt.Println(line)
	}

}
