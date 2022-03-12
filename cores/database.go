package cores

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseMysql struct {
	DB *gorm.DB
}

func NewDatabaseMysql() *DatabaseMysql {

	return &DatabaseMysql{}
}

func (d *DatabaseMysql) ConnectDB() error {
	dsn := "root@tcp(localhost:3306)/belajar"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	d.DB = db

	if err != nil {
		return fmt.Errorf("could not connect to databse, %s", err.Error())
	}

	return nil
}

func (d *DatabaseMysql) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("could not close databse, %s", err.Error())
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("could not close databse, %s", err.Error())
	}

	return nil
}
