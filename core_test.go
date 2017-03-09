package mailjet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	req, err := createRequest("GET", apiBase, nil, nil)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if req.Method != "GET" {
		t.Fatal("Wrong method:", req.Method)
	}
	if req.URL.String() != apiBase {
		t.Fatal("Wrong URL:", req.URL.String())
	}
	ua := fmt.Sprintf("%s/%s;%s",
		UserAgentBase,
		UserAgentVersion,
		runtime.Version(),
	)
	if req.Header["User-Agent"] == nil || req.Header["User-Agent"][0] != ua {
		t.Fatal("Wrong User-agent:", req.Header["User-Agent"][0])
	}
}

func TestConvertPayload(t *testing.T) {
	type Test struct {
		ID           int64 `mailjet:"read_only"`
		Name         string
		Email        string
		Address      string            `json:",omitempty" mailjet:"read_only"`
		TextPart     string            `json:"Text-Part,omitempty"`
		Header       map[string]string `json:",omitempty"`
		MjCampaignID int64             `json:"Mj-CampaignID,omitempty" mailjet:"read_only"`
	}
	test := &Test{
		ID:       -42,
		Email:    "ex@mple.com",
		TextPart: "This is text",
	}
	resMap := make(map[string]interface{})
	resMap["Email"] = "ex@mple.com"
	resMap["Text-Part"] = "This is text"
	body, err := convertPayload(test, []string{"Email", "TextPart"})
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	res, _ := json.Marshal(resMap)
	if !bytes.Equal(body, res) {
		t.Fatal("Wrong body:", string(body), string(res))
	}

	resMap["Name"] = ""
	body, err = convertPayload(&test, nil)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	res, _ = json.Marshal(resMap)
	if !bytes.Equal(body, res) {
		t.Fatal("Wrong body:", string(body), string(res))
	}
}

func TestBuildUrl(t *testing.T) {
	info := &Request{
		Resource: "contactslist",
		ID:       1,
		Action:   "managemanycontacts",
		ActionID: 5,
	}
	expected := "https://api.mailjet.com/v3/REST/contactslist/1/managemanycontacts/5"
	res := buildURL(apiBase, info)
	if res != expected {
		t.Fatal("Fail to build URL:", res)
	}
}

func TestReadJsonEmptyResult(t *testing.T) {
	type TestStruct struct {
		Email string
	}
	var data []TestStruct
	body := `{"Count":0,"Data":[],"Total":0}`
	_, _, err := readJSONResult(strings.NewReader(body), &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
}

func TestReadJsonResult(t *testing.T) {
	type TestStruct struct {
		Email string
	}
	var data []TestStruct
	body := `{"Count":2,"Data":[{"Email":"qwe@qwe.com"},{"Email":"aze@aze.com"}],"Total":1}`
	count, total, err := readJSONResult(strings.NewReader(body), &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count != 2 {
		t.Fatalf("Wrong count: %d != %d", count, 2)
	}
	if total != 1 {
		t.Fatalf("Wrong total: %d != %d", total, 2)
	}
	if data != nil {
		if data[0].Email != "qwe@qwe.com" {
			t.Fatalf("Fail to unmarshal JSON: %s != %s", data[0].Email, "qwe@qwe.com")
		}
		if data[1].Email != "aze@aze.com" {
			t.Fatalf("Fail to unmarshal JSON: %s != %s", data[1].Email, "aze@aze.com")
		}
	} else {
		t.Fatal("Fail to unmarshal JSON: empty res")
	}
}
