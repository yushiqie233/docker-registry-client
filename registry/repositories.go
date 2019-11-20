package registry

type repositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func (registry *Registry) Repositories() ([]string, error) {
	url := registry.url("/v2/_catalog")
	var err error //We create this here, otherwise url will be rescoped with :=
	var response repositoriesResponse

	registry.Logf("registry.repositories url=%s", url)
	url, err = registry.getPaginatedJson(url, &response)
	switch err {
	case ErrNoMorePages:
		return response.Repositories, nil
	case nil:
		return response.Repositories, nil
	default:
		return nil, err
	}

}
