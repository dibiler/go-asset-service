package main

import (
	"fmt"
	"os"

	"github.com/dibiler/go-asset-service/internal/assets"
	"github.com/dibiler/go-asset-service/pkg/utils"
)

func main() {
	data, err := utils.LoadJson(assets.Asset)("data.assets.json")

	if err != nil {
		panic(err)
	}

	service := assets.NewService(data)

	args := os.Args

	if len(args) > 1 {
		filter := args[1]
		result := service.FilterByType(filter)

		fmt.Println("Filtered Assets:")
		for _, a := range result {
			fmt.Printf("- %s (%s)\n", a.Name, a.Type)
		}
		return
	}

	all := service.getAll()

	fmt.Println("All Assets:")
	for _, a := range all {
		fmt.Printf("- %s (%s)\n", a.Name, a.Type)
	}

}
