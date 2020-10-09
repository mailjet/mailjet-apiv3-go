// Package resources provides mailjet resources properties. This is an helper and
// not a mandatory package.
package resources

import (
	"bytes"
	"time"
)

//
// Resources Properties
//

// Aggregategraphstatistics: Aggregated campaign statistics grouped over intervals.
type Aggregategraphstatistics struct {
	BlockedCount        float64 `mailjet:"read_only"`
	BlockedStdDev       float64 `mailjet:"read_only"`
	BouncedCount        float64 `mailjet:"read_only"`
	BouncedStdDev       float64 `mailjet:"read_only"`
	CampaignAggregateID int     `mailjet:"read_only"`
	ClickedCount        float64 `mailjet:"read_only"`
	ClickedStdDev       float64 `mailjet:"read_only"`
	OpenedCount         float64 `mailjet:"read_only"`
	OpenedStdDev        float64 `mailjet:"read_only"`
	RefTimestamp        int     `mailjet:"read_only"`
	SentCount           float64 `mailjet:"read_only"`
	SentStdDev          float64 `mailjet:"read_only"`
	SpamComplaintCount  float64 `mailjet:"read_only"`
	SpamcomplaintStdDev float64 `mailjet:"read_only"`
	UnsubscribedCount   float64 `mailjet:"read_only"`
	UnsubscribedStdDev  float64 `mailjet:"read_only"`
}

// Apikey: Manage your Mailjet API Keys.
// API keys are used as credentials to access the API and SMTP server.
type Apikey struct {
	ACL             string           `json:",omitempty"`
	APIKey          string           `mailjet:"read_only"`
	CreatedAt       *RFC3339DateTime `mailjet:"read_only"`
	ID              int64            `mailjet:"read_only"`
	IsActive        bool             `json:",omitempty"`
	IsMaster        bool             `mailjet:"read_only"`
	Name            string
	QuarantineValue int      `mailjet:"read_only"`
	Runlevel        RunLevel `mailjet:"read_only"`
	SecretKey       string   `mailjet:"read_only"`
	TrackHost       string   `mailjet:"read_only"`
	UserID          int64    `mailjet:"read_only"`
}

// Apikeyaccess: Access rights description on API keys for subaccounts/users.
type Apikeyaccess struct {
	AllowedAccess  string           `json:",omitempty"`
	APIKeyID       int64            `json:",omitempty"`
	APIKeyALT      string           `json:",omitempty"`
	CreatedAt      *RFC3339DateTime `json:",omitempty"`
	CustomName     string           `json:",omitempty"`
	ID             int64            `mailjet:"read_only"`
	IsActive       bool             `json:",omitempty"`
	LastActivityAt *RFC3339DateTime `json:",omitempty"`
	RealUserID     int64            `json:",omitempty"`
	RealUserALT    string           `json:",omitempty"`
	Subaccount     *SubAccount      `json:",omitempty"`
	UserID         int64            `json:",omitempty"`
	UserALT        string           `json:",omitempty"`
}

// Apikeytotals: Global counts for an API Key, since its creation.
type Apikeytotals struct {
	BlockedCount       int64 `mailjet:"read_only"`
	BouncedCount       int64 `mailjet:"read_only"`
	ClickedCount       int64 `mailjet:"read_only"`
	DeliveredCount     int64 `mailjet:"read_only"`
	LastActivity       int64 `mailjet:"read_only"`
	OpenedCount        int64 `mailjet:"read_only"`
	ProcessedCount     int64 `mailjet:"read_only"`
	QueuedCount        int64 `mailjet:"read_only"`
	SpamcomplaintCount int64 `mailjet:"read_only"`
	UnsubscribedCount  int64 `mailjet:"read_only"`
}

// Apitoken: Access token for API, used to give access to an API Key in conjunction with our IFrame API.
type Apitoken struct {
	ACL           string `json:",omitempty"`
	AllowedAccess string
	APIKeyID      int64            `json:",omitempty"`
	APIKeyALT     string           `json:",omitempty"`
	CatchedIP     string           `json:"CatchedIp,omitempty"`
	CreatedAt     *RFC3339DateTime `json:",omitempty"`
	FirstUsedAt   *RFC3339DateTime `json:",omitempty"`
	ID            int64            `mailjet:"read_only"`
	IsActive      bool             `json:",omitempty"`
	Lang          string           `json:",omitempty"`
	LastUsedAt    *RFC3339DateTime `json:",omitempty"`
	SentData      string           `json:",omitempty"`
	Timezone      string           `json:",omitempty"`
	Token         string           `json:",omitempty"`
	TokenType     string
	ValidFor      int `json:",omitempty"`
}

// Axtesting: AX testing object
type Axtesting struct {
	ContactListID   int64            `json:",omitempty"`
	ContactListALT  string           `json:",omitempty"`
	CreatedAt       *RFC3339DateTime `json:",omitempty"`
	Deleted         bool             `json:",omitempty"`
	ID              int64            `mailjet:"read_only"`
	Mode            AXTestMode       `json:",omitempty"`
	Name            string           `json:",omitempty"`
	Percentage      float64          `json:",omitempty"`
	RemainderAt     *RFC3339DateTime `json:",omitempty"`
	SegmentationID  int64            `json:",omitempty"`
	SegmentationALT string           `json:",omitempty"`
	Starred         bool             `json:",omitempty"`
	StartAt         *RFC3339DateTime `json:",omitempty"`
	Status          string           `json:",omitempty"`
	StatusCode      int              `mailjet:"read_only"`
	StatusString    string           `json:",omitempty"`
	WinnerClickRate float64          `json:",omitempty"`
	WinnerID        int              `json:",omitempty"`
	WinnerMethod    WinnerMethod     `json:",omitempty"`
	WinnerOpenRate  float64          `json:",omitempty"`
	WinnerSpamRate  float64          `json:",omitempty"`
	WinnerUnsubRate float64          `json:",omitempty"`
}

// Batchjob: Batch jobs running on the Mailjet infrastructure.
type Batchjob struct {
	AliveAt     int64  `json:",omitempty"`
	APIKeyID    int64  `json:",omitempty"`
	APIKeyALT   string `json:",omitempty"`
	Blocksize   int    `json:",omitempty"`
	Count       int    `json:",omitempty"`
	Current     int    `json:",omitempty"`
	Data        *BaseData
	Errcount    int   `json:",omitempty"`
	ErrTreshold int   `json:",omitempty"`
	ID          int64 `mailjet:"read_only"`
	JobEnd      int64 `json:",omitempty"`
	JobStart    int64 `json:",omitempty"`
	JobType     string
	Method      string `json:",omitempty"`
	RefID       int64  `json:"RefID,omitempty"`
	RequestAt   int64  `json:",omitempty"`
	Status      string `json:",omitempty"`
	Throttle    int    `json:",omitempty"`
}

// Bouncestatistics: Statistics on the bounces generated by emails sent on a given API Key.
type Bouncestatistics struct {
	BouncedAt        *RFC3339DateTime `mailjet:"read_only"`
	CampaignID       int64            `mailjet:"read_only"`
	CampaignALT      string           `mailjet:"read_only"`
	ContactID        int64            `mailjet:"read_only"`
	ContactALT       string           `mailjet:"read_only"`
	ID               int64            `mailjet:"read_only"`
	IsBlocked        bool             `mailjet:"read_only"`
	IsStatePermanent bool             `mailjet:"read_only"`
	StateID          int64            `mailjet:"read_only"`
}

// Campaign: Historical view of sent emails, both transactional and marketing.
// Each e-mail going through Mailjet is attached to a Campaign.
// This object is automatically generated by Mailjet.
type Campaign struct {
	CampaignType            int              `mailjet:"read_only"`
	ClickTracked            int64            `mailjet:"read_only"`
	CreatedAt               *RFC3339DateTime `mailjet:"read_only"`
	CustomValue             string           `mailjet:"read_only"`
	FirstMessageID          int64            `mailjet:"read_only"`
	FromID                  int64            `mailjet:"read_only"`
	FromALT                 string           `mailjet:"read_only"`
	FromEmail               string           `mailjet:"read_only"`
	FromName                string           `mailjet:"read_only"`
	HasHtmlCount            int64            `mailjet:"read_only"`
	HasTxtCount             int64            `mailjet:"read_only"`
	ID                      int64            `mailjet:"read_only"`
	IsDeleted               bool             `json:",omitempty"`
	IsStarred               bool             `json:",omitempty"`
	ListID                  int64            `mailjet:"read_only"`
	ListALT                 string           `mailjet:"read_only"`
	NewsLetterID            int64            `mailjet:"read_only"`
	OpenTracked             int64            `mailjet:"read_only"`
	SegmentationID          int64            `mailjet:"read_only"`
	SegmentationALT         string           `mailjet:"read_only"`
	SendEndAt               *RFC3339DateTime `mailjet:"read_only"`
	SendStartAt             *RFC3339DateTime `mailjet:"read_only"`
	SpamassScore            float64          `mailjet:"read_only"`
	Status                  string           `mailjet:"read_only"`
	Subject                 string           `mailjet:"read_only"`
	UnsubscribeTrackedCount int64            `mailjet:"read_only"`
}

// Campaignaggregate: User defined campaign aggregates
type Campaignaggregate struct {
	CampaignIDS      string           `json:",omitempty"`
	ContactFilterID  int64            `json:",omitempty"`
	ContactFilterALT string           `json:",omitempty"`
	ContactsListID   int64            `json:",omitempty"`
	ContactsListALT  string           `json:",omitempty"`
	Final            bool             `mailjet:"read_only"`
	FromDate         *RFC3339DateTime `json:",omitempty"`
	ID               int64            `mailjet:"read_only"`
	Keyword          string           `json:",omitempty"`
	Name             string           `json:",omitempty"`
	SenderID         int64            `json:",omitempty"`
	SenderALT        string           `json:",omitempty"`
	ToDate           *RFC3339DateTime `json:",omitempty"`
}

// Campaigndraft: Newsletter and CampaignDraft objects are differentiated by the EditMode values.
type Campaigndraft struct {
	AXFractionName     string           `json:",omitempty"`
	AXTesting          *Axtesting       `json:",omitempty"`
	CampaignID         int64            `json:",omitempty"`
	CampaignALT        string           `json:",omitempty"`
	ContactsListID     int64            `json:",omitempty"`
	ContactsListALT    string           `json:",omitempty"`
	CreatedAt          *RFC3339DateTime `json:",omitempty"`
	Current            int64            `json:",omitempty"`
	DeliveredAt        *RFC3339DateTime `json:",omitempty"`
	EditMode           string           `json:",omitempty"`
	ID                 int64            `mailjet:"read_only"`
	IsStarred          bool             `json:",omitempty"`
	IsTextPartIncluded bool             `json:",omitempty"`
	Locale             string
	ModifiedAt         *RFC3339DateTime `json:",omitempty"`
	Preset             string           `json:",omitempty"`
	ReplyEmail         string           `json:",omitempty"`
	SegmentationID     int64            `json:",omitempty"`
	SegmentationALT    string           `json:",omitempty"`
	Sender             string
	SenderEmail        string
	SenderName         string `json:",omitempty"`
	Status             string `mailjet:"read_only"`
	Subject            string
	TemplateID         int64  `json:",omitempty"`
	TemplateALT        string `json:",omitempty"`
	Title              string `json:",omitempty"`
	URL                string `json:"Url,omitempty"`
	Used               bool   `json:",omitempty"`
}

// CampaigndraftSchedule:
type CampaigndraftSchedule struct {
	Date *RFC3339DateTime
}

// CampaigndraftTest:
type CampaigndraftTest struct {
	Recipients []Recipient
}

// CampaigndraftDetailcontent:
type CampaigndraftDetailcontent struct {
	TextPart    string      `json:"Text-part,omitempty"`
	HtmlPart    string      `json:"Html-part,omitempty"`
	MJMLContent string      `json:",omitempty"`
	Headers     interface{} `json:",omitempty"`
}

// Campaigngraphstatistics: API Campaign statistics grouped over intervals
type Campaigngraphstatistics struct {
	Clickcount int64 `mailjet:"read_only"`
	ID         int64 `mailjet:"read_only"`
	Opencount  int64 `mailjet:"read_only"`
	Spamcount  int64 `mailjet:"read_only"`
	Tick       int64 `mailjet:"read_only"`
	Unsubcount int64 `mailjet:"read_only"`
}

// Campaignoverview: Returns a list of campaigns, including the AX campaigns
type Campaignoverview struct {
	ClickedCount   int64  `mailjet:"read_only"`
	DeliveredCount int64  `mailjet:"read_only"`
	EditMode       string `mailjet:"read_only"`
	EditType       string `mailjet:"read_only"`
	ID             int64  `mailjet:"read_only"`
	IDType         string `mailjet:"read_only"`
	OpenedCount    int64  `mailjet:"read_only"`
	ProcessedCount int64  `mailjet:"read_only"`
	SendTimeStart  int64  `mailjet:"read_only"`
	Starred        bool   `mailjet:"read_only"`
	Status         int    `mailjet:"read_only"`
	Subject        string `mailjet:"read_only"`
	Title          string `mailjet:"read_only"`
}

// Campaignstatistics: Statistics related to emails processed by Mailjet, grouped in a Campaign.
type Campaignstatistics struct {
	AXTesting           *Axtesting       `mailjet:"read_only"`
	BlockedCount        int64            `mailjet:"read_only"`
	BouncedCount        int64            `mailjet:"read_only"`
	CampaignID          int64            `mailjet:"read_only"`
	CampaignALT         string           `mailjet:"read_only"`
	CampaignIsStarred   bool             `mailjet:"read_only"`
	CampaignSendStartAt *RFC3339DateTime `mailjet:"read_only"`
	CampaignSubject     string           `mailjet:"read_only"`
	ClickedCount        int64            `mailjet:"read_only"`
	ContactListName     string           `mailjet:"read_only"`
	DeliveredCount      int64            `mailjet:"read_only"`
	LastActivityAt      *RFC3339DateTime `mailjet:"read_only"`
	NewsLetterID        int64            `mailjet:"read_only"`
	OpenedCount         int64            `mailjet:"read_only"`
	ProcessedCount      int64            `mailjet:"read_only"`
	QueuedCount         int64            `mailjet:"read_only"`
	SegmentName         string           `mailjet:"read_only"`
	SpamComplaintCount  int64            `mailjet:"read_only"`
	UnsubscribedCount   int64            `mailjet:"read_only"`
}

// Clickstatistics: Click statistics for messages.
type Clickstatistics struct {
	ClickedAt    string `mailjet:"read_only"`
	ClickedDelay int64  `mailjet:"read_only"`
	ContactID    int64  `mailjet:"read_only"`
	ContactALT   string `mailjet:"read_only"`
	ID           int64  `mailjet:"read_only"`
	MessageID    int64  `mailjet:"read_only"`
	URL          string `json:"Url" mailjet:"read_only"`
	UserAgent    string `mailjet:"read_only"`
}

// Contact: Manage the details of a Contact.
type Contact struct {
	CreatedAt         *RFC3339DateTime `mailjet:"read_only"`
	DeliveredCount    int64            `mailjet:"read_only"`
	Email             string
	ID                int64            `mailjet:"read_only"`
	IsOptInPending    bool             `mailjet:"read_only"`
	IsSpamComplaining bool             `mailjet:"read_only"`
	LastActivityAt    *RFC3339DateTime `mailjet:"read_only"`
	LastUpdateAt      *RFC3339DateTime `mailjet:"read_only"`
	Name              string           `json:",omitempty"`
	UnsubscribedAt    *RFC3339DateTime `mailjet:"read_only"`
	UnsubscribedBy    string           `mailjet:"read_only"`
}

// ContactManagecontactslists: Managing the lists for a single contact. POST is supported.
type ContactManagecontactslists struct {
	ContactsLists []ContactsListAction
}

// ContactManagemanycontacts: Uploading many contacts and returns a job_id.
// To monitor the upload issue a GET request to: APIBASEURL/contact/managemanycontacts/:job_id
type ContactManagemanycontacts struct {
	ContactsLists []ContactsListAction
	Contacts      []AddContactAction
}

// Contactdata: This resource can be used to examine and manipulate the associated extra static data of a contact.
type Contactdata struct {
	ContactID int64        `json:",omitempty"`
	Data      KeyValueList `json:",omitempty"`
	ID        int64        `mailjet:"read_only"`
}

// Contactfilter: A list of filter expressions for use in newsletters.
type Contactfilter struct {
	Description string `json:",omitempty"`
	Expression  string `json:",omitempty"`
	ID          int64  `mailjet:"read_only"`
	Name        string `json:",omitempty"`
	Status      string `json:",omitempty"`
}

// Contacthistorydata: This resource can be used to examine the associated extra historical data of a contact.
type Contacthistorydata struct {
	ContactID  int64            `json:",omitempty"`
	ContactALT string           `json:",omitempty"`
	CreatedAt  *RFC3339DateTime `json:",omitempty"`
	Data       string           `json:",omitempty"`
	ID         int64            `mailjet:"read_only"`
	Name       string           `json:",omitempty"`
}

// Contactmetadata: Definition of available extra data items for contacts.
type Contactmetadata struct {
	Datatype  string
	ID        int64 `mailjet:"read_only"`
	Name      string
	NameSpace string `mailjet:"read_only"`
}

// Contactslist: Manage your contact lists. One Contact might be associated to one or more ContactsList.
type Contactslist struct {
	Address         string           `mailjet:"read_only"`
	CreatedAt       *RFC3339DateTime `mailjet:"read_only"`
	ID              int64            `mailjet:"read_only"`
	IsDeleted       bool             `json:",omitempty"`
	Name            string           `json:",omitempty"`
	SubscriberCount int              `mailjet:"read_only"`
}

// ContactslistManageContact: An action for adding a contact to a contact list.
// The API will internally create the new contact if it does not exist, add or update the name and properties.
// The properties have to be defined before they can be used.
// The API then adds the contact to the contact list with active=true and unsub=specified value
// if it is not already in the list, or updates the entry with these values.
// On success, the API returns a packet with the same format but with all properties available for that contact.
// Only POST is supported.
type ContactslistManageContact struct {
	Email      string
	Name       string
	Action     string
	Properties JSONObject
}

// ContactslistManageManyContacts: Multiple contacts can be uploaded asynchronously using that action.
// Only POST is supported.
type ContactslistManageManyContacts struct {
	Action   string
	Contacts []AddContactAction
}

// ContactslistImportList: Import the contacts of another contact list into the current list and apply the specified action on imported contacts.
// In case of conflict, the contact original subscription state is overridden. Returns the ID of the job to monitor.
type ContactslistImportList struct {
	Action string
	ListID int64
}

// Contactslistsignup: Contacts list signup request.
type Contactslistsignup struct {
	ConfirmAt  int64  `json:",omitempty"`
	ConfirmIP  string `json:"ConfirmIp"`
	ContactID  int64  `json:",omitempty"`
	ContactALT string `json:",omitempty"`
	Email      string
	ID         int64  `mailjet:"read_only"`
	ListID     int64  `json:",omitempty"`
	ListALT    string `json:",omitempty"`
	SignupAt   int64  `json:",omitempty"`
	SignupIP   string `json:"SignupIp,omitempty"`
	SignupKey  string `json:",omitempty"`
	Source     string
	SourceId   int64 `json:",omitempty"`
}

// Contactstatistics: View message statistics for a given contact.
type Contactstatistics struct {
	BlockedCount          int64            `mailjet:"read_only"`
	BouncedCount          int64            `mailjet:"read_only"`
	ClickedCount          int64            `mailjet:"read_only"`
	ContactID             int64            `mailjet:"read_only"`
	ContactALT            string           `mailjet:"read_only"`
	DeliveredCount        int64            `mailjet:"read_only"`
	LastActivityAt        *RFC3339DateTime `mailjet:"read_only"`
	MarketingContacts     int64            `mailjet:"read_only"`
	OpenedCount           int64            `mailjet:"read_only"`
	ProcessedCount        int64            `mailjet:"read_only"`
	QueuedCount           int64            `mailjet:"read_only"`
	SpamComplaintCount    int64            `mailjet:"read_only"`
	UnsubscribedCount     int64            `mailjet:"read_only"`
	UserMarketingContacts int64            `mailjet:"read_only"`
}

// Csvimport: A wrapper for the CSV importer
type Csvimport struct {
	AliveAt         *RFC3339DateTime `mailjet:"read_only"`
	ContactsListID  int64            `json:",omitempty"`
	ContactsListALT string           `json:",omitempty"`
	Count           int              `mailjet:"read_only"`
	Current         int              `mailjet:"read_only"`
	DataID          int64
	Errcount        int              `mailjet:"read_only"`
	ErrTreshold     int              `json:",omitempty"`
	ID              int64            `mailjet:"read_only"`
	ImportOptions   string           `json:",omitempty"`
	JobEnd          *RFC3339DateTime `mailjet:"read_only"`
	JobStart        *RFC3339DateTime `mailjet:"read_only"`
	Method          string           `json:",omitempty"`
	RequestAt       *RFC3339DateTime `mailjet:"read_only"`
	Status          string           `json:",omitempty"`
}

// Dns: Sender Domain properties.
type Dns struct {
	DKIMRecordName           string           `mailjet:"read_only"`
	DKIMRecordValue          string           `mailjet:"read_only"`
	DKIMStatus               string           `mailjet:"read_only"`
	Domain                   string           `mailjet:"read_only"`
	ID                       int64            `mailjet:"read_only"`
	IsCheckInProgress        bool             `mailjet:"read_only"`
	LastCheckAt              *RFC3339DateTime `mailjet:"read_only"`
	OwnerShipToken           string           `mailjet:"read_only"`
	OwnerShipTokenRecordName string           `mailjet:"read_only"`
	SPFRecordValue           string           `mailjet:"read_only"`
	SPFStatus                string           `mailjet:"read_only"`
}

type DnsCheck struct {
	DKIMErrors             []string `mailjet:"read_only"`
	DKIMStatus             string   `mailjet:"read_only"`
	DKIMRecordCurrentValue string   `mailjet:"read_only"`
	SPFRecordCurrentValue  string   `mailjet:"read_only"`
	SPFErrors              []string `mailjet:"read_only"`
	SPFStatus              string   `mailjet:"read_only"`
}

// Domainstatistics: View Campaign/Message/Click statistics grouped per domain.
type Domainstatistics struct {
	BlockedCount       int64  `mailjet:"read_only"`
	BouncedCount       int64  `mailjet:"read_only"`
	ClickedCount       int64  `mailjet:"read_only"`
	DeliveredCount     int64  `mailjet:"read_only"`
	Domain             string `mailjet:"read_only"`
	ID                 int64  `mailjet:"read_only"`
	OpenedCount        int64  `mailjet:"read_only"`
	ProcessedCount     int64  `mailjet:"read_only"`
	QueuedCount        int64  `mailjet:"read_only"`
	SpamComplaintCount int64  `mailjet:"read_only"`
	UnsubscribedCount  int64  `mailjet:"read_only"`
}

// Eventcallbackurl: Manage event-driven callback URLs, also called webhooks,
// used by the Mailjet platform when a specific action is triggered
type Eventcallbackurl struct {
	APIKeyID  int64  `json:",omitempty"`
	APIKeyALT string `json:",omitempty"`
	EventType string `json:",omitempty"`
	ID        int64  `mailjet:"read_only"`
	IsBackup  bool   `json:",omitempty"`
	Status    string `json:",omitempty"`
	URL       string `json:"Url"`
	Version   int    `json:",omitempty"`
}

// Geostatistics: Message click/open statistics grouped per country
type Geostatistics struct {
	ClickedCount int64  `mailjet:"read_only"`
	Country      string `mailjet:"read_only"`
	OpenedCount  int64  `mailjet:"read_only"`
}

// Graphstatistics: API Campaign/message/click statistics grouped over intervals.
type Graphstatistics struct {
	BlockedCount       int64  `mailjet:"read_only"`
	BouncedCount       int64  `mailjet:"read_only"`
	ClickedCount       int64  `mailjet:"read_only"`
	DeliveredCount     int64  `mailjet:"read_only"`
	OpenedCount        int64  `mailjet:"read_only"`
	ProcessedCount     int64  `mailjet:"read_only"`
	QueuedCount        int64  `mailjet:"read_only"`
	RefTimestamp       string `mailjet:"read_only"`
	SendtimeStart      int64  `mailjet:"read_only"`
	SpamcomplaintCount int64  `mailjet:"read_only"`
	UnsubscribedCount  int64  `mailjet:"read_only"`
}

// Listrecipient: Manage the relationship between a contact and a contactslists.
type Listrecipient struct {
	ContactID      int64            `mailjet:"read_only"`
	ContactALT     string           `mailjet:"read_only"`
	ID             int64            `mailjet:"read_only"`
	IsActive       bool             `json:",omitempty"`
	IsUnsubscribed bool             `json:",omitempty"`
	ListID         int64            `json:",omitempty"`
	ListALT        string           `json:",omitempty"`
	UnsubscribedAt *RFC3339DateTime `json:",omitempty"`
}

// Listrecipientstatistics: View statistics on Messages sent to the recipients of a given list.
type Listrecipientstatistics struct {
	BlockedCount       int64            `mailjet:"read_only"`
	BouncedCount       int64            `mailjet:"read_only"`
	ClickedCount       int64            `mailjet:"read_only"`
	Data               KeyValueList     `mailjet:"read_only"`
	DeliveredCount     int64            `mailjet:"read_only"`
	LastActivityAt     *RFC3339DateTime `mailjet:"read_only"`
	ListRecipientID    int64            `mailjet:"read_only"`
	OpenedCount        int64            `mailjet:"read_only"`
	ProcessedCount     int64            `mailjet:"read_only"`
	QueuedCount        int64            `mailjet:"read_only"`
	SpamComplaintCount int64            `mailjet:"read_only"`
	UnsubscribedCount  int64            `mailjet:"read_only"`
}

// Liststatistics: View Campaign/message/click statistics grouped by ContactsList.
type Liststatistics struct {
	ActiveCount             int64            `mailjet:"read_only"`
	ActiveUnsubscribedCount int64            `mailjet:"read_only"`
	Address                 string           `mailjet:"read_only"`
	BlockedCount            int64            `mailjet:"read_only"`
	BouncedCount            int64            `mailjet:"read_only"`
	ClickedCount            int64            `mailjet:"read_only"`
	CreatedAt               *RFC3339DateTime `mailjet:"read_only"`
	DeliveredCount          int64            `mailjet:"read_only"`
	ID                      int64            `mailjet:"read_only"`
	IsDeleted               bool             `mailjet:"read_only"`
	LastActivityAt          *RFC3339DateTime `mailjet:"read_only"`
	Name                    string           `mailjet:"read_only"`
	OpenedCount             int64            `mailjet:"read_only"`
	SpamComplaintCount      int64            `mailjet:"read_only"`
	SubscriberCount         int              `mailjet:"read_only"`
	UnsubscribedCount       int64            `mailjet:"read_only"`
}

// Message: Allows you to list and view the details of a Message (an e-mail) processed by Mailjet
type Message struct {
	ArrivedAt          *RFC3339DateTime `json:",omitempty"`
	AttachmentCount    int              `json:",omitempty"`
	AttemptCount       int              `json:",omitempty"`
	CampaignID         int64            `json:",omitempty"`
	CampaignALT        string           `json:",omitempty"`
	ContactID          int64            `json:",omitempty"`
	ContactALT         string           `json:",omitempty"`
	Delay              float64          `json:",omitempty"`
	Destination        Destination
	FilterTime         int                `json:",omitempty"`
	FromID             int64              `json:",omitempty"`
	FromALT            string             `json:",omitempty"`
	ID                 int64              `mailjet:"read_only"`
	IsClickTracked     bool               `json:",omitempty"`
	IsHTMLPartIncluded bool               `json:",omitempty"`
	IsOpenTracked      bool               `json:",omitempty"`
	IsTextPartIncluded bool               `json:",omitempty"`
	IsUnsubTracked     bool               `json:",omitempty"`
	MessageSize        int64              `json:",omitempty"`
	SpamassassinScore  float64            `json:",omitempty"`
	SpamassRules       []SpamAssassinRule `json:",omitempty"`
	StateID            int64              `json:",omitempty"`
	StatePermanent     bool               `json:",omitempty"`
	Status             string             `json:",omitempty"`
}

// Messagehistory: Event history of a message.
type Messagehistory struct {
	Comment   string `mailjet:"read_only"`
	EventAt   int64  `mailjet:"read_only"`
	EventType string `mailjet:"read_only"`
	State     string `mailjet:"read_only"`
	UserAgent string `json:"Useragent" mailjet:"read_only"`
}

// Messageinformation: API Key campaign/message information.
type Messageinformation struct {
	CampaignID        int64              `mailjet:"read_only"`
	CampaignALT       string             `mailjet:"read_only"`
	ClickTrackedCount int64              `mailjet:"read_only"`
	ContactID         int64              `mailjet:"read_only"`
	ContactALT        string             `mailjet:"read_only"`
	CreatedAt         *RFC3339DateTime   `mailjet:"read_only"`
	ID                int64              `mailjet:"read_only"`
	MessageSize       int64              `mailjet:"read_only"`
	OpenTrackedCount  int64              `mailjet:"read_only"`
	QueuedCount       int64              `mailjet:"read_only"`
	SendEndAt         *RFC3339DateTime   `mailjet:"read_only"`
	SentCount         int64              `mailjet:"read_only"`
	SpamAssassinRules []SpamAssassinRule `mailjet:"read_only"`
	SpamAssassinScore float64            `mailjet:"read_only"`
}

// Messagesentstatistics: API Key Statistical campaign/message data.
type Messagesentstatistics struct {
	ArrivalTs      *RFC3339DateTime `mailjet:"read_only"`
	Blocked        bool             `mailjet:"read_only"`
	Bounce         bool             `mailjet:"read_only"`
	BounceDate     *RFC3339DateTime `mailjet:"read_only"`
	BounceReason   string           `mailjet:"read_only"`
	CampaignID     int64            `mailjet:"read_only"`
	CampaignALT    string           `mailjet:"read_only"`
	Click          bool             `mailjet:"read_only"`
	CntRecipients  int64            `mailjet:"read_only"`
	ComplaintDate  *RFC3339DateTime `mailjet:"read_only"`
	ContactID      int64            `mailjet:"read_only"`
	ContactALT     string           `mailjet:"read_only"`
	Details        string           `mailjet:"read_only"`
	FBLSource      string           `mailjet:"read_only"`
	MessageID      int64            `mailjet:"read_only"`
	Open           bool             `mailjet:"read_only"`
	Queued         bool             `mailjet:"read_only"`
	Sent           bool             `mailjet:"read_only"`
	Spam           bool             `mailjet:"read_only"`
	StateID        int64            `mailjet:"read_only"`
	StatePermanent bool             `mailjet:"read_only"`
	Status         string           `mailjet:"read_only"`
	ToEmail        string           `mailjet:"read_only"`
	Unsub          bool             `mailjet:"read_only"`
}

// Messagestate: Message state reference.
type Messagestate struct {
	ID        int64  `mailjet:"read_only"`
	RelatedTo string `json:",omitempty"`
	State     string
}

// MessageStatistics: API key Campaign/Message statistics.
type MessageStatistics struct {
	AverageClickDelay   float64 `mailjet:"read_only"`
	AverageClickedCount float64 `mailjet:"read_only"`
	AverageOpenDelay    float64 `mailjet:"read_only"`
	AverageOpenedCount  float64 `mailjet:"read_only"`
	BlockedCount        int64   `mailjet:"read_only"`
	BouncedCount        int64   `mailjet:"read_only"`
	CampaignCount       int64   `mailjet:"read_only"`
	ClickedCount        int64   `mailjet:"read_only"`
	DeliveredCount      int64   `mailjet:"read_only"`
	OpenedCount         int64   `mailjet:"read_only"`
	ProcessedCount      int64   `mailjet:"read_only"`
	QueuedCount         int64   `mailjet:"read_only"`
	SpamComplaintCount  int64   `mailjet:"read_only"`
	TransactionalCount  int64   `mailjet:"read_only"`
	UnsubscribedCount   int64   `mailjet:"read_only"`
}

// Metadata: Mailjet API meta data.
type Metadata struct {
	APIVersion       string             `mailjet:"read_only"`
	Actions          []ResourceAction   `mailjet:"read_only"`
	Description      string             `mailjet:"read_only"`
	Filters          []ResourceFilter   `mailjet:"read_only"`
	IsReadOnly       bool               `mailjet:"read_only"`
	Name             string             `mailjet:"read_only"`
	Properties       []ResourceProperty `mailjet:"read_only"`
	PublicOperations string             `mailjet:"read_only"`
	SortInfo         []struct {
		AllowDescending bool   `mailjet:"read_only"`
		PropertyName    string `mailjet:"read_only"`
	} `mailjet:"read_only"`
	UniqueKey string `mailjet:"read_only"`
}

type ResourceAction struct {
	Description      string             `mailjet:"read_only"`
	IsGlobalAction   bool               `mailjet:"read_only"`
	Name             string             `mailjet:"read_only"`
	Parameters       []ResourceFilter   `mailjet:"read_only"`
	Properties       []ResourceProperty `mailjet:"read_only"`
	PublicOperations string             `mailjet:"read_only"`
}

type ResourceFilter struct {
	DataType     string `mailjet:"read_only"`
	DefaultValue string `mailjet:"read_only"`
	Description  string `mailjet:"read_only"`
	IsRequired   bool   `mailjet:"read_only"`
	Name         string `mailjet:"read_only"`
	ReadOnly     bool   `mailjet:"read_only"`
}

type ResourceProperty struct {
	DataType     string `mailjet:"read_only"`
	DefaultValue string `mailjet:"read_only"`
	Description  string `mailjet:"read_only"`
	IsRequired   bool   `mailjet:"read_only"`
	Name         string `mailjet:"read_only"`
	ReadOnly     bool   `mailjet:"read_only"`
}

// Metasender: Management of domains used for sending messages.
// A domain or address must be registered and validated before being used.
// See the related Sender object if you wish to register a given e-mail address.
type Metasender struct {
	CreatedAt   *RFC3339DateTime `json:",omitempty"`
	Description string           `json:",omitempty"`
	Email       string
	Filename    string `mailjet:"read_only"`
	ID          int64  `mailjet:"read_only"`
	IsEnabled   bool   `json:",omitempty"`
}

// Myprofile: Manage user profile data such as address, payment information etc.
type Myprofile struct {
	AddressCity           string           `json:",omitempty"`
	AddressCountry        string           `json:",omitempty"`
	AddressPostalCode     string           `json:",omitempty"`
	AddressState          string           `json:",omitempty"`
	AddressStreet         string           `json:",omitempty"`
	BillingEmail          string           `json:",omitempty"`
	BirthdayAt            *RFC3339DateTime `json:",omitempty"`
	CompanyName           string           `json:",omitempty"`
	CompanyNumOfEmployees string           `json:",omitempty"`
	ContactPhone          string           `json:",omitempty"`
	EstimatedVolume       int              `json:",omitempty"`
	Features              string           `json:",omitempty"`
	Firstname             string           `json:",omitempty"`
	ID                    int64            `mailjet:"read_only"`
	Industry              string           `json:",omitempty"`
	JobTitle              string           `json:",omitempty"`
	Lastname              string           `json:",omitempty"`
	UserID                int64            `json:",omitempty"`
	UserALT               string           `json:",omitempty"`
	VAT                   float64          `mailjet:"read_only"`
	VATNumber             string           `json:",omitempty"`
	Website               string           `json:",omitempty"`
}

// Newsletter: Newsletter data.
type Newsletter struct {
	AXFraction           float64          `json:",omitempty"`
	AXFractionName       string           `json:",omitempty"`
	AXTesting            *Axtesting       `json:",omitempty"`
	Callback             string           `json:",omitempty"`
	CampaignID           int64            `mailjet:"read_only"`
	CampaignALT          string           `mailjet:"read_only"`
	ContactsListID       int64            `json:",omitempty"`
	ContactsListALT      string           `json:",omitempty"`
	CreatedAt            *RFC3339DateTime `json:",omitempty"`
	DeliveredAt          *RFC3339DateTime `json:",omitempty"`
	EditMode             string           `json:",omitempty"`
	EditType             string           `json:",omitempty"`
	Footer               string           `json:",omitempty"`
	FooterAddress        string           `json:",omitempty"`
	FooterWYSIWYGType    int              `json:",omitempty"`
	HeaderFilename       string           `json:",omitempty"`
	HeaderLink           string           `json:",omitempty"`
	HeaderText           string           `json:",omitempty"`
	HeaderURL            string           `json:"HeaderUrl,omitempty"`
	ID                   int64            `mailjet:"read_only"`
	IP                   string           `json:"Ip,omitempty"`
	IsHandled            bool             `json:",omitempty"`
	IsStarred            bool             `json:",omitempty"`
	IsTextPartIncluded   bool             `json:",omitempty"`
	Locale               string
	ModifiedAt           *RFC3339DateTime `json:",omitempty"`
	Permalink            string           `json:",omitempty"`
	PermalinkHost        string           `json:",omitempty"`
	PermalinkWYSIWYGType int              `json:",omitempty"`
	PolitenessMode       int              `json:",omitempty"`
	ReplyEmail           string           `json:",omitempty"`
	SegmentationID       int64            `json:",omitempty"`
	SegmentationALT      string           `json:",omitempty"`
	Sender               string
	SenderEmail          string
	SenderName           string `json:",omitempty"`
	Status               string `json:",omitempty"`
	Subject              string
	TemplateID           int64  `json:",omitempty"`
	TestAddress          string `json:",omitempty"`
	Title                string `json:",omitempty"`
	URL                  string `json:"Url,omitempty"`
}

// NewsletterDetailcontent: An action to upload the content of the newsletter
type NewsletterDetailcontent struct {
	TextPart string `json:"Text-part,omitempty"`
	HtmlPart string `json:"Html-part,omitempty"`
}

// NewsletterSchedule: An action to schedule a newsletters.
type NewsLetterSchedule struct {
	Date *RFC3339DateTime
}

// NewsletterTest: An action to test a newsletter.
type NewsletterTest struct {
	Recipients []Recipient
}

// Newslettertemplate: Manages a Newsletter Template Properties.
type Newslettertemplate struct {
	CategoryID           int64            `json:",omitempty"`
	CreatedAt            *RFC3339DateTime `json:",omitempty"`
	Footer               string           `json:",omitempty"`
	FooterAddress        string           `json:",omitempty"`
	FooterWYSIWYGType    int              `json:",omitempty"`
	HeaderFilename       string           `json:",omitempty"`
	HeaderLink           string           `json:",omitempty"`
	HeaderText           string           `json:",omitempty"`
	HeaderURL            string           `json:"HeaderUrl,omitempty"`
	ID                   int64            `mailjet:"read_only"`
	Locale               string
	Name                 string `json:",omitempty"`
	Permalink            string `json:",omitempty"`
	PermalinkWYSIWYGType int    `json:",omitempty"`
	SourceNewsLetterID   int64  `json:",omitempty"`
	Status               string `json:",omitempty"`
}

// Newslettertemplatecategory: Manage categories for your newsletters.
// Allows you to group newsletters by category.
type Newslettertemplatecategory struct {
	Description      string
	ID               int64 `mailjet:"read_only"`
	Locale           string
	ParentCategoryID int64
	Value            string
}

// Openinformation: Retrieve informations about messages opened at least once by their recipients.
type Openinformation struct {
	ArrivedAt     *RFC3339DateTime `mailjet:"read_only"`
	CampaignID    int64            `mailjet:"read_only"`
	CampaignALT   string           `mailjet:"read_only"`
	ContactID     int64            `mailjet:"read_only"`
	ContactALT    string           `mailjet:"read_only"`
	ID            int64            `mailjet:"read_only"`
	MessageID     int64            `mailjet:"read_only"`
	OpenedAt      *RFC3339DateTime `mailjet:"read_only"`
	UserAgent     string           `mailjet:"read_only"`
	UserAgentFull string           `mailjet:"read_only"`
}

// Openstatistics: Retrieve statistics on e-mails opened at least once by their recipients.
type Openstatistics struct {
	OpenedCount    int64   `mailjet:"read_only"`
	OpenedDelay    float64 `mailjet:"read_only"`
	ProcessedCount int64   `mailjet:"read_only"`
}

// Parseroute: ParseRoute description
type Parseroute struct {
	APIKeyID  int64  `json:",omitempty"`
	APIKeyALT string `json:",omitempty"`
	Email     string `json:",omitempty"`
	ID        int64  `mailjet:"read_only"`
	URL       string `json:"Url"`
}

// Preferences: User preferences in key=value format.
type Preferences struct {
	ID      int64 `mailjet:"read_only"`
	Key     string
	UserID  int64  `json:",omitempty"`
	UserALT string `json:",omitempty"`
	Value   string `json:",omitempty"`
}

// Preset: The preset object contains global and user defined presets (styles) independent from templates or newsletters.
// Access is similar to template and depends on OwnerType, Owner. No versioning is done. Presets are never referenced by their ID.
// The preset value is copied into the template or newsletter.
type Preset struct {
	Author      string `json:",omitempty"`
	Copyright   string `json:",omitempty"`
	Description string `json:",omitempty"`
	ID          int64  `mailjet:"read_only"`
	Name        string `json:",omitempty"`
	OwnerID     int64  `json:",omitempty"`
	OwnerType   string `json:",omitempty"`
	Preset      string `json:",omitempty"`
}

// Sender: Manage an email sender for a single API key.
// An e-mail address or a complete domain (*) has to be registered and validated before being used to send e-mails.
// In order to manage a sender available across multiple API keys, see the related MetaSender resource.
type Sender struct {
	CreatedAt       *RFC3339DateTime `mailjet:"read_only"`
	DNS             string           `mailjet:"read_only"` // deprecated
	DNSID           int64            `mailjet:"read_only"`
	Email           string
	EmailType       string `json:",omitempty"`
	Filename        string `mailjet:"read_only"`
	ID              int64  `mailjet:"read_only"`
	IsDefaultSender bool   `json:",omitempty"`
	Name            string `json:",omitempty"`
	Status          string `mailjet:"read_only"`
}

// Senderstatistics: API Key sender email address message/open/click statistical information.
type Senderstatistics struct {
	BlockedCount       int64            `mailjet:"read_only"`
	BouncedCount       int64            `mailjet:"read_only"`
	ClickedCount       int64            `mailjet:"read_only"`
	DeliveredCount     int64            `mailjet:"read_only"`
	LastActivityAt     *RFC3339DateTime `mailjet:"read_only"`
	OpenedCount        int64            `mailjet:"read_only"`
	ProcessedCount     int64            `mailjet:"read_only"`
	QueuedCount        int64            `mailjet:"read_only"`
	SenderID           int64            `mailjet:"read_only"`
	SenderALT          string           `mailjet:"read_only"`
	SpamComplaintCount int64            `mailjet:"read_only"`
	UnsubscribedCount  int64            `mailjet:"read_only"`
}

// SenderValidate: validation result for a sender or domain
type SenderValidate struct {
	Errors           map[string]string `mailjet:"read_only"`
	ValidationMethod string            `mailjet:"read_only"`
	GlobalError      string            `mailjet:"read_only"`
}

// Template: template description
type Template struct {
	Author      string   `json:",omitempty"`
	Categories  []string `json:",omitempty"`
	Copyright   string   `json:",omitempty"`
	Description string   `json:",omitempty"`
	EditMode    int      `json:",omitempty"`
	ID          int64    `mailjet:"read_only"`
	IsStarred   bool     `json:",omitempty"`
	Name        string   `json:",omitempty"`
	OwnerId     int      `mailjet:"read_only"`
	OwnerType   string   `json:",omitempty"`
	Presets     string   `json:",omitempty"`
	Previews    []int64  `mailjet:"read_only"`
	Purposes    []string `json:",omitempty"`
}

// TemplateDetailcontent: GET, POST are supported to read, create, modify and delete the content of a template
type TemplateDetailcontent struct {
	TextPart    string      `json:"Text-part,omitempty"`
	HtmlPart    string      `json:"Html-part,omitempty"`
	MJMLContent MJMLContent `json:",omitempty"`
	Headers     interface{} `json:",omitempty"`
}

// MJMLContent: Structure of Passport template.
type MJMLContent struct {
	tagName    string
	attributes map[string]interface{}
	id         string
}

// Toplinkclicked: Top links clicked historgram.
type Toplinkclicked struct {
	ClickedCount int64  `mailjet:"read_only"`
	ID           int64  `mailjet:"read_only"`
	LinkID       int64  `json:"LinkId"mailjet:"read_only"`
	URL          string `json:"Url" mailjet:"read_only"`
}

// Trigger: Triggers for outgoing events.
type Trigger struct {
	AddedTs int64  `json:",omitempty"`
	APIKey  int    `json:",omitempty"`
	Details string `json:",omitempty"`
	Event   string `json:",omitempty"`
	ID      int64  `mailjet:"read_only"`
	User    int    `json:",omitempty"`
}

// User: User account definition for Mailjet.
type User struct {
	ACL               string           `json:",omitempty"`
	CreatedAt         *RFC3339DateTime `mailjet:"read_only"`
	Email             string           `json:",omitempty"`
	ID                int64            `mailjet:"read_only"`
	LastIp            string
	LastLoginAt       *RFC3339DateTime `json:",omitempty"`
	Locale            string           `json:",omitempty"`
	MaxAllowedAPIKeys int              `mailjet:"read_only"`
	Timezone          string           `json:",omitempty"`
	Username          string
	WarnedRatelimitAt *RFC3339DateTime `json:",omitempty"`
}

// Useragentstatistics: View statistics on User Agents.
// See total counts or filter per Campaign or Contacts List.
// API Key message Open/Click statistical data grouped per user agent (browser).
type Useragentstatistics struct {
	Count         int64  `mailjet:"read_only"`
	DistinctCount int64  `mailjet:"read_only"`
	Platform      string `mailjet:"read_only"`
	UserAgent     string `mailjet:"read_only"`
}

// Widget: Manage settings for Widgets.
// Widgets are small registration forms that you may include on your website to ease the process of subscribing to a Contacts List.
// Mailjet widget definitions.
type Widget struct {
	CreatedAt  int64  `json:",omitempty"`
	FromID     int64  `json:",omitempty"`
	FromALT    string `json:",omitempty"`
	ID         int64  `mailjet:"read_only"`
	IsActive   bool   `json:",omitempty"`
	ListID     int64  `json:",omitempty"`
	ListALT    string `json:",omitempty"`
	Locale     string
	Name       string           `json:",omitempty"`
	Replyto    string           `json:",omitempty"`
	Sendername string           `json:",omitempty"`
	Subject    string           `json:",omitempty"`
	Template   *MessageTemplate `json:",omitempty"`
}

// Widgetcustomvalue: Specifics settings for a given Mailjet Widget. See Widget.Mailjet widget settings.
type Widgetcustomvalue struct {
	APIKeyID  int64  `json:",omitempty"`
	APIKeyALT string `json:",omitempty"`
	Display   bool   `json:",omitempty"`
	ID        int64  `mailjet:"read_only"`
	Name      string
	Value     string `json:",omitempty"`
	WidgetID  int64
}

type AXTestMode string

const (
	AXTestAutomatic = AXTestMode("automatic")
	AXTestManual    = "manual"
)

type RunLevel string

const (
	RunNormal   = RunLevel("Normal")
	RunSoftlock = "Softlock"
	RunHardlock = "Hardlock"
)

type WinnerMethod string

const (
	WinnerOpenRate  = WinnerMethod("OpenRate")
	WinnerClickRate = "ClickRate"
	WinnerSpamRate  = "SpamRate"
	WinnerUnsubRate = "UnsubRate"
	WinnerMJScore   = "MJScore"
)

type BaseData struct {
	DataAsString string
	DataType     string
}

type SubAccount struct {
	IsActive       bool
	CreatedAt      *RFC3339DateTime
	Email          string
	FirstIP        string `json:"FirstIp"`
	Firstname      string
	LastIP         string `json:"LastIp"`
	LastLoginAt    *RFC3339DateTime
	Lastname       string
	NewPasswordKey string
	Password       string
}

type Destination struct {
	Domain         string
	IsNeverBlocked bool
}

type SpamAssassinRule struct {
	HitCount int64
	Name     string
	Score    float64
}

type MessageTemplate struct {
	Category    string
	ContentHtml string
	ContentText string
	DefaultType string
	Footer      string
	FromAddr    string
	FromName    string
	Name        string
	Locale      string
	ReplyTo     string
	Subject     string
	Subtitle    string
	Template    string
	Title       string
}

type KeyValueList []map[string]string

type JSONObject interface{}

type ContactsListAction struct {
	ListID int64
	Action string
}

type AddContactAction struct {
	Email      string
	Name       string
	Properties JSONObject
}

type Recipient struct {
	Email string
	Name  string
}

type RFC3339DateTime struct {
	time.Time
}

func (dt *RFC3339DateTime) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, `" `)
	if b == nil {
		return nil
	}

	dt.Time, err = time.Parse(time.RFC3339, string(b))

	return err
}

func (dt *RFC3339DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dt.Format(`"` + time.RFC3339 + `"`)), nil
}
