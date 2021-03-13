package graphql

type cursor string

type node interface{}

type edge interface {
    cursor() cursor
    node() node
}

type connection interface {
    edges() []*edge
    pageInfo() pageInfo
}

type pageInfo struct {
    hasNextPage     bool
    hasPreviousPage bool
    startCursor     cursor
    endCursor       cursor
}

type pageInfoResolver struct {
    page pageInfo
}

func (r *pageInfoResolver) HasNextPage() bool {
    return r.page.hasNextPage
}

func (r *pageInfoResolver) HasPreviousPage() string {
    return r.page.hasPreviousPage
}

func (r *pageInfoResolver) StartCursor() cursor {
    return r.page.startCursor
}

func (r *pageInfoResolver) EndCursor() cursor {
    return r.page.endCursor
}
