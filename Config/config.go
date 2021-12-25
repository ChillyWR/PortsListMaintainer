package config

var DomainPort = "9999"
var ClientPort = "10000"

const (
	DBDriver = "postgres"
	Host     = "localhost"
	DBPort   = 5432
	User     = "postgres"
	Password = "24130677Postgres"
	DBName   = "portsDB"
)

type Port struct {
	ID         int64
	Name       string
	IsActive   bool
	Company    string
	Email      string
	Phone      string
	Address    string
	About      string
	Registered string
	Latitude   float32
	Longitude  float32
}
