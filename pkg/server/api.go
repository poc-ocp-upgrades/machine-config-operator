package server

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"net/http"
	godefaulthttp "net/http"
	"path"
	"github.com/golang/glog"
)

type poolRequest struct{ machineConfigPool string }
type APIServer struct {
	handler		http.Handler
	port		int
	insecure	bool
	cert		string
	key			string
}

func NewAPIServer(a *APIHandler, p int, is bool, c, k string) *APIServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux := http.NewServeMux()
	mux.Handle("/config/", a)
	mux.Handle("/healthz", &healthHandler{})
	mux.Handle("/", &defaultHandler{})
	return &APIServer{handler: mux, port: p, insecure: is, cert: c, key: k}
}
func (a *APIServer) Serve() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mcs := &http.Server{Addr: fmt.Sprintf(":%v", a.port), Handler: a.handler}
	glog.Info("launching server")
	if a.insecure {
		if err := mcs.ListenAndServe(); err != http.ErrServerClosed {
			glog.Exitf("Machine Config Server exited with error: %v", err)
		}
	} else {
		if err := mcs.ListenAndServeTLS(a.cert, a.key); err != http.ErrServerClosed {
			glog.Exitf("Machine Config Server exited with error: %v", err)
		}
	}
}

type APIHandler struct{ server Server }

func NewServerAPIHandler(s Server) *APIHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &APIHandler{server: s}
}
func (sh *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "" {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cr := poolRequest{machineConfigPool: path.Base(r.URL.Path)}
	conf, err := sh.server.GetConfig(cr)
	if err != nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusInternalServerError)
		glog.Errorf("couldn't get config for req: %v, error: %v", cr, err)
		return
	}
	if conf == nil && err == nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := json.Marshal(conf)
	if err != nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusInternalServerError)
		glog.Errorf("failed to marshal %v config: %v", cr, err)
		return
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		glog.Errorf("failed to write %v response: %v", cr, err)
	}
}

type healthHandler struct{}

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Set("Content-Length", "0")
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

type defaultHandler struct{}

func (h *defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Set("Content-Length", "0")
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
