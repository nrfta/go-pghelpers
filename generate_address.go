package pghelpers

import "fmt"

// GenerateAddress returns a string with DB connection information
func (c *PostgresConfig) GenerateAddress() string {
	return c.generateAddress(false)
}
func (c *PostgresConfig) GenerateReadOnlyAddress() string {
	return c.generateAddress(true)
}

func (c *PostgresConfig) generateAddress(ro bool) string {
	var host string
	if ro {
		host = c.HostReadOnly
	}
	if host == "" {
		host = c.Host
	}

	addr := fmt.Sprintf("host=%s port=%d dbname=%s",
		host,
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
	return c.url(false)
}
func (c PostgresConfig) UrlReadOnly() string {
	return c.url(true)
}

func (c PostgresConfig) url(ro bool) string {
	var sslMode string
	if c.SSLEnabled {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	var host string
	if ro {
		host = c.HostReadOnly
	}
	if host == "" {
		host = c.Host
	}

	url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		c.Username, c.Password, host, c.Port, c.Database, sslMode)

	if c.ApplicationName != "" {
		url += fmt.Sprintf("&fallback_application_name=%s", c.ApplicationName)
	}

	return url
}
