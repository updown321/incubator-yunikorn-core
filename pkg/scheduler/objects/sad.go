package objects

import (
	"github.com/apache/incubator-yunikorn-core/pkg/common/resources"
	k8score "k8s.io/kubernetes/pkg/apis/core"
	"log"
	"sort"
	"strconv"
	"strings"
)

type tree struct {
	self     int
	parent   int
	children []int
	visited  bool
}

type aNode struct {
	NodeID            string
	Address           k8score.NodeAddress
	availableResource *resources.Resource
	totalResource     *resources.Resource
	topology_line_num int
	topologyTree      tree
	topologyStarNum   int
	local             bool
}

type aPod struct {
	//Address string
	Address k8score.PodIP
}

func start(nodes []*aNode, pod aPod) {
	// nodes.availableResource.Resource["VCORE"]
	// nodes.availableResource.Resource["MEMORY"]
	// nodes.Address
	nodes = compareIP(nodes, pod)

	local := make([]*aNode, 0)
	nlocal := make([]*aNode, 0)
	for i := range nodes {
		if nodes[i].local == true {
			local = append(local, nodes[i])
		} else {
			nlocal = append(nlocal, nodes[i])
		}
	}

	first := make([]tree, 20)
	for i := 0; i < len(nodes); i += 25 {
		first[i/25] = nodes[i].topologyTree
	}
	for i := range first {
		log.Println("firstTree ", i, first[i])
	}

	nodes = sortStarTopology(local, nlocal, local[0].topologyStarNum)
	//nodes = sortLineTopology(nodes)
	//nodes = sortTreeTopology(nodes, local[0].topologyTree.self, first)

	log.Println("sad: want ", pod)
	for i := range nodes {
		log.Println("sad: ", "["+strconv.Itoa(i+1)+"]", nodes[i].NodeID, " / ", nodes[i].Address.Address, " / ", nodes[i].availableResource)
	}

}

func sadSort(nodes []*aNode) []*aNode {
	sort.SliceStable(nodes, func(i, j int) bool {
		l := nodes[i]
		r := nodes[j]
		return resources.CompUsageShares(r.availableResource.Clone(), l.availableResource.Clone()) > 0
	})

	return nodes
}

func sortTreeTopology(nodes []*aNode, localSelf int, first []tree) (result []*aNode) {
	queue := make([]int, 2)
	queue[0] = localSelf
	queue[1] = -1
	first[localSelf].visited = true
	var j int
	var cnt int
	j, cnt = 0, 1
	for cnt < 20 {
		//queue[j]!=0 ~ queue[j]!=0
		for ; queue[j] != -1; j++ {

			//one queue[j]
			for i := range first {
				if first[i].self == queue[j] {

					//have parent and children
					if first[i].parent != -1 && first[i].children[0] != -1 {

						//parent
						if first[first[i].parent].visited == false {
							queue = append(queue, first[i].parent)
							cnt++
						}

						//children
						for c := range first[i].children {
							if first[first[i].children[c]].visited == false {
								queue = append(queue, first[i].children[c])
								cnt++
							}
						}

						// have only children
					} else if first[i].parent == -1 && first[i].children[0] != -1 {
						for c := range first[i].children {
							if first[first[i].children[c]].visited == false {
								queue = append(queue, first[i].children[c])
								cnt++
							}
						}

						// have only parent
					} else if first[i].parent != -1 && first[i].children[0] == -1 {
						if first[first[i].parent].visited == false {
							queue = append(queue, first[i].parent)
							cnt++
						}

					}

					first[i].visited = true
				}
			}
		}
		queue = append(queue, -1)
		j++
	}

	tmp := make([]*aNode, 0)
	for i := range queue {
		if queue[i] != -1 {
			for n := range nodes {
				if nodes[n].topologyTree.self == queue[i] {
					tmp = append(tmp, nodes[n])
				}
			}
		} else {
			sadSort(tmp)

			for t := range tmp {
				result = append(result, tmp[t])
			}
			tmp = make([]*aNode, 0)
		}
	}

	return
}

func sortLineTopology(nodes []*aNode) (result []*aNode) {
	// 20 group, 0~19
	var max int
	var r int
	var l int
	max = 19

	queue := make([]int, 0)

	for n := range nodes {
		if nodes[n].local == true {
			queue = append(queue, nodes[n].topology_line_num, -1)
			r = nodes[n].topology_line_num + 1
			l = nodes[n].topology_line_num - 1
			break
		}
	}

	for {
		if r <= max {
			queue = append(queue, r)
			r++
		}
		if l >= 0 {
			queue = append(queue, l)
			l--
		}
		queue = append(queue, -1)

		if l == -1 && r > max {
			break
		}
	}

	tmp := make([]*aNode, 0)
	for i := range queue {
		if queue[i] != -1 {
			for n := range nodes {
				if nodes[n].topology_line_num == queue[i] {
					tmp = append(tmp, nodes[n])
				}
			}
		} else {
			sadSort(tmp)

			for t := range tmp {
				result = append(result, tmp[t])
			}
			tmp = make([]*aNode, 0)
		}
	}

	return
}

func sortStarTopology(local []*aNode, nlocal []*aNode, localNum int) (result []*aNode) {
	sadSort(local)
	result = append(local)

	if localNum == 1 {
		sadSort(nlocal)
		for i := range nlocal {
			result = append(result, nlocal[i])
		}

	} else {
		tmp1 := make([]*aNode, 0)
		tmp2 := make([]*aNode, 0)

		for i := range nlocal {
			if nlocal[i].topologyStarNum == 1 {
				tmp1 = append(tmp1, nlocal[i])
			} else {
				tmp2 = append(tmp2, nlocal[i])
			}
		}

		sadSort(tmp1)
		sadSort(tmp2)

		for i := range tmp1 {
			result = append(result, tmp1[i])
		}
		for i := range tmp2 {
			result = append(result, tmp2[i])
		}
	}

	return
}

func compareIP(nodes []*aNode, pod aPod) []*aNode {
	ai := strings.Split(pod.Address.IP, ".")
	for i := range nodes {
		ni := strings.Split(nodes[i].Address.Address, ".")

		if ai[0] == ni[0] && ai[1] == ni[1] && ai[2] == ni[2] {
			nodes[i].local = true
		}
	}

	return nodes
}
