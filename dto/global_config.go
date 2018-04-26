package dto

type GlobalConfig struct {
	Log struct {
		LogPath string `yaml:"logpath"`
		LogName string `yaml:"logname"`
	}

	Mysql struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		UserName  string `yaml:"username"`
		PassWord  string `yaml:"password"`
		DataBase  string `yaml:"database"`
		Charset   string `yaml:"charset"`
		ParseTime string `yaml:"parseTime"`
	}
}

var globalConf *GlobalConfig

func GetGloabConf() *GlobalConfig {
	return globalConf
}

func SetGlobalConf(confVar *GlobalConfig) {
	globalConf = confVar
}
