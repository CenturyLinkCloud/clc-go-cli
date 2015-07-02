package config

func SetConfigPathFunc(f func() (string, error)) {
	getConfigPath = f
}
