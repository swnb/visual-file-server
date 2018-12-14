package config

// Conf define all methods that other module require
type Conf interface {
	Get(string) interface{}
	GetString(string) string
	GetBool(string) bool
	GetInt(string) int
	GetFloat64(string) float64
	IsSet(key string) bool
}

var conf Conf

func init() {
	// init config instance for easy use
	conf = Config{}
}

// Config implement Conf ;; any struct implement this Conf can replace it
type Config struct {
}

// Get return the config value with unknown type
func Get(key string) interface{} { return conf.Get(key) }
func (Config) Get(key string) interface{} {
	value := selfConfig.Get(key)
	if value != nil {
		return value
	}
	return baseConfig.Get(key)
}

// GetString return the config value with default type string
func GetString(key string) string { return conf.GetString(key) }
func (Config) GetString(key string) string {
	value := selfConfig.GetString(key)
	if value != "" {
		return value
	}
	return baseConfig.GetString(key)
}

// GetBool return the config value with default type bool
func GetBool(key string) bool { return conf.GetBool(key) }
func (Config) GetBool(key string) bool {
	if selfConfig.IsSet(key) {
		return selfConfig.GetBool(key)
	}
	return baseConfig.GetBool(key)
}

// GetInt return the config value with default type int
func GetInt(key string) int { return conf.GetInt(key) }
func (Config) GetInt(key string) int {
	value := selfConfig.GetInt(key)
	if value != 0 {
		return value
	}
	return baseConfig.GetInt(key)
}

// GetFloat64 return the config value with default type float64
func GetFloat64(key string) float64 { return conf.GetFloat64(key) }
func (Config) GetFloat64(key string) float64 {
	value := selfConfig.GetFloat64(key)
	if value != 0 {
		return value
	}
	return baseConfig.GetFloat64(key)
}

// IsSet  whether config key is set or not
func IsSet(key string) bool { return conf.IsSet(key) }
func (Config) IsSet(key string) bool {
	if !selfConfig.IsSet(key) {
		return selfConfig.IsSet(key)
	}
	return true
}

// New create new config instance
func New() Conf {
	return Config{}
}
