package main

import (
	"fmt"
	"io"
	"os"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdexec "k8s.io/kubectl/pkg/cmd/exec"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func mainKubectl() {
	var inStream io.Reader = os.Stdin
	var outStream io.Writer = os.Stdout
	var errStream io.Writer = os.Stderr
	ioStreams := genericclioptions.IOStreams{In: inStream, Out: outStream, ErrOut: errStream}

	apiServer := "https://api.kube.yourcompany.com:6443"
	bearerToken := "your-bearer-token"
	namespace := "mynamespace"
	insecure := true

	kubeConfigFlags := genericclioptions.NewConfigFlags(false).WithDeprecatedPasswordFlag()
	kubeConfigFlags.APIServer = &apiServer
	kubeConfigFlags.BearerToken = &bearerToken
	kubeConfigFlags.Namespace = &namespace
	kubeConfigFlags.Insecure = &insecure
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)

	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)

	command := cmdexec.NewCmdExec(f, ioStreams)
	command.SetArgs([]string{
		"mynamespace-mydeployment-7bjkw", "-it",
		"--", "ps", "-ef",
	})

	if command == nil {
		os.Exit(1)
	}
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
