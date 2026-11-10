// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"net"
	"strconv"
	"strings"
)

type Config struct {
	Info          *Info          `yaml:"info"`
	Logger        *Logger        `yaml:"logger"`
	Configuration *Configuration `yaml:"configuration"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
	HttpVersion int    `yaml:"http-version,omitempty"`
}

type Logger struct {
	LogLevel string `yaml:"level,omitempty"`
}

type Configuration struct {
	NfStreams          []NFStream       `yaml:"nfStreams,omitempty"`
	AnalyticsStream    *AnalyticsStream `yaml:"analyticsStream,omitempty"`
	ApiServer          ServerAddr       `yaml:"apiServer,omitempty"`
	PrometheusServer   ServerAddr       `yaml:"prometheusServer,omitempty"`
	DebugProfile       ServerAddr       `yaml:"debugProfileServer,omitempty"`
	UserAppApiServer   ServerAddr       `yaml:"userAppApiServer,omitempty"`
	RocEndPoint        ServerAddr       `yaml:"rocEndPoint,omitempty"`
	MetricFuncEndPoint ServerAddr       `yaml:"metricFuncEndPoint,omitempty"`
	ControllerFlag     bool             `yaml:"controllerFlag,omitempty"`
}

type ServerAddr struct {
	Addr         string `yaml:"addr,omitempty"`
	Path         string `yaml:"path,omitempty"`
	Port         int    `yaml:"port,omitempty"`
	PollInterval int    `yaml:"pollInterval,omitempty"`
}

// HostPort returns Addr and Port joined as "host:port", with an IPv6 address
// in brackets.
func (s ServerAddr) HostPort() string {
	return joinHostPort(s.Addr, s.Port)
}

type Urls struct {
	Uri  string `yaml:"uri,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

// HostPort returns Uri and Port joined as "host:port", with an IPv6 address
// in brackets.
func (u Urls) HostPort() string {
	return joinHostPort(u.Uri, u.Port)
}

// joinHostPort joins a host name, IPv4 address or IPv6 address with a port.
// An IPv6 address that is already in brackets, which used to be the only way
// to configure one, is accepted as well and not bracketed twice.
func joinHostPort(host string, port int) string {
	if len(host) > 1 && strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}
	return net.JoinHostPort(host, strconv.Itoa(port))
}

type NFStream struct {
	Topic Topic  `yaml:"topic,omitempty"`
	Urls  []Urls `yaml:"urls,omitempty"`
}

type Topic struct {
	TopicName string `yaml:"topicName,omitempty"`
}

type Groups struct {
	Analytics  string `yaml:"analytics,omitempty"`
	MongoDB    string `yaml:"mongodb,omitempty"`
	RestApis   string `yaml:"restapi,omitempty"`
	Prometheus string `yaml:"prometheus,omitempty"`
}

type AnalyticsStream struct {
	TopicName string   `yaml:"topicName,omitempty"`
	Urls      []string `yaml:"urls,omitempty"`
	Enable    bool     `yaml:"enable,omitempty"`
}
