package user

type Storage interface {
	Store(user User) error
	All() ([]User, error)
}

type Cacher interface {
	Cache(user User) error
}
