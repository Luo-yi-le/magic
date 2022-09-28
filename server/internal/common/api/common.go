package api

import (
	"magic/pkg/biz"
	"magic/pkg/ctx"
	"magic/pkg/utils"
)

type Common struct {
}

func (i *Common) RasPublicKey(rc *ctx.ReqCtx) {
	publicKeyStr, err := utils.GetRsaPublicKey()
	biz.ErrIsNilAppendErr(err, "rsa生成公私钥失败")
	rc.ResData = publicKeyStr
}
