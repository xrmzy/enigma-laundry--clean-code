package manager

import (
	"database/sql"
	"enigma-laundry-clean-code/config"
	"fmt"
)

type InfraManager interface {
	Connection() *sql.DB
}

type infraManager struct {
	db   *sql.DB
	conf *config.Config
}

func (i *infraManager) openConnection() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.conf.Host, i.conf.Port, i.conf.User, i.conf.Password, i.conf.Name)

	db, err := sql.Open(i.conf.Driver, dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	i.db = db
	return nil
}

func (i *infraManager) Connection() *sql.DB {
	return i.db
}

func NewInfraManager(conf *config.Config) (InfraManager, error) {
	conn := &infraManager{conf: conf}
	if err := conn.openConnection(); err != nil {
		return nil, err
	}
	return conn, nil
}
