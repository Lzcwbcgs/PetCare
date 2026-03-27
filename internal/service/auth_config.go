package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func authConfigString(ctx context.Context, pattern string, def string) string {
	value, err := g.Cfg().GetWithEnv(ctx, pattern, def)
	if err != nil || value == nil {
		return def
	}
	return value.String()
}

func authConfigInt(ctx context.Context, pattern string, def int) int {
	value, err := g.Cfg().GetWithEnv(ctx, pattern, def)
	if err != nil || value == nil {
		return def
	}
	return value.Int()
}
