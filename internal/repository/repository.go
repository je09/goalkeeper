package repository

type Authorization interface {
	CreateUser(user User) error
	GetUser(username, password string) (User, error)
}

type ActivityJournal interface {
	Create() error
	GetAll()
	GetAttacks()
	GetUserActivities()
}
