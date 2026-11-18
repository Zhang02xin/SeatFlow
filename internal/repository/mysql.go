package repository

import (
	"errors"
	"fmt"

	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMySQL(dsn string) (*gorm.DB, error) {
	cfg, err := mysqldriver.ParseDSN(dsn)
	if err != nil {
		return nil, errors.New("MYSQL_DSN 格式无效")
	}

	db, err := gorm.Open(mysql.New(mysql.Config{DSNConfig: cfg}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接 MySQL: %w", err)
	}
	return db, nil
}
