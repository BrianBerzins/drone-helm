package main

import (
	"fmt"
	"testing"
)

func testConfig() Config {
	return Config{
		KubeConfig:      "kubeconfig",
		TillerNamespace: "test-tiller-namespace",
		Namespace:       "test-namespace",
		Context:         "test-context",
		Release:         "test-release",
		Chart:           "test-chart",
		Values:          "test-value",
		Set:             "test-key1=test-value1,test-key2=test-value2",
		RolloutStatus:   "test-rollout-status",
	}
}

// {kubectl} config view --output jsonpath='{.contexts[?(@.name == "{context}")].context.namespace}'
func TestKubectlNamespacePath(test *testing.T) {
	if kubectlNamespaceCmd(testConfig()).Path != kubectl {
		test.Fail()
	}
}
func TestKubectlNamespaceConfig(test *testing.T) {
	if kubectlNamespaceCmd(testConfig()).Args[1] != "config" {
		test.Fail()
	}
}
func TestKubectlNamespaceView(test *testing.T) {
	if kubectlNamespaceCmd(testConfig()).Args[2] != "view" {
		test.Fail()
	}
}
func TestKubectlNamespaceOutput(test *testing.T) {
	if kubectlNamespaceCmd(testConfig()).Args[3] != "--output" {
		test.Fail()
	}
}
func TestKubectlNamespaceJson(test *testing.T) {
	if kubectlNamespaceCmd(testConfig()).Args[4] != "jsonpath={.contexts[?(@.name == \"test-context\")].context.namespace}" {
		fmt.Println(kubectlNamespaceCmd(testConfig()).Args[4])
		test.Fail()
	}
}

// {helm} --kube-context {context} --tiller-namespace {namespace} init --skip-refresh --upgrade
func TestHelmInitPath(test *testing.T) {
	if helmInitCmd(testConfig()).Path != helm {
		test.Fail()
	}
}
func TestHelmInitKubeContext(test *testing.T) {
	if helmInitCmd(testConfig()).Args[1] != "--kube-context" {
		test.Fail()
	}
}
func TestHelmInitContext(test *testing.T) {
	if helmInitCmd(testConfig()).Args[2] != "test-context" {
		test.Fail()
	}
}
func TestHelmInitKubeTillerNamespace(test *testing.T) {
	if helmInitCmd(testConfig()).Args[3] != "--tiller-namespace" {
		test.Fail()
	}
}
func TestHelmInitKubeNamespace(test *testing.T) {
	if helmInitCmd(testConfig()).Args[4] != "test-tiller-namespace" {
		test.Fail()
	}
}
func TestHelmInitInit(test *testing.T) {
	if helmInitCmd(testConfig()).Args[5] != "init" {
		test.Fail()
	}
}
func TestHelmInitClient(test *testing.T) {
	if helmInitCmd(testConfig()).Args[6] != "--client-only" {
		test.Fail()
	}
}
func TestHelmInitSkipRefresh(test *testing.T) {
	if helmInitCmd(testConfig()).Args[7] != "--skip-refresh" {
		test.Fail()
	}
}
func TestHelmInitUpgrade(test *testing.T) {
	if helmInitCmd(testConfig()).Args[8] != "--upgrade" {
		test.Fail()
	}
}

// {kubectl} --context {context} rollout status --watch=false {status}
func TestRolloutStatusPath(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Path != kubectl {
		test.Fail()
	}
}
func TestRolloutStatusContext(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Args[1] != "--context" {
		test.Fail()
	}
}
func TestRolloutStatusContextValue(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Args[2] != "test-context" {
		test.Fail()
	}
}
func TestRolloutStatusRollout(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Args[3] != "rollout" {
		test.Fail()
	}
}
func TestRolloutStatusStatus(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Args[4] != "status" {
		test.Fail()
	}
}
func TestRolloutStatusWait(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Args[5] != "--watch=false" {
		test.Fail()
	}
}
func TestRolloutStatusValue(test *testing.T) {
	if kubectlRolloutStatus(testConfig()).Args[6] != "test-rollout-status" {
		test.Fail()
	}
}

// {helm} --kube-context {context} --tiller-namespace {namespace} upgrade {release} {chart} --install --namespace {namespace} --values {values} --set {set} --wait
func TestHelmUpgradePath(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Path != helm {
		test.Fail()
	}
}
func TestHelmUpgradeKubeContext(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[1] != "--kube-context" {
		test.Fail()
	}
}
func TestHelmUpgradeContext(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[2] != "test-context" {
		test.Fail()
	}
}
func TestHelmUpgradeKubeTillerNamespace(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[3] != "--tiller-namespace" {
		test.Fail()
	}
}
func TestHelmUpgradeKubeNamespace(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[4] != "test-tiller-namespace" {
		test.Fail()
	}
}
func TestHelmUpgradeUpgrade(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[5] != "upgrade" {
		test.Fail()
	}
}
func TestHelmUpgradeRelease(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[6] != "test-release" {
		test.Fail()
	}
}
func TestHelmUpgradeChart(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[7] != "test-chart" {
		test.Fail()
	}
}
func TestHelmUpgradeInstall(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[8] != "--install" {
		test.Fail()
	}
}
func TestHelmUpgradeNamespace(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[9] != "--namespace" {
		test.Fail()
	}
}
func TestHelmUpgradeNamespaceValue(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[10] != "test-namespace" {
		test.Fail()
	}
}
func TestHelmUpgradeValues(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[11] != "--values" {
		test.Fail()
	}
}
func TestHelmUpgradeValuesValue(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[12] != "test-value" {
		test.Fail()
	}
}
func TestHelmUpgradeSet(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[13] != "--set" {
		test.Fail()
	}
}
func TestHelmUpgradeSetValues(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[14] != "test-key1=test-value1,test-key2=test-value2" {
		test.Fail()
	}
}
