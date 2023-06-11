package controllers

import (
	database "jdudp/dbtabases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client = database.InitDB()

type Data struct {
	Userid         string
	Groupid        string
	Nasname        string
	Ports          int
	Starospassword string
	Records        string
	tableId        string
}
type Datas struct {
	Userid         string
	Groupid        string
	Nasname        string
	Ports          int
	Starospassword string
	Records        string
}

func ErrorPage(c *gin.Context) {
	var title = c.Query("title")
	var href = c.Query("href")
	var err = c.Query("err")
	c.HTML(http.StatusOK, "error.html", gin.H{"title": title, "href": href, "err": err})
}

var rdb *redis.Client

func initRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "120.55.54.57:6379",
		Password: "zhutao@123",
		DB:       1,
	})
	_, err = rdb.Ping().Result()

	if err != nil {
		return err
	}
	return nil
}

func FindAll(context *gin.Context) {
	userid := context.PostForm("userid")
	nas := context.PostForm("nas")
	DataDb := client.Database("user_log").Collection(nas)
	filter := bson.M{"userid": userid, "nasname": nas}
	var findOptions = options.FindOptions{}
	findOptions.SetSort(bson.M{"_id": -1})
	res, err := DataDb.Find(context, filter, &findOptions)
	if err != nil {
		panic(err)
	}
	var persons []Datas
	err = res.All(context, &persons)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"code": 20000,
		"data": persons,
	})
}
