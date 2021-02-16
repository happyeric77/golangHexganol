package main

import (
	"project5/serializer/msgpack"
	"project5/serializer/json"
	"testing"
	"project5/repository/redis"
	"project5/repository/mongo"
	"project5/shortener"
	"time"
)

func TestRedis(t *testing.T){
	repo, err := redis.NewRedisRepository("redis://127.0.0.1:6379")
	if err != nil {
		t.Log(err)
	}

	dummyRedirect := shortener.Redirect{
		Code: "redisdummyTest0",
		URL: "redisdummy.com",
		CreatedAt: time.Now().Unix(),
	}

	err = repo.Store(&dummyRedirect)
	if err != nil {
		t.Fatal(err)
	}

	foundRedirect, err := repo.Find("redisdummyTest0")
	if err!= nil {
		t.Fatal(err)
	}

	t.Log(foundRedirect.URL, foundRedirect.Code, foundRedirect.CreatedAt)

}

func TestMongo(t *testing.T){
	repo, err := mongo.NewMongoRepository("mongodb://127.0.0.1:27017/", "dummy", 3)
	if err != nil {
		t.Log(err)
	}

	dummyRedirect := shortener.Redirect{
		Code: "mongodummyTest0",
		URL: "mongodummy.com",
		CreatedAt: time.Now().Unix(),
	}

	err = repo.Store(&dummyRedirect)
	if err != nil {
		t.Fatal(err)
	}

	foundRedirect, err := repo.Find("mongodummyTest0")
	if err!= nil {
		t.Fatal(err)
	}

	t.Log(foundRedirect.URL, foundRedirect.Code, foundRedirect.CreatedAt)

}

func TestJsonSierializer(t *testing.T){
	dummyRedirect := shortener.Redirect{
		Code: "jsondummyTest0",
		URL: "jsondummy.com",
		CreatedAt: time.Now().Unix(),
	}
	r := json.Redirect{}

	encoded, err := r.Encode(&dummyRedirect)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Encoded: ", encoded)

	decoded, err := r.Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Decoded: ", decoded)
}

func TestMsgpackSierializer(t *testing.T){
	dummyRedirect := shortener.Redirect{
		Code: "msgpackdummyTest0",
		URL: "msgpackdummy.com",
		CreatedAt: time.Now().Unix(),
	}
	r := msgpack.Redirect{}

	encoded, err := r.Encode(&dummyRedirect)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Encoded: ", encoded)

	decoded, err := r.Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Decoded: ", decoded)
}