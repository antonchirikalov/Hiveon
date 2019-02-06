package service

import (
	. "hiveon-api/model"
	. "hiveon-api/repository"
)

type UserService interface {
	GetUserWallet(fid int) []UserWallet
	SaveUserWallet(fid int, wallet string, coin string)
}

type userService struct {
 userRepository IUserRepository
}


func NewUserService() UserService{
	return &userService{userRepository:NewUserRepository()}
}

func (u *userService) GetUserWallet(fid int) []UserWallet {
	return u.userRepository.GetUserWallets(fid)
}

func (u *userService) SaveUserWallet(fid int, wallet string, coin string) {
	w:= UserWallet{Fid:fid, Wallet:wallet, Coin:coin}
	u.userRepository.SaveUserWallet(w)
	return
}



