package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"os"
	"sync"
)

var ctx = context.Background()

func main() {
	file, err := os.Open("/Users/fanjindong/Downloads/FakeDau.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := csv.NewReader(file)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "r-bp11deher6a4iovqr5pd.redis.rds.aliyuncs.com:6379",
		Password: "Neoclub2018", // no password set
		DB:       0,             // use default DB
	})

	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go readAndWriter(wg, r, rdb)
	}
	wg.Wait()
}

func readAndWriter(wg *sync.WaitGroup, r *csv.Reader, rdb *redis.Client) {
	data := make(map[string]*[]interface{})
	length := 4069

	defer wg.Done()
	defer func() {
		for k, v := range data {
			if len(*v) > 0 {
				err := rdb.SAdd(ctx, k, *v...).Err()
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()
	_, _ = r.Read() // 第一行是标题
	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		uid, key, mark := record[0], record[1], record[2]
		// mark字段说明： 0代表正常的日活用户 1代表频道31的新增用户 其它的都是补的数据
		if mark == "1" || mark == "0" {
			continue
		}
		uids, ok := data[key]
		if !ok {
			value := make([]interface{}, 0, length)
			uids = &value
			data[key] = uids
		}

		*uids = append(*uids, uid)
		if len(*uids) > length {
			err := rdb.SAdd(ctx, key, *uids...).Err()
			if err != nil {
				panic(err)
			}
			delete(data, key)
			fmt.Println(key, "SAdd success")
		}
	}
}
