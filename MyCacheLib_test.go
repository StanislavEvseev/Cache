package MyCacheLib

import (
	//"fmt"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	MyCache := NewCache(time.Duration(6*time.Second), 7) //время жизни - 6 секунд, размер - 7 значений. Можно задать любые.
	MyCache.Set("A", "1")
	time.Sleep(3 * time.Second)
	if !MyCache.cleaneractive {
		t.Error("!!!Cleaner goroutine doesn't started!!!", MyCache) //чистильщик не запустился
	}
	MyCache.Set("B", "2")
	time.Sleep(3 * time.Second)
	MyCache.Set("C", "3")
	time.Sleep(3 * time.Second)
	MyCache.Set("D", "4")
	time.Sleep(3 * time.Second)
	MyCache.Set("E", "5")
	time.Sleep(3 * time.Second)
	MyCache.Set("F", "6")
	if len(MyCache.values) == 0 {
		t.Error("!!!Cache empty!!!", MyCache) //значения не попадают в кэш
	}
	if len(MyCache.values) > 5 {
		t.Error("!!!Time deletion failed!!!", MyCache) //удаление по сроку годности не работает
	}
	MyCache.Set("G", "7")
	MyCache.Set("H", "8")
	MyCache.Set("I", "9")
	MyCache.Set("J", "10")
	if len(MyCache.values) > 3 {
		t.Error("!!!Cache overflow!!!", MyCache) //очистка по переполнению не работает
	}
	time.Sleep(10 * time.Second)
	if MyCache.cleaneractive {
		t.Error("!!!Cleaner goroutine doesn't stopped!!!", MyCache) //чистильщик не завершился
	}
}
