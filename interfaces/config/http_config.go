package config

type HttpConfig struct {
    Port int
}

func LoadHttpConfig() HttpConfig {
    return HttpConfig{ Port: 8080 }
}
