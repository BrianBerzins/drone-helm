package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

const kubectl = "/usr/local/bin/kubectl"
const kubeConfig = "/root/.kube/config"
const helm = "/usr/local/bin/helm"

type (
	// configuration for the plugin
	Config struct {
		KubeConfig      string
		TillerNamespace string
		Context         string
		Namespace       string
		Release         string
		Chart           string
		Values          string
		Set             string
		RolloutStatus   string
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

	// get namespace to deploy into
	namespace, namespaceErr := outputFromCmd(kubectlNamespaceCmd(config))
	if namespaceErr != nil {
		return namespaceErr
	}
	config.Namespace = namespace

	// if the tiller-namespace was not specified, use the namespace obtained from the kube config
	if config.TillerNamespace == "" {
		fmt.Fprintf(os.Stdout, "tiller namespace was not specified, using namespace %s", config.Namespace)
		config.TillerNamespace = config.Namespace
	}

	// helm initialize
	initError := executeCmd(helmInitCmd(config))
	if initError != nil {
		return initError
	}

	// helm upgrade
	upgradeError := executeCmd(helmUpgradeCmd(config))
	if upgradeError != nil {
		return upgradeError
	}

	// wait for any deployments
	if config.RolloutStatus != "" {
		lastOutput := ""
		logCommand(kubectlRolloutStatus(config))
		for {
			output, err := outputFromCmd(kubectlRolloutStatus(config))
			if err != nil {
				return err
			}
			currentOutput := string(output)
			// only log output if it has changed
			if currentOutput != lastOutput {
				lastOutput = currentOutput
				fmt.Fprint(os.Stdout, lastOutput)
				if strings.Contains(lastOutput, "successfully rolled out") {
					return nil // successful
				}
			} else {
				time.Sleep(time.Second) // don't spam it too hard
			}
		}
	}

	// success
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
// {helm} --kube-context {context} --tiller-namespace {namespace} init --client --skip-refresh --upgrade
func helmInitCmd(config Config) *exec.Cmd {
	return exec.Command(helm,
		"--kube-context",
		config.Context,
		"--tiller-namespace",
		config.TillerNamespace,
		"init",
		"--client-only",
		"--skip-refresh",
		"--upgrade")
}

// wait for a rollout to fully complete
// {kubectl} --context {context} rollout status --watch=false {rollout-status}
func kubectlRolloutStatus(config Config) *exec.Cmd {
	return exec.Command(kubectl,
		"--context",
		config.Context,
		"rollout",
		"status",
		"--watch=false",
		config.RolloutStatus)
}

// upgrade our release based on the configuration parameters
// {helm} --kube-context {context} --tiller-namespace {namespace} upgrade {release} {chart} --namespace {namespace} --install --values {values} --set {set}
func helmUpgradeCmd(config Config) *exec.Cmd {
	return exec.Command(helm,
		"--kube-context",
		config.Context,
		"--tiller-namespace",
		config.TillerNamespace,
		"upgrade",
		config.Release,
		config.Chart,
		"--install",
		"--namespace",
		config.Namespace,
		"--values",
		config.Values,
		"--set",
		config.Set)
}

func outputFromCmd(cmd *exec.Cmd) (string, error) {
	cmd.Stderr = os.Stderr
	bytes, err := cmd.Output()
	output := string(bytes)
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
