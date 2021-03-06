package main

import (
	"database/sql"
	"flag"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"startones"
	"contents_feeder_redis/find_adv_phone"
	"contents_feeder_redis/find_all_img"
	"contents_feeder_redis/find_names"
	"contents_feeder_redis/find_regions"
	"contents_feeder_redis/find_all_keywords"
)

var siteFlag = flag.String("site", "", "must be test.com www.test.com")

func main() {
	flag.Parse() // Scan the arguments list

	golog, config := startones.Start()

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

	redisprotocol := config.Redis.Prot
	redishost := config.Redis.Host

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c.Close()

	site := *siteFlag

	if site != ""  {

		names := find_names.FindAll(golog, *db)
		phones := find_adv_phone.FindAll(golog, *db)
		regions := find_regions.FindAll(golog, *db)
		keywords := find_all_keywords.FindAllKeywords(golog,c,"fi_FI","porno")

		var ch find_all_img.Characters
		//
		ch.Find_all_img("/home/juno/git/imagehoster_redis/upload/img")

		ch.Add_name_phone_region_keyword(golog, names, phones, regions,keywords)

		ch.Find_age(golog, *db)

		ch.Find_free_paragraph(golog, c, "fi_FI", "porno")

		ch.Create_local_charters(golog, site)

	} else {

		fmt.Println("Need paramater --site=wwww.test.com ")

	}

}
