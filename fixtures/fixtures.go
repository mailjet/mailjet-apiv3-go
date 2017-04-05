// Package fixtures provid fake data so we can mock the `Client` struct in mailjet_client.go
package fixtures

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/mailjet/mailjet-apiv3-go/resources"
)

// Fixtures definition
type Fixtures struct {
	data map[interface{}][]byte
}

// New loads fixtures in memory by iterating through its fixture method
func New() *Fixtures {
	f := new(Fixtures)
	f.data = make(map[interface{}][]byte)
	fix := reflect.ValueOf(f)
	for i := 0; i < fix.NumMethod(); i++ {
		method := fix.Method(i)
		if method.Type().NumIn() == 0 && method.Type().NumOut() == 2 {
			values := method.Call([]reflect.Value{})
			reflect.ValueOf(f.data).SetMapIndex(values[0], values[1])
		}
	}

	return f
}

func (f *Fixtures) Read(v interface{}) error {
	for t, val := range f.data {
		if reflect.ValueOf(v).Type().String() == reflect.ValueOf(t).Type().String() {
			json.Unmarshal(val, v)
			return nil
		}
	}
	return errors.New("not found")
}

// User fixture info
func (f *Fixtures) User() (*[]resources.User, []byte) {
	return &[]resources.User{}, []byte(`[{"ID":24, "Email": "passenger@mailjet.com", "LastIP": "127.0.0.1", "Username": "passenger", "MaxAllowedAPIKeys": 5}]`)
}

// Contact fixture info
func (f *Fixtures) Contact() (*[]resources.Contact, []byte) {
	return &[]resources.Contact{}, []byte(`[{"ID":42, "Email": "contact@mailjet.com", "DeliveredCount": 42}]`)
}

// ContactList fixture info
func (f *Fixtures) ContactList() (*[]resources.Contactslist, []byte) {
	return &[]resources.Contactslist{}, []byte(`[{"ID":84, "Address": "contact@mailjet.com", "Name": "John Doe", "SubscriberCount": 1000}]`)
}

// ListRecipient fixture info
func (f *Fixtures) ListRecipient() (*[]resources.Listrecipient, []byte) {
	return &[]resources.Listrecipient{}, []byte(`[{"ID":168}]`)
}

// Sender fixture info
func (f *Fixtures) Sender() (*[]resources.Sender, []byte) {
	return &[]resources.Sender{}, []byte(`[{"ID":336, "Name": "Mansa Musa", "Status": "Active", "EmailType": "transactional"}]`)
}
