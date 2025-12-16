package go_strings

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "simple string",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "string with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "unicode characters",
			input:    "привет",
			expected: "тевирп",
		},
		{
			name:     "mixed unicode and ascii",
			input:    "hello мир",
			expected: "рим olleh",
		},
		{
			name:     "emoji",
			input:    "👋🌍",
			expected: "🌍👋",
		},
		{
			name:     "palindrome",
			input:    "racecar",
			expected: "racecar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStripNonAlphanumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only alphanumeric",
			input:    "hello123",
			expected: "hello123",
		},
		{
			name:     "with spaces",
			input:    "hello world 123",
			expected: "hello world 123",
		},
		{
			name:     "with punctuation",
			input:    "hello, world! 123.",
			expected: "hello world 123",
		},
		{
			name:     "with special characters",
			input:    "hello@world#123$%^&*()",
			expected: "helloworld123",
		},
		{
			name:     "unicode letters",
			input:    "привет мир 123",
			expected: "привет мир 123",
		},
		{
			name:     "mixed unicode and punctuation",
			input:    "Привет, мир! 123.",
			expected: "Привет мир 123",
		},
		{
			name:     "only special characters",
			input:    "!@#$%^&*()_+-=[]{}|;':\",./<>?",
			expected: "",
		},
		{
			name:     "newlines and tabs",
			input:    "hello\nworld\t123",
			expected: "helloworld123",
		},
		{
			name:     "emoji and unicode",
			input:    "hello 👋 мир 123 🌍",
			expected: "hello  мир 123 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StripNonAlphanumeric(tt.input)
			if result != tt.expected {
				t.Errorf("StripNonAlphanumeric(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStripBlanks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no spaces",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "single space",
			input:    "hello world",
			expected: "helloworld",
		},
		{
			name:     "multiple spaces",
			input:    "hello  world  test",
			expected: "helloworldtest",
		},
		{
			name:     "leading and trailing spaces",
			input:    "  hello world  ",
			expected: "helloworld",
		},
		{
			name:     "only spaces",
			input:    "     ",
			expected: "",
		},
		{
			name:     "unicode with spaces",
			input:    "привет мир",
			expected: "приветмир",
		},
		{
			name:     "mixed unicode and ascii with spaces",
			input:    "hello мир test",
			expected: "helloмирtest",
		},
		{
			name:     "tabs and newlines (should remain)",
			input:    "hello\tworld\ntest",
			expected: "hello\tworld\ntest",
		},
		{
			name:     "emoji with spaces",
			input:    "👋 🌍 🌟",
			expected: "👋🌍🌟",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StripBlanks(tt.input)
			if result != tt.expected {
				t.Errorf("StripBlanks(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkReverse(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkReverseUnicode(b *testing.B) {
	s := "Быстрая коричневая лиса прыгает через ленивую собаку"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkStripNonAlphanumeric(b *testing.B) {
	s := "Hello, world! 123 @#$%^&*()_+"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StripNonAlphanumeric(s)
	}
}

func BenchmarkStripBlanks(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StripBlanks(s)
	}
}

// Example tests
func ExampleReverse() {
	result := Reverse("hello")
	fmt.Println(result)
	// Output: olleh
}

func ExampleStripNonAlphanumeric() {
	result := StripNonAlphanumeric("Hello, world! 123")
	fmt.Println(result)
	// Output: Hello world 123
}

func ExampleStripBlanks() {
	result := StripBlanks("hello world test")
	fmt.Println(result)
	// Output: helloworldtest
}
