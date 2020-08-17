package env

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

var initialized = false

type EnvVar struct {
	key          string
	desc         string
	defaultValue interface{}
}

func Var(key string) *EnvVar {
	if initialized == false {
		initialized = true
		Formatter := new(logrus.TextFormatter)
		Formatter.TimestampFormat = "02-01-2006 15:04:05"
		Formatter.FullTimestamp = true
		logrus.SetFormatter(Formatter)
	}
	return &EnvVar{key: key}
}

func (p *EnvVar) Print(v interface{}) {
	logrus.Infoln(p.desc+":", p.key, v)
}

func (p *EnvVar) Desc(val string) *EnvVar {
	p.desc = val
	return p
}

func (p *EnvVar) Default(val interface{}) *EnvVar {
	p.defaultValue = val
	return p
}

func (p *EnvVar) Min(val int) *EnvVar {
	if p.defaultValue == nil {
		p.defaultValue = val
	} else if p.defaultValue.(int) < val {
		p.defaultValue = val
	}
	return p
}

func (p *EnvVar) Max(val int) *EnvVar {
	if p.defaultValue == nil {
		p.defaultValue = val
	} else if p.defaultValue.(int) > val {
		p.defaultValue = val
	}
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
