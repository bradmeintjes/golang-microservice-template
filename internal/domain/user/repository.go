package user

type Storager interface {
	Store(user User) error
}

type Cacher interface {
	Cache(user User) error
}
