package haproxy

import (
	"github.com/rgbkrk/interlocarina"
)

const (
	pluginName        = "haproxy"
	pluginVersion     = "0.1"
	pluginDescription = "haproxy load balancer and reverse proxy"
	pluginUrl         = "https://github.com/rgbkrk/interlocarina/tree/master/plugins/haproxy"
)

var (
	pluginInfo = &interlock.PluginInfo{
		Name:        pluginName,
		Version:     pluginVersion,
		Description: pluginDescription,
		Url:         pluginUrl,
	}
)
