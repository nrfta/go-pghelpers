package pghelpers

// PostgresConfig holds the information for connecting with a postgres database
type PostgresConfig struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	SSLEnabled      bool   `mapstructure:"ssl_enabled"`
	MigrationsTable string `mapstructure:"migrations_table"`
}
