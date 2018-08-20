# Captcha

an example for use captcha

```
package controllers

import (
	"github.com/izi-global/izigo"
	"github.com/izi-global/izigo/cache"
	"github.com/izi-global/izigo/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	// use izigo cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

type MainController struct {
	izigo.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
}

func (this *MainController) Post() {
	this.TplName = "index.tpl"

	this.Data["Success"] = cpt.VerifyReq(this.Ctx.Request)
}
```

template usage

```
{{.Success}}
<form action="/" method="post">
	{{create_captcha}}
	<input name="captcha" type="text">
</form>
```
