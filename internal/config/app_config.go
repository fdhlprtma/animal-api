package config

type AppConfig struct {
	AppEnv         string
	AllowedOrigins []string
}

var Config = AppConfig{
	// Pilihan: "development", "strict", "super_strict"
	AppEnv: "strict",

	AllowedOrigins: []string{
		"http://localhost",
	},
}
