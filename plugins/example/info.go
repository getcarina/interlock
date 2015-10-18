package example

import (
	"github.com/rgbkrk/interlocarina"
)

const (
	pluginName        = "example"
	pluginVersion     = "0.1"
	pluginDescription = "example plugin"
	pluginUrl         = "https://github.com/rgbkrk/interlocarina/tree/master/plugins/example"
)

var (
	pluginInfo = &interlock.PluginInfo{
		Name:        pluginName,
		Version:     pluginVersion,
		Description: pluginDescription,
		Url:         pluginUrl,
	}
)
