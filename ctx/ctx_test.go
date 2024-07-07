package ctx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type spyStore struct {
	response  string
	cancelled bool
}

func (s *spyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *spyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("should return valid response", func(t *testing.T) {
		data := "hello, world"
		store := spyStore{response: data}

		svr := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("expected %q, but got %q", data, res.Body.String())
		}

		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})
	t.Run("tells store to cancel workd if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := spyStore{
			response: data,
		}
		svr := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}
