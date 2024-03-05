package gee

import "strings"

type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中一部分
	children []*node //子节点
	isWild   bool    //是否模糊匹配

}

// 匹配路径的
//
//	/p->/p  /p->/:name====/p(能匹配上的第一个孩子节点)
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// （能匹配上的所有孩子节点）
// 所有匹配成功的，查找路径
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 路由的注册与匹配
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		//匹配给最后一节url加完整路径
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		//没有匹配上的孩子结点，给爸爸加上孩子结点
		n.children = append(n.children, child)
	}
	//再给下一节url部分添加注册
	child.insert(pattern, parts, height+1)
}

//匹配

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		//因为*只能有一个
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	//  p  :name
	for _, child := range children {
		//沿着p这条路径匹配走或者沿:name走
		result := child.search(parts, height+1)
		//result匹配上的话就不是nil就说明已经找到了
		if result != nil {
			return result
		}
	}
	return nil
}

// 让list包含n能扩展的所有url结点
func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
