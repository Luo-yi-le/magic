package config

import "magic/pkg/utils/assert"

type Jwt struct {
	Key        string `yaml:"key"`
	ExpireTime uint64 `yaml:"expire-time"` // 过期时间，单位分钟
}

func (j *Jwt) Valid() {
	assert.IsTrue(j.ExpireTime != 0, "config.yml之 [jwt.expire-time] 不能为空")
}
