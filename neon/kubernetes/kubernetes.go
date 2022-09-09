package kubernetes

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var cs *kubernetes.Clientset

type KubeResource interface {
	v1.Pod | v1.Service
}

func InitKubernetes(inCluster bool, kubePath string) {
	if inCluster {
		initInCluster()
	} else {
		initOutOfCluster(kubePath)
	}
}

func initInCluster() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	cs = clientset
}

func initOutOfCluster(kubePath string) {
	var kubeconfig string
	if kubePath != "" {
		kubeconfig = filepath.Join(kubePath, "config")
	} else {
		home := homedir.HomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	cs = clientset
}

func Pods(c *gin.Context, namespace string) podRepo {
	if cs == nil {
		errors.NewInternal("kubernetes client not initialized", nil).Abort(c)
	}
	return podRepo{
		namespace: namespace,
		c:         c,
		cs:        cs,
	}
}

func Services(c *gin.Context, namespace string) serviceRepo {
	if cs == nil {
		errors.NewInternal("kubernetes client not initialized", nil).Abort(c)
	}
	return serviceRepo{
		namespace: namespace,
		c:         c,
		cs:        cs,
	}
}
