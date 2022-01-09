package config

const (
	// Version roughly represents the applications current version.
	Version = "0.1"
	// Title is the name of the application used throughout the configuration process.
	Title = "bobbago"
)

var (
	// GenConfig when toggled will cause, upon initialization to write its default config to the cwd and then exit.
	GenConfig = false
	// NoColor when true will disable the banner and any colored console output.
	NoColor bool
)

// "habbo"
var (
	HabboBind string
	HabboPort int
	MaxConns  int
)

// "http"
var (
	// HTTPBind is defined via our toml configuration file. It is the address that our HTTP server listens on.
	HTTPBind string
	// HTTPPort is defined via our toml configuration file. It is the port that our HTTP server listens on.
	HTTPPort string
)

// "data"
var (
	// DataDirectory is where we store our database files.
	DataDirectory string
)
