package config

func SetConfigPathFunc(f func() (string, error)) {
	GetPath = f
}
