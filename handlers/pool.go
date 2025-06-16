package handlers

import (
	"context"
	"sync"
)

// Pool manages a pool of workers for processing requests
type Pool struct {
	workers int
	wg      sync.WaitGroup
}

// NewPool creates a new worker pool
func NewPool() *Pool {
	return &Pool{
		workers: 10, // Default number of workers
	}
}

// Size returns the current number of workers in the pool
func (p *Pool) Size() int {
	return p.workers
}

// WithContext executes a function in the worker pool with the given context
func (p *Pool) WithContext(ctx context.Context, fn func() error) error {
	p.wg.Add(1)
	defer p.wg.Done()

	// Create a channel to receive the result
	result := make(chan error, 1)

	// Execute the function in a goroutine
	go func() {
		result <- fn()
	}()

	// Wait for either the context to be done or the function to complete
	select {
	case err := <-result:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
