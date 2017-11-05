package resource

type ResourceIndetifier uint8

type ResultCode uint8

const (
	//Success status code
	OC_OK               ResultCode = 203
	OC_RESOURCE_CREATED ResultCode = 201
	OC_RESOURCE_DELETED ResultCode = 202
	OC_RESOURCE_CHANGED ResultCode = 204

	//Error status code
	OC_INVALID_URI   ResultCode = 401
	OC_INVALID_QUERY ResultCode = 400
	OC_INVALID_IP    ResultCode = 403
	OC_INVALID_PORT   ResultCode = 405
	OC_INVALID_CALLBACK ResultCode = 402
	OC_INVALID_METHOD    ResultCode = 406
	OC_NO_RESOURCE  ResultCode = 404
	OC_RESOURCE_EXIT ResultCode = 407
)

type PropertyAttribute struct {
	propertyKey   string
	propertyValue interface{}
}
