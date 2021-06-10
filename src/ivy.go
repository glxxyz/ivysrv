package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"robpike.io/ivy/config"
	"robpike.io/ivy/exec"
	"robpike.io/ivy/parse"
	"robpike.io/ivy/run"
	"robpike.io/ivy/scan"
	"robpike.io/ivy/value"
	"strings"
)

func ivyHandler(w http.ResponseWriter, r *http.Request) {
	withoutSlash := r.RequestURI[1:]
	unescaped, err := url.PathUnescape(withoutSlash)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, "404 Not Found\n\n%v\n", err)
		return
	}

	if unescaped == "" {
		_, _ = fmt.Fprintln(w, "ivysrv")
		return
	}

	result, err := runCommand(unescaped)
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintf(w, "400 Bad Request\n\n%v", err)
		return
	}

	_, err = fmt.Fprintf(w, "%v", result)
	if err != nil {
		w.WriteHeader(500)
		_, _ = fmt.Fprintf(w, "500 Internal Server Error\n\n%v", err)
		return
	}
}

// runCommand executes command as an ivy program and returns the result as a string.
func runCommand(command string) (string, error) {
	var conf config.Config
	var result strings.Builder
	var errBuff strings.Builder

	conf.SetOutput(&result)
	conf.SetErrOutput(&errBuff)
	context := exec.NewContext(&conf)

	runContext(context, command)

	var err error
	if errBuff.Len() > 0 {
		err = errors.New(errBuff.String())
	}

	return result.String(), err
}

// runContext executes command as an ivy program on a given context.
func runContext(context value.Context, command string) {
	reader := strings.NewReader(command)
	scanner := scan.New(context, "<input>", reader)
	parser := parse.NewParser("<input>", scanner, context)
	run.Run(parser, context, false)
}
