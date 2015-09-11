package config

type Config struct {
	User              string
	Password          string
	DefaultFormat     string
	Profiles          map[string]Profile
	DefaultDataCenter string
}

type Profile struct {
	User     string
	Password string
}
