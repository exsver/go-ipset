package ipset

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Config struct {
	// Path to ipset bin
	// Examples:
	//   "/usr/sbin/ipset"
	Path string
	// Set name
	Set string
	// Logger - debug logger
	Logger *log.Logger
	// IgnoreExist adds "-exist" option
	IgnoreExist bool
}

func (c *Config) String() string {
	return fmt.Sprintf("Path: '%s', Set: '%s'", c.Path, c.Set)
}

func NewConfig(path string, set string) (*Config, error) {
	return &Config{
		Path:   path,
		Set:    set,
		Logger: log.New(io.Discard, "", 0),
	}, nil
}

func (c *Config) SetLogger(logger *log.Logger) {
	c.Logger = logger
}

func (c *Config) SetIgnoreExist(ignore bool) {
	c.IgnoreExist = ignore
}

func (c *Config) do(args []string) error {
	stdout, stderr, err := c.exec(args)
	if err != nil {
		return err
	}

	if stderr != "" {
		return errors.New(stderr)
	}

	if stdout != "" {
		return errors.New(stdout)
	}

	return nil
}

func (c *Config) get(args []string) (string, error) {
	out, err := c.execCombined(args)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (c *Config) exec(args []string) (string, string, error) {
	c.Logger.Printf("exec '%s %s'", c.Path, strings.Join(args, " "))
	cmd := exec.CommandContext(context.Background(), c.Path, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		c.Logger.Printf("exec error '%s' '%s' '%s'", strings.TrimSpace(stdout.String()), strings.TrimSpace(stderr.String()), err.Error())
		return stdout.String(), stderr.String(), fmt.Errorf("error while executing command: %v", err)
	}

	return stdout.String(), stderr.String(), nil
}

func (c *Config) execCombined(args []string) (string, error) {
	c.Logger.Printf("exec '%s %s'", c.Path, strings.Join(args, " "))
	cmd := exec.CommandContext(context.Background(), c.Path, args...)

	combinedOutput, err := cmd.CombinedOutput()
	combinedOutputString := strings.TrimSpace(string(combinedOutput))

	if err != nil {
		c.Logger.Printf("exec error '%s' '%s'", combinedOutputString, err.Error())
	}

	return combinedOutputString, err
}
