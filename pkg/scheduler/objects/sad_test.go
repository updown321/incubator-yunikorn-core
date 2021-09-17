package objects

import (
	k8score "k8s.io/kubernetes/pkg/apis/core"
	"log"
	"strconv"
	"testing"

	"github.com/apache/incubator-yunikorn-core/pkg/common/resources"
)

func TestSad(t *testing.T) {
	nodelist := make([]*aNode, 500)
	var pod aPod
	var v resources.Quantity
	var m resources.Quantity

	for i := 0; i < 25; i++ {
		v = resources.Quantity((i + 1) * 4)
		m = resources.Quantity(100 - v)
		nodelist[i] = anewNodeRes("node-"+strconv.Itoa(i), map[string]resources.Quantity{"vcore": v, "memory": m}, "120.11.11."+strconv.Itoa(i+1), 0, tree{self: 0, parent: -1, children: []int{1, 2}, visited: false}, 1)
		nodelist[i+25] = anewNodeRes("node-"+strconv.Itoa(i+25), map[string]resources.Quantity{"vcore": v, "memory": m}, "125.21.11."+strconv.Itoa(i+1), 1, tree{self: 1, parent: 0, children: []int{3, 4}, visited: false}, 0)
		nodelist[i+50] = anewNodeRes("node-"+strconv.Itoa(i+50), map[string]resources.Quantity{"vcore": v, "memory": m}, "130.11.44."+strconv.Itoa(i+1), 2, tree{self: 2, parent: 0, children: []int{5, 6}, visited: false}, 0)
		nodelist[i+75] = anewNodeRes("node-"+strconv.Itoa(i+75), map[string]resources.Quantity{"vcore": v, "memory": m}, "135.53.20."+strconv.Itoa(i+1), 3, tree{self: 3, parent: 1, children: []int{7, 8}, visited: false}, 0)
		nodelist[i+100] = anewNodeRes("node-"+strconv.Itoa(i+100), map[string]resources.Quantity{"vcore": v, "memory": m}, "140.111.11."+strconv.Itoa(i+1), 4, tree{self: 4, parent: 1, children: []int{9, 10}, visited: false}, 0)
		nodelist[i+125] = anewNodeRes("node-"+strconv.Itoa(i+125), map[string]resources.Quantity{"vcore": v, "memory": m}, "145.15.141."+strconv.Itoa(i+1), 5, tree{self: 5, parent: 2, children: []int{11, 12}, visited: false}, 0)
		nodelist[i+150] = anewNodeRes("node-"+strconv.Itoa(i+150), map[string]resources.Quantity{"vcore": v, "memory": m}, "150.511.311."+strconv.Itoa(i+1), 6, tree{self: 6, parent: 2, children: []int{13, 14}, visited: false}, 0)
		nodelist[i+175] = anewNodeRes("node-"+strconv.Itoa(i+175), map[string]resources.Quantity{"vcore": v, "memory": m}, "155.112.121."+strconv.Itoa(i+1), 7, tree{self: 7, parent: 3, children: []int{15, 16}, visited: false}, 0)
		nodelist[i+200] = anewNodeRes("node-"+strconv.Itoa(i+200), map[string]resources.Quantity{"vcore": v, "memory": m}, "160.114.11."+strconv.Itoa(i+1), 8, tree{self: 8, parent: 3, children: []int{17, 18}, visited: false}, 0)
		nodelist[i+225] = anewNodeRes("node-"+strconv.Itoa(i+225), map[string]resources.Quantity{"vcore": v, "memory": m}, "165.101.101."+strconv.Itoa(i+1), 9, tree{self: 9, parent: 4, children: []int{19}, visited: false}, 0)

		nodelist[i+250] = anewNodeRes("node-"+strconv.Itoa(i+250), map[string]resources.Quantity{"vcore": v, "memory": m}, "170.911.181."+strconv.Itoa(i+1), 10, tree{self: 10, parent: 4, children: []int{-1}, visited: false}, 0)
		nodelist[i+275] = anewNodeRes("node-"+strconv.Itoa(i+275), map[string]resources.Quantity{"vcore": v, "memory": m}, "175.191.113."+strconv.Itoa(i+1), 11, tree{self: 11, parent: 5, children: []int{-1}, visited: false}, 0)
		nodelist[i+300] = anewNodeRes("node-"+strconv.Itoa(i+300), map[string]resources.Quantity{"vcore": v, "memory": m}, "180.441.166."+strconv.Itoa(i+1), 12, tree{self: 12, parent: 5, children: []int{-1}, visited: false}, 0)
		nodelist[i+325] = anewNodeRes("node-"+strconv.Itoa(i+325), map[string]resources.Quantity{"vcore": v, "memory": m}, "185.661.501."+strconv.Itoa(i+1), 13, tree{self: 13, parent: 6, children: []int{-1}, visited: false}, 0)
		nodelist[i+350] = anewNodeRes("node-"+strconv.Itoa(i+350), map[string]resources.Quantity{"vcore": v, "memory": m}, "190.561.201."+strconv.Itoa(i+1), 14, tree{self: 14, parent: 6, children: []int{-1}, visited: false}, 0)
		nodelist[i+375] = anewNodeRes("node-"+strconv.Itoa(i+375), map[string]resources.Quantity{"vcore": v, "memory": m}, "195.110.011."+strconv.Itoa(i+1), 15, tree{self: 15, parent: 7, children: []int{-1}, visited: false}, 0)
		nodelist[i+400] = anewNodeRes("node-"+strconv.Itoa(i+400), map[string]resources.Quantity{"vcore": v, "memory": m}, "200.121.210."+strconv.Itoa(i+1), 16, tree{self: 16, parent: 7, children: []int{-1}, visited: false}, 0)
		nodelist[i+425] = anewNodeRes("node-"+strconv.Itoa(i+425), map[string]resources.Quantity{"vcore": v, "memory": m}, "205.141.151."+strconv.Itoa(i+1), 17, tree{self: 17, parent: 8, children: []int{-1}, visited: false}, 0)
		nodelist[i+450] = anewNodeRes("node-"+strconv.Itoa(i+450), map[string]resources.Quantity{"vcore": v, "memory": m}, "210.131.121."+strconv.Itoa(i+1), 18, tree{self: 18, parent: 8, children: []int{-1}, visited: false}, 0)
		nodelist[i+475] = anewNodeRes("node-"+strconv.Itoa(i+475), map[string]resources.Quantity{"vcore": v, "memory": m}, "215.111.110."+strconv.Itoa(i+1), 19, tree{self: 19, parent: 9, children: []int{-1}, visited: false}, 0)

	}

	pod.Address.IP = "150.511.311.26"
	start(nodelist, pod)
	log.Println("sad: " + "start success!")
}

func anewNodeRes(nodeID string, availableMap map[string]resources.Quantity, address string, topologyNumber int, forTree tree, starNum int) *aNode {
	available := resources.NewResourceFromMap(availableMap)
	total := resources.NewResourceFromMap(map[string]resources.Quantity{"vcore": 100, "memory": 100})
	return anewNodeInternal(nodeID, available, total, address, topologyNumber, forTree, starNum)
}

func anewNodeInternal(nodeID string, available *resources.Resource, total *resources.Resource, address string, topologyNumber int, forTree tree, starNum int) *aNode {
	return &aNode{
		NodeID:            nodeID,
		Address:           k8score.NodeAddress{Type: "NodeExternalIP", Address: address},
		availableResource: available,
		totalResource:     total,
		topology_line_num: topologyNumber,
		topologyTree:      forTree,
		local:             false,
		topologyStarNum:   starNum,
	}
}
