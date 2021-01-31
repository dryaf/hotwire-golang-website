package flags

import "github.com/alecthomas/kong"

// API api related flags passing in env variables
type API struct {
	Version kong.VersionFlag
	AppName string `help:"Stage the name of the service." env:"APP_NAME"`
	Stage   string `help:"Stage the software is deployed." env:"STAGE"`
	Branch  string `help:"Branch used to build software." env:"BRANCH"`
}

// ServerAPI server api related flags passing in env variables
type ServerAPI struct {
	API
	Port     string `help:"Port number to bind our TLS listener." env:"PORT" default:"9443"`
	CertFile string `help:"Certificate used to bind our TLS listener." env:"CERT_FILE" default:".certs/hotwire.localhost.pem"`
	KeyFile  string `help:"Private Key used to bind our TLS listener." env:"KEY_FILE" default:".certs/hotwire.localhost.key"`
	Local    bool   `env:"LOCAL"`
}
