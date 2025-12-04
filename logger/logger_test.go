package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintln(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer
	
	// Set the logger output to our buffer
	SetOutput(&buf)
	
	// Test the Println function
	Println("Test message", "with multiple", "parts")
	
	output := buf.String()
	
	// Check that the output contains the message parts
	if !strings.Contains(output, "Test message") {
		t.Errorf("Output should contain 'Test message', got: %s", output)
	}
	if !strings.Contains(output, "with multiple") {
		t.Errorf("Output should contain 'with multiple', got: %s", output)
	}
	if !strings.Contains(output, "parts") {
		t.Errorf("Output should contain 'parts', got: %s", output)
	}
	
	// Check that the output contains color codes
	if !strings.Contains(output, "\033[36m") { // Cyan color for timestamp
		t.Errorf("Output should contain cyan color code for timestamp")
	}
	if !strings.Contains(output, "\033[32m") { // Green color for message
		t.Errorf("Output should contain green color code for message")
	}
	if !strings.Contains(output, "\033[0m") { // Reset color
		t.Errorf("Output should contain color reset code")
	}
	
	// Check that the output contains a timestamp in the expected format
	if !strings.Contains(output, "/") || !strings.Contains(output, ":") {
		t.Errorf("Output should contain a timestamp with '/' and ':', got: %s", output)
	}
}
