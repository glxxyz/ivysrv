package main

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
)

func TestHandlerSuccess(t *testing.T) {
	parameters := []struct {
		input, expected string
	}{
		{"2+2", "4\n"},
		{"iota 10", "1 2 3 4 5 6 7 8 9 10\n"},
		{"iota%2010", "1 2 3 4 5 6 7 8 9 10\n"},
		{"'literal'", "literal\n"},
		{"1/2+1/4", "3/4\n"},
		{"1%2F2%2B1%2f4", "3/4\n"},
	}

	for i, param := range parameters {
		w := successResponseWriter{t: t}
		r := http.Request{RequestURI: "/" + param.input}
		ivyHandler(&w, &r)
		actual := w.buff.String()
		if actual != param.expected {
			t.Errorf("Test[%v] expected: %v, actual: %v", i, param.expected, actual)
		}
	}
}

func TestHandlerFailure(t *testing.T) {
	parameters := []struct {
		input    string
		expected int
	}{
		{"blah", 400},
		{"%ZZ", 404},
	}

	for i, param := range parameters {
		w := failureResponseWriter{t: t}
		r := http.Request{RequestURI: "/" + param.input}
		ivyHandler(&w, &r)
		actual := w.statusCode
		if actual != param.expected {
			t.Errorf("Test[%v] expected: %v, actual: %v", i, param.expected, actual)
		}
	}
}

func TestBadWriter(t *testing.T) {
	w := badResponseWriter{t: t}
	r := http.Request{RequestURI: "/2+2"}
	expected := 500
	ivyHandler(&w, &r)
	actual := w.statusCode
	if actual != expected {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

type successResponseWriter struct {
	t    *testing.T
	buff bytes.Buffer
}

func (w *successResponseWriter) Header() http.Header {
	w.t.Error("Unexpected call to Header()")
	return nil
}

func (w *successResponseWriter) Write(bytes []byte) (int, error) {
	return w.buff.Write(bytes)
}

func (w *successResponseWriter) WriteHeader(int) {
	w.t.Error("Unexpected call to WriteHeader()")
}

type failureResponseWriter struct {
	t          *testing.T
	buff       bytes.Buffer
	statusCode int
}

func (w *failureResponseWriter) Header() http.Header {
	w.t.Error("Unexpected call to Header()")
	return nil
}

func (w *failureResponseWriter) Write(bytes []byte) (int, error) {
	return w.buff.Write(bytes)
}

func (w *failureResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode != 0 {
		w.t.Errorf("Unexpected multiple calls to WriteHeader() previous: %v current: %v", w.statusCode, statusCode)
		return
	}
	w.statusCode = statusCode
}

type badResponseWriter struct {
	t          *testing.T
	statusCode int
}

func (w *badResponseWriter) Header() http.Header {
	w.t.Error("Unexpected call to Header()")
	return nil
}

func (w *badResponseWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func (w *badResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode != 0 {
		w.t.Errorf("Unexpected multiple calls to WriteHeader() previous: %v current: %v", w.statusCode, statusCode)
		return
	}
	w.statusCode = statusCode
}
