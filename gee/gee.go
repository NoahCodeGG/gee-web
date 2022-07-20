package gee

import (
	"net/http"
)

// HandlerFunc 用于用户定义路由映射处理方法
type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router        // 路由映射表
	groups []*RouterGroup // 所有路由分组
}

type RouterGroup struct {
	prefix      string        // 分组前缀
	middlewares []HandlerFunc // 支持中间件
	parent      *RouterGroup  // 父级分组
	engine      *Engine       // 所有分组共用一个 Engine 实例
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 添加路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 添加 GET 方法路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 添加 POST 方法路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run 启动一个 Http 服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
