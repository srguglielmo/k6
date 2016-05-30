package simple

import (
	log "github.com/Sirupsen/logrus"
	"github.com/loadimpact/speedboat"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/context"
	"time"
)

type Runner struct {
	Client *fasthttp.Client
}

func New() *Runner {
	return &Runner{
		Client: &fasthttp.Client{
			MaxIdleConnDuration: time.Duration(0),
		},
	}
}

func (r *Runner) RunVU(ctx context.Context, t speedboat.Test) {
	for {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(res)

		req.SetRequestURI(t.URL)

		startTime := time.Now()
		if err := r.Client.Do(req, res); err != nil {
			log.WithError(err).Error("Request error")
		}
		duration := time.Since(startTime)

		log.WithField("duration", duration).Info("Duration")

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}