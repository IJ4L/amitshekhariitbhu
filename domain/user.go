package domain

type User struct {
	ID       uint   `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}