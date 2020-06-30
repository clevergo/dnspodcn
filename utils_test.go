// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDomain(t *testing.T) {
	cases := []struct {
		domain, expected string
	}{
		{"example.com", "example.com"},
		{"example.com.", "example.com"},
		{"abc.example.com", "abc.example.com"},
		{"abc.example.com.", "abc.example.com"},
	}
	for _, test := range cases {
		assert.Equal(t, test.expected, formatDomain(test.domain))
	}
}

func TestFormatSubdomain(t *testing.T) {
	cases := []struct {
		domain, subdomain, expected string
	}{
		{"example.com", "foo.example.com", "foo"},
		{"example.com", "bar.example.com.", "bar"},
	}
	for _, test := range cases {
		assert.Equal(t, test.expected, formatSubdomain(test.domain, test.subdomain))
	}
}

func TestFormatTTL(t *testing.T) {
	cases := []struct {
		ttl      time.Duration
		expected string
	}{
		{60 * time.Second, "60"},
		{time.Hour, "3600"},
	}
	for _, test := range cases {
		assert.Equal(t, test.expected, formatTTL(test.ttl))
	}
}
