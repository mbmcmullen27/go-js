package main

import (
	"context"
	// "flag"
	"fmt"
	// "path/filepath"
	// "time"
	// "encoding/json"
	"syscall/js"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
)

func configure() *kubernetes.Clientset {
	fmt.Print("\nConfiguring...")
	// var kubeconfig *string
	array := js.Global().Get("kubeconfig")
	buf := make([]byte, array.Get("length").Int())
	js.CopyBytesToGo(buf, array)
	
	// fmt.Printf("%v",buf)

	fmt.Print("\nCreate rest config from byte array...")
	// config, err := clientcmd.Load(buf)
	config, err := clientcmd.RESTConfigFromKubeConfig(buf)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("\nno errors???")
	fmt.Print("\n%s",config.String())


	clientset := kubernetes.NewForConfigOrDie(config)
	
	return clientset
}

func main() {
	fmt.Print("\ngetting those pods...")
	clientset := configure()
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})

	// fmt.Printf("%d\n",len(pods.Items))
	fmt.Printf("%s\n", pods)
	
	length := len(pods.Items)

	for i:=0; i<length; i++ {
		if err != nil {
			panic(err.Error())
		}

		name:=pods.Items[i].ObjectMeta.Name
		// data, _ := json.Marshal(pods.Items[i])
		fmt.Printf("%s\n", name)
		// fmt.Printf("%s\n", pods)

	}
}

