package main

import (
	"fmt"
	"log"

	"golang.org/x/sync/singleflight"
)

var getTokenGroup singleflight.Group

type GetTokenTask struct {
	Region   string
	Business string
	callback func() (interface{}, error)
}

func (t *GetTokenTask) key() string {
	return fmt.Sprintf("%s|%s", t.Region, t.Business)
}

func (t *GetTokenTask) Do() string {
	key := t.key()
	v, err, _ := getTokenGroup.Do(key, t.callback)
	if err != nil {
		log.Printf("[GetTokenTask] [%s] get token err: %v", key, err)
		return ""
	}
	token, ok := v.(string)
	if !ok {
		log.Printf("[GetTokenTask] [%s] convert token to string err", key)
		return ""
	}
	log.Printf("[GetTokenTask] [%s] got token: %s", key, token)
	return token
}

func newGetTokenTask(region string, business string) *GetTokenTask {
	return &GetTokenTask{
		Region:   region,
		Business: business,
		callback: func() (interface{}, error) {
			return getToken(region, business)
		},
	}
}
