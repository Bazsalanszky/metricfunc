// SPDX-FileCopyrightText: 2026 Balázs Toldi
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"net/url"
	"testing"

	"github.com/omec-project/metricfunc/config"
)

func TestRocServiceUrl(t *testing.T) {
	for _, tc := range []struct {
		addr     string
		want     string
		wantHost string
	}{
		{addr: "aether-roc.svc", want: "http://aether-roc.svc:31194", wantHost: "aether-roc.svc"},
		{addr: "10.0.0.1", want: "http://10.0.0.1:31194", wantHost: "10.0.0.1"},
		{addr: "fd00::1", want: "http://[fd00::1]:31194", wantHost: "fd00::1"},
	} {
		got := rocServiceUrl(config.ServerAddr{Addr: tc.addr, Port: 31194})
		if got != tc.want {
			t.Errorf("rocServiceUrl(%q) = %q, want %q", tc.addr, got, tc.want)
			continue
		}
		u, err := url.Parse(got + "/aether-roc-api/targets")
		if err != nil {
			t.Errorf("url.Parse(%q): %v", got, err)
			continue
		}
		if u.Hostname() != tc.wantHost || u.Port() != "31194" {
			t.Errorf("%q parses to host %q port %q, want %q 31194", got, u.Hostname(), u.Port(), tc.wantHost)
		}
	}
}
