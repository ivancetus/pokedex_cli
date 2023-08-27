package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Millisecond)
	if c.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	c := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
	}

	for _, cas := range cases {
		c.Add(cas.inputKey, cas.inputVal)
		actual, ok := c.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("value not match, %v vs %v", string(actual), string(cas.inputVal))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond*2)

	_, ok := c.Get(keyOne)
	if ok {
		t.Errorf("%s, should have been reaped", keyOne)
	}

}
