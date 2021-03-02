package subscribe

type Service interface {
	Create(inpput SubscribeInput) (Subscribe, error)
	Find() ([]Subscribe, error)
	DeleteSubscribe(input SubscribeDetail) (Subscribe, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(inpput SubscribeInput) (Subscribe, error) {
	subscribe := Subscribe{}
	subscribe.Email = inpput.Email

	newSubscribe, err := s.repository.Save(subscribe)
	if err != nil {
		return newSubscribe, err
	}
	return newSubscribe, nil
}

func (s *service) Find() ([]Subscribe, error) {
	getSubscribe, err := s.repository.FindAll()
	if err != nil {
		return getSubscribe, err
	}
	return getSubscribe, nil
}

func (s *service) DeleteSubscribe(input SubscribeDetail) (Subscribe, error) {
	subscribe, err := s.repository.Delete(input.ID)
	if err != nil {
		return subscribe, err
	}
	return subscribe, nil
}
