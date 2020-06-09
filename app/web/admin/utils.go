package admin

import (
	"net/http"
	"time"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"
)

type JSONResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Path     string    `json:"path"`
	Occurred time.Time `json:"occurred"`
}

func Modules(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.Title(util.KeyModules)
		ctx.Breadcrumbs = adminBC(ctx, util.KeyModules, util.KeyModules)
		return act.T(admintemplates.ModulesList(ctx, w))
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.Title(util.KeyRoutes)
		ctx.Breadcrumbs = adminBC(ctx, util.KeyRoutes, util.KeyRoutes)
		return act.T(admintemplates.RoutesList(ctx, w))
	})
}

func adminAct(w http.ResponseWriter, r *http.Request, f func(*web.RequestContext) (string, error)) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		if ctx.Profile.Role != util.RoleAdmin {
			if act.IsContentTypeJSON(act.GetContentType(r)) {
				ae := JSONResponse{Status: "error", Message: "you are not an administrator", Path: r.URL.Path, Occurred: time.Now()}
				return act.RespondJSON(w, "", ae, ctx.Logger)
			}
			msg := "you're not an administrator, silly!"
			return act.FlashAndRedir(false, msg, "home", w, r, ctx)
		}
		return f(ctx)
	})
}

func adminBC(ctx *web.RequestContext, action string, name string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.AdminLink()), util.KeyAdmin)
	bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(action)), name)...)
	return bc
}
