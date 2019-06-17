package pghelpers

// PostgresConfig holds the information for connecting with a postgres database
type PostgresConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Database   string
	SSLEnabled bool `mapstructure:"ssl_enabled"`
	Tracing    PostgresTracingConfig
}

// PostgresTracingConfig is the config options for enabling tracing
type PostgresTracingConfig struct {
	// Enabled enables tracing on database calls
	Enabled bool

	// CreateRowsNextSpan will enable the creation of spans
	// on each rows.Next call
	CreateRowsNextSpan bool `mapstructure:"create_rows_next_span"`

	// CreateRowsCloseSpan will enable the creation of spans
	// on each rows.Close call
	CreateRowsCloseSpan bool `mapstructure:"create_rows_close_span"`

	// CreateRowsAffectedSpan will enable the creation of spans
	// on each rows affected call
	CreateRowsAffectedSpan bool `mapstructure:"create_rows_affected_span"`

	// CreateLastInsertedIDSpan will enable the creation of spans
	// on a last insert id query call
	CreateLastInsertedIDSpan bool `mapstructure:"create_last_inserted_id_span"`

	// AddQueryAttribute will enable the SQL query to be added to the
	// query span attributes
	AddQueryAttribute bool `mapstructure:"add_query_attribute"`

	// AddQueryParamsAttributes will enable the SQL query params to be added to
	// the query span attributes
	AddQueryParamsAttributes bool `mapstructure:"add_query_params_attributes"`
}
