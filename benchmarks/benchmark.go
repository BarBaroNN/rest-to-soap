package benchmarks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// BenchmarkConfig holds benchmark configuration
type BenchmarkConfig struct {
	ProxyURL     string
	DirectURL    string
	RequestCount int
	Concurrency  int
	Timeout      time.Duration
}

// BenchmarkResult holds benchmark results
type BenchmarkResult struct {
	TotalTime    time.Duration
	SuccessCount int
	ErrorCount   int
	MinLatency   time.Duration
	MaxLatency   time.Duration
	AvgLatency   time.Duration
}

// RunBenchmark runs a benchmark comparing proxy vs direct calls
func RunBenchmark(cfg BenchmarkConfig) (proxyResult, directResult BenchmarkResult, err error) {
	// Run proxy benchmark
	proxyResult, err = runBenchmark(cfg.ProxyURL, cfg)
	if err != nil {
		return
	}

	// Run direct benchmark
	directResult, err = runBenchmark(cfg.DirectURL, cfg)
	return
}

func runBenchmark(url string, cfg BenchmarkConfig) (BenchmarkResult, error) {
	client := &http.Client{
		Timeout: cfg.Timeout,
	}

	// Create request body
	body := map[string]interface{}{
		"test": "value",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return BenchmarkResult{}, err
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if err != nil {
		return BenchmarkResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Run benchmark
	start := time.Now()
	results := make(chan time.Duration, cfg.RequestCount)
	errors := make(chan error, cfg.RequestCount)

	// Create worker pool
	sem := make(chan struct{}, cfg.Concurrency)
	for i := 0; i < cfg.RequestCount; i++ {
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
			reqStart := time.Now()
			resp, err := client.Do(req)
			if err != nil {
				errors <- err
				return
			}
			resp.Body.Close()
			results <- time.Since(reqStart)
		}()
	}

	// Collect results
	var latencies []time.Duration
	for i := 0; i < cfg.RequestCount; i++ {
		select {
		case latency := <-results:
			latencies = append(latencies, latency)
		case err := <-errors:
			return BenchmarkResult{}, err
		}
	}

	// Calculate statistics
	var min, max, sum time.Duration
	if len(latencies) > 0 {
		min = latencies[0]
		max = latencies[0]
		for _, l := range latencies {
			if l < min {
				min = l
			}
			if l > max {
				max = l
			}
			sum += l
		}
	}

	return BenchmarkResult{
		TotalTime:    time.Since(start),
		SuccessCount: len(latencies),
		ErrorCount:   cfg.RequestCount - len(latencies),
		MinLatency:   min,
		MaxLatency:   max,
		AvgLatency:   sum / time.Duration(len(latencies)),
	}, nil
}

// BenchmarkProxyVsDirect runs a benchmark comparing proxy vs direct calls
func BenchmarkProxyVsDirect(b *testing.B) {
	cfg := BenchmarkConfig{
		ProxyURL:     "http://localhost:8080/api/soap/example",
		DirectURL:    "http://example.com/soap",
		RequestCount: 1000,
		Concurrency:  100,
		Timeout:      30 * time.Second,
	}

	proxyResult, directResult, err := RunBenchmark(cfg)
	if err != nil {
		b.Fatal(err)
	}

	// Print results
	fmt.Printf("Proxy Results:\n")
	fmt.Printf("  Total Time: %v\n", proxyResult.TotalTime)
	fmt.Printf("  Success Count: %d\n", proxyResult.SuccessCount)
	fmt.Printf("  Error Count: %d\n", proxyResult.ErrorCount)
	fmt.Printf("  Min Latency: %v\n", proxyResult.MinLatency)
	fmt.Printf("  Max Latency: %v\n", proxyResult.MaxLatency)
	fmt.Printf("  Avg Latency: %v\n", proxyResult.AvgLatency)

	fmt.Printf("\nDirect Results:\n")
	fmt.Printf("  Total Time: %v\n", directResult.TotalTime)
	fmt.Printf("  Success Count: %d\n", directResult.SuccessCount)
	fmt.Printf("  Error Count: %d\n", directResult.ErrorCount)
	fmt.Printf("  Min Latency: %v\n", directResult.MinLatency)
	fmt.Printf("  Max Latency: %v\n", directResult.MaxLatency)
	fmt.Printf("  Avg Latency: %v\n", directResult.AvgLatency)

	// Calculate overhead
	overhead := float64(proxyResult.AvgLatency-directResult.AvgLatency) / float64(directResult.AvgLatency) * 100
	fmt.Printf("\nProxy Overhead: %.2f%%\n", overhead)
}
