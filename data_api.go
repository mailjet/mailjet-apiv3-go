package mailjet

import "strings"

// ListData issues a GET to list the specified data resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *Client) ListData(resource string, resp interface{}, options ...RequestOptions) (count, total int, err error) {
	url := buildDataURL(mj.apiBase, &DataRequest{SourceType: resource})
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return count, total, err
	}

	return mj.httpClient.Send(req).Read(resp).Call()
}

// GetData issues a GET to view a resource specifying an id
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
// Without an specified SourceTypeID in MailjetDataRequest, it is the same as ListData.
func (mj *Client) GetData(mdr *DataRequest, res interface{}, options ...RequestOptions) (err error) {
	url := buildDataURL(mj.apiBase, mdr)
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return err
	}

	_, _, err = mj.httpClient.Send(req).Read(res).Call()
	return err
}

// PostData issues a POST to create a new data resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *Client) PostData(fmdr *FullDataRequest, res interface{}, options ...RequestOptions) (err error) {
	url := buildDataURL(mj.apiBase, fmdr.Info)
	req, err := createRequest("POST", url, fmdr.Payload, nil, options...)
	if err != nil {
		return err
	}

	headers := map[string]string{"Content-Type": "application/json"}
	if fmdr.Info.MimeType != "" {
		contentType := strings.Replace(fmdr.Info.MimeType, ":", "/", 1)
		headers = map[string]string{"Content-Type": contentType}
	}

	_, _, err = mj.httpClient.Send(req).With(headers).Read(res).Call()
	return err
}

// PutData is used to update a data resource.
// Fields to be updated must be specified by the string array onlyFields.
// If onlyFields is nil, all fields except these with the tag read_only, are updated.
// Filters can be add via functional options.
func (mj *Client) PutData(fmr *FullDataRequest, onlyFields []string, options ...RequestOptions) (err error) {
	url := buildDataURL(mj.apiBase, fmr.Info)
	req, err := createRequest("PUT", url, fmr.Payload, onlyFields, options...)
	if err != nil {
		return err
	}

	headers := map[string]string{"Content-Type": "application/json"}
	_, _, err = mj.httpClient.Send(req).With(headers).Call()

	return err
}

// DeleteData is used to delete a data resource.
func (mj *Client) DeleteData(mdr *DataRequest) (err error) {
	url := buildDataURL(mj.apiBase, mdr)
	req, err := createRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}

	_, _, err = mj.httpClient.Send(req).Call()

	return err
}
