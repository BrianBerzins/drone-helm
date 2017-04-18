package main

import (
	"fmt"
	"testing"
)

func testConfig() Config {
	return Config{
		KubeConfig: "kubeconfig",
		Context:    "test-context",
		Release:    "test-release",
		Chart:      "test-chart",
		Values:     "[\"test-values1\", \"test-values2\"]",
		Set:        "test-key1=test-value1,test-key2=test-value2",
		Namespace:  "test-namespace",
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
	if helmInitCmd(testConfig()).Args[4] != "test-namespace" {
		test.Fail()
	}
}
func TestHelmInitInit(test *testing.T) {
	if helmInitCmd(testConfig()).Args[5] != "init" {
		test.Fail()
	}
}
func TestHelmInitSkipRefresh(test *testing.T) {
	if helmInitCmd(testConfig()).Args[6] != "--skip-refresh" {
		test.Fail()
	}
}
func TestHelmInitUpgrade(test *testing.T) {
	if helmInitCmd(testConfig()).Args[7] != "--upgrade" {
		test.Fail()
	}
}

// {kubectl} --context {context} rollout status deployment/tiller-deploy
func TestTillerRolloutPath(test *testing.T) {
	if tillerRolloutCmd(testConfig()).Path != kubectl {
		test.Fail()
	}
}
func TestTillerRolloutContext(test *testing.T) {
	if tillerRolloutCmd(testConfig()).Args[1] != "--context" {
		test.Fail()
	}
}
func TestTillerRolloutContextValue(test *testing.T) {
	if tillerRolloutCmd(testConfig()).Args[2] != "test-context" {
		test.Fail()
	}
}
func TestTillerRolloutRollout(test *testing.T) {
	if tillerRolloutCmd(testConfig()).Args[3] != "rollout" {
		test.Fail()
	}
}
func TestTillerRolloutStatus(test *testing.T) {
	if tillerRolloutCmd(testConfig()).Args[4] != "status" {
		test.Fail()
	}
}
func TestTillerRolloutDeployment(test *testing.T) {
	if tillerRolloutCmd(testConfig()).Args[5] != "deployment/tiller-deploy" {
		test.Fail()
	}
}

// {helm} --kube-context {context} --tiller-namespace {namespace} upgrade --install {release} {chart} --values {values} --set {set} --wait
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
	if helmUpgradeCmd(testConfig()).Args[4] != "test-namespace" {
		test.Fail()
	}
}
func TestHelmUpgradeUpgrade(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[5] != "upgrade" {
		test.Fail()
	}
}
func TestHelmUpgradeInstall(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[6] != "--install" {
		test.Fail()
	}
}
func TestHelmUpgradeRelease(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[7] != "test-release" {
		test.Fail()
	}
}
func TestHelmUpgradeChart(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[8] != "test-chart" {
		test.Fail()
	}
}
func TestHelmUpgradeValues(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[9] != "--values" {
		test.Fail()
	}
}
func TestHelmUpgradeValuesValue(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[10] != "[\"test-values1\", \"test-values2\"]" {
		test.Fail()
	}
}
func TestHelmUpgradeSet(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[11] != "--set" {
		test.Fail()
	}
}
func TestHelmUpgradeSetValues(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[12] != "test-key1=test-value1,test-key2=test-value2" {
		test.Fail()
	}
}
func TestHelmUpgradeWait(test *testing.T) {
	if helmUpgradeCmd(testConfig()).Args[13] != "--wait" {
		test.Fail()
	}
}

// context
// namespace
// release
// chart
// values (files)
// set (specific values)
