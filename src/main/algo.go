package main

import (
	"math"
	"math/rand"
	"sort"
)

type MatchAlgo interface {
	Match(players []Player) [][]string
}
type GraphAlgo struct {
	min int
	res [][]Team
}

func (g *GraphAlgo) Match(players []Player) [][]string {
	//给每个用户分配一个ID(位运算，简化后续的冲突计算逻辑)
	mp := make(map[int]Player)
	for i := 0; i < len(players); i++ {
		mp[1<<i] = players[i]
	}
	//穷举所有组合(组合直接用上面生成的ID计算)
	compose := make([]Team, 0, len(players)*(len(players)-1)/2)
	for i := 0; i < len(players)-1; i++ {
		for j := i + 1; j < len(players); j++ {
			compose = append(compose, Team{1 << i, 1 << j, players[i].Score + players[j].Score})
		}
	}
	if len(compose)%2 == 1 {
		//如果是奇数的话随机抽取两个人作为一个新的组合加入(这两个人可以多打一场)
		p := rand.Intn(len(compose))
		compose = append(compose, compose[p])
	}
	//构造图
	graph := make([][]int, len(compose))
	for i := 0; i < len(compose); i++ {
		graph[i] = make([]int, len(compose))
		for j := 0; j < len(compose); j++ {
			//没有冲突意味着交集为0
			if compose[i].TeamId()&compose[j].TeamId() == 0 {
				w := abs(compose[i].Score - compose[j].Score)
				graph[i][j] = w
			} else {
				graph[i][j] = -1
			}
		}
	}
	matchSize := len(compose) / 2
	g.backtrace(compose, graph, make([][]Team, matchSize), 0, matchSize, 0)
	//输出结果
	//TODO:排序
	r := make([][]string, matchSize)
	for idx, m := range g.res {
		r[idx] = []string{mp[m[0].P1].Name, mp[m[0].P2].Name, mp[m[1].P1].Name, mp[m[1].P2].Name}
	}
	return r
}

func (g *GraphAlgo) backtrace(compose []Team, graph [][]int, path [][]Team, idx int, length int, weight int) {
	if idx >= length {
		if weight < g.min {
			g.min = weight
			g.res = make([][]Team, 0, length)
			for _, t := range path {
				g.res = append(g.res, []Team{t[0], t[1]})
			}
		}
		return
	}
	//提前剪枝，不然O(n!)的复杂度会极其恐怖
	if weight >= g.min {
		return
	}
	//选择两个组合
	//这里需要使用启发式搜索，优先搜索权重少的边，不然很容易超时
	edges := make([]Edge, 0)
	for i := 0; i < len(compose)-1; i++ {
		for j := 0; j < len(compose); j++ {
			if graph[i][j] >= 0 {
				edges = append(edges, Edge{i, j, graph[i][j]})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	for _, edge := range edges {
		//更新图,i,j下所有的线取消
		backup1 := make([]int, len(compose))
		backup2 := make([]int, len(compose))
		backup3 := make([]int, len(compose))
		backup4 := make([]int, len(compose))
		//横线处理
		for k := range compose {
			backup1[k] = graph[edge.Start][k]
			graph[edge.Start][k] = -1
			backup2[k] = graph[edge.End][k]
			graph[edge.End][k] = -1
		}
		//竖线处理
		for k := range compose {
			//横竖线相交的地方需要特殊处理
			if k != edge.Start && k != edge.End {
				backup4[k] = graph[k][edge.End]
				backup3[k] = graph[k][edge.Start]
			}
			graph[k][edge.Start] = -1
			graph[k][edge.End] = -1
		}
		path[idx] = []Team{compose[edge.Start], compose[edge.End]}
		g.backtrace(compose, graph, path, idx+1, length, weight+graph[edge.Start][edge.End])
		//还原现场
		for k := range compose {
			graph[edge.Start][k] = backup1[k]
			graph[edge.End][k] = backup2[k]
		}
		for k := range compose {
			if k != edge.Start && k != edge.End {
				graph[k][edge.Start] = backup3[k]
				graph[k][edge.End] = backup4[k]
			}
		}
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func MakeAlgo() *GraphAlgo {
	return &GraphAlgo{min: math.MaxInt32, res: make([][]Team, 0)}
}

type Team struct {
	P1    int
	P2    int
	Score int
}

func (t *Team) TeamId() int {
	return t.P1 + t.P2
}

type Player struct {
	Name  string
	Score int
}

type Edge struct {
	Start  int
	End    int
	Weight int
}
