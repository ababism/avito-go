package config

import (
	"avito-go/pkg/metrics"
	"avito-go/pkg/mylogger"
	"avito-go/pkg/xapp"
	"avito-go/pkg/xconfig"
	"avito-go/pkg/xdb/xpostgres"
	"avito-go/pkg/xhttp"
	"avito-go/pkg/xshutdown"
	"avito-go/pkg/xtracer"
	"avito-go/services/banner/internal/daemons/cacherefresher"
	"avito-go/services/banner/internal/repository/cache"
	"github.com/spf13/viper"
	"log"
	//"avito-go/services/banner/internal/repository/postgre"
)

type Config struct {
	App              *xapp.Config           `mapstructure:"app"`
	Http             *xhttp.Config          `mapstructure:"http"`
	Logger           *mylogger.Config       `mapstructure:"logger"`
	Metrics          *metrics.Config        `mapstructure:"metrics"`
	GracefulShutdown *xshutdown.Config      `mapstructure:"graceful_shutdown"`
	Tracer           *xtracer.Config        `mapstructure:"tracer"`
	CacheRefresher   *cacherefresher.Config `mapstructure:"cache_refresher"`
	Cache            *cache.Config          `mapstructure:"cache"`
	Postgres         *xpostgres.Config      `mapstructure:"postgres"`
}

func NewConfig(filePath string, appName string) (*Config, error) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file: %v", err)
	}

	// Загрузка конфигурации в структуру Config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error while unmarshalling config file: %v", err)
	}

	// Замена значений из переменных окружения, если они заданы
	xconfig.ReplaceWithEnv(&config, appName)
	return &config, nil
}
