package dao

import (
	"fmt"
	"testing"
)

func TestGetUserByPhone(t *testing.T) {
	user, _ := GetUserByPhone("15010729356")
	fmt.Println("user = ", user)
}
