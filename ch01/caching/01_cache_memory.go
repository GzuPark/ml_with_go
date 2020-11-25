package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	// 기본 만료 시간 5분, 30초마다 만료된 항목 제거하는 cache 생성
	c := cache.New(5*time.Minute, 30*time.Second)

	// cache에 (key, value) 넣음
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}
}
