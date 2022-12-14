package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/app"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/mini"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	App     app.RouterGroup
	Mini    mini.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
