package main

import (
	"os/exec"

	"github.com/pkg/errors"
)

func execute(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.Wrapf(err, "out: %s", string(stdout))
	}

	return string(stdout), nil
}
