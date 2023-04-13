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

type CacheElem struct {
	value  string
	expire time.Time
}

type Cache struct {
	mu     sync.RWMutex
	ttl    time.Duration
	size   int
	values map[string]CacheElem
}

// NewCache инициализирует кэш
func NewCache(ttl time.Duration, Size int) *Cache {
	var c Cache
	c.ttl = ttl                           //время жизни
	c.values = make(map[string]CacheElem) //здесь будут значения
	c.size = Size                         //предельный размер
	go c.cleaner()                        //запускает процесс периодической уборки старых значений с периодом, равным заданному времени жизни элемента
	return &c
}

// Set заносит элемент в кэш. Если кэш переполнен, перед этим его данные предварительно очищаются
func (c *Cache) Set(Key string, Value string) {
	var elem CacheElem
	if len(c.values) == c.size {
		c.PurgeAll() //кэш очищается при переполнении
	}
	//c.PurgeExpired()
	c.mu.Lock()
	elem.value = Value
	expire := time.Now()       //берём время добавления элемента (текущее)
	expire = expire.Add(c.ttl) //прибавляем время жизни, заданное в конструкторе
	elem.expire = expire       //записываем срок годности элемента
	c.values[Key] = elem
	defer c.mu.Unlock()
}

// Get извлекает значение элемента из кэша и проверяет элемент на наличие
func (c *Cache) Get(Key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	k, Exist := c.values[Key]
	Value := k.value
	return Value, Exist
}

// Очищает данные кэша полностью
func (c *Cache) PurgeAll() {
	c.mu.Lock()
	for k := range c.values {
		delete(c.values, k)
	}
	defer c.mu.Unlock()
}

// Очищает данные кэша от старых значений
func (c *Cache) PurgeExpired() {
	c.mu.Lock()
	checktime := time.Now()
	for k := range c.values {
		e := c.values[k]
		t := e.expire
		fmt.Println(t)
		if t.Before(checktime) {
			delete(c.values, k)
		}
	}
	defer c.mu.Unlock()
}

func (c *Cache) cleaner() { //периодически чистит кэш от старых элементов, период берётся из параметров кэша
	for {
		time.Sleep(c.ttl)
		c.PurgeExpired()
	}
}

func main() { //Проверяет работу кэша
	MyCache := NewCache(time.Duration(6*time.Second), 20) //время жизни - 10 секунды, размер - 20 значений. Можно задать любые.
	MyCache.Set("A", "1")
	fmt.Println(MyCache.values)
	time.Sleep(3 * time.Second)
	MyCache.Set("B", "2")
	fmt.Println(MyCache.values)
	time.Sleep(3 * time.Second)
	MyCache.Set("C", "3")
	fmt.Println(MyCache.values)
	time.Sleep(3 * time.Second)
	MyCache.Set("D", "4")
	fmt.Println(MyCache.values)
	time.Sleep(3 * time.Second)
	MyCache.Set("E", "5")
	fmt.Println(MyCache.values)
	time.Sleep(3 * time.Second)
	MyCache.Set("F", "6")
	fmt.Println(MyCache.values)
}
