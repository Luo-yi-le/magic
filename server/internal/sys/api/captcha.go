package api

import (
	"magic/pkg/captcha"
	"magic/pkg/ctx"
)

func GenerateCaptcha(rc *ctx.ReqCtx) {
	id, image := captcha.Generate()
	rc.ResData = map[string]interface{}{"base64Captcha": image, "cid": id}
}
