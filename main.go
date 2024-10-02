package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
			log.Fatal("Error loading .env file")
			os.Exit(1)
		}
	}

	// Get the target URL from environment variable
	target := os.Getenv("TARGET_URL")
	if target == "" {
		log.Fatal("TARGET_URL is not set in the environment")
	}

	// Parse the target URL
	url, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Define a custom transport for optimization
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// Create a custom client with the transport and set timeout
	client := &http.Client{
		Transport: transport,
	}

	// Assign the custom client to the proxy
	proxy.Transport = client.Transport

	// Proxy handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying request: %s %s", r.Method, r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	// Start the proxy server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	log.Printf("Proxy server is running on port %s, forwarding to %s", port, target)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
