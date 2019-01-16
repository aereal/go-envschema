package envloader

import (
	"os"
	"strings"

	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
)

type envLoader struct {
	goLoader gojsonschema.JSONLoader
}

func New() gojsonschema.JSONLoader {
	envs := getEnvs()
	return &envLoader{
		goLoader: gojsonschema.NewGoLoader(envs),
	}
}

func (l *envLoader) JsonSource() interface{} {
	return l.goLoader.JsonSource()
}

func (l *envLoader) LoadJSON() (interface{}, error) {
	return l.goLoader.LoadJSON()
}

func (l *envLoader) JsonReference() (gojsonreference.JsonReference, error) {
	return l.goLoader.JsonReference()
}

func (l *envLoader) LoaderFactory() gojsonschema.JSONLoaderFactory {
	return l.goLoader.LoaderFactory()
}

func getEnvs() map[string]string {
	return environAsMap(os.Environ())
}

func environAsMap(pairs []string) map[string]string {
	envs := map[string]string{}
	for _, env := range pairs {
		pair := strings.SplitN(env, "=", 2)
		envs[pair[0]] = pair[1]
	}
	return envs
}
