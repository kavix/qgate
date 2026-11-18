// webhook is the Kubernetes validating admission webhook that rejects pods
// whose images do not satisfy the signature policy.
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	// TODO: /validate -> parse AdmissionReview, verify signatures, apply policy.
	log.Println("webhook listening on :8443 (TLS wiring TODO)")
	log.Fatal(http.ListenAndServe(":8443", nil))
}
