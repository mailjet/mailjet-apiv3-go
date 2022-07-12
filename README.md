[mailjet]: http://www.mailjet.com
[api_credential]: https://app.mailjet.com/account/api_keys
[doc]: http://dev.mailjet.com/guides/?go#


![alt text](https://www.mailjet.com/images/email/transac/logo_header.png "Mailjet")

# Official Mailjet Go Client

[![Build Status](https://travis-ci.org/mailjet/mailjet-apiv3-go.svg?branch=master)](https://travis-ci.org/mailjet/mailjet-apiv3-go)
[![GoDoc](https://godoc.org/github.com/mailjet/mailjet-apiv3-go?status.svg)](https://godoc.org/github.com/mailjet/mailjet-apiv3-go)
[![Go Report Card](https://goreportcard.com/badge/mailjet/mailjet-apiv3-go)](https://goreportcard.com/report/mailjet/mailjet-apiv3-go)
![Current Version](https://img.shields.io/badge/version-3.0-green.svg)

## Overview

This repository contains the official Go wrapper for the Mailjet API.

Check out all the resources and Go code examples in the [Official Documentation][doc].

## Table of contents

- [Compatibility](#compatibility)
- [Installation](#installation)
- [Authentication](#authentication)
	- [Functional test](#functional-test)
- [Make your first call](#make-your-first-call)
- [Client / Call configuration specifics](#client--call-configuration-specifics)
  - [API versioning](#api-versioning)
	- [Send emails through proxy](#send-email-through-proxy)
- [Request examples](#request-examples)
  - [POST request](#post-request)
    - [Simple POST request](#simple-post-request)
    - [Using actions](#using-actions)
  - [GET request](#get-request)
    - [Retrieve all objects](#retrieve-all-objects)
    - [Use filtering](#use-filtering)
    - [Retrieve a single object](#retrieve-a-single-object)
  - [PUT request](#put-request)
  - [DELETE request](#delete-request)
  - [Response](#response)
  - [API resources helpers](#api-resources-helpers)
- [Contribute](#contribute)

## Compatibility

Our library requires Go version 1.13 or higher.

**NOTE: Backward compatibility has been broken with the `v3.0` release which includes versioned paths required by go modules (See [Releasing Modules](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher)).**

**Pin your dependencies to the `2.4.5` tag if you are not ready for `v3.0`**

### Installation

Get package:

```
go get github.com/mailjet/mailjet-apiv3-go/v4
```

And create a new MailjetClient:

```go
// Import the Mailjet wrapper
import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
	"os"
)
```

## Authentication

The Mailjet Email API uses your API and Secret keys for authentication. [Grab][api_credential] and save your Mailjet API credentials.

```bash
export MJ_APIKEY_PUBLIC='your API key'
export MJ_APIKEY_PRIVATE='your API secret'
```

Then initialize your Mailjet client:

```go
// Get your environment Mailjet keys and connect
publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

mj := mailjet.NewMailjetClient(publicKey, secretKey)
```

### Functional test

In the `tests` folder you will find a small program using the wrapper. It can be used to check whether the Mailjet API keys in your environment are valid and active.

```
go run main.go
```

## Make your first call

Here's an example on how to send an email:

```go
package main

import (
	"fmt"
	"log"
    "os"
  
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "pilot@mailjet.com",
				Name:  "Mailjet Pilot",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "passenger1@mailjet.com",
					Name:  "passenger 1",
				},
			},
			Subject:  "Your email flight plan!",
			TextPart: "Dear passenger 1, welcome to Mailjet! May the delivery force be with you!",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you!",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
```

## Client / Call configuration specifics

### Base URL

The default base domain name for the Mailjet API is `https://api.mailjet.com`. You can modify this base URL by adding a different URL in the client configuration for your call:

```go
mailjetClient := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"), "https://api.us.mailjet.com")
```

If your account has been moved to Mailjet's **US architecture**, the URL value you need to set is `https://api.us.mailjet.com`.

### Send emails through proxy

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

// Set the http client with the given proxy url
func setupProxy(proxyURLStr string) *http.Client {
	proxyURL, err := url.Parse(proxyURLStr)
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{}
	client.Transport = tr

	return client
}

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")
	proxyURL  := os.Getenv("HTTP_PROXY")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	// Here we inject our http client configured with our proxy
	client := setupProxy(proxyURL)
	mj.SetClient(client)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "qwe@qwe.com",
				Name:  "Bob Patrick",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "qwe@qwe.com",
				},
			},
			Subject:  "Hello World!",
			TextPart: "Hi there !",
		},
	}

	messages := &mailjet.MessagesV31{Info: messagesInfo}

	res, err := mj.SendMailV31(messages)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}
```

## Request examples

### POST request

#### Simple POST request

```go
/*
Create a new contact.
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	var data []resources.Contact
	mr := &mailjet.Request{
		Resource: "contact",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.Contact{
			Email:                   "passenger@mailjet.com",
			IsExcludedFromCampaigns: true,
			Name:                    "New Contact",
		},
	}
	err := mailjetClient.Post(fmr, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}
```

#### Using actions

```go
/*
Create : Manage a contact subscription to a list
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	var data []resources.ContactManagecontactslists
	mr := &mailjet.Request{
		Resource: "contact",
		ID:       423, // replace with your contact ID here
		Action:   "managecontactslists",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.ContactManagecontactslists{
			ContactsLists: []resources.ContactsListAction{ // replace with your contact lists here
				{
					ListID: 432,
					Action: "addnoforce",
				},
				{
					ListID: 553,
					Action: "addforce",
				},
			},
		},
	}
	err := mailjetClient.Post(fmr, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}
```

### GET request

#### Retrieve all objects

```go
/*
Retrieve all contacts:
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	var data []resources.Contact
	_, _, err := mailjetClient.List("contact", &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}
```

#### Use filtering

```go
/*
Retrieve all contacts that are not in the campaign exclusion list :
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	var data []resources.Contact
	_, _, err := mailjetClient.List("contact", &data, mailjet.Filter("IsExcludedFromCampaigns", "false"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}
```

#### Retrieve a single object

```go
/*
Retrieve a specific contact ID :
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	var data []resources.Contact
	mr := &mailjet.Request{
		Resource: "contact",
		ID:       5234, // replace with your contact ID here
	}
	err := mailjetClient.Get(mr, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}
```

### PUT request

A `PUT` request in the Mailjet API will work as a `PATCH` request - the update will affect only the specified properties. The other properties of an existing resource will neither be modified, nor deleted. It also means that all non-mandatory properties can be omitted from your payload.

Here's an example of a `PUT` request:

```go
/*
Update the contact properties for a contact:
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	mr := &mailjet.Request{
		Resource: "contactdata",
		ID:       325, // replace with your contact ID here
		//AltID: "user1@example.com", // alternatively you can use contact's email
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.Contactdata{
			Data: resources.KeyValueList{
				{
					"Name":  "name",
					"Value": "John",
				},
				{
					"Name":  "country",
					"Value": "Canada",
				},
			},
		},
	}
	err := mailjetClient.Put(fmr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
```

### DELETE request

Upon a successful DELETE request the response will not include a response body, but only a 204 No Content response code.

Here's an example of a DELETE request:

```go
/*
Delete an email template:
*/
package main

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	mr := &mailjet.Request{
		Resource: "template",
		ID:       423, // replace with your template ID here
	}

	err := mailjetClient.Delete(mr)
	if err != nil {
		fmt.Println(err)
	}
}
```

## Contribute

Mailjet loves developers. You can be part of this project!

This wrapper is a great introduction to the open source world, check out the code!

Feel free to ask anything, and contribute:

- Fork the project.
- Create a new branch.
- Implement your feature or bug fix.
- Add documentation to it.
- Commit, push, open a pull request and voila.

If you have suggestions on how to improve the guides, please submit an issue in our [Official API Documentation repo](https://github.com/mailjet/api-documentation).
