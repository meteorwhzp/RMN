package resource

type Resource struct {
	m_uri        string
	m_resourceId ResourceIndetifier
	m_properties map[string]PropertyAttribute
}

type ResourceList struct {
	resourceList map[string]*Resource
}

