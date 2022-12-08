package test

import (
	"cron/api"
	"cron/pkg"
	"fmt"
	"testing"
)

func TestApiResp(*testing.T) {
	a := api.Response{
		Code: 1,
		Msg:  "fff",
		Data: nil,
	}
	a.Page = 1

	fmt.Println(1111, pkg.MarshalToString(a))
}
