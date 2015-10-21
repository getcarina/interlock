package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/rackerlabs/libcarina"
	"github.com/getcarina/interlock"
	"github.com/getcarina/interlock/plugins"
	"github.com/getcarina/interlock/version"
)

func waitForInterrupt() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	for _ = range sigChan {
		return
	}
}

func cmdStart(c *cli.Context) {
	username := c.GlobalString("username")
	apiKey := c.GlobalString("api-key")
	clusterName := c.GlobalString("clustername")
	endpoint := c.GlobalString("endpoint")

	if clusterName == "" {
		log.Fatalf("cluster name must not be empty")
	}

	config := &interlock.Config{}

	client, err := libcarina.NewClusterClient(endpoint, username, apiKey)

	if err != nil {
		log.Fatalf("error getting access to the cluster: %s", err)
	}

	swarmURL, tlsConfig, err := client.GetDockerConfig(clusterName)
	if err != nil {
		log.Fatalf("error retrieving tls config and swarm URL: %s", err)
	}

	config.SwarmUrl = swarmURL
	config.EnabledPlugins = c.GlobalStringSlice("plugin")

	m := NewManager(config, tlsConfig)

	log.Infof("interlock running version=%s", version.FullVersion())
	if err := m.Run(); err != nil {
		log.Fatal(err)
	}

	waitForInterrupt()

	log.Infof("shutting down")
	if err := m.Stop(); err != nil {
		log.Fatal(err)
	}
}

func cmdListPlugins(c *cli.Context) {
	allPlugins := plugins.GetPlugins()
	w := tabwriter.NewWriter(os.Stdout, 8, 1, 3, ' ', 0)

	fmt.Fprintln(w, "NAME\tVERSION\tDESCRIPTION\tURL")

	for _, p := range allPlugins {
		i := p.Info()
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			i.Name,
			i.Version,
			i.Description,
			i.Url,
		)
	}
	w.Flush()
}

func cmdInfo(c *cli.Context) {
	haproxyOut, hErr := exec.Command("/usr/sbin/haproxy", "-v").Output()
	if hErr != nil {
		log.Fatal(hErr)
	}

	hData := strings.Split(string(haproxyOut), "\n")

	nginxOut, nErr := exec.Command("/usr/sbin/nginx", "-v").CombinedOutput()
	if nErr != nil {
		log.Fatal(nErr)
	}

	nData := strings.Split(string(nginxOut), "\n")

	fmt.Println("interlock " + version.FullVersion())
	fmt.Println(" " + string(hData[0]))
	fmt.Println(" " + string(nData[0]))
}
