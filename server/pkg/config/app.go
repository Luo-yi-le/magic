package config

import "fmt"

const (
	AppName = "magic"
	Version = "v1.2.11"
)

func GetAppInfo() string {
	return fmt.Sprintf("[%s:%s]", AppName, Version)
}
