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
	return addr
}
