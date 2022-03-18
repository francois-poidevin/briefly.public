package config

// Configuration contains conectivity settings
type Configuration struct {
	Log struct {
		Level string `toml:"level" default:"warn" comment:"Log level: debug, info, warn, error, dpanic, panic, and fatal"`
	} `toml:"Log" comment:"###############################\n Logs Settings \n##############################"`

	Briefly_public struct {
		Bbox string `toml:"bbox" default:"43.52,1.32^43.70,1.69" comment:"tracking bbox (Lat/Lon)"`
	} `toml:"Briefly_public" comment:"###############################\n Briefly.public Settings \n##############################"`
}
