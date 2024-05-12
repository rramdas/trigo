package registrar

import "sync"

type Registrar struct {
	values map[string]interface{}
	sync.Mutex
}

func (r *Registrar) Register(name string, value interface{}) {
	if r == nil {
		return
	}

	r.Lock()
	defer r.Unlock()
	r.values[name] = value
}
