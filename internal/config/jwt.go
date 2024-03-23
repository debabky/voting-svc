package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"time"
)

type JWTConfiger interface {
	JWTConfig() *JWTConfig
}

type JWTConfig struct {
	SecretKey             string        `fig:"secret_key,required"`
	AccessExpirationTime  time.Duration `fig:"access_expiration_time,required"`
	RefreshExpirationTime time.Duration `fig:"refresh_expiration_time,required"`
}

type jwt struct {
	once   comfig.Once
	getter kv.Getter
}

func NewJWTConfiger(getter kv.Getter) JWTConfiger {
	return &jwt{
		getter: getter,
	}
}

func (i *jwt) JWTConfig() *JWTConfig {
	return i.once.Do(func() interface{} {
		var result JWTConfig

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(i.getter, "jwt")).
			Please()
		if err != nil {
			panic(err)
		}

		return &result
	}).(*JWTConfig)
}
