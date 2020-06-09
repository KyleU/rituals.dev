package admin

import (
	"github.com/kyleu/rituals.dev/app/web/act"
	"net/http"

	"github.com/kyleu/rituals.dev/app/database/gql"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func GraphiQL(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return act.EResp(err)
		}

		ctx.Breadcrumbs = adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		ctx.Title = "GraphiQL"
		return act.T(templates.GraphiQL(ctx, w))
	})
}

func GraphQLVoyagerQuery(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return act.EResp(err)
		}

		bc := adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		link := util.AdminLink(util.KeyVoyager, "query")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link), "query")...)
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return act.T(templates.GraphQLVoyager(gql.QueryName, ctx, w))
	})
}

func GraphQLVoyagerMutation(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return act.EResp(err)
		}

		bc := adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		link := util.AdminLink(util.KeyVoyager, "mutation")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link), "mutation")...)
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return act.T(templates.GraphQLVoyager(gql.MutationName, ctx, w))
	})
}
