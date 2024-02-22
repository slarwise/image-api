package main

import (
	"context"
	"testing"
)

func TestListTags(t *testing.T) {
	actual, err := getTags(context.TODO(), "ghcr.io/slarwise/version-checker")
	if err != nil {
		t.Fatalf("Got unexpected error: %s", err)
	}
	t.Log(actual)
}
