package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"math/rand"
)

var provides = make(map[string]Provider)

//全局sesssion管理器
type Manager struct {
	cookieName  string     //私有cookine的name
	lock        sync.Mutex //互斥锁
	provider    Provider
	maxLifeTime int64 //最大生存时间
}

//生成全局惟一的sessionID
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//新建一个session管理器
func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provide, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("provide不存在")
	}

	return &Manager{provider: provide, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

type Session interface {
	Set(key, value interface{}) error //设置session变量
	Get(key interface{}) interface{}  // 获取session变量
	Delete(key interface{}) error     // 删除session变量
	SessionID() string                // back current sessionID
}

//抽象接口,用来表达session管理器底层的存储结构
type Provider interface {
	SessionInit(sid string) (Session, error) //初始化session,返回新的session变量
	SessionRead(sid string) (Session, error) //通过sid来返回session的变量
	SessionDestroy(sid string) error         //同来销毁sid对应的session的变量
	SessionGC(maxLifeTime int64)             //根据maxlifetime来删除session
}

//注册存储session	的结构的函数
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

//判断用户是否持有session
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

//然后在init函数中初始化
func main() {
	//生成一个新的session
	globalSessions, err := NewManager("memory", "gosessionid", 3600)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(globalSessions)
}
