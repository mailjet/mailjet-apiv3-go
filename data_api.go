package mailjet

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

// ListData issues a GET to list the specified data resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *Client) ListData(resource string, res interface{}, options ...RequestOptions) (count, total int, err error) {
	url := buildDataURL(&DataRequest{SourceType: resource})
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

	return readJSONResult(resp.Body, res)
}

// GetData issues a GET to view a resource specifying an id
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
// Without an specified SourceTypeID in MailjetDataRequest, it is the same as ListData.
func (mj *Client) GetData(mdr *DataRequest, res interface{}, options ...RequestOptions) (err error) {
	url := buildDataURL(mdr)
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
		contentType := resp.Header["Content-Type"][0]
		if contentType == "application/json" {
			err = json.NewDecoder(resp.Body).Decode(&res)
		} else if contentType == "text/csv" {
			res, err = csv.NewReader(resp.Body).ReadAll()
		}
	}
	return err
}

// PostData issues a POST to create a new data resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *Client) PostData(fmdr *FullDataRequest, res interface{}, options ...RequestOptions) (err error) {
	url := buildDataURL(fmdr.Info)
	req, err := createRequest("POST", url, fmdr.Payload, nil, options...)
	if err != nil {
		return err
	}
	if fmdr.Info.MimeType != "" {
		contentType := strings.Replace(fmdr.Info.MimeType, ":", "/", 1)
		req.Header.Add("Content-Type", contentType)
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

// PutData is used to update a data resource.
// Fields to be updated must be specified by the string array onlyFields.
// If onlyFields is nil, all fields except these with the tag read_only, are updated.
// Filters can be add via functional options.
func (mj *Client) PutData(fmr *FullDataRequest, onlyFields []string, options ...RequestOptions) (err error) {
	url := buildDataURL(fmr.Info)
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
func (mj *Client) DeleteData(mdr *DataRequest) (err error) {
	url := buildDataURL(mdr)
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
