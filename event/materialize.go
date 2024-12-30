package event

type Materialize struct {
	mytype  EventType
	payload interface{}
}

func NewMaterialize(p interface{}) *Materialize {
	return &Materialize{
		mytype:  MaterializeEvent,
		payload: p,
	}
}

func (ev *Materialize) GetType() EventType {
	return ev.mytype
}

func (ev *Materialize) GetPayload() interface{} {
	return ev.payload
}
