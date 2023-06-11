package controllers

import (
	"context"
	database "jdudp/dbtabases"
	"jdudp/logics"
	"jdudp/models"
	"log"
	"net/http"
	"strconv"
	"strings"

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

// 处理日志
func Get() {
	var model []models.UserNas
	var tableIds = "2"
	var context = context.TODO()
	initRedisClient()
	if model, tableIdr, err := logics.NasList(model, tableIds); err != nil {
		log.Fatal(err)
	} else {
		for _, value := range model {
			for i := 0; i <= 500; i++ {
				// val := rdb.LRange(value.Nasname, 0, -1).Val()
				// fmt.Println(tableIdr)
				// fmt.Println(val)
				val, _ := rdb.LPop(value.Nasname).Result()
				if val != "" {
					comma := strings.Index(val, ":")
					pos := strings.Index(val[comma:], "message")
					vl1 := val[comma+pos:]
					posn := strings.Index(vl1, "prospector")
					if posn < 3 {
						valout := vl1[10:]
						if strings.Contains(valout, strconv.Itoa(value.Ports)) == true {
							DataDb := client.Database("user_log").Collection(value.Nasname)
							data := Data{Userid: value.Userid, Records: valout, Groupid: value.Groupid, Nasname: value.Nasname, Ports: value.Ports, Starospassword: value.Starospassword, tableId: tableIdr}
							_, err := DataDb.InsertOne(context, data)
							if err != nil {
								log.Fatal(err)
							}
						}
					} else {
						for _, value2 := range model {
							vl2 := vl1[:posn-3]
							valout := vl2[10:]
							if strings.Contains(valout, strconv.Itoa(value2.Ports)) == true && value2.Nasname == value.Nasname {
								DataDb := client.Database("user_log").Collection(value2.Nasname)
								data := Data{Userid: value2.Userid, Records: valout, Groupid: value2.Groupid, Nasname: value2.Nasname, Ports: value2.Ports, Starospassword: value2.Starospassword, tableId: tableIdr}
								_, err := DataDb.InsertOne(context, data)
								if err != nil {
									log.Fatal(err)
								}
							}
						}
					}
				} else {
					break
				}
			}
		}
	}
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
