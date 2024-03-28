package nullable_value

type Value[Val any] struct {
	val     Val
	present bool
}

func New[Val any](val Val, present bool) Value[Val] {
	return Value[Val]{
		val:     val,
		present: present,
	}
}

func (val Value[Val]) Value() Val {
	return val.val
}

func (val Value[Val]) Valid() bool {
	return val.present
}
