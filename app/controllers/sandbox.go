package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/sandbox"

	"github.com/kyleu/rituals.dev/app/web"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.KeyPluralTitle(util.KeySandbox)
		ctx.Breadcrumbs = append(aboutBC(ctx), web.BreadcrumbsSimple(ctx.Route(util.KeySandbox), util.KeySandbox)...)
		return tmpl(templates.SandboxList(sandbox.AllSandboxes, ctx, w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sb := sandbox.FromString(key)
		if sb == nil {
			return "", util.IDError(util.KeySandbox, key)
		}
		content, err := sb.Resolve(ctx)
		if err != nil {
			return "", errors.Wrap(err, "error running sandbox ["+key+"]")
		}

		ctx.Title = sb.Title + " Sandbox"
		bc := append(aboutBC(ctx), web.BreadcrumbsSimple(ctx.Route(util.KeySandbox), util.KeySandbox)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.KeySandbox), util.KeySandbox)...)
		bc = append(bc, web.Breadcrumb{Path: ctx.Route(util.KeySandbox+".run", util.KeyKey, key), Title: key})
		ctx.Breadcrumbs = bc

		return tmpl(templates.SandboxRun(sb, util.ToJSON(content), ctx, w))
	})
}
