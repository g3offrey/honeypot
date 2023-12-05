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

func RegisterCatch(catchRepo CatchRepository, r *http.Request) error {
	c := Catch{
		IPAddress: r.RemoteAddr,
		CatchedOn: r.URL.Path,
		At:        time.Now(),
	}

	err := catchRepo.Create(c)
	if err != nil {
		return fmt.Errorf("can't create catch: %w", err)
	}

	return nil
}
