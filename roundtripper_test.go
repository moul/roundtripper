package roundtripper_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"moul.io/roundtripper"
)

func Example() {
	client := http.Client{
		Transport: &roundtripper.Transport{
			ExtraHeader: http.Header{
				"Authorization": []string{"Bearer LoremIpsumDolorSitAmet..."},
			},
		},
	}
	client.Get("...")
}

func TestTransport(t *testing.T) {
	client := http.Client{
		Transport: &roundtripper.Transport{
			ExtraHeader: http.Header{
				"Authorization": []string{"Bearer LoremIpsumDolorSitAmet..."},
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := http.Header{
			"Authorization":   []string{"Bearer LoremIpsumDolorSitAmet..."},
			"Accept-Encoding": []string{"gzip"},
			"User-Agent":      []string{"Go-http-client/1.1"},
		}
		require.Equal(t, r.Header, expected)
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := client.Get(ts.URL)
	require.NoError(t, err)

	greeting, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	res.Body.Close()
	require.Equal(t, greeting, []byte("Hello, client\n"))

}
