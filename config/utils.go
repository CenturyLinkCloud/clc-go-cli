package config

func SetConfigPathFunc(f func() string) {
	GetClcHome = f
}
