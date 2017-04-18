package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const kubectl = "/usr/local/bin/kubectl"
const kubeConfig = "/root/.kube/config"
const helm = "/usr/local/bin/helm"

type (
	// configuration for the plugin
	Config struct {
		KubeConfig string
		Context    string
		Release    string
		Chart      string
		Values     string
		Set        string
		Namespace  string
	}
	// plugin default
	Plugin struct {
		Config Config
	}
)

// Exec default method
func (plugin *Plugin) Exec() error {
	config := plugin.Config

	// check required fields
	if config.KubeConfig == "" {
		return errors.New("KUBE_CONFIG not provided")
	}
	if config.Chart == "" {
		return errors.New("chart not provided")
	}
	if config.Release == "" {
		return errors.New("release not provided")
	}

	// write the kube config file
	writeError := writeKubeConfig(config)
	if writeError != nil {
		return writeError
	}

	// get namespace from kube config file
	namespace, namespaceError := outputFromCmd(kubectlNamespaceCmd(config))
	if namespaceError != nil {
		return namespaceError
	}
	config.Namespace = namespace

	// helm initialize
	initError := executeCmd(helmInitCmd(config))
	if initError != nil {
		return initError
	}

	// wait for helm deployment to complete
	rolloutError := executeCmd(tillerRolloutCmd(config))
	if rolloutError != nil {
		return rolloutError
	}

	// helm upgrade
	upgradeError := executeCmd(helmUpgradeCmd(config))
	if upgradeError != nil {
		return upgradeError
	}

	return nil
}

// write the kube config file from environment variable the default location
func writeKubeConfig(config Config) error {
	fmt.Fprintf(os.Stdout, "writing kube config to %s\n", kubeConfig)
	return ioutil.WriteFile(kubeConfig, []byte(config.KubeConfig), 0644)
}

// get the namespace for tiller from the kube config (using the context to figure out which one)
// {kubectl} config view --output jsonpath='{.contexts[?(@.name == "{context}")].context.namespace}'
func kubectlNamespaceCmd(config Config) *exec.Cmd {
	return exec.Command(kubectl,
		"config",
		"view",
		"--output",
		fmt.Sprintf("jsonpath={.contexts[?(@.name == \"%s\")].context.namespace}", config.Context))
}

// do the helm init in case it is not already there
// {helm} --kube-context {context} --tiller-namespace {namespace} init --skip-refresh --upgrade
func helmInitCmd(config Config) *exec.Cmd {
	return exec.Command(helm,
		"--kube-context",
		config.Context,
		"--tiller-namespace",
		config.Namespace,
		"init",
		"--skip-refresh",
		"--upgrade")
}

// wait for the tiller deployment to be fully rolled out
// {kubectl} --context {context} rollout status deployment/tiller-deploy
func tillerRolloutCmd(config Config) *exec.Cmd {
	return exec.Command(kubectl,
		"--context",
		config.Context,
		"rollout",
		"status",
		"deployment/tiller-deploy")
}

// upgrade our release based on the configuration parameters
// {helm} --kube-context {context} --tiller-namespace {namespace} upgrade --install {release} {chart} --values {values} --set {set} --wait
func helmUpgradeCmd(config Config) *exec.Cmd {
	return exec.Command(helm,
		"--kube-context",
		config.Context,
		"--tiller-namespace",
		config.Namespace,
		"upgrade",
		"--install",
		config.Release,
		config.Chart,
		"--values",
		config.Values,
		"--set",
		config.Set,
		"--wait")
}

func outputFromCmd(cmd *exec.Cmd) (string, error) {
	logCommand(cmd) // log the command to be run
	cmd.Stderr = os.Stderr
	bytes, err := cmd.Output()
	output := string(bytes)
	fmt.Fprintf(os.Stdout, "%s\n", output)
	return output, err
}

func executeCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout // log standard out
	cmd.Stderr = os.Stderr // log standard err
	logCommand(cmd)        // log the command to be run
	return cmd.Run()
}

func logCommand(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}
