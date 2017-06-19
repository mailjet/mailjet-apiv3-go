package mailjet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// DebugLevel defines the verbosity of the debug.
var DebugLevel int

// These are the different level of debug.
const (
	LevelNone      = iota // No debug.
	LevelDebug            // Debug without body.
	LevelDebugFull        // Debug with body.
)

// User-Agent is formated as "UserAgentBase/UserAgentVersion;runtime.Version()".
const (
	UserAgentBase    = "mailjet-api-v3-go"
	UserAgentVersion = "2.4.5"
)

const (
	apiBase  = "https://api.mailjet.com/v3"
	apiPath  = "REST"
	dataPath = "DATA"
)

// createRequest is the main core function.
func createRequest(method string, url string,
	payload interface{}, onlyFields []string,
	options ...RequestOptions) (req *http.Request, err error) {

	body, err := convertPayload(payload, onlyFields)
	if err != nil {
		return req, fmt.Errorf("creating request: %s\n", err)
	}
	req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return req, fmt.Errorf("creating request: %s\n", err)
	}
	for _, option := range options {
		option(req)
	}
	userAgent(req)
	req.Header.Add("Accept", "application/json")
	return req, err
}

// converPayload returns payload casted in []byte.
// If the payload is a structure, it's encoded to JSON.
func convertPayload(payload interface{}, onlyFields []string) (body []byte, err error) {
	if payload != nil {
		switch t := payload.(type) {
		case string:
			body = []byte(t)
		case []byte:
			body = t
		default:
			v := reflect.Indirect(reflect.ValueOf(payload))
			if v.Kind() == reflect.Ptr {
				return convertPayload(v.Interface(), onlyFields)
			} else if v.Kind() == reflect.Struct {
				body, err = json.Marshal(buildMap(v, onlyFields))
				if err != nil {
					return body, err
				}
			}
		}
		if DebugLevel == LevelDebugFull {
			log.Println("Body:", string(body))
		}
	}
	return body, err
}

// buildMap returns a map with fields specified in onlyFields (all fields if nil)
// and without the read_only fields.
func buildMap(v reflect.Value, onlyFields []string) map[string]interface{} {
	res := make(map[string]interface{})
	if onlyFields != nil {
		for _, onlyField := range onlyFields {
			fieldType, exist := v.Type().FieldByName(onlyField)
			if exist {
				addFieldToMap(true, fieldType, v.FieldByName(onlyField), res)
			}
		}
	} else {
		for i := 0; i < v.NumField(); i++ {
			addFieldToMap(false, v.Type().Field(i), v.Field(i), res)
		}
	}
	return res
}

func addFieldToMap(onlyField bool, fieldType reflect.StructField,
	fieldValue reflect.Value, res map[string]interface{}) {
	if fieldType.Tag.Get("mailjet") != "read_only" {
		name, second := parseTag(fieldType.Tag.Get("json"))
		if name == "" {
			name = fieldType.Name
		}
		if !onlyField && second == "omitempty" &&
			isEmptyValue(fieldValue) {
			return
		}
		res[name] = fieldValue.Interface()
	}
}

func parseTag(tag string) (string, string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tag[idx+1:]
	}
	return tag, ""
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// userAgent add the User-Agent value to the request header.
func userAgent(req *http.Request) {
	ua := fmt.Sprintf("%s/%s;%s",
		UserAgentBase,
		UserAgentVersion,
		runtime.Version(),
	)
	req.Header.Add("User-Agent", ua)
}

func buildURL(baseURL string, info *Request) string {
	tokens := []string{baseURL, apiPath, info.Resource}
	if info.ID != 0 {
		id := strconv.FormatInt(info.ID, 10)
		tokens = append(tokens, id)
	} else if info.AltID != "" {
		tokens = append(tokens, string(info.AltID))
	}
	if info.Action != "" {
		tokens = append(tokens, info.Action)
	}
	if info.ActionID != 0 {
		actionID := strconv.FormatInt(info.ActionID, 10)
		tokens = append(tokens, actionID)
	}
	return strings.Join(tokens, "/")
}

func buildDataURL(baseURL string, info *DataRequest) string {
	tokens := []string{baseURL, dataPath, info.SourceType}
	if info.SourceTypeID != 0 {
		id := strconv.FormatInt(info.SourceTypeID, 10)
		tokens = append(tokens, id)
	}
	if info.DataType != "" {
		tokens = append(tokens, info.DataType)
		if info.MimeType != "" {
			tokens = append(tokens, info.MimeType)
		}
	}
	if info.DataTypeID != 0 {
		DataTypeID := strconv.FormatInt(info.DataTypeID, 10)
		tokens = append(tokens, DataTypeID)
	} else if info.LastID == true {
		tokens = append(tokens, "LAST")
	}
	return strings.Join(tokens, "/")
}

// readJsonResult decodes the API response, returns Count and Total values
// and stores the Data in the value pointed to by data.
func readJSONResult(r io.Reader, data interface{}) (int, int, error) {
	var res RequestResult
	res.Data = data

	jsonBlob, err := ioutil.ReadAll(r) // ReadAll and store in jsonBlob (mandatory if we want to unmarshal two times)
	if err != nil {
		return 0, 0, fmt.Errorf("Error reading API response: %s", err)
	}
	if DebugLevel == LevelDebugFull {
		log.Println("Body: ", string(jsonBlob)) // DEBUG
	}

	err = json.Unmarshal(jsonBlob, &res) // First try with the RequestResult struct
	if err != nil {
		return 0, 0, fmt.Errorf("Error decoding API response: %s", err)
	} else if _, ok := data.(**SentResult); ok { // Send API case
		err = json.Unmarshal(jsonBlob, data) // Trying directly with struct specified in parameter
		if err != nil {
			return 0, 0, fmt.Errorf("Error decoding API response: %s", err)
		}
		return 0, 0, nil // Count and Total are undetermined
	}
	return res.Count, res.Total, nil
}

// NbAttempt defines the number of attempt
// for a request as long as StatusCode == 500.
var NbAttempt = 5

// doRequest is called to execute the request. Authentification is set
// with the public key and the secret key specified in MailjetClient.
func (c *HTTPClient) doRequest(req *http.Request) (resp *http.Response, err error) {
	debugRequest(req) //DEBUG
	req.SetBasicAuth(c.apiKeyPublic, c.apiKeyPrivate)
	for attempt := 0; attempt < NbAttempt; attempt++ {
		if resp != nil {
			resp.Body.Close()
		}
		resp, err = c.client.Do(req)
		if err != nil || (resp != nil && resp.StatusCode != 500) {
			break
		}
	}
	defer debugResponse(resp) //DEBUG
	if err != nil {
		return resp, fmt.Errorf("Error getting %s: %s", req.URL, err)
	}
	err = checkResponseError(resp)
	return resp, err
}

// checkResponseError returns response error if the statuscode is < 200 or >= 400.
func checkResponseError(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		var mailjetErr RequestError
		err := json.NewDecoder(resp.Body).Decode(&mailjetErr)

		if err != nil {
			return fmt.Errorf("Unexpected server response code: %d: %s", resp.StatusCode, err)
		}
		return fmt.Errorf("Unexpected server response code: %d: %s (%s)",
			resp.StatusCode, mailjetErr.ErrorMessage, mailjetErr.ErrorInfo)
	}
	return nil
}

// debugRequest is a custom dump of the request.
// Method used, final URl called, and Header content are logged.
func debugRequest(req *http.Request) {
	if DebugLevel > LevelNone && req != nil {
		log.Printf("Method used is: %s\n", req.Method)
		log.Printf("Final URL is: %s\n", req.URL)
		log.Printf("Header is: %s\n", req.Header)
	}
}

// debugResponse is a custom dump of the response.
// Status and Header content are logged.
func debugResponse(resp *http.Response) {
	if DebugLevel > LevelNone && resp != nil {
		log.Printf("Status is: %s\n", resp.Status)
		log.Printf("Header is: %s\n", resp.Header)
	}
}
