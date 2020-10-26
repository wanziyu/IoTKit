package main

import (
	"IoTKit/timedb/influx"
	"fmt"
	"log"
)

func main() {
	conn := influx.ConnInflux()
	fmt.Println(conn)

	// insert
	influx.WritesPoints(conn)

	// 获取10条数据并展示
	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, err := influx.QueryDB(conn, qs)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range res[0].Series[0].Values {
		for j, value := range row {
			log.Printf("j:%d value:%v\n", j, value)
		}
	}
}