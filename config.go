package psqlx

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

type Config struct {
	Network  string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var rx = regexp.MustCompile(`postgres:\/\/(.*):(.*)@(.*):(\d+)\/(.*)\?sslmode=disable`)

func (c *Config) LoadFromStr(dsn string) error {
	sub := rx.FindStringSubmatch(dsn)
	if len(sub) != 6 {
		return errors.New("invalid connection string")
	}

	c.Network = "tcp"
	c.Host = sub[3]
	c.Port = sub[4]
	c.Username = sub[1]
	c.Password = sub[2]
	c.Database = sub[5]

	return nil
}

func (c *Config) LoadFromEnv() {
	c.Network = os.Getenv("DB_NETWORK")
	if c.Network == "" {
		c.Network = "tcp"
	}

	c.Host = os.Getenv("DB_HOST")
	if c.Host == "" {
		c.Host = "db"
	}

	c.Port = os.Getenv("DB_PORT")
	if c.Port == "" {
		c.Port = "5432"
	}

	c.Username = os.Getenv("DB_USERNAME")
	if c.Username == "" {
		c.Username = "postgres"
	}

	c.Password = os.Getenv("DB_PASSWORD")
	if c.Password == "" {
		c.Password = "postgres"
	}

	c.Database = os.Getenv("DB_DATABASE")
	if c.Database == "" {
		c.Database = "postgres"
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
