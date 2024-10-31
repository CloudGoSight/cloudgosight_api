package conf

import (
	"github.com/CloudGoSight/cloudgosight_api/util"
	"github.com/go-ini/ini"
	"github.com/go-playground/validator/v10"
)

type system struct {
	// 注释的作用 https://blog.csdn.net/luduoyuan/article/details/130737221
	// validator https://segmentfault.com/a/1190000023725115
	Mode          string `validate:"eq=master|eq=slave"`
	Listen        string `validate:"required"`
	Debug         bool
	SessionSecret string
	HashIDSalt    string
	GracePeriod   int    `validate:"gte=0"`
	ProxyHeader   string `validate:"required_with=Listen"`
}

type ssl struct {
	CertPath string `validate:"omitempty,required"`
	KeyPath  string `validate:"omitempty,required"`
	Listen   string `validate:"required"`
}

type unix struct {
	Listen string
	Perm   uint32
}

type slave struct {
	Secret          string `validate:"omitempty,gte=64"`
	CallbackTimeout int    `validate:"omitempty,gte=1"`
	SignatureTTL    int    `validate:"omitempty,gte=1"`
}

type datebase struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBFile      string
	Port        int
	Charset     string
	UnixSocket  bool
}

type redis struct {
	Network  string
	Server   string
	User     string
	Password string
	DB       string
}

// 跨域配置
type cors struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	SameSite         string
	Secure           bool
}

const defaultConf = `[System]
Debug = false
Mode = master
Listen = :5212
SessionSecret = {SessionSecret}
HashIDSalt = {HashIDSalt}
`

var cfg *ini.File

func Init(path string) {

	if path == "" || !util.Exists(path) {
		// 新建
		confContent := util.Replace(map[string]string{
			"{SessionSecret}": util.RandStringRunes(64),
			"{HashIDSalt}":    util.RandStringRunes(64),
		}, defaultConf)
		f, err := util.CreateNestedFile(path)
		if err != nil {
			util.Log().Error("Create nested file error: %s", err)
		}

		//写文件
		_, err = f.WriteString(confContent)
		if err != nil {
			util.Log().Error("Write nested file error: %s", err)
		}
		f.Close()
	}

	var err error
	cfg, err = ini.Load(path)
	if err != nil {
		util.Log().Error("Load nested file error: %s", err)
	}

	sections := map[string]interface{}{
		"Database":   DatabaseConfig,
		"SSL":        SSLConfig,
		"Redis":      RedisConfig,
		"System":     SystemConfig,
		"UnixSocket": UnixConfig,
		"CORS":       CORSConfig,
		"Slave":      SlaveConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			util.Log().Error("Map section error %q: %s", sectionName, err)
		}
	}

	for _, key := range cfg.Section("OptionOverwrite").Keys() {
		OptionOverwrite[key.Name()] = key.Value()
	}

	if !SystemConfig.Debug {
		util.Level = util.LevelInfo
		util.GlobalLogger = nil
		util.Log()
	}
}

func mapSection(section string, confStruct interface{}) error {
	err := cfg.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}
	return nil
}
