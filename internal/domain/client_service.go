package domain

type ClientRepository interface {
	GetAllClients() ([]Client, error)
	CreateClient(client Client) error
}

type ClientService struct {
	repo ClientRepository
}

func NewClientService(repo ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) GetAllClients() ([]Client, error) {
	return s.repo.GetAllClients()
}

func (s *ClientService) CreateClient(client Client) error {
	return s.repo.CreateClient(client)
}
