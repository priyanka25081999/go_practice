package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func InClusterConfig() {

	fmt.Println("*****Enter - InClusterConfig*****")
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("Error occured during in-cluster config: ", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error occured during pod list : ", err)
	}

	// pods is of list type
	fmt.Println("Pods from default namespace..")
	for _, pod := range pods.Items {
		fmt.Printf("\n%s", pod.Name)
	}
}
