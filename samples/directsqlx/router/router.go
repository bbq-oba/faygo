/**
 * desc : 系统功能路由注册
 * author:畅雨
 * date:  2016.05.16
 * history:
 *
 */
package router

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/faygo/ext/db/directsqlx"
	"github.com/henrylee2cn/faygo/samples/directsql/handler"
)

// Register the system router in a chain style
func Route(frame *faygo.Framework) {
	faygo.SetUpload("./upload", false, false)
	faygo.SetStatic("./static", false, false)
	// Home page
	frame.NamedAPI("Home", "GET", "/", handler.Index())
	// bos 执行SQL定义的路由
	frame.NamedAPI("DirectSQLX", "POST", "/bosx/*path", directsql.DirectSQL())
	frame.NamedGET("DirectSQLX ModelSql Reload", "/bomx/reloadall", directsql.DirectSQLReloadAll())
	frame.NamedGET("DirectSQLX ModelSql Reload", "/bomx/reload/*path", directsql.DirectSQLReloadModel())
	frame.NamedAPI("Pongo2", "GET", "/pongo2", handler.Pongo2())
	frame.NamedStaticFS("render", "/tpl", faygo.RenderFS(
		"./view",
		".tpl", // "*"
		faygo.Map{"title": "tpl page"},
	))
}