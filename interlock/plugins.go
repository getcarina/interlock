package main

// interlock plugins
import (
	_ "github.com/getcarina/interlock/plugins/example"
	_ "github.com/getcarina/interlock/plugins/haproxy"
	_ "github.com/getcarina/interlock/plugins/nginx"
	//_ "github.com/getcarina/interlock/plugins/stats"
)
