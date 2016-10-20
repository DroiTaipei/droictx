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
	m := c.HeaderMap()
	v, ok := m["X-Droi-AppID"]
	if !ok || v != "ZXC123ASDQWE" {
		t.Error("AppID not match Aid")
	}

	v, ok = m["X-Droi-ReqID"]
	if !ok || v != "1029384756" {
		t.Error("ReqID not match Rid")
	}
}

func TestHeaderSet(t *testing.T) {
	c := Context{}
	c.HeaderSet("X-Droi-AppID", "ZXC123ASDQWE")
	c.HeaderSet("X-Droi-ReqID", "1029384756")
	if v, ok := c.GetString("Aid"); !ok || v != "ZXC123ASDQWE" {
		t.Error("AppID not match Aid")
	}
	if v, ok := c.GetString("Rid"); !ok || v != "1029384756" {
		t.Error("ReqID not match Rid")
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
		return nil
	})

	c := GetContextFromPeeker(peeker)
	if v, ok := c.GetString("Aid"); !ok || v != "ZXC123ASDQWE" {
		t.Error("AppID not match Aid")
	}
	if v, ok := c.GetString("Rid"); !ok || v != "1029384756" {
		t.Error("ReqID not match Rid")
	}
}
