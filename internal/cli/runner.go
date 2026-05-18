package cli

import (
	"fmt"
	"strings"

	"github.com/dibiler/go-asset-service/internal/assets"
)

func Run(cmd Request, service assets.AssetService) []string {
	var results []string
	switch cmd.method {
	case "filter":
		result := service.FilterBy(func(asset assets.Asset) bool {
			return asset.Type == cmd.value
		})

		if len(result) == 0 {
			results = append(results, "No Assets were found for type: "+cmd.value)
			break
		}
		results = append(results, "Filtered Assets:")
		results = append(results, ListFormatter(result)...)
	case "filterby":
		var filterFunction func(assets.Asset) bool
		switch cmd.filter {
		case "name":
			filterFunction = func(asset assets.Asset) bool {
				return asset.Name == cmd.value
			}
		case "type":
			filterFunction = func(asset assets.Asset) bool {
				return asset.Type == cmd.value
			}
		case "location":
			filterFunction = func(asset assets.Asset) bool {
				return asset.Location == cmd.value
			}
		default:
			results = append(results, fmt.Sprintf("filter parameter [%s] not in: name, location, type", cmd.filter))
			return results
		}
		result := service.FilterBy(filterFunction)

		if len(result) == 0 {
			results = append(results, fmt.Sprintf("No Assets were found for -filter=[%s] and -value=[%s]", cmd.filter, cmd.value))
			break
		}
		results = append(results, "Filtered Assets:")
		results = append(results, ListFormatter(result)...)

	case "list":
		all := service.GetAll()

		results = append(results, "All Assets:")
		results = append(results, ListFormatter(all)...)

	case "countbytype":
		if cmd.value == "" {
			results = append(results, "can't count an empty type. Please add value parameter")
			break
		}
		count := service.CountByType(cmd.value)
		results = append(results, fmt.Sprintf("Number of assets with type[%s]: %d", cmd.value, count))
	default:
		results = append(results, fmt.Sprintf("The -method introduced [%s] does not match any of the defined methods: %s.", cmd.method, strings.Join(allowedMethods, ", ")))
	}
	return results
}

func ListFormatter(list []assets.Asset) []string {
	var results []string
	for _, a := range list {
		results = append(results, fmt.Sprintf("- %s (%s)", a.Name, a.Type))
	}
	return results
}
