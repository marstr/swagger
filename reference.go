package swagger

type Reference string

func (ref Reference) Resolve() interface{} {
	return nil
}
