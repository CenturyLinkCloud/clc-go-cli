package cache_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/centurylinkcloud/clc-go-cli/autocomplete/cache"
	"github.com/centurylinkcloud/clc-go-cli/proxy"
)

func TestCache(t *testing.T) {
	proxy.Config()
	defer proxy.CloseConfig()

	// Touch empty cache.
	got, exist := cache.Get("key")
	if exist {
		t.Errorf("Invalid result\nExpected no options\nGot: %v", got)
	}

	// Lets write something.
	cache.LONG_AUTOCOMPLETE_REFRESH_TIMEOUT = 100
	heroes := []string{"John Snow", "Sam Tarly"}
	cache.Put("key", heroes)
	got, exist = cache.Get("key")
	if !exist {
		t.Errorf("Invalid result\nOptions have to exist but they don't")
	}
	if !reflect.DeepEqual(got, heroes) {
		t.Errorf("Invalid result\nExpected %v\nGot %v", heroes, got)
	}

	// Test that the items expire.
	cache.LONG_AUTOCOMPLETE_REFRESH_TIMEOUT = 1
	time.Sleep(time.Second)
	_, exist = cache.Get("key")
	if exist {
		t.Errorf("Invalid result - the cache entry has not expired but has been expected to")
	}
}
