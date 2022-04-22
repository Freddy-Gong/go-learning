package knowledege

import (
	"fmt"

	"github.com/go-redis/redis"
)

//Redis的用处
/*
1. cache缓存
2. 简单的队列
3. 排行榜
*/
//也是一个数据库连接池
var redisDb *redis.Client

func InitRedisDb() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//进行账号校验
	_, err = redisDb.Ping().Result()
	return
}

func Redis() {
	err := InitRedisDb()
	if err != nil {
		return
	}
	fmt.Println("连接redis成功")
}

//基本使用 Set Get的例子
func redisExample() {
	err := redisDb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val, err := redisDb.Get("score").Result()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := redisDb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exit")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

//zset 一个有序的集合
func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	//ZADD
	num, err := redisDb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	//把Golang的分数加10
	newScore, err := redisDb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
	//取分数最高的3个
	ret, err := redisDb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
