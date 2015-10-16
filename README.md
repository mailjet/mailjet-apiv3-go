[doc]: http://dev.mailjet.com/
[api_credential]: https://app.mailjet.com/account/api_keys
[issues]: https://github.com/mailjet/mailjet-apiv3-go/issues
[go_documentation]:http://dev.mailjet.com/guides/?go

# Mailjet-apiv3-go

This GO library provides client functionality for version 3 of the [Mailjet API][doc].

## Getting Started

Every code examples can be find on the [Mailjet Documentation][go_documentation]

### Prerequisites

Make sure to have the following requirements:
* A Mailjet API Key
* A Mailjet API Secret Key
* A Go installation (v. >= 1.3)

API key and an API secret can be found [here][api_credential].

Get cosy with Mailjet and save your credentials in your environment:
```
export MJ_APIKEY_PUBLIC='your api key'
export MJ_APIKEY_PRIVATE='your api secret'
```

### Installation

Get package:
```
go get github.com/mailjet/mailjet-apiv3-go
```

And create a new MailjetClient:
``` go
// Import the mailjet wrapper
import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"os"
)

[...]

// Get your environment Mailjet keys and connect
publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

mj := mailjet.NewMailjetClient(publicKey, secretKey)

```

It's ready to use !

## Examples

### List resources
``` go
package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"fmt"
	"os"
)

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	var res []resources.Metadata
	count, total, err := mj.List("metadata", &res)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Count: %d\nTotal: %d\n", count, total)

	fmt.Println("Resources:")
	for _, resource := range res {
		fmt.Println(resource.Name)
	}
}
```

### Create a resource
``` go
package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"fmt"
	"os"
)

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	var senders []resources.Sender
	fmr := &FullMailjetRequest{
		Info:    &MailjetRequest{Resource: "sender"},
		Payload: &resources.Sender{Name: "Default", Email: "qwe@qwe.com"},
	}
	err := mj.Post(fmr, &senders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if senders != nil {
		fmt.Printf("Data struct: %+v\n", senders[0])
	}
}
```

### Update a resource
``` go
package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"fmt"
	"os"
)

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	fmr := &FullMailjetRequest{
		Info:    &MailjetRequest{Resource: "sender", AltID: "qwe@qwe.com"},
		Payload: &resources.Sender{Name: "Bob", IsDefaultSender: true},
	}
	err := mj.Put(fmr, []string{"Name", "IsDefaultSender"})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("Success")
	}
}
```

### View a resource
``` go
package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"fmt"
	"os"
)

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	var senders []resources.Sender
	info := &MailjetRequest {
		Resource: "sender",
		AltID:    "qwe@qwe.com",
	}
	err := mj.Get(info, &senders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if senders != nil {
		fmt.Printf("Sender struct: %+v\n", senders[0])
	}
}
```

### Delete a resource
``` go
package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"fmt"
	"os"
)

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	info := &MailjetRequest{
		Resource: "sender",
		AltID:    "qwe@qwe.com",
	}
	err := mj.Delete(info)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("Success")
	}
}
```

### Send a mail
``` go
package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"fmt"
	"os"
)

func main() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	param := &mailjet.MailjetSendMail{
		FromEmail: "qwe@qwe.com",
		FromName: "Bob Patrick",
		Recipients: []mailjet.MailjetRecipient{
			mailjet.MailjetRecipient{
				Email: "qwe@qwe.com",
			},
		},
		Subject: "Hello World!",
		TextPart: "Hi there !",
	}
	res, err := mj.SendMail(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}
```

## Contribute

We welcome any contribution.

Please make sure you follow this step by step guide before contributing :

* Fork the project.
* Create a topic branch.
* Implement your feature or bug fix.
* Add documentation for your feature or bug fix.
* Commit and push your changes.
* Submit a pull request

Submit yours issues, [here][issues]!.

