package mailjet

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

// List issues a GET to list the specified data resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *MailjetClient) ListData(resource string, res interface{}, options ...MailjetOptions) (count, total int, err error) {
	url := buildDataUrl(&MailjetDataRequest{SourceType: resource})
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return count, total, err
	}
	resp, err := mj.doRequest(req)
	if err != nil {
		return count, total, err
	} else if resp == nil {
		return count, total, fmt.Errorf("empty response")
	}
	defer resp.Body.Close()
	return readJsonResult(resp.Body, res)
}

// Get issues a GET to view a resource specifying an id
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
// Without an specified SourceTypeID in MailjetDataRequest, it is the same as ListData.
func (mj *MailjetClient) GetData(mdr *MailjetDataRequest, res interface{}, options ...MailjetOptions) (err error) {
	url := buildDataUrl(mdr)
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return err
	}
	resp, err := mj.doRequest(req)
	if err != nil {
		return err
	} else if resp == nil {
		return fmt.Errorf("empty response")
	}
	defer resp.Body.Close()

	if resp.Header["Content-Type"] != nil {
		content_type := resp.Header["Content-Type"][0]
		if content_type == "application/json" {
			err = json.NewDecoder(resp.Body).Decode(&res)
		} else if content_type == "text/csv" {
			res, err = csv.NewReader(resp.Body).ReadAll()
		}
	}
	return err
}

// Post issues a POST to create a new data resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *MailjetClient) PostData(fmdr *FullMailjetDataRequest, res interface{}, options ...MailjetOptions) (err error) {
	url := buildDataUrl(fmdr.Info)
	req, err := createRequest("POST", url, fmdr.Payload, nil, options...)
	if err != nil {
		return err
	}
	if fmdr.Info.MimeType != "" {
		content_type := strings.Replace(fmdr.Info.MimeType, ":", "/", 1)
		req.Header.Add("Content-Type", content_type)
	} else {
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := mj.doRequest(req)
	if err != nil {
		return err
	} else if resp == nil {
		return fmt.Errorf("empty response")
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&res)
}

// Put is used to update a data resource.
// Fields to be updated must be specified by the string array onlyFields.
// If onlyFields is nil, all fields except these with the tag read_only, are updated.
// Filters can be add via functional options.
func (mj *MailjetClient) PutData(fmr *FullMailjetDataRequest, onlyFields []string, options ...MailjetOptions) (err error) {
	url := buildDataUrl(fmr.Info)
	req, err := createRequest("PUT", url, fmr.Payload, onlyFields, options...)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := mj.doRequest(req)
	if resp != nil {
		resp.Body.Close()
	}

	return err
}

// DeleteData is used to delete a data resource.
func (mj *MailjetClient) DeleteData(mdr *MailjetDataRequest) (err error) {
	url := buildDataUrl(mdr)
	r, err := createRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}
	resp, err := mj.doRequest(r)
	if resp != nil {
		resp.Body.Close()
	}
	return err
}
