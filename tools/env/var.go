package env

import (
	"log"
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

func (p *EnvVar) Print(v interface{}) {
	log.Println("ENV:", p.key, v, "("+p.desc+")")
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
	if v == "" {
		v = p.defaultValue.(string)
	}
	p.Print(v)
	return v
}

func (p *EnvVar) Int() int {
	v, err := strconv.Atoi(os.Getenv(p.key))
	if err == nil {
		v = p.defaultValue.(int)
	}
	p.Print(v)
	return v
}

func (p *EnvVar) Bool() bool {
	v, err := strconv.ParseBool(os.Getenv(p.key))
	if err == nil {
		v = p.defaultValue.(bool)
	}
	p.Print(v)
	return v
}

func (p *EnvVar) Duration() time.Duration {
	value := os.Getenv(p.key)
	v := time.Duration(0)
	if value == "" {
		if p.defaultValue != nil {
			v = time.Duration(p.defaultValue.(int))
		}
	} else {
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			v = time.Duration(i)
		}
	}
	p.Print(v)
	return v
}
