package local

import (
    "os"
    "time"
)

// Handler TODO
type Handler struct {
    config Config
    file   *os.File
}

// New TODO
func New(
    config Config,
) *Handler {
    return &Handler{
        config: config,
    }
}

func (h *Handler) getBase(t time.Time) string {
    pre := h.config.Base.Prefix
    pattern := t.Format(h.config.Base.Format)
    post := h.config.Base.Postfix
    ext := h.config.Base.Extension

    return pre + pattern + post + "." + ext
}

func (h *Handler) getPath(t time.Time) string {
    loc := h.config.Path.Location
    pattern := t.Format(h.config.Path.Format)

    // Add in some path logic here for slashes and such
    return loc + pattern
}
