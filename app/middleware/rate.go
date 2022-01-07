package middleware

import "time"

// 令牌桶中间件
// 每个用户 每分钟限制访问50次请求


type Rate struct {
	period time.Duration
	count  int
}
