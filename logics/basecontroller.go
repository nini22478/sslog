package logics

import (
	databases "jdudp/dbtabases"
	"jdudp/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = databases.Connection()
var dbh = databases.ConnectionH()
var Eloquent *gorm.DB
var err error

func OpenMySQL(username, password, host, port, database string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	// db, err := sql.Open("mysql", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func OpenDb(Id string) *gorm.DB {
	var model []models.History
	var MysqlUser string
	var MysqlPassword string
	var MysqlDns string
	var MysqlPorts string
	var MysqlFrom string
	// Id, _ = strconv.Atoi(Id)
	dbh.Table("history.history_product").Where("id = ?", Id).Find(&model)
	for _, value := range model {
		MysqlUser = value.MysqlUser
		MysqlPassword = value.MysqlPassword
		MysqlDns = value.MysqlDns
		MysqlPorts = strconv.Itoa(value.MysqlPorts)
		MysqlFrom = value.MysqlFrom
	}
	dbs := OpenMySQL(MysqlUser, MysqlPassword, MysqlDns, MysqlPorts, MysqlFrom)
	return dbs
}
