package server

import (
	"net/http"
	"sync"
)

type Server struct {
	sync.Mutex
	HttpClient  *http.Client
	TestEnvFlag bool
}
