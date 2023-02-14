package usecase

import (
	"awesomeProject/kofeslivka/pkg/entity"
)

type (
	Usecase interface {
		CreateUser(*entity.RawMaterial) (int64, error)
	}

	Repository interface {
		CreateUser(*entity.RawMaterial) (int64, error)
	}
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateUser(user *entity.RawMaterial) (int64, error) {
	uid, error := u.repository.CreateUser(user)
	return uid, error
}

/*func (u *usecase) DeleteUser(user *entity.DeleteUser) string {
	b := u.repository.DeleteUser(user)
	return b
}

func (u *usecase) GetFriends(a int) (b string, err error) {
	b, err = u.repository.GetFriends(a)
	return b, err
}

func (u *usecase) UpdateAge(user *entity.UpdateUser) string {
	s := u.repository.UpdateAge(user)
	return s
}*/
