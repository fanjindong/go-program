package main

import (
	"fmt"
	"go-program/script/gpool"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/imroc/req"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	database           = "User"
	JPUSH_APPKEY       = "e63fbaa1cb2232d74bbfaae6"
	JPUSH_MASTERSECRET = "46fcc2ee1df583503f39dfb0"
	JPUSH_DEVICE_URL   = "https://device.jpush.cn"
	JPUSH_TOKEN        = "ZTYzZmJhYTFjYjIyMzJkNzRiYmZhYWU2OjQ2ZmNjMmVlMWRmNTgzNTAzZjM5ZGZiMA=="
	gPoolCount         = 50
)

var mongoURL string
var msgQueueURL string
var sumCount int
var count int32

type User struct {
	ID  bson.ObjectId `bson:"_id"` //在数据库中命名为_id，会代替默认生成的_id，为主键
	App struct {
		PushTag []string `bson:"pushTag"`
	} `bson:"app"`
	Birth string `bson:"birth"`
}

func pushMessage(uid int, gp *gpool.Pool) {
	c := atomic.AddInt32(&count, 1)
	defer gp.Done()
	deviceID := getJpushRegistrationID(uid)
	if deviceID == "" {
		fmt.Printf("%d/%d, uid: %d, deviceId: %s, resp: %s\n", c, sumCount, uid, deviceID, "")
		return
	}
	body := req.BodyJSON(map[string]interface{}{
		"queueName": "actor-user",
		"actorName": "init_user_push_tag",
		"args":      [3]interface{}{uid, deviceID, []string{}},
	})

	r, _ := req.Post(msgQueueURL, body)
	resp, _ := r.ToString()
	fmt.Printf("%d/%d, uid: %d, deviceId: %s, resp: %s\n", c, sumCount, uid, deviceID, resp)
}

func getJpushRegistrationID(uid int) string {
	url := fmt.Sprintf("%s/v3/aliases/%d?platform=android,ios", JPUSH_DEVICE_URL, uid)
	r, err := req.Get(url, req.Header{"Authorization": fmt.Sprintf("Basic %s", JPUSH_TOKEN)})
	if err != nil {
		return ""
	}
	result := make(map[string][]string)
	err = r.ToJSON(&result)
	if err != nil || len(result["registration_ids"]) == 0 {
		return ""
	}
	return result["registration_ids"][0]
}

func main() {
	env := os.Getenv("ENV")
	if env == "loc" {
		mongoURL = "mongodb://192.168.3.3:27017"
		msgQueueURL = "http://192.168.3.3:20001/v1/task"
	} else {
		mongoURL = "mongodb://root:Neoclub2018@dds-bp1fdc32bfa36a841.mongodb.rds.aliyuncs.com:3717,dds-bp1fdc32bfa36a842.mongodb.rds.aliyuncs.com:3717,dds-bp1fdc32bfa36a843.mongodb.rds.aliyuncs.com:3717,dds-bp1fdc32bfa36a844.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-10457077&readPreference=secondaryPreferred"
		msgQueueURL = "http://msg-queue.uki.im:20001/v1/task"
	}
	session, _ := mgo.Dial(mongoURL)
	collection := session.DB("uki").C("user")
	query := collection.Find(bson.M{"recentAt": bson.M{"$gte": time.Now().Unix() - 41*86400}})
	sumCount, _ = query.Count()

	iter := query.Select(bson.M{"_id": 1, "app": 1, "birth": 1}).Sort("recentAt").Iter()
	user := &User{}
	userIDArray := make([]int, sumCount)
	gpool := gpool.New(gPoolCount)
	for iter.Next(user) { //当遍历完结果后为false
		if len(user.App.PushTag) == 0 && user.Birth != "" {
			uid, _ := strconv.Atoi(user.ID.Hex())
			userIDArray = append(userIDArray, uid)
		} else {
			atomic.AddInt32(&count, 1)
		}
	}
	for uid := range userIDArray {
		if uid == 0 {
			fmt.Println(uid)
		}
		// gpool.Add(1)
		// go pushMessage(uid, gpool)
	}
	gpool.Wait()
}
