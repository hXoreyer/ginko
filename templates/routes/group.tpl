package routes

import (
	"{{.Name}}/middlewares"

	"github.com/gin-gonic/gin"
)

type group struct {
	name        string
	middlewares []gin.HandlerFunc
	handle      *gin.RouterGroup
	initFunc    func(router *gin.RouterGroup)
	children    []group
}

type Group struct {
	groups []group
}

// 递归初始化路由组
func (g *Group) initRouter(groups []group) {
	for i := range groups {
		grp := &groups[i] // 使用指针以确保修改到原始数据
		if grp.handle != nil && grp.initFunc != nil {
			grp.initFunc(grp.handle)
		}
		if len(grp.children) > 0 {
			g.initRouter(grp.children)
		}
	}
}

func (g *Group) Init() {
	g.initRouter(g.groups)
}

// 初始化单个组及其子组
func initGroup(r *gin.Engine, parent *gin.RouterGroup, grp *group) {
	var currentGroup *gin.RouterGroup
	if parent == nil {
		currentGroup = r.Group(grp.name, grp.middlewares...)
	} else {
		currentGroup = parent.Group(grp.name, grp.middlewares...)
	}

	grp.handle = currentGroup

	for i := range grp.children {
		initGroup(r, currentGroup, &grp.children[i])
	}
}

// 返回包含所有分组的 Group 对象
func Groups(r *gin.Engine) Group {
	groupList := defineGroups()
	ret := Group{}
	for i := range groupList {
		initGroup(r, nil, &groupList[i])
	}
	ret.groups = groupList

	return ret
}

// 定义分组及其路由
func defineGroups() []group {
	return []group{
		{
			name:     "api",
			initFunc: SetApiGroupRoutes,
			middlewares: []gin.HandlerFunc{
				middlewares.Cors(),
			},
			children: []group{
				{
					name:     "v1",
					initFunc: SetV1GroupRoutes,
					middlewares: []gin.HandlerFunc{
						middlewares.Cors(),
					},
				},
			},
		},
	}
}
