package user

import "fmt"

type UserService struct {
	users map[int]User
}

func NewUserService() UserService {
	return UserService{
		users: make(map[int]User),
	}
}

func (us *UserService) Add(u User) error {
	if _, ok := us.users[u.Id]; ok {
		return fmt.Errorf("user already exists")
	}

	us.users[u.Id] = u
	return nil
}

func (us *UserService) Get(id int) (User, error) {
	user, ok := us.users[id]

	if !ok {
		return user, fmt.Errorf("user %v does not exists", id)
	}

	return user, nil
}
