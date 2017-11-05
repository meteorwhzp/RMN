package resource

import logger "github.com/shengkehua/xlog4go"

func (rl *ResourceList) RegisterResource(uri string) ResultCode {
	if _, ok := rl.resourceList[uri]; !ok {
		rl.resourceList[uri] = &Resource{m_uri: uri}
		return OC_RESOURCE_CREATED
	}else {
		return OC_RESOURCE_EXIT
	}
}

func (rl *ResourceList) DeleteResource(uri string) ResultCode {
	if _, ok := rl.resourceList[uri]; !ok {
		return OC_NO_RESOURCE
	}
	delete(rl.resourceList, uri)
	logger.Info("resource %s is deleted", uri)
	return OC_RESOURCE_DELETED
}

func (rl *ResourceList) RetrieveResource(uri string) (res *Resource, status ResultCode) {
	if res, ok := rl.resourceList[uri]; !ok {
		return res, OC_NO_RESOURCE
	}
	return res, OC_OK
}

func ResourceDiscovery() {

}