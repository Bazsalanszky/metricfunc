// SPDX-FileCopyrightText: 2026 Balázs Toldi
//
// SPDX-License-Identifier: Apache-2.0

package reader

import (
	"slices"
	"testing"

	"github.com/omec-project/metricfunc/config"
)

func TestMakeUrlsFromUriPort(t *testing.T) {
	got := makeUrlsFromUriPort([]config.Urls{
		{Uri: "sd-core-kafka-headless", Port: 9092},
		{Uri: "10.0.0.1", Port: 9092},
		{Uri: "fd00::1", Port: 9092},
	})
	want := []string{"sd-core-kafka-headless:9092", "10.0.0.1:9092", "[fd00::1]:9092"}
	if !slices.Equal(got, want) {
		t.Errorf("makeUrlsFromUriPort() = %v, want %v", got, want)
	}
}
