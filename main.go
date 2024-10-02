package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Initialize Zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Load .env file if not running in Railway environment
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
			logger.Fatal("Error loading .env file", zap.Error(err))
			os.Exit(1)
		}
	}

	// Get the target URL from the environment variable
	target := os.Getenv("TARGET_URL")
	if target == "" {
		logger.Fatal("TARGET_URL is not set in the environment")
	}

	// Parse the target URL
	url, err := url.Parse(target)
	if err != nil {
		logger.Fatal("Invalid target URL", zap.String("url", target), zap.Error(err))
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
		logger.Info("Proxying request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("remote_addr", r.RemoteAddr),
		)
		proxy.ServeHTTP(w, r)
	})

	// Start the proxy server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	logger.Info("Proxy server is running", zap.String("port", port), zap.String("target_url", target))
	logger.Fatal("Server failed", zap.Error(http.ListenAndServe(":"+port, nil)))
}
