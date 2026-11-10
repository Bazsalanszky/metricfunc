// SPDX-FileCopyrightText: 2026 Balázs Toldi
//
// SPDX-License-Identifier: Apache-2.0

package config

import "testing"

func TestHostPort(t *testing.T) {
	for _, tc := range []struct {
		name string
		host string
		port int
		want string
	}{
		{name: "host name", host: "sd-core-kafka-headless", port: 9092, want: "sd-core-kafka-headless:9092"},
		{name: "IPv4", host: "10.0.0.1", port: 9092, want: "10.0.0.1:9092"},
		{name: "IPv6", host: "fd00::1", port: 9092, want: "[fd00::1]:9092"},
		{name: "bracketed IPv6", host: "[fd00::1]", port: 9092, want: "[fd00::1]:9092"},
		{name: "empty host", host: "", port: 9092, want: ":9092"},
		{name: "zero port", host: "10.0.0.1", port: 0, want: "10.0.0.1:0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := (ServerAddr{Addr: tc.host, Port: tc.port}).HostPort(); got != tc.want {
				t.Errorf("ServerAddr{%q, %d}.HostPort() = %q, want %q", tc.host, tc.port, got, tc.want)
			}
			if got := (Urls{Uri: tc.host, Port: tc.port}).HostPort(); got != tc.want {
				t.Errorf("Urls{%q, %d}.HostPort() = %q, want %q", tc.host, tc.port, got, tc.want)
			}
		})
	}
}
