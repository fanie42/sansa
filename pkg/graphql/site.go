package graphql

import "github.com/graph-gophers/graphql-go"

type siteResolver struct {
    site *site.Service
    systems *system.Service
}

func (r *siteResolver) ID() graphql.ID {
    return graphql.ID(r.site.ID)
}

func (r *siteResolver) Name() string {
    return r.site.Name
}

func (r *siteResolver) Abbr() string {
    return r.site.Abbr
}

func (r *siteResolver) Systems() *siteConnectionResolver {
    return &siteConnectionResolver{}
}

type sitesConnectionResolver struct {
    sites *site.Service
}

func (r *siteConnectionResolver) Edges(
    ctx context.Context,
) (*[]*siteEdgeResolver, error) {
    sites := r.sites.Fetch()

    resolvers := []*siteEdgeResolver{}
    for _, site := range sites {
        resolver := &siteEdgeResolver{
            site: &siteResolver{
                systems: systems,
                site: 
            },
        }
        resolvers = append(resolvers, resolver)
    }

    return resolvers
}

func (r *siteConnectionResolver) PageInfo() *pageInfoResolver {
    return &pageInfoResolver{

    }
}

type siteEdgeResolver struct{
    site *siteResolver
}

func (r *siteEdgeResolver) Cursor(
    ctx context.Context,
) string {}

func (r *siteEdgeResolver) Node(
    ctx context.Context,
) *siteResolver {
    return &siteResolver{
        systems: systemConnectionResolver
        site: 
    }
}