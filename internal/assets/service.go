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
