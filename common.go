package mapOperation

import "sync"

//jkihjoihoihgosadja
/*对map进行常规安全的增删查改/初始化操作：加读写锁*/
//安全的Map
/*type SynchronizedMap struct{
	rw *sync.RWMutex
	data map[interface{}]interface{}
}
//生成一个初始化的SynchronizedMap
func NewSynchronizedMap() *SynchronizedMap{
	return &SynchronizedMap{
		rw:new(sync.RWMutex),
		data:make(map[interface{}]interface{}),
	}
}
//存储操作
func (sm *SynchronizedMap) Put(k,v interface{}){
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k]=v
}
//删除操作
func (sm *SynchronizedMap) delete(k interface{}){
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data,k)
}
//获取操作
func (sm *SynchronizedMap) Get(k interface{}) interface{}{
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}
//遍历Map,并把遍历的值给回调函数，可以让调用者控制做任何事情
func (sm *SynchronizedMap) Each(cb func(interface{},interface{})){
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	for  k,v := range sm.data{
		cb(k,v)
	}
}*/

//第二遍
type SynchronizedMap struct{
	rw *sync.RWMutex
	data map[interface{}]interface{}
}
//初始化
func NewSynchronizedMap() *SynchronizedMap{
	return &SynchronizedMap{
		rw:new(sync.RWMutex),
		data:make(map[interface{}]interface{}),
	}
}
func (sm *SynchronizedMap) Push(k,v interface{}){
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k]=v
}
func (sm *SynchronizedMap) delete(k interface{}){
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data,k)
}
func (sm *SynchronizedMap) Get(k interface{}) interface{}{
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}
func (sm *SynchronizedMap) Each(cb func(interface{},interface{})){
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	for k,v := range sm.data{
		cb(k,v)
	}
}
