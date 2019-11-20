package env

import (
	"os"
	"strconv"
	"time"
)

type EnvVar struct {
	key          string
	desc         string
	defaultValue interface{}
}

func Var(key string) *EnvVar {
	return &EnvVar{key: key}
}

func (p *EnvVar) Desc(val string) *EnvVar {
	p.desc = val
	return p
}

func (p *EnvVar) Default(val interface{}) *EnvVar {
	p.defaultValue = val
	return p
}

func (p *EnvVar) String() string {
	v := os.Getenv(p.key)
	if v != "" {
		return v
	}
	return p.defaultValue.(string)
}

func (p *EnvVar) Int() int {
	v, err := strconv.Atoi(os.Getenv(p.key))
	if err != nil {
		return v
	}
	return p.defaultValue.(int)
}

func (p *EnvVar) Bool() bool {
	v, err := strconv.ParseBool(os.Getenv(p.key))
	if err != nil {
		return v
	}
	return p.defaultValue.(bool)
}

func (p *EnvVar) Duration() time.Duration {
	value := os.Getenv(p.key)
	if value != "" {
		v, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return time.Duration(v)
		}
	}

	if p.defaultValue != nil {
		return time.Duration(p.defaultValue.(int))
	}

	return time.Duration(0)
}
