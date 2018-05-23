package json

import (
	"encoding/json"
)

// JSON type
type JSON struct {
	Raw interface{}
}

// NewJSON func
func NewJSON(raw interface{}) *JSON {
	return &JSON{Raw: raw}
}

func (j *JSON) k(key string) interface{} {
	return j.Raw.(map[string]interface{})[key]
}

// HasKey func
func (j *JSON) HasKey(key string) bool {
	return j.K(key).Raw != nil
}

// SetKey func
func (j *JSON) SetKey(key string, value interface{}) {
	j.K(key).Raw.(map[string]interface{})[key] = value
}

// PutKey func
func (j *JSON) PutKey(key string, value interface{}) {
	j.Raw.(map[string]interface{})[key] = value
}

// DelKey func
func (j *JSON) DelKey(key string) {
	delete(j.Raw.(map[string]interface{}), key)
}

// K func
func (j *JSON) K(key string) *JSON {
	return &JSON{Raw: j.k(key)}
}

// String func
func (j *JSON) String() string {
	return j.Raw.(string)
}

// SafeString func
func (j *JSON) SafeString() string {
	defer func() { recover() }()
	return j.String()
}

// Int func
func (j *JSON) Int() int {
	return int(j.Float64())
}

// SafeInt func
func (j *JSON) SafeInt() int {
	defer func() { recover() }()
	return j.Int()
}

// Float64 func
func (j *JSON) Float64() float64 {
	return j.Raw.(float64)
}

// SafeFloat64 func
func (j *JSON) SafeFloat64() float64 {
	defer func() { recover() }()
	return j.Float64()
}

// Array func
func (j *JSON) Array() []*JSON {
	array := []*JSON{}
	for _, r := range j.Raw.([]interface{}) {
		array = append(array, &JSON{Raw: r})
	}
	return array
}

// Dict func
func (j *JSON) Dict() map[string]*JSON {
	dict := map[string]*JSON{}
	for k, r := range j.Raw.(map[string]interface{}) {
		dict[k] = &JSON{Raw: r}
	}
	return dict
}

// SafeDecode func
func (j *JSON) SafeDecode(v interface{}) interface{} {
	b, _ := json.Marshal(j.Raw)
	if err := json.Unmarshal(b, v); err != nil {
		return nil
	}
	return v
}
