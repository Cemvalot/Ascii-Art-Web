package config

type Config struct {
	Port         string
	TemplatePath string
	BannerPath   string
	StaticPath   string
}

func Conf() *Config {
	return &Config{
		Port:         ":8080",
		TemplatePath: "templates/index.html",
		BannerPath:   "banners",
		StaticPath:   "static",
	}
}
