package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// use of flag module
	// Here, we are expecting the kubeconfig file
	// if anyone doesn't specify the file then we can pick the config file from default location
	// which is ~/.kube/config, this is the first way
	// Now, another way, if we are trying to access cluster inside the cluster itself, in that case, we won't have
	// access to the kubeconfig file, bcz we are running inside k8s cluster
	kubeconfig := flag.String("kubeconfig", "/home/priyankasalunke/.kube/config", "Specify the location of kubeconfig file")

	// now as we have a kubeconfig file, we can create a client who can communicate to the api server
	// of the cluster using this kubeconfig file
	// here, all the resource methods are exposed in the clientcmd
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	//fmt.Println("Config:", config)

	// Now, we want to create an actual client sets, which can be done using kubernetes module
	// this clientset we can use to interact with the resources (i.e pod resources)
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	// Now, we want to list pods
	// here, 2nd arg for the list function is listoptions, which comes from apimachinery. so let's import it
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error occured during pod list : ", err)
	}

	// pods is of list type
	fmt.Println("Pods from default namespace..")
	for _, pod := range pods.Items {
		fmt.Printf("\n%s", pod.Name)
	}

	// Get deployments
	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error occured during deployment list : ", err)
	}

	// deployment is of list type
	fmt.Println("\n---------------")
	fmt.Println("Deployments from default namespace..")
	for _, deploy := range deployments.Items {
		fmt.Printf("\n%s", deploy.Name)
	}

	// let's see the incluster config
	// For this, we also need to create a deployment. From local machine, we can not access the cluster.
	// we require service host and service port to access the cluster
	// This is just for the reference
	InClusterConfig()
}
