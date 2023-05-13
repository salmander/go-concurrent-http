package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// LogCapturer is a custom io.Writer that captures the log output
type LogCapturer struct {
	buf bytes.Buffer
}

// Write implements the io.Writer interface
func (c *LogCapturer) Write(p []byte) (n int, err error) {
	return c.buf.Write(p)
}

func TestMyHandler(t *testing.T) {
	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodGet, "/my-endpoint", nil)
	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	MyHandler(rr, req)

	// Assert the response status code is 204
	if rr.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rr.Code)
	}
}

func TestMyConcurrentFunction(t *testing.T) {
	// Create a LogCapturer instance and replace the default logger output with it
	capturer := &LogCapturer{}
	log.SetOutput(capturer)

	//done := make(chan bool)

	// Create a test server with a handler that returns a sample response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"current_weather": {"temperature": 25.5}}`))
	}))
	defer ts.Close()

	// Set the test server URL as the API endpoint
	URL = ts.URL

	// Call the concurrent function
	MyConcurrentFunction()

	// Wait for the concurrent function to complete
	//<-done

	// Get the captured log output
	logOutput := capturer.buf.String()

	// Verify the expected log line
	expectedLogLine := "Temperature: 25.5"
	if !bytes.Contains([]byte(logOutput), []byte(expectedLogLine)) {
		t.Errorf("expected log line not found: %s in %s", expectedLogLine, logOutput)
	}
}

func TestHandlerIntegration(t *testing.T) {
	// Create a new HTTP test server with the handler
	srv := httptest.NewServer(http.HandlerFunc(MyHandler))
	defer srv.Close()

	// Make an HTTP request to the app
	resp, err := http.Get(srv.URL + "/my-endpoint")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Assert the response status code is 204
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, resp.StatusCode)
	}
}
