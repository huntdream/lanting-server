package app

import (
	"database/sql"

	"github.com/huntdream/lanting-server/config"
)

// Global context
var (
	DB     *sql.DB
	Config config.Configuration
)
