package main

import (
	"strings"
	"time"

	"github.com/pkg/errors"
)

func add(filename string) error {
	_, err := execute("git", "add", filename)
	if err != nil {
		return err
	}

	return nil
}

func commitInDate(date time.Time, allowEmpty bool) error {
	dateStr := date.Format(time.RFC3339)

	args := []string{
		"commit", "--date", dateStr, "-am", "this is a stub commit",
	}

	if allowEmpty {
		args = append(args, "--allow-empty")
	}

	extra, err := execute("git", args...)
	if err != nil {
		return errors.Wrapf(err, "extra: %s", extra)
	}

	return nil
}

func rebaseDates() error {
	_, err := execute("git", "rebase", "--committer-date-is-author-date", "anchor-tag")
	if err != nil {
		return err
	}

	return nil
}

func pushCurrent(shouldForce bool) error {
	args := []string{
		"push",
	}
	if shouldForce {
		args = append(args, "-f")
	}
	_, err := execute("git", args...)
	if err != nil {
		return err
	}

	return nil
}

func currentBranch() (string, error) {
	result, err := execute("git", "branch", "--show-current")
	if err != nil {
		return "", err
	}

	return strings.Trim(result, "\n\t "), nil
}
