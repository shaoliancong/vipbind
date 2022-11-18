package controller

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"net"
	"strings"
)

func Iplist(ip string) bool {
	interface_list, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	var byName *net.Interface
	var addrList []net.Addr
	var oneAddrs []string
	for _, i := range interface_list {
		byName, err = net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addrList, err = byName.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, oneAddr := range addrList {
			oneAddrs = strings.SplitN(oneAddr.String(), "/", 2)
			//fmt.Println(oneAddrs[0])
			if oneAddrs[0] == ip {
				return true
			}
		}
	}
	return false
}

func Updatelabel(clientset *kubernetes.Clientset, name, label string, ) {
	node, err := clientset.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("get node[%s] err is %s", name+"\n", err)
		return
	}
	labels := node.Labels
	fmt.Println("old label:", node.Labels)
	if labels["vipbind"] == label {
		return
	} else {
		labels["vipbind"] = label
		patchData := map[string]interface{}{"metadata": map[string]map[string]string{"labels": labels}}
		playLoadBytes, _ := json.Marshal(patchData)

		newNode, err := clientset.CoreV1().Nodes().Patch(name, types.StrategicMergePatchType, playLoadBytes)
		if err != nil {
			fmt.Printf("[Up//datePodByPodSn] %v pod Patch fail %v\n", name, err)
			return
		}
		fmt.Println("new label:", newNode.Labels)
	}
}
