package main

import (
	"context"
	"flag"
	"fmt"
	hermesclient "github.com/william-lbn/hermes/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 如果没有设置 kubeconfig 文件路径，默认加载 ~/.kube/config
	kubeconfig := flag.String("kubeconfig", "/Users/williamlee/pml-dev/2025.02.11/hermes/config", "location to your kubeconfig file")
	flag.Parse()

	// 加载 kubeconfig 配置文件
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %s", err.Error())
	}

	hermesclient := hermesclient.NewForConfigOrDie(config)

	list, err := hermesclient.HermesV1().SubscriberRules("hypermonitor").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return
	}

	fmt.Println(list.Items)

}
