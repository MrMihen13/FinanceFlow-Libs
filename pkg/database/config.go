package database

type ConnConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  bool
}
