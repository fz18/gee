package gee

import "strings"

// trie树以及相关插入，查询方法

// define tree's node
type node struct {
	pattern  string  // 完整路由
	part     string  // 当前节点
	children []*node // 子节点
	isWild   bool    // 当前节点是否精确匹配 part 含有 : 或 * 不是精确匹配， True
}

func (n *node) Insert(pattern string, parts []string, height int) {
	// 递归出口
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 找位置插入
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 如果没找到就新建一个
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// 插入
	child.Insert(pattern, parts, height+1)
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) Search(parts []string, height int) *node {
	// 递归找
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		res := child.Search(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
