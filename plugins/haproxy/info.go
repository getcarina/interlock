package haproxy

import (
	"github.com/getcarina/interlock"
)

const (
	pluginName        = "haproxy"
	pluginVersion     = "0.1"
	pluginDescription = "haproxy load balancer and reverse proxy"
	pluginUrl         = "https://github.com/getcarina/interlock/tree/master/plugins/haproxy"
)

var (
	pluginInfo = &interlock.PluginInfo{
		Name:        pluginName,
		Version:     pluginVersion,
		Description: pluginDescription,
		Url:         pluginUrl,
	}
)
