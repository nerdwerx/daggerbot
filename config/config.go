package config

const Version = "0.1.0"

var (
	Debug  bool
	Guilds map[string]Guild
	Prefix string
	Token  string
)
