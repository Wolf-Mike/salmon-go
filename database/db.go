package database

import "github.com/jmoiron/sqlx"

var (
	//DB is sqlx.DB对象
	DB *sqlx.DB
)

//Init 初始化数据库
func Init(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)

	if err != nil {
		return err
	}

	DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}
