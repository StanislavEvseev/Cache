package main

// Реализовать Cacher для Cache
// Методы вызываются из разных горутин и должны быть thread-safe
// Чистить кеш не надо, TTL задаётся в конструкторе
// По истечению TTL поведение такое, как будто значения в кеше не было записано

import (
	"fmt"
	"sync"
	"time"
)

type Cacher interface {
	Get(string) (string, bool)
	Set(string, string)
}

type Cache struct {
	mu     sync.RWMutex
	ttl    time.Duration
	size   int
	values map[string]string
}

// NewCache инициализирует кэш
func NewCache(ttl time.Duration, Size int) *Cache {
	var c Cache
	c.ttl = ttl                        //время жизни
	c.values = make(map[string]string) //здесь будут значения
	c.size = Size                      //предельный размер
	return &c
}

// Set заносит элемент в кэш. Если кэш переполнен, при этом его данные предварительно очищаются
func (c *Cache) Set(Key string, Value string) {
	if len(c.values) == c.size {
		c.Purge()                      //кэш очищается при переполнении
		time.AfterFunc(c.ttl, c.Purge) //кэш также очищается через заданное время независимо от наполнения
	}
	c.mu.Lock()
	c.values[Key] = Value
	defer c.mu.Unlock()
}

// Get извлекает значение элемента из кэша и проверяет элемент на наличие
func (c *Cache) Get(Key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	Value, Exist := c.values[Key]
	return Value, Exist
}

// Очищает данные кэша
func (c *Cache) Purge() {
	c.mu.Lock()
	for k := range c.values {
		delete(c.values, k)
	}
	defer c.mu.Unlock()
}

func main() { //Проверяет работу кэша
	MyCache := NewCache(time.Duration(3*time.Second), 3) //время жизни - 3 секунды, размер - 3 значения. Можно задать любые.
	MyCache.Set("A", "1")
	fmt.Println(MyCache.values)
	MyCache.Set("B", "2")
	fmt.Println(MyCache.values)
	MyCache.Set("C", "3")
	fmt.Println(MyCache.values)
	MyCache.Set("D", "4") //здесь кэш должен очиститься по переполнению
	fmt.Println(MyCache.values)
	time.Sleep(4 * time.Second) //здесь кэш должен очиститься по таймеру
	fmt.Println(MyCache.values)
}
