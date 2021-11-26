package config

type Database struct {
	Host       string
	Port       string
	UserName   string
	Password   string
	SchemaName string
}

type DBConfig struct {
	Db Database `mapstructure:"db"`
}

type URLS struct {
	Name string
}

type URLSConfig struct {
	URLS URLS `mapstructure:"urls"`
}

type AppConfig struct {
	Port string
}
