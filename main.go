package main

// Реализовать Cacher для Cache
// Методы вызываются из разных горутин и должны быть thread-safe
// Чистить кеш не надо, TTL задаётся в конструкторе
// По истечению TTL поведение такое, как будто значения в кеше не было записано

import (
	"MyCacheLib"
	"fmt"
	"time"
)

func main() { //демонстрирует работу кэша
	MyCache := MyCacheLib.NewCache(time.Duration(6*time.Second), 20) //время жизни - 6 секунд, размер - 20 значений. Можно задать любые.
	MyCache.Set("A", "1")
	fmt.Println(MyCache)
	time.Sleep(3 * time.Second)
	MyCache.Set("B", "2")
	fmt.Println(MyCache)
	time.Sleep(3 * time.Second)
	MyCache.Set("C", "3")
	fmt.Println(MyCache)
	time.Sleep(3 * time.Second)
	MyCache.Set("D", "4")
	fmt.Println(MyCache)
	time.Sleep(3 * time.Second)
	MyCache.Set("E", "5")
	fmt.Println(MyCache)
	time.Sleep(3 * time.Second)
	MyCache.Set("F", "6")
	fmt.Println(MyCache)
	time.Sleep(1 * time.Minute)
	fmt.Println(MyCache)
}
