package config

// Configuration contains conectivity settings
type Configuration struct {
	Log struct {
		Level         string `toml:"level" default:"warning" comment:"Log level: trace, debug, info, warning, error, panic, and fatal"`
		JSONFormatter bool   `toml:"jsonformatter" default:"false" comment:"Allow to display logs in Json format if true"`
	} `toml:"Log" comment:"###############################\n Logs Settings \n##############################"`

	Briefly_public struct {
		REST struct {
			ListenPort string `toml:"listenPort" default:":8080" comment:"On which port REST HTTP service will listen"`
		} `toml:"REST" comment:"###############################\n REST API settings \n##############################"`

		Briefly struct {
			Adress string `toml:"adress" default:"localhost:5556" comment:"URL and Port for the Briefly gRPC server"`
		} `toml:"Briefly" comment:"###############################\n Briefly gRPC API settings \n##############################"`
	} `toml:"Briefly_public" comment:"###############################\n Briefly.public Settings \n##############################"`
}
