package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type CookiesConfiger interface {
	CookiesConfig() *CookiesConfig
}

type CookiesConfig struct {
	Domain string `fig:"domain,required"`
	Secure bool   `fig:"secure,required"`
}

type cookies struct {
	once   comfig.Once
	getter kv.Getter
}

func NewCookiesConfiger(getter kv.Getter) CookiesConfiger {
	return &cookies{
		getter: getter,
	}
}

func (c *cookies) CookiesConfig() *CookiesConfig {
	return c.once.Do(func() interface{} {
		var result CookiesConfig
		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "cookies")).
			Please()

		if err != nil {
			panic(errors.WithMessage(err, "failed to figure out"))
		}

		return &result
	}).(*CookiesConfig)
}
