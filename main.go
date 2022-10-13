package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func metrics(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received request for metrics")

	data, _ := ioutil.ReadFile("gauge.data")
	prometheusDummy := fmt.Sprintf("# HELP resque_processed_jobs lots of jobs\n# TYPE resque_processed_jobs gauge\nresque_processed_jobs %s", data)
	fmt.Fprintf(w, prometheusDummy)
}

func setMetric(w http.ResponseWriter, req *http.Request) {
	count := req.URL.Query().Get("count")

	ioutil.WriteFile("gauge.data", []byte(count), 0777)
	fmt.Fprintf(w, "metric set")
}

func main() {
	fmt.Println("im running")
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/set_metric", setMetric)

	http.ListenAndServe(":8090", nil)
}
