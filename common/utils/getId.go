package utils

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetFromId(c *app.RequestContext) (uint64, error) {
	value, exists := c.Get("from_id")
	if exists != true {
		return 0, errors.New("token解析失败")
	}
	fromId := uint64(value.(uint))
	return fromId, nil
}
