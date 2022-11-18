package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"
	"vipbind/controller"
)

func main(){
	conf, iniErr := ini.Load("/etc/conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ip := conf.Section("k8s").Key("ip").String()
	port := conf.Section("k8s").Key("port").String()
	vip := conf.Section("vip").Key("vip").String()
	name := conf.Section("k8s").Key("name").String()
	// Create configuration
	master := "https://"+ip+":"+port
	config, err := clientcmd.BuildConfigFromFlags(master, clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// Create client
	clientset, err := kubernetes.NewForConfig(config)
	//informer
	//factory := informers.NewSharedInformerFactory(clientset,0)
	//nodeinformer := factory.Core().V1().Nodes()
	//controller.Updatelabel(clientset,name,"no")
	for {
		time.Sleep(10 * time.Second)
		if controller.Iplist(vip) {
			controller.Updatelabel(clientset,name,"yes")
		}else {
			controller.Updatelabel(clientset,name,"no")
		}
	}
}
