package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

type CodeError struct {
	Code int
	Err  string
}

func (ce CodeError) Error() string {
	return ce.Err
}

type Logger struct {
	l   *slog.Logger
	ctx context.Context
}

func (l Logger) Write(p []byte) (n int, err error) {
	ss := strings.SplitSeq(string(p), "\n")
	for s := range ss {
		if s == "" {
			continue
		}

		s = strings.TrimSpace(s)

		level := slog.LevelInfo

		if strings.Contains(strings.ToLower(s), "warn") {
			level = slog.LevelWarn
		}

		if strings.Contains(strings.ToLower(s), "error") {
			level = slog.LevelError
		}

		l.l.Log(l.ctx, level, s)
	}
	return len(p), nil
}

func codeError(code int, error string, args ...string) CodeError {
	if args == nil {
		return CodeError{
			Code: code,
			Err:  error,
		}
	}
	return CodeError{
		Code: code,
		Err:  fmt.Sprintf(error, args),
	}
}

func splitOnMarker(s []string, marker string) (left, right []string) {
	for i, v := range s {
		if v == marker {
			left = s[:i]
			right = s[i+1:]
			return
		}
	}
	return s, nil
}

func args() (tags map[string]string, cmd []string, err error) {
	tags = map[string]string{}

	c, cmd := splitOnMarker(os.Args[1:], "--")
	if len(cmd) == 0 {
		return nil, nil, codeError(64, "no command provided")
	}

	for i := 0; i < len(c); i++ {
		if !strings.HasPrefix(c[i], "--") {
			return nil, nil, codeError(64, "malformed arguments: value without tag")
		}

		// Handle --tag=value
		if strings.Contains(c[i], "=") {
			sc := strings.Split(c[i], "=")
			if len(sc) != 2 {
				return nil, nil, codeError(64, "malformed arguments: tag without value after =")
			}

			t := strings.TrimSpace(strings.TrimLeft(sc[0], "-"))
			v := strings.TrimSpace(sc[1])
			tags[t] = v
			continue
		}

		// Handle --tag value
		if len(c) < i+2 {
			return nil, nil, codeError(64, "malformed arguments: tag without value")
		}

		if strings.HasPrefix(c[i+1], "-") {
			return nil, nil, codeError(64, "malformed arguments: two tags after each other")
		}

		t := strings.TrimSpace(strings.TrimLeft(c[i], "-"))
		v := strings.TrimSpace(c[i+1])
		tags[t] = v

		i++
		continue
	}

	return tags, cmd, nil
}

func run(ctx context.Context, l *slog.Logger) error {
	tags, cmd, err := args()
	if err != nil {
		return err
	}

	attrs := []any{}
	for t, v := range tags {
		attrs = append(attrs, t, v)
	}

	l = l.With(attrs...)

	cmdName := cmd[0]
	cmdArgs := []string{}

	if len(cmd) > 1 {
		cmdArgs = cmd[1:]
	}

	c := exec.CommandContext(ctx, cmdName, cmdArgs...)

	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()
	err = c.Start()
	if err != nil {
		return err
	}

	ioLogger := Logger{l: l, ctx: ctx}

	go io.Copy(ioLogger, stdout)
	go io.Copy(ioLogger, stderr)

	err = c.Wait()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return codeError(exitErr.ExitCode(), "command exited with non-zero code")
		}
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	err := run(ctx, logger)
	if err != nil {
		cerr, ok := err.(CodeError)
		if ok {
			logger.ErrorContext(ctx, cerr.Err, "code", cerr.Code)
			os.Exit(cerr.Code)
		}

		logger.ErrorContext(ctx, "an error occured", "error", err.Error())
		os.Exit(1)
	}
}
