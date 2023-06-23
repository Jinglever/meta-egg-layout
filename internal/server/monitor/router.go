// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.5.12-EE
// Author: meta-egg
// Generated at: 2023-06-09 16:24

package monitor

import (
	"net/http"
	"net/http/pprof"
)

func (s *Server) initRouter() {
	router := http.NewServeMux()
	router.HandleFunc("/debug/pprof", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.HandleFunc("/debug/pprof/{action}", pprof.Index)

	s.Router = router
}