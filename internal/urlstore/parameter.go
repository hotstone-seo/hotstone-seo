package urlstore

// Parameter of node
type Parameter struct {
	keys   []string
	values []string
	m      map[string]string
}

// NewParameter return new instance of Parameter
func NewParameter(keys, values []string) *Parameter {
	m := make(map[string]string)
	for i, key := range keys {
		m[key] = values[i]
	}
	return &Parameter{
		keys:   keys,
		values: values,
		m:      m,
	}
}

// Map of parameter
func (p *Parameter) Map() map[string]string {
	return p.m
}

// Keys of parameter
func (p *Parameter) Keys() []string {
	return p.keys
}

// Values of parameter
func (p *Parameter) Values() []string {
	return p.values
}

// Empty return true if no keys
func (p *Parameter) Empty() bool {
	return len(p.keys) < 1
}
