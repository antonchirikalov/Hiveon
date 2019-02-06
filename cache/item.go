package cache

import "time"

type Item struct {
	data    string
	expires time.Time
}

func (item *Item) expired() bool {
	return item.expires.Before(time.Now())
}
