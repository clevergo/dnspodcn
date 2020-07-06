// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		id, token string
	}{
		{"foo", "bar"},
		{"fizz", "buzz"},
	}

	for _, test := range cases {
		c := NewClient(test.id, test.token)
		assert.Equal(t, test.id, c.AppID)
		assert.Equal(t, test.token, c.AppToken)
		assert.Equal(t, apiBaseURL, c.BaseURL)
	}
}
