package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/gorilla/mux"
)

func AuthSubmit(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if !ctx.App.Auth().Enabled {
			return "", auth.ErrorAuthDisabled
		}
		prv := auth.ProviderFromString(mux.Vars(r)[npncore.KeyKey])
		ref := r.Header.Get("Referer")
		state := "/"
		if len(ref) > 0 {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				state = u.Path
			}
		}

		u := ctx.App.Auth().URLFor(state, prv)
		if len(u) == 0 {
			return npncontroller.ENew(prv.Title + " is disabled")
		}
		return u, nil
	})
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if !ctx.App.Auth().Enabled {
			return "", auth.ErrorAuthDisabled
		}
		_, _ = ctx.App.User().SaveProfile(ctx.Profile)
		prv := auth.ProviderFromString(mux.Vars(r)[npncore.KeyKey])
		code, ok := r.URL.Query()["code"]
		if !ok || len(code) == 0 {
			return npncontroller.ENew("no auth code provided")
		}
		stateS, ok := r.URL.Query()["state"]
		u := "/"
		if ok && len(stateS) > 0 && strings.HasPrefix(stateS[0], "/") {
			u = stateS[0]
		}
		record, err := ctx.App.Auth().Handle(ctx.Profile, prv, code[0])
		if err != nil {
			return npncontroller.EResp(err)
		}

		msg := "signed in as " + record.Name
		return npncontroller.FlashAndRedir(true, msg, u, w, r, ctx)
	})
}

func AuthSignout(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if !ctx.App.Auth().Enabled {
			return "", auth.ErrorAuthDisabled
		}
		id, err := npnweb.IDFromParams(npncore.KeyAuth, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err, npncore.IDErrorString(npncore.KeyAuth, ""))
		}

		err = ctx.App.Auth().Delete(*id)
		if err != nil {
			return npncontroller.EResp(err, "unable to delete auth record")
		}

		ref := r.Header.Get("Referer")
		if len(ref) == 0 {
			ref = "/"
		}

		return ref, nil
	})
}
