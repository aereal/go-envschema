[![Build Status](https://travis-ci.org/aereal/go-envschema.png?branch=master)][travis]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/aereal/go-envschema?status.svg)][godoc]

# go-envschema

Validate the runtime environment against given [JSON schema][json-schema].

## Synopsis

```go
import (
	"github.com/aereal/go-envschema/loader"
	"github.com/xeipuuv/gojsonschema"
)

type config struct {
	Addr string `json:"LISTEN_ADDR"`
}

func run() error {
	// LISTEN_ADDR=localhost:8000

	loader, err := loader.New(gojsonschema.NewReferenceLoader("file://./config-schema.json"))
	if err != nil {
		return err
	}
	cfg := config{}
	if err := loader.Load(&cfg); err != nil {
		return err
	}
	// cfg.Addr // => "localhost:8000"

	return nil
}
```

## Motivation

[The Twelve-Factor app][twelve-factor-app] that is container-aware application pattern mainly written by Heroku says <q>the twelve-factor app stores config in environment variables</q>.

That pattern is almost alive to this days but that is lacking in validation system in comparison with common config file pattern.

So all we need is validating the environment variables against the schema and mapping it into runtime structure mechanism.

refs. https://this.aereal.org/entry/container-apps-and-json-schema (Blog entry in Japanese)

## Related works

- [go-playground/validator][go-playground-validator]: focused on validating struct, no environment mapping

[travis]: https://travis-ci.org/aereal/go-envschema
[license]: https://github.com/aereal/go-envschema/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/aereal/go-envschema
[twelve-factor-app]: https://12factor.net
[json-schema]: https://json-schema.org
[go-playground-validator]: https://github.com/go-playground/validator
