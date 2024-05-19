package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	fmt.Println("hello, fdataflow")

	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.NoExpiration)
	foo, ok := c.Get("foo")
	if ok {
		fmt.Println("foo=", foo)
	}
}
