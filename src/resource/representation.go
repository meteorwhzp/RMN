package resource

type Representation struct {
	m_uri string
	m_values map[string]PropertyAttribute
}

func (rep *Representation) GetValue(str string, val interface{}) bool {
	if v, ok := rep.m_values[str]; !ok {
		val = v.propertyValue
		return true
	} else {
		return false
	}
}
