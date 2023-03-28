package usecase

import (
	"awesomeProject/kofeslivka/pkg/entity"
)

type (
	Usecase interface {
		AddIngredient(*entity.RawMaterial) (int64, error)
		MakeDesert(desert *entity.Desert) (string, error)
	}

	Repository interface {
		AddIngredient(*entity.RawMaterial) (int64, error)
		MakeDesert(desert *entity.Desert) (string, error)
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

func (u *usecase) AddIngredient(user *entity.RawMaterial) (int64, error) {
	uid, error := u.repository.AddIngredient(user)
	return uid, error
}
func (u *usecase) MakeDesert(des *entity.Desert) (string, error) {
	uid, error := u.repository.MakeDesert(des)
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
