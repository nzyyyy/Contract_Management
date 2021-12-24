package service

type CreateLogRequest struct {
	UserId  int
	Content string
}
type DeleteLogRequest struct {
	UserId int
	LogId  int
}

func (svc *Service) CreateLog(params *CreateLogRequest) error {
	return svc.dao.CreateLog(params.UserId, params.Content)
}

func (svc *Service) DeleteLog(params *DeleteLogRequest) error {
	return svc.dao.DeleteLog(params.LogId)
}
