package nimble_test

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/meshenka/nimble"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("integration test")
	}

	addr, teardown := setup(t)
	t.Cleanup(teardown)
	time.Sleep(300 * time.Millisecond)
	t.Run("get a random hero", func(t *testing.T) {
		req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/heros", http.NoBody)
		require.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer res.Body.Close() //nolint:errcheck

		require.Equal(t, http.StatusOK, res.StatusCode)

		data, err := io.ReadAll(res.Body)
		require.NoError(t, err)
		resp := new(Response)
		require.NoError(t, json.Unmarshal(data, resp))
		t.Log(resp)
		assert.NotZero(t, resp.Sentence)
	})

	t.Run("classes", func(t *testing.T) {
		t.Run("get all", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/classes", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusOK, res.StatusCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)
		})
		t.Run("get one", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/classes/Shepherd", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusOK, res.StatusCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)
		})

	})
	t.Run("ancestries", func(t *testing.T) {
		t.Run("get all", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/ancestries", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusOK, res.StatusCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)
		})
		t.Run("get one", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/ancestries/Elf", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusOK, res.StatusCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)
		})
	})
	t.Run("backgrounds", func(t *testing.T) {
		t.Run("get all", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/backgrounds", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusOK, res.StatusCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)
		})
		t.Run("get one", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/backgrounds/Ear to the Ground", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusOK, res.StatusCode)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)
		})
		t.Run("get non existing", func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/api/backgrounds/NO_MATCH", http.NoBody)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close() //nolint:errcheck

			require.Equal(t, http.StatusNotFound, res.StatusCode)
		})
	})
}

type Response struct {
	Sentence string `json:"sentence"`
}

func setup(t *testing.T) (string, context.CancelFunc) {
	t.Helper()

	ctx, cancel := context.WithCancel(t.Context())
	httpAddr := addr()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		require.NoError(t, nimble.Serve(ctx,
			nimble.WithApplicationServer(httpAddr),
			nimble.WithLogLevel("debug"),
		))
	}()

	return httpAddr, func() { cancel(); wg.Wait() }
}

func addr() string {
	lst, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer lst.Close() //nolint:errcheck
	return lst.Addr().String()
}
