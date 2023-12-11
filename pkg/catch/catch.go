package catch

import (
	"fmt"
	"net/http"
	"time"
)

type Catch struct {
	IPAddress string
	CatchedOn string
	At        time.Time
}

type CatchRepository interface {
	Create(c Catch) error
}

func getIP(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		return forwardedFor
	}

	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	return r.RemoteAddr
}

func RegisterCatch(catchRepo CatchRepository, r *http.Request) error {
	c := Catch{
		IPAddress: getIP(r),
		CatchedOn: r.URL.Path,
		At:        time.Now(),
	}

	err := catchRepo.Create(c)
	if err != nil {
		return fmt.Errorf("can't create catch: %w", err)
	}

	return nil
}
