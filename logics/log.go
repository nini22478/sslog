package logics

import (
	"context"
	"jdudp/models"
	"log"
	"strconv"
	"strings"

	database "jdudp/dbtabases"

	"github.com/go-redis/redis"
)

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

var rdb *redis.Client

var client = database.InitDB()

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
	if model, tableIdr, err := NasList(model, tableIds); err != nil {
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
