package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var client *mongo.Client

func InitDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://120.55.54.57:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions) // 连接数据库
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil) // 检查数据库是否连接成功
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var db *gorm.DB
var err error

func Connection() *gorm.DB {
	dsn := "backstage:Zhutao@908@tcp(123.57.148.218:3306)/feisuvpn?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var dbh *gorm.DB

func ConnectionH() *gorm.DB {
	dsn := "backstage:Zhutao@908@tcp(123.57.148.218:3306)/history?charset=utf8&parseTime=True&loc=Local"
	dbh, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return dbh
}
