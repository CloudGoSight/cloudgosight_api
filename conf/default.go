package conf

var RedisConfig = &redis{
	Network:  "tcp",
	Server:   "127.0.0.1:6379",
	Password: "",
	DB:       "0",
}

var DatabaseConfig = &datebase{
	Type:       "UNSET",
	Charset:    "utf8",
	DBFile:     "cloudgosight.db",
	Port:       3306,
	UnixSocket: false,
}

var SystemConfig = &system{
	Debug:       false,
	Mode:        "master",
	Listen:      ":5212",
	ProxyHeader: "X-Forwarded-For",
}

var CORSConfig = &cors{
	AllowedOrigins:   []string{"UNSET"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
	AllowedHeaders:   []string{"Cookie", "X-Cr-Policy", "Authorization", "Content-Type", "Content-Length", "X-Cr-Path", "X-Cr-FileName"},
	AllowCredentials: false,
	ExposedHeaders:   nil,
	SameSite:         "Default",
	Secure:           false,
}

var SlaveConfig = &slave{
	CallbackTimeout: 20,
	SignatureTTL:    60,
}

var SSLConfig = &ssl{
	Listen:   ":443",
	CertPath: "",
	KeyPath:  "",
}

var UnixConfig = &unix{
	Listen: "",
}

var OptionOverwrite = map[string]interface{}{}
