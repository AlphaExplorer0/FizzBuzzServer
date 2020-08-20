package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Config config
type Config struct {
	BindIP   string
	BindPort int
}

// Init init
func (c *Config) Init() error {

	fmt.Println("Loading configuration")

	// DEFAULT

	// Rest server
	c.BindIP = "0.0.0.0"
	c.BindPort = 8080

	// Env variables - Allow configuration loading from ENV variables named FIZZBUZZ_${KEY}
	if err := decodeFromEnv(c, "fizzbuzz"); err != nil {
		return err
	}

	return nil
}

func decodeFromEnv(config interface{}, prefixes ...string) error {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	if configValue.Kind() != reflect.Struct {
		return errors.New("invalid config, should be struct")
	}

	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		var (
			envName     string
			fieldStruct = configType.Field(i)
			field       = configValue.Field(i)
		)
		if !field.CanAddr() || !field.CanInterface() {
			continue
		}

		// PREFIX_FIELD_VALUE
		envName = strings.ToUpper(strings.Join(append(prefixes, fieldStruct.Name), "_"))

		// Load From Shell ENV
		value, found := os.LookupEnv(envName)
		if found {
			if err := decodeField(field, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func getPrefixForStruct(prefixes []string, fieldStruct *reflect.StructField) []string {
	if fieldStruct.Anonymous {
		return prefixes
	}
	return append(prefixes, fieldStruct.Name)
}

func decodeField(field reflect.Value, value string) error {
	k := field.Kind()

	if k >= reflect.Int && k <= reflect.Uint64 {
		return decodeInt(value, field)
	}
	switch k {
	case reflect.String:
		field.Set(reflect.ValueOf(value))
		return nil
	default:
		return fmt.Errorf("Unsupported type")
	}
}

func decodeInt(strValue string, v reflect.Value) error {
	k := v.Kind()

	bitSize := 0 // 0 stands for reflect.Int and reflect.Uint
	switch k {
	case reflect.Int8, reflect.Uint8:
		bitSize = 8
	case reflect.Int16, reflect.Uint16:
		bitSize = 16
	case reflect.Int32, reflect.Uint32:
		bitSize = 32
	case reflect.Int64, reflect.Uint64:
		bitSize = 64
	}

	if k >= reflect.Int && k <= reflect.Int64 {
		i, err := strconv.ParseInt(strValue, 0, bitSize)
		if err != nil {
			return fmt.Errorf("unable to parse int value %q: %q", strValue, err.Error())
		}
		v.SetInt(i)
	} else if k >= reflect.Uint && k <= reflect.Uint64 {
		u, err := strconv.ParseUint(strValue, 0, bitSize)
		if err != nil {
			return fmt.Errorf("unable to parse uint value %q: %q", strValue, err.Error())
		}
		v.SetUint(u)
	}

	return nil
}
