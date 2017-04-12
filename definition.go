package swagger

// Definitions allows for Document level scope of reusable objects.
type Definitions map[string]Schema

// Visit alerts a Visitor to execute on this set of Definitions.
func (defs Definitions) Visit(caller Visitor) {
	caller.VisitDefinitions(defs)
}
