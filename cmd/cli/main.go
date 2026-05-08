package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dibiler/go-asset-service/internal/assets"
	"github.com/dibiler/go-asset-service/pkg/utils"
)

var methods = [...]string{"list", "filter", "countByType", "filterBy"}

func main() {
	data, err := utils.LoadJSON[assets.Asset]("data/assets.json")

	if err != nil {
		fmt.Println("Failed to load assets:", err)
		os.Exit(1)
	}

	service := assets.NewService(data)

	args := os.Args

	if len(args) > 1 {
		method := args[1]

		switch method {
		case "filter":
			if len(args) < 3 {
				fmt.Println("Missing Type to filter by")
				return
			}
			filter := args[2]
			result := service.FilterByType(filter)

			if len(result) == 0 {
				fmt.Println("No Assets were found.")
				return
			}
			fmt.Println("Filtered Assets:")
			for _, a := range result {
				fmt.Printf("- %s (%s)\n", a.Name, a.Type)
			}
		case "filterBy":
			if len(args) < 4 {
				fmt.Println("Missing field to filter by or filtered value.")
				return
			}
			filterField := strings.ToLower(args[2])
			filterValue := args[3]
			var filterFunction func(assets.Asset) bool
			switch filterField {
			case "name":
				filterFunction = func(asset assets.Asset) bool {
					return asset.Name == filterValue
				}
			case "type":
				filterFunction = func(asset assets.Asset) bool {
					return asset.Type == filterValue
				}
			case "location":
				filterFunction = func(asset assets.Asset) bool {
					return asset.Location == filterValue
				}
			}
			result := service.FilterBy(filterFunction)

			if len(result) == 0 {
				fmt.Println("No Assets were found.")
				return
			}
			fmt.Println("Filtered Assets:")
			for _, a := range result {
				fmt.Printf("- %s (%s)\n", a.Name, a.Type)
			}
		case "list":
			all := service.GetAll()

			fmt.Println("All Assets:")
			for _, a := range all {
				fmt.Printf("- %s (%s)\n", a.Name, a.Type)
			}
		case "countByType":
			if len(args) < 3 {
				fmt.Println("Missing Type to Count by")
				return
			}
			assetsType := args[2]

			count := service.CountByType(assetsType)
			fmt.Printf("Number of assets with type [%s]: %d\n", assetsType, count)
		default:
			fmt.Println("The method introduce does not match any of the defined methods.")
		}
	} else {
		fmt.Println("Add a method, current available methods are:")
		for _, a := range methods {
			fmt.Printf("- %s\n", a)
		}
	}

}
