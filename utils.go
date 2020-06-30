// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"strconv"
	"strings"
	"time"
)

// formatDomain removes the dot suffix of a domain.
func formatDomain(domain string) string {
	if domain[len(domain)-1] == '.' {
		return domain[:len(domain)-1]
	}

	return domain
}

// formatSubdomain removes the dot suffix and root domain of subdomain.
func formatSubdomain(domain, subdomain string) string {
	return formatDomain(strings.TrimSuffix(formatDomain(subdomain), domain))
}

// formatTTL returns a string that represents the seconds of ttl.
func formatTTL(ttl time.Duration) string {
	return strconv.Itoa(int(ttl.Seconds()))
}
