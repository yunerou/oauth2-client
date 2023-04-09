package comhttp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/yunerou/oauth2-client/singleton"
)

// https://stackoverflow.com/questions/39527847/is-there-middleware-for-go-http-client

// This type implements the http.RoundTripper interface
type LoggingRoundTripper struct {
	Ctx     context.Context
	Proxied http.RoundTripper
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	maxLengthJSON := singleton.GetViper().GetInt("MAX_LENGTH_LOG_IN_BYTE")
	const (
		infoLv string = "info"
		warnLv string = "warn"
	)
	level := infoLv

	startTime := time.Now()
	logfn := singleton.GetLogger(lrt.Ctx).With()

	defer func() {
		apiWaitDuration := time.Since(startTime)
		logfn = logfn.
			Str("url", req.URL.String()).
			Str("method", req.Method).
			Str("process_duration", apiWaitDuration.String())
		l := logfn.Logger()
		if level == infoLv {
			l.Info().Send()
		}
		if level == warnLv {
			l.Warn().Send()
		}
	}()

	if req.Body != nil {
		var reqb bytes.Buffer
		tee := io.TeeReader(req.Body, &reqb)
		reader1, _ := io.ReadAll(tee)
		if len(reader1) > maxLengthJSON {
			logfn = logfn.Str("reqBody", string(reader1[:maxLengthJSON]))
			// upload to s3 if need
		} else {
			logfn = logfn.RawJSON("reqBody", reader1)
		}
		req.Body = io.NopCloser(&reqb)
	}

	// Send the request, get the response (or the error)
	res, e = lrt.Proxied.RoundTrip(req)

	// Handle the result.
	if e != nil {
		logfn = logfn.Err(e)
	} else {
		if res.StatusCode != http.StatusOK {
			level = warnLv
		}
		logfn = logfn.
			Int("resStatus", res.StatusCode)
		if res.Body != nil {
			var resb bytes.Buffer
			tee := io.TeeReader(res.Body, &resb)
			reader1, _ := io.ReadAll(tee)
			if len(reader1) > maxLengthJSON {
				logfn = logfn.Str("resBody", string(reader1[:maxLengthJSON]))
				// upload to s3 if need
			} else {
				logfn = logfn.RawJSON("resBody", reader1)
			}

			res.Body = io.NopCloser(&resb)
		}
	}
	return
}
