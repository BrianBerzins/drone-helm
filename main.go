package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "helm plugin"
	app.Usage = "helm plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kube-config",
			Usage:  "kubernetes config",
			EnvVar: "PLUGIN_KUBE_CONFIG,KUBE_CONFIG",
		},
		cli.StringFlag{
			Name:   "tiller-namespace",
			Usage:  "namespace where tiller resides",
			EnvVar: "PLUGIN_TILLER_NAMESPACE,TILLER_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "context",
			Usage:  "context (from kube/config to use for this command)",
			EnvVar: "PLUGIN_CONTEXT,CONTEXT",
		},
		cli.StringFlag{
			Name:   "release",
			Usage:  "Kubernetes helm release",
			EnvVar: "PLUGIN_RELEASE,RELEASE",
		},
		cli.StringFlag{
			Name:   "chart",
			Usage:  "Kubernetes helm chart",
			EnvVar: "PLUGIN_CHART,CHART",
		},
		cli.StringFlag{
			Name:   "values",
			Usage:  "files with chart values (e.g. [\"file1.yml\",\"file2.yml\"])",
			EnvVar: "PLUGIN_VALUES,VALUES",
		},
		cli.StringFlag{
			Name:   "set",
			Usage:  "key value pairs that override values (e.g. test-key1=test-value1,test-key2=test-value2)",
			EnvVar: "PLUGIN_SET,SET",
		},
		cli.StringFlag{
			Name:   "rollout-status",
			Usage:  "a kubernetes resource to wait for before completing",
			EnvVar: "PLUGIN_ROLLOUT_STATUS,ROLLOUT_STATUS",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(context *cli.Context) error {
	if context.String("env-file") != "" {
		_ = godotenv.Load(context.String("env-file"))
	}
	plugin := Plugin{
		Config: Config{
			KubeConfig:      context.String("kube-config"),
			TillerNamespace: context.String("tiller-namespace"),
			Context:         context.String("context"),
			Release:         context.String("release"),
			Chart:           context.String("chart"),
			Values:          context.String("values"),
			Set:             context.String("set"),
			RolloutStatus:   context.String("rollout-status"),
		},
	}
	return plugin.Exec()
}
