package initial

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/douban-girls/server/app/schema"
	"github.com/go-redis/redis"
	"github.com/revel/revel"
)

// DB postgres instance
var DB *sql.DB

// Redis redis client instance
var Redis *redis.Client

func InitDB() {
	config := revel.Config
	username, _ := config.String("db.username")
	pwd, _ := config.String("db.pwd")
	dbname, _ := config.String("db.dbname")
	inDocker, _ := config.Bool("inDocker")
	revel.INFO.Println(inDocker)

	host := "localhost"
	if inDocker {
		host = "db"
	}

	dbPath := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, pwd, dbname)
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		panic(err)
	}
	log.Println(db)

	if _, err := db.Exec(schema.GetSchema()); err != nil {
		revel.INFO.Println(err.Error())
	}

	DB = db

	revel.INFO.Println("DB connected")
}

func InitRedis() {

	inDocker, _ := revel.Config.Bool("inDocker")
	host := "localhost"
	if inDocker {
		host = "redis"
	}

	host += ":6379"

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	revel.INFO.Println("redis connected")
	Redis = client
}
