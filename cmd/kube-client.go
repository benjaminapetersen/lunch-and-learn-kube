package main

import (
	"flag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)

	pod, err := clientset.CoreV1().Pods("book").Get("example", metav1.GetOptions{}
}
