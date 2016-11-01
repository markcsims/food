package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/quick"

	"github.com/markcsims/bbcfood/tools"
)

func newServer(t *testing.T) *httptest.Server {
	testLogger := &tools.TestLogger{T: t}
	testConfig := &appConfig{}
	tsdConfig := tools.StatsDConfig{Log: testLogger}
	testStatsD, _ := tools.NewStatsD(tsdConfig)

	return httptest.NewServer(newRouter(testConfig, testLogger, testStatsD))
}

func TestHelloRouter(t *testing.T) {
	server := newServer(t)
	response, err := http.Get(server.URL + "/hello")

	if err != nil {
		t.Fatal("Bugger, i got an error doing a request", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Error("Expected a 200 but i got", response.StatusCode)
	}
}

func TestTheDomainWithProperties(t *testing.T) {
	assertion := func(name string) bool {
		result := Hello(name)
		return strings.Contains(result, name) && strings.Contains(result, "Hello")
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkTheDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if Hello("Chris") != "Hello, Chris" {
			b.Error()
		}
	}
}
