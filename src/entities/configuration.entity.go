package entities

type Configuration struct {
	Api      ApiConfiguration
	Database DatabaseConfiguration
	Jwt      JwtConfiguration
}
type ApiConfiguration struct {
	Environment string
	Port        int
	Debug       bool
}
type DatabaseConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}
type JwtConfiguration struct {
	Secret string
}
