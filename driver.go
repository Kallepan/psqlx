package psqlx

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

const driverName = "psqlx"

func init() {
	sql.Register(driverName, NewDriver())
}

type Driver struct {
	connector *Connector
}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	connector, err := d.OpenConnector(name)
	if err != nil {
		return nil, err
	}

	return connector.Connect(context.TODO())
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	cfg := &Config{}
	cfg.LoadFromStr(name)

	return NewConnector(cfg), nil
}
