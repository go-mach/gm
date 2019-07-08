package config

type (
	service struct {
		Group   string
		Name    string
		Version string
	}

	db struct {
		Type     string
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Log      bool
	}

	mgmtEndpoint struct {
		Port            int
		BaseRoutingPath string
	}

	mgmtHealth struct {
		Path string
		Full bool
	}
	management struct {
		Endpoint mgmtEndpoint
		Health   mgmtHealth
	}

	log struct {
		Path     string
		Filename string
		Console  struct {
			Enabled       bool
			DisableColors bool
			Colors        bool
		}
		Level           string
		JSON            bool
		MaxSize         int
		MaxBackups      int
		MaxAge          int
		Compress        bool
		LocalTime       bool
		TimestampFormat string
		FullTimestamp   bool
		ForceFormatting bool
	}

	api struct {
		Endpoint struct {
			Port            int
			BaseRoutingPath string
		}
		Security struct {
			Enabled bool
			Jwt     struct {
				Secret     string
				Expiration struct {
					Enabled bool
					Minutes int32
				}
			}
		}
	}

	// Ldap configuration
	Ldap struct {
		Base   string
		Host   string
		Port   int
		UseSSL bool
		Bind   struct {
			DN       string
			Password string
		}
		UserFilter  string
		GroupFilter string
		Attributes  []string
	}

	// Configuration describe the type for the configuration file
	Configuration struct {
		Service    service
		API        api
		DB         db
		Management management
		Log        log
		Ldap       Ldap
	}
)
