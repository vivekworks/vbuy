package vbuy

const (
    Production  = "prod"
    Development = "dev"
    QA          = "qa"
    Stage       = "stg"
)

type Config struct {
    HTTP   *HTTPConfig   `mapstructure:"http"`
    DB     *DBConfig     `mapstructure:"db"`
    Logger *LoggerConfig `mapstructure:"log"`
}

type HTTPConfig struct {
    Host    string `mapstructure:"host"`
    Port    string `mapstructure:"port"`
    Timeout struct {
        Idle       int `mapstructure:"idle"`
        Read       int `mapstructure:"read"`
        ReadHeader int `mapstructure:"readheader"`
        Write      int `mapstructure:"write"`
        Shutdown   int `mapstructure:"shutdown"`
    } `mapstructure:"timeout"`
}

type DBConfig struct {
    Name    string `mapstructure:"name"`
    Host    string `mapstructure:"host"`
    Port    int    `mapstructure:"port"`
    SslMode string `mapstructure:"sslmode"`
}

type LoggerConfig struct {
    Level string `mapstructure:"level"`
}
