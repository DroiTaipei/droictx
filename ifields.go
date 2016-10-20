package droictx

const (
	HTTPHeaderAppID       = "X-Droi-AppID"
	HTTPHeaderAppIDMode   = "X-Droi-AidMode"
	HTTPHeaderDeviceID    = "X-Droi-DeviceID"
	HTTPHeaderUserID      = "X-Droi-UserID"
	HTTPHeaderRequestID   = "X-Droi-ReqID"
	HTTPHeaderPlatformKey = "X-Droi-Platform-Key"
	HTTPHeaderAPIKey      = "X-Droi-ApiKey"
	ShortAppID            = "Aid"
	ShortAppIDMode        = "Aidm"
	ShortDeviceID         = "Did"
	ShortUserID           = "Uid"
	ShorRequestID         = "Rid"
	ShortPlatformKey      = "XPk"
	ShortAPIKey           = "Ak"
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
		ShortAppID:       HTTPHeaderAppID,
		ShortAppIDMode:   HTTPHeaderAppIDMode,
		ShortDeviceID:    HTTPHeaderDeviceID,
		ShortUserID:      HTTPHeaderUserID,
		ShorRequestID:    HTTPHeaderRequestID,
		ShortPlatformKey: HTTPHeaderPlatformKey,
		ShortAPIKey:      HTTPHeaderAPIKey,
	}

	hKMap = map[string]string{
		HTTPHeaderAppID:       ShortAppID,
		HTTPHeaderAppIDMode:   ShortAppIDMode,
		HTTPHeaderDeviceID:    ShortDeviceID,
		HTTPHeaderUserID:      ShortUserID,
		HTTPHeaderRequestID:   ShorRequestID,
		HTTPHeaderPlatformKey: ShortPlatformKey,
		HTTPHeaderAPIKey:      ShortAPIKey,
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
	var v string
	for hk, sk := range hKMap {
		v = string(p.Peek(hk))
		c.Set(sk, v)
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

func (c *Context) HeaderSet(headerField, headerValue string) {
	if sk, ok := hKMap[headerField]; ok {
		c.Set(sk, headerValue)
	}
	return
}
