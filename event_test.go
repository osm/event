package event

import (
	"sync"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a new event handler hub
	h := NewHub()

	// Make sure we actually pass a function
	err := h.Handle("foo", "bar")
	if err == nil {
		t.Error("event handlers should only accept functions")
	}

	// Make sure we don't accept handlers with more than one parameter
	err = h.Handle("foo", func(a, b string) {})
	if err == nil {
		t.Error("event handlers should only accept functions with one parameter")
	}

	// Make sure that we accept handlers with one parameter
	err = h.Handle("foo", func(a string) {})
	if err != nil {
		t.Error("event handlers should accept functions with one parameter of any type")
	}
}

func TestSend(t *testing.T) {
	// Create a new event handler hub
	h := NewHub()

	// Create a event handler
	h.Handle("foo", func(a string) {})

	// Make sure that we get an error if we try to send an event that don't have any handlers
	err := h.Send("bar", "foo")
	if err == nil {
		t.Error("we should get an error if we send an event that don't have any handlers")
	}

	// We should not get any errors when we send an event that has a handler
	err = h.Send("foo", "foo")
	if err != nil {
		t.Error("we should not get errors when we send events that has handlers")
	}
}

func TestHub(t *testing.T) {
	// Create a new event handler hub
	h := NewHub()

	// Create a new wait group
	var wg sync.WaitGroup

	// Wait for the event handlers
	wg.Add(3)

	// Add two event handlers for the "foo" event
	h.Handle("foo", func(s string) {
		wg.Done()
	})

	h.Handle("foo", func(s string) {
		wg.Done()
	})

	h.Handle("foo", func(i int) {
		wg.Done()
	})

	// Send a string to the "foo" event handlers
	h.Send("foo", "test")

	// Send an int to the "foo" event handler
	h.Send("foo", 1)

	// Wait for all handlers to finish
	wg.Wait()
}
