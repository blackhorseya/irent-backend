package ginprom

import (
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricsPath = "/metrics"
)

var (
	httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server",
		Name:      "request_seconds",
		Help:      "Histogram of response latency (seconds) of http handlers.",
	}, []string{"method", "code", "uri"})
)

func init() {
	prometheus.Register(httpHistogram)
}

type handlerPath struct {
	sync.Map
}

func (h *handlerPath) get(handler string) string {
	value, ok := h.Load(handler)
	if !ok {
		return ""
	}

	return value.(string)
}

func (h *handlerPath) set(ri gin.RouteInfo) {
	h.Store(ri.Handler, ri.Path)
}

// GinPrometheus declare gin server for prometheus
type GinPrometheus struct {
	engine  *gin.Engine
	ignored map[string]bool
	pathMap *handlerPath
	updated bool
}

// Option declare GinPrometheus configuration
type Option func(*GinPrometheus)

// Ignore added ignore path
func Ignore(path ...string) Option {
	return func(gp *GinPrometheus) {
		for _, p := range path {
			gp.ignored[p] = true
		}
	}
}

// New serve caller to create a GinPrometheus
func New(e *gin.Engine, options ...Option) *GinPrometheus {
	if e == nil {
		return nil
	}

	gp := &GinPrometheus{
		engine: e,
		ignored: map[string]bool{
			metricsPath: true,
		},
		pathMap: &handlerPath{},
	}
	for _, option := range options {
		option(gp)
	}

	return gp
}

func (gp *GinPrometheus) updatePath() {
	gp.updated = true
	for _, ri := range gp.engine.Routes() {
		gp.pathMap.set(ri)
	}
}

// Middleware return middleware for gin
func (gp *GinPrometheus) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !gp.updated {
			gp.updatePath()
		}

		if gp.ignored[c.Request.URL.String()] {
			c.Next()
			return
		}
		begin := time.Now()

		c.Next()

		httpHistogram.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
			gp.pathMap.get(c.HandlerName()),
		).Observe(time.Since(begin).Seconds())
	}
}
