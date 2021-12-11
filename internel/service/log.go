package service

type CreateLogRequest struct {
	UserId  int
	Content string
}

func (svc *Service) CreateLog(params *CreateLogRequest) error {
	return svc.dao.CreateLog(params.UserId, params.Content)
}
