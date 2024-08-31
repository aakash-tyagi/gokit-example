package main

import (
	"net/http"
	"os"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/log"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (output int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.Count(s)
	return
}
