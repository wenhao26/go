package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// 数据结构体
type Test struct {
	Id    string `bson:"_id"`
	Name  string `bson:"name"`
	Level int    `bson:"level"`
}

var db *mongo.Database           // database 话柄
var collection *mongo.Collection // collection 话柄

func ConnectToDB(uri, name string, timeout time.Duration) (*mongo.Database, error) {
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 返回 client
	return client.Database(name), nil
}

func AddOne(t *Test) {
	objId, err := collection.InsertOne(context.TODO(), &t)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("录入数据成功，objId:", objId)
}

func Del(m bson.M) {
	deleteResult, err := collection.DeleteOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
}

func EditOne(t *Test, m bson.M) {
	update := bson.M{"$set": t}
	updateResult, err := collection.UpdateOne(context.Background(), m, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

func Update(t *Test, m bson.M) {
	update := bson.M{"$set": t}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), m, update, updateOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

func Sectle(m bson.M) {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var t Test
		if err = cur.Decode(&t); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find name=primitive.Regex{xx}: ", t)
	}
	_ = cur.Close(context.Background())
}

func GetOne(m bson.M) {
	var one Test
	err := collection.FindOne(context.Background(), m).Decode(&one)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.FindOne: ", one)
}

func GetList(m bson.M) {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []*Test
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	_ = cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", all)
	for _, one := range all {
		log.Println("Id:", one.Id, " - name:", one.Name, " - level:", one.Level)
	}
}

func Count() {
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	log.Println("collection.CountDocuments:", count)
}

func main() {
	uri := "mongodb://127.0.0.1:27017"
	name := "test"
	maxTime := time.Duration(2)   // 连接超时时间
	table := "cs_community_event" // 表名

	db, err := ConnectToDB(uri, name, maxTime)
	if err != nil {
		panic("链接数据库有误!")
	}

	collection = db.Collection(table)

	t := Test{
		Id:    "1",
		Name:  "zngw",
		Level: 55,
	}

	// 添加一条数据
	AddOne(&t)

	// EditOne 编辑一条数据
	t.Name = "guoke"
	EditOne(&t, bson.M{"_id": 1})

	// 删除一条数据
	Del(bson.M{"_id": 1})

	// 更新数据 - 存在更新，不存在就新增
	Update(&t, bson.M{"_id": "1"})

	// Sectle 模糊查询
	Sectle(bson.M{"name": primitive.Regex{Pattern: "guo"}})

	// 准确搜索一条数据
	GetOne(bson.M{"name": "guoke"})

	// 统计collection的数据总数
	Count()

	// GetList 获取多条数据
	GetList(bson.M{"level": 55})
}
