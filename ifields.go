package droictx

const (
	HTTPHeaderAppID            = "X-Droi-AppID"
	HTTPHeaderAppIDMode        = "X-Droi-AidMode"
	HTTPHeaderDeviceID         = "X-Droi-DeviceID"
	HTTPHeaderUserID           = "X-Droi-UserID"
	HTTPHeaderRequestID        = "X-Droi-ReqID"
	HTTPHeaderPlatformKey      = "X-Droi-Platform-Key"
	HTTPHeaderAPIKey           = "X-Droi-Api-Key"
	HTTPHeaderServiceAppID     = "X-Droi-Service-AppID"
	HTTPHeaderServiceAppIDMode = "X-Droi-SAidMode"
	ShortAppID                 = "Aid"
	ShortAppIDMode             = "Aidm"
	ShortDeviceID              = "Did"
	ShortUserID                = "Uid"
	ShorRequestID              = "Rid"
	ShortPlatformKey           = "XPk"
	ShortAPIKey                = "Ak"
	ShortServiceAppID          = "SAid"
	ShortServiceAppIDMode      = "SAidm"
)

// This interface designed from fasthttp *RequestHeader
type Peeker interface {
	Peek(key string) []byte
}

var (
	sKMap, hKMap map[string]string
)

func init() {
	sKMap = map[string]string{
		ShortAppID:            HTTPHeaderAppID,
		ShortAppIDMode:        HTTPHeaderAppIDMode,
		ShortDeviceID:         HTTPHeaderDeviceID,
		ShortUserID:           HTTPHeaderUserID,
		ShorRequestID:         HTTPHeaderRequestID,
		ShortPlatformKey:      HTTPHeaderPlatformKey,
		ShortAPIKey:           HTTPHeaderAPIKey,
		ShortServiceAppID:     HTTPHeaderServiceAppID,
		ShortServiceAppIDMode: HTTPHeaderServiceAppIDMode,
	}

	hKMap = map[string]string{
		HTTPHeaderAppID:            ShortAppID,
		HTTPHeaderAppIDMode:        ShortAppIDMode,
		HTTPHeaderDeviceID:         ShortDeviceID,
		HTTPHeaderUserID:           ShortUserID,
		HTTPHeaderRequestID:        ShorRequestID,
		HTTPHeaderPlatformKey:      ShortPlatformKey,
		HTTPHeaderAPIKey:           ShortAPIKey,
		HTTPHeaderServiceAppID:     ShortServiceAppID,
		HTTPHeaderServiceAppIDMode: ShortServiceAppIDMode,
	}
}

// retrun a map, key is HTTP Header Field, value is Short Field
func IFieldHeaderKeyMap() (keyMap map[string]string) {
	return hKMap
}

// retrun a map, key is Short Field, value is HTTP Header Field
func IFieldShortKeyMap() (keyMap map[string]string) {
	return sKMap
}

func GetContextFromPeeker(p Peeker) Context {
	c := Context{}
	var v []byte
	for hk, sk := range hKMap {
		v = p.Peek(hk)
		if len(v) > 0 {
			c.Set(sk, string(v))
		}
	}
	return c
}

// retrun a map, key HTTP Header Field, value is the field value stored in Context
func (c *Context) HeaderMap() (ret map[string]string) {
	ret = make(map[string]string)
	for sk, hk := range sKMap {
		v, _ := c.GetString(sk)
		ret[hk] = v
	}
	return
}

// Set Context with Header Key, Store in Context with Short Key
func (c *Context) HeaderSet(headerField, headerValue string) {
	if sk, ok := hKMap[headerField]; ok {
		c.Set(sk, headerValue)
	}
	return
}
