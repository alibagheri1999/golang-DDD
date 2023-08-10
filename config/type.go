package config

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	Mysql MysqlConfig `mapstructure:"mysql"`
}

type AppConfig struct {
	Name            string `mapstructure:"HOST_IP" validate:"required"`
	Env             string `mapstructure:"ENV" validate:"required"`
	Port            int    `mapstructure:"APP_PORT" validate:"required"`
	ApplyMigrations bool   `mapstructure:"APPLY_MIGRATION" validate:"required"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     string `mapstructure:"DB_PORT" validate:"required"`
	Username string `mapstructure:"DB_USER" validate:"required"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
}
