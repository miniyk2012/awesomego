package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	RegionAmerica = "America"
	RegionEurope  = "Europe"
	RegionAsia    = "Asia"
	RegionAfrica  = "Africa"
)

var mpRegionWaitTime = map[string]time.Duration{
	RegionAmerica: 5 * time.Second,
	RegionEurope:  3 * time.Second,
	RegionAfrica:  4 * time.Second,
	RegionAsia:    2 * time.Second,
}

func GetToken(region string, business string) (string, error) {
	waitTime, ok := mpRegionWaitTime[region]
	if !ok || waitTime == 0 {
		return "", errors.New("unsupported region: " + region)
	}
	log.Printf("[getToken] region: %s, business: %s, wait-time: %v", region, business, waitTime)
	time.Sleep(waitTime)
	return fmt.Sprintf("%s|%s|%d", region, business, time.Now().UnixMilli()), nil
}
