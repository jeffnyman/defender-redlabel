package game

import (
	"github.com/jeffnyman/defender-redlabel/logger"
	"github.com/jeffnyman/defender-redlabel/types"
)

var idCounter types.EntityID

func init() {
	idCounter = 0
}

type Entity struct {
	Id         types.EntityID
	Class      types.EntityType
	components map[types.CmpType]types.ICmp
	active     bool
	engine     *Engine
	parent     types.EntityID
	child      types.EntityID
	paused     bool
}

func NewEntity(engine *Engine, class types.EntityType) *Entity {
	rv := &Entity{
		Id:         idCounter,
		Class:      class,
		components: map[types.CmpType]types.ICmp{},
		engine:     engine,
		active:     false,
		parent:     idCounter,
		child:      idCounter,
	}

	engine.AddEntity(rv)
	idCounter++

	return rv
}

func (e *Entity) GetID() types.EntityID {
	return e.Id
}

func (e *Entity) GetEngine() types.IEngine {
	return e.engine
}

func (e *Entity) Active() bool {
	return e.active
}

func (e *Entity) SetActive(s bool) {
	logger.Debug("Entity %d set active %t ", e.Id, s)

	e.active = s
}

func (e *Entity) Paused() bool {
	return e.paused
}

func (e *Entity) SetPaused(s bool) {
	e.paused = s
}

func (e *Entity) AddComponent(c types.ICmp) {
	logger.Debug("Entity %d add component %s", e.Id, c.Type())

	e.components[c.Type()] = c
	e.engine.AddComponent(e, c)

}

func (e *Entity) RemoveComponent(ct types.CmpType) {
	logger.Debug("Entity %d remove component %s", e.Id, ct.String())

	delete(e.components, ct)
	e.engine.RemoveComponent(e, ct)
}

func (e *Entity) HasComponent(c types.CmpType) bool {
	_, ok := e.components[c]

	return ok
}

func (e *Entity) GetComponent(c types.CmpType) types.ICmp {
	rv, ok := e.components[c]

	if !ok {
		return nil
	}

	return rv
}

func (e *Entity) GetComponents() map[types.CmpType]types.ICmp {
	return e.components
}

func (e *Entity) Parent() types.EntityID {
	return e.parent
}

func (e *Entity) Child() types.EntityID {
	return e.child
}

func (e *Entity) SetParent(pe types.EntityID) {
	e.parent = pe
}

func (e *Entity) SetChild(ce types.EntityID) {
	e.child = ce
}

func (e *Entity) GetClass() types.EntityType {
	return e.Class
}
