// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValuesSet(t *testing.T) {
	cases := []struct {
		key, value string
	}{
		{"foo", ""},
		{"foo", "bar"},
		{"foo", "fizz"},
		{"foo", "buzz"},
	}
	vs := Values{}
	for _, test := range cases {
		vs.Set(test.key, test.value)
		assert.Equal(t, test.value, vs[test.key])
	}
}

func TestValuesGet(t *testing.T) {
	vs := Values{
		"foo":  "bar",
		"fizz": "buzz",
	}
	assert.Equal(t, "", vs.Get("nil"))
	assert.Equal(t, "bar", vs.Get("foo"))
	assert.Equal(t, "buzz", vs.Get("fizz"))
}

func TestValuesDel(t *testing.T) {
	vs := Values{
		"foo": "bar",
	}
	vs.Del("foo")
	_, ok := vs["foo"]
	assert.False(t, ok)
}

func TestValuesEncode(t *testing.T) {
	vs := Values{}
	assert.Equal(t, "", vs.Encode())

	vs.Set("foo", "bar")
	assert.Equal(t, "foo=bar", vs.Encode())

	email := "fizz@buzz.com"
	vs.Set("email", email)
	assert.Equal(t, fmt.Sprintf("%s=%s&%s", "email", url.QueryEscape(email), "foo=bar"), vs.Encode())
}
