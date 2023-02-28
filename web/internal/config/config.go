package config

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "12345"
  dbname   = "igcms"
)

type Config struct {
  Host string
  Port int32
  User string
  Password string
  Dbname string
}

func NewConfig() *Config{
  return &Config{
    Host: host,
    Port: port,
    User: user,
    Password: password,
    Dbname: dbname,
  }
}
