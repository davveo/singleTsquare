package config

type DataBaseConfig struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"tsquare"`
	Password     string `default:"tsquare"`
	DatabaseName string `default:"tsquare"`
	MaxIdleConns int    `default:"10"`
	MaxOpenConns int    `default:"10"`
}

type SessionConfig struct {
}

type Config struct {
	Database      DataBaseConfig
	Session       SessionConfig
	ServerPort    int  `default:"8080"`
	IsDevelopment bool `default:"True"`
}
