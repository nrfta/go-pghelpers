package pghelpers

import "fmt"

// GenerateAddress returns a string with DB connection information
func (c *PostgresConfig) GenerateAddress() string {
	addr := fmt.Sprintf("host=%s port=%d dbname=%s",
		c.Host,
		c.Port,
		c.Database,
	)
	if c.Username != "" {
		addr = addr + fmt.Sprintf(" user=%s", c.Username)
	}
	if c.Password != "" {
		addr = addr + fmt.Sprintf(" password=%s", c.Password)
	}

	if c.SSLEnabled {
		addr = addr + " sslmode=require"
	} else {
		addr = addr + " sslmode=disable"
	}

	if c.ApplicationName != "" {
		addr = addr + fmt.Sprintf(" application_name=%s", c.ApplicationName)
	}
	return addr
}

// URL returns a 'postgres://user:pass@host:port/database?sslmode=' style address
func (c PostgresConfig) URL() string {
	var sslMode string
	if c.SSLEnabled {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, sslMode)
}
