package service

func (svc *Service) RemoveUserRole(userId, roleId int) error {
	return svc.dao.RemoveUserRole(userId, roleId)
}
