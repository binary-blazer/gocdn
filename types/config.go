package types

type Config struct {
	Webserver 	Webserver `json:"webserver"`
	Database  	Database  `json:"database"`
	UploadTypes []string	`json:"uploadTypes"`
	BigTypes    []string	`json:"bigTypes"`
	MaxUploadSize  int	`json:"maxUploadSize"`
}

type Webserver struct {
	Port int	`json:"port"`
}

type Database struct {
	Host     string	`json:"host"`
	Port     int	`json:"port"`
	User     string	`json:"user"`
	Password string	`json:"password"`
	Database string	`json:"database"`
}