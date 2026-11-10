// SPDX-FileCopyrightText: 2026 Balázs Toldi
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"testing"

	"github.com/omec-project/metricfunc/config"
)

const (
	testPodIPv4 = "10.42.0.5"
	testPodIPv6 = "2001:db8:600::5"
)

func TestPodIP(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want string
	}{
		{name: "IPv4", in: testPodIPv4, want: testPodIPv4},
		{name: "IPv4-mapped IPv6", in: "::ffff:" + testPodIPv4, want: testPodIPv4},
		{name: "IPv6", in: testPodIPv6, want: testPodIPv6},
		{name: "IPv6 in canonical form", in: "2001:DB8:600:0:0:0:0:5", want: testPodIPv6},
		{name: "unset", in: "", want: ""},
		{name: "not an address", in: "metricfunc-0", want: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := podIP(tc.in); got != tc.want {
				t.Errorf("podIP(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestUserAppServiceUrl(t *testing.T) {
	for _, tc := range []struct {
		addr string
		want string
	}{
		{addr: "userAppapp", want: "http://userAppapp:9301/rogue"},
		{addr: "10.0.0.1", want: "http://10.0.0.1:9301/rogue"},
		{addr: "fd00::1", want: "http://[fd00::1]:9301/rogue"},
	} {
		if got := userAppServiceUrl(config.ServerAddr{Addr: tc.addr, Port: 9301, Path: "/rogue"}); got != tc.want {
			t.Errorf("userAppServiceUrl(%q) = %q, want %q", tc.addr, got, tc.want)
		}
	}
}
