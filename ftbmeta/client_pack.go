package ftbmeta

type PackService service

func (s *PackService) GetPacks() ([]*PackInfo, error) {
	request, err := s.client.NewRequest("GET", "packs/")
	if err != nil {
		return nil, err
	}

	var response []*PackInfo
	_, err = s.client.Do(request, &response)
	return response, err
}

func (s *PackService) GetPack(pack string) (*Pack, error) {
	request, err := s.client.NewRequest("GET", "pack/" + pack + "/")
	if err != nil {
		return nil, err
	}

	var response Pack
	_, err = s.client.Do(request, &response)
	return &response, err
}

func (s *PackService) GetVersion(pack string, version string) (*Version, error) {
	request, err := s.client.NewRequest("GET", "pack/" + pack + "/" + version + "/")
	if err != nil {
		return nil, err
	}

	var response Version
	_, err = s.client.Do(request, &response)
	return &response, err
}
