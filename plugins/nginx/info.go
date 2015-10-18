package nginx

import (
	"github.com/rgbkrk/interlocarina"
)

const (
	pluginName        = "nginx"
	pluginVersion     = "0.1"
	pluginDescription = "nginx plugin"
	pluginUrl         = "https://github.com/rgbkrk/interlocarina/tree/master/plugins/nginx"
)

var (
	pluginInfo = &interlock.PluginInfo{
		Name:        pluginName,
		Version:     pluginVersion,
		Description: pluginDescription,
		Url:         pluginUrl,
	}
)
