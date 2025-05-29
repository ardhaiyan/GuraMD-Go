package helpers

import "os"

var (
	Public = true
	Name   = os.Getenv("Name_Bot")
	Owner  = os.Getenv("Owner_Number")
)

func SetName(name string) {
	Name = name
}
