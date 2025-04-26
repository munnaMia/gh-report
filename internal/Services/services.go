package services

type GhService struct{}

func (g GhService) Fetch(userName string) (string, error) {
	return "", nil
}

func (g GhService) Show() {}
