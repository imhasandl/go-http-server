package main
   
import (
   "log"
	"fmt"
	"sync/atomic"
	"net/http"
) 

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main(){
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig {
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.handleReset)
	
	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}	

	log.Printf("Serving files from %s\n You are at the port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", )))
}

func (cfg *apiConfig) middlewareMetricsInc(next, http.Handler) http.Handler {
	return nil
}