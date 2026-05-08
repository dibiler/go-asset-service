package assets

type Service struct {
	assets []Asset
}

func NewService(data []Asset) *Service {
	return &Service{assets: data}
}

func (s *Service) GetAll() []Asset {
	return s.assets
}

func (s *Service) FilterByType(assetType string) []Asset {
	var result []Asset

	for _, a := range s.assets {
		if a.Type == assetType {
			result = append(result, a)
		}
	}

	return result
}

func (s *Service) FilterByLocation(assetLocation string) []Asset {
	var result []Asset

	for _, a := range s.assets {
		if a.Location == assetLocation {
			result = append(result, a)
		}
	}
	return result
}
func (s *Service) FilterByName(assetName string) []Asset {
	var result []Asset

	for _, a := range s.assets {
		if a.Name == assetName {
			result = append(result, a)
		}
	}
	return result
}
func (s *Service) CountByType(assetType string) int {
	var count int = 0

	for _, a := range s.assets {
		if a.Type == assetType {
			count++
		}
	}
	return count
}
func (s *Service) FilterBy(predicate func(Asset) bool) []Asset {
	var result []Asset

	for _, a := range s.assets {
		if predicate(a) {
			result = append(result, a)
		}
	}

	return result
}
