package user

type User struct {
	Name     string;
	Age      int32;
	Location string;
}

func (user User) GetUserName() string {
	return user.Name
}