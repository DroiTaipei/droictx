package droictx

import (
	"testing"
)

// Mock Peeker
type mockPeeker func(key string) []byte

func (p mockPeeker) Peek(key string) []byte { return p(key) }

func TestHeaderMap(t *testing.T) {
	c := Context{}
	c.Set("Aid", "ZXC123ASDQWE")
	c.Set("Rid", "1029384756")
	c.Set("Ak", "abcdefg123")
	c.Set("SAid", "4RRFFV3edc")
	m := c.HeaderMap()
	v, ok := m["X-Droi-AppID"]
	if !ok || v != "ZXC123ASDQWE" {
		t.Error("AppID not match Aid")
	}

	v, ok = m["X-Droi-ReqID"]
	if !ok || v != "1029384756" {
		t.Error("ReqID not match Rid")
	}
	v, ok = m["X-Droi-Api-Key"]
	if !ok || v != "abcdefg123" {
		t.Error("ApiKey not match Rid")
	}
	v, ok = m["X-Droi-Service-AppID"]
	if !ok || v != "4RRFFV3edc" {
		t.Error("SAid not match Rid")
	}
}

func TestHeaderSet(t *testing.T) {
	c := Context{}
	c.HeaderSet("X-Droi-AppID", "ZXC123ASDQWE")
	c.HeaderSet("X-Droi-ReqID", "1029384756")
	c.HeaderSet("X-Droi-Api-Key", "abcdefg123")
	c.HeaderSet("X-Droi-Service-AppID", "4RRFFV3edc")
	if v, ok := c.GetString("Aid"); !ok || v != "ZXC123ASDQWE" {
		t.Error("AppID not match Aid")
	}
	if v, ok := c.GetString("Rid"); !ok || v != "1029384756" {
		t.Error("ReqID not match Rid")
	}
	if v, ok := c.GetString("Ak"); !ok || v != "abcdefg123" {
		t.Error("APIKey not match Ak")
	}
	if v, ok := c.GetString("SAid"); !ok || v != "4RRFFV3edc" {
		t.Error("Service Aid not match SAid")
	}
}

func TestPeeker(t *testing.T) {

	peeker := mockPeeker(func(key string) []byte {
		if key == "X-Droi-AppID" {
			return []byte("ZXC123ASDQWE")
		}
		if key == "X-Droi-ReqID" {
			return []byte("1029384756")
		}
		if key == "X-Droi-Api-Key" {
			return []byte("abcdefg123")
		}
		if key == "X-Droi-Service-AppID" {
			return []byte("4RRFFV3edc")
		}
		return []byte{}
	})

	c := GetContextFromPeeker(peeker)
	if v, ok := c.GetString("Aid"); !ok || v != "ZXC123ASDQWE" {
		t.Error("AppID not match Aid")
	}
	if v, ok := c.GetString("Rid"); !ok || v != "1029384756" {
		t.Error("ReqID not match Rid")
	}
	if v, ok := c.GetString("Ak"); !ok || v != "abcdefg123" {
		t.Error("APIKey not match Ak")
	}
	if v, ok := c.GetString("SAid"); !ok || v != "4RRFFV3edc" {
		t.Error("Service AppID not match SAid")
	}

	// Test Empty Field
	m := c.Map()
	if _, ok := m["Aidm"]; ok {
		t.Error("Aidm should not exists")
	}
}
