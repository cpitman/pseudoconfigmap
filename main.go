package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	configMapName := os.Getenv("CONFIG_MAP_NAME")
	if configMapName == "" {
		panic("Environment Variable CONFIG_MAP_NAME must be set")
	}
	listOptions := metav1.ListOptions{FieldSelector: "metadata.name=" + configMapName}

	watchInterface, err := clientset.CoreV1().ConfigMaps("").Watch(listOptions)

	if errors.IsNotFound(err) {
		fmt.Printf("ConfigMap not found\n")
	} else if err != nil {
		panic(err.Error())
	}

	for {
		event := <-watchInterface.ResultChan()
		if event.Type == watch.Added || event.Type == watch.Modified {
			fmt.Printf("ConfigMap modified")
		}
	}
}
