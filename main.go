package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xuperchain/contract-sdk-go/code"
	"github.com/xuperchain/contract-sdk-go/driver"
)

// ErrNotFound is returned when key is not found
var errNotFound = "Key not found"

type loveContract struct {
	ctx code.Context
}

//初始化方法,接口要求,必须实现
func (love *loveContract) Initialize(ctx code.Context) code.Response {
	return code.OK(nil)
}

func main() {
	driver.Serve(new(loveContract))
}

func (love *loveContract) Boy(ctx code.Context) code.Response {
	love.ctx = ctx
	args := ctx.Args()
	boy := string(args["boy"])
	girl := string(args["girl"])
	return love.checkLove(boy, girl)
}

func (love *loveContract) Girl(ctx code.Context) code.Response {
	love.ctx = ctx
	args := ctx.Args()
	boy := string(args["boy"])
	girl := string(args["girl"])
	return love.checkLove(girl, boy)
}

func (love *loveContract) checkLove(suitor string, loved string) code.Response {
	if suitor == "" || loved == "" {
		return code.Errors("Missing key: boy or girl")
	}
	operator := love.ctx.Caller()
	if operator != suitor {
		return code.Errors("operator is error")
	}
	suitorLover, err := love.beloved(suitor)
	if err != nil {
		return code.Error(err)
	}

	if suitorLover == loved {
		return code.OK([]byte("You are true love"))
	} else if suitorLover != "" {
		return code.Errors("You are Neptune")
	}

	loveder, err := love.beloved(loved)
	if err != nil {
		return code.Error(err)
	}
	if loveder == suitor {
		return code.OK([]byte("You are true love"))
	} else if loveder != "" {
		return code.Errors("Sorry, you are green")
	}
	err = love.ctx.PutObject([]byte(suitor), []byte(loved))
	if err != nil {
		return code.Error(err)
	}
	return code.OK([]byte("You are true love"))
}

func (love *loveContract) beloved(spouse string) (string, error) {
	spouseBytes, err := love.ctx.GetObject([]byte(spouse))
	if err != nil {
		if strings.Contains(err.Error(), errNotFound) {
			return "", nil
		} else {
			return "", fmt.Errorf("beloved-->e.ctx.GetObject:%w", err)
		}

	} else {
		loved := ""
		err := json.Unmarshal(spouseBytes, &loved)
		return loved, err
	}

}
