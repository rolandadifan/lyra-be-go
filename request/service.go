package request

type Service interface {
	CreateData(input RequestInput) (Request, error)
	FindData() ([]Request, error)
	DestroyData(input RequestDetail) (Request, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateData(input RequestInput) (Request, error) {
	request := Request{}
	request.Name = input.Name
	request.Request = input.Request

	newRequest, err := s.repository.Save(request)
	if err != nil {
		return newRequest, err
	}
	return newRequest, nil
}

func (s *service) FindData() ([]Request, error) {
	request, err := s.repository.Find()
	if err != nil {
		return request, err
	}
	return request, nil
}

func (s *service) DestroyData(input RequestDetail) (Request, error) {
	request, err := s.repository.Delete(input.ID)
	if err != nil {
		return request, err
	}
	return request, nil

}
