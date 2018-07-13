// Package config provides a way to take configuration from the environment and
// produce a configuration struct for the application to consume.  Only the run
// package and maybe test files should import this package.
//
// The Values struct in this package is the single source of truth of all
// configuration that comes out of the environment into the application.  No
// other part of the system should look at environment variables or config files
// or CLI flags.
package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

const (
	// This defines the filename (without extension) for the config file that
	// viper will read.
	configFilenamePrefix = "my-server" // TODO: update
	// This defines the prefix on environment variables that viper will read.
	configEnvVarPrefix = "MYSERVER" // TODO: update
)

// Values defines all the values required to configure the server. This
// struct defines how configuration values are read from the environment.
type Values struct {
	// TODO: add configuration values here.

	// Tip: you can add unnamed struct types inline for ease of reading:
	//
	// Database struct {
	//     Port int
	// }
	//
	// The above can then be configured as MYSERVER_DATABASE_PORT in the
	// environment or as expected to Unmarshal from a configuration file.
}

// configPaths defines where viper looks for config files for the application.
// You may want to update this depending on where you want the app to look for
// configuration files.
func configPaths(v *viper.Viper) {
	v.AddConfigPath(".")
	v.AddConfigPath("/etc/my-server/") // TODO: update
	if home := os.Getenv("HOME"); home != "" {
		v.AddConfigPath(home) // this is mainly for developers for running tests.
	}

}

// Init initializes the configuration for the application.
func Init(v *viper.Viper) *Values {
	return initVals(v, configFilenamePrefix, configEnvVarPrefix)
}

func initVals(v *viper.Viper, filename, envPrefix string) *Values {
	if debug, _ := strconv.ParseBool((os.Getenv("VIPER_DEBUG"))); debug {
		jww.SetLogThreshold(jww.LevelDebug)
		jww.SetLogOutput(os.Stderr)
	}
	v.SetEnvPrefix(envPrefix)
	v.SetConfigName(filename)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	configPaths(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("can't read config file: %v", err))
		}
	}
	vals := &Values{}
	bindEnvs(v, *vals)
	if err := v.Unmarshal(vals); err != nil {
		panic(fmt.Errorf("viper failed to unmarshal environment into struct: %v", err))
	}
	return vals
}

// bindEnvs iterates through a struct's fields and calls viper.BindEnv with each
// field's name.  This is unfortunately necessary because viper's Unmarshal
// method does not respect the AutomaticEnv flag, and thus only unmarshals
// environment variables that have been first bound.
// see https://github.com/spf13/viper/issues/188
func bindEnvs(v *viper.Viper, strct interface{}, parts ...string) {
	ifv := reflect.ValueOf(strct)
	ift := reflect.TypeOf(strct)
	for i := 0; i < ift.NumField(); i++ {
		fieldv := ifv.Field(i)
		t := ift.Field(i)
		name := strings.ToLower(t.Name)

		// this is a viperism, the struct tag label for defining what
		// environment variable is used by viper's Unmarshal.
		tag, ok := t.Tag.Lookup("mapstructure")
		if ok {
			name = tag
		}
		path := append(parts, name)
		switch fieldv.Kind() {
		case reflect.Struct:
			bindEnvs(v, fieldv.Interface(), path...)
		default:
			v.BindEnv(strings.Join(path, "."))
		}
	}
}
