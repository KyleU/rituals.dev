package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
	"golang.org/x/oauth2/slack"
	"time"

	"emperror.dev/errors"

	"github.com/gofrs/uuid"
)

type Provider struct {
	Key      string
	Title    string
	Icon     string
	Endpoint oauth2.Endpoint
	Scopes   []string
}

type Providers []*Provider

func (p Providers) Names() []string {
	ret := make([]string, len(p))
	for i, prv := range p {
		ret[i] = prv.Title
	}
	return ret
}

var ProviderGitHub = Provider{Key: "github", Title: "GitHub", Icon: "github-alt", Endpoint: github.Endpoint, Scopes: githubScopes}
var ProviderGoogle = Provider{Key: "google", Title: "Google", Icon: "google", Endpoint: google.Endpoint, Scopes: googleScopes}
var ProviderSlack = Provider{Key: "slack", Title: "Slack", Icon: "hashtag", Endpoint: slack.Endpoint, Scopes: slackScopes}
var ProviderAmazon = Provider{Key: "amazon", Title: "Amazon", Icon: "cart", Endpoint: amazon.Endpoint, Scopes: amazonScopes}
var ProviderMicrosoft = Provider{Key: "microsoft", Title: "Microsoft", Icon: "world", Endpoint: microsoft.AzureADEndpoint(""), Scopes: microsoftScopes}

var AllProviders = Providers{&ProviderGitHub, &ProviderGoogle, &ProviderSlack, &ProviderAmazon, &ProviderMicrosoft}

func ProviderFromString(s string) *Provider {
	for _, t := range AllProviders {
		if t.Key == s {
			return t
		}
	}
	return &ProviderGitHub
}

type Display struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
}

type Displays []*Display

type Record struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Provider   *Provider
	ProviderID string
	Expires    *time.Time
	Name       string
	Email      string
	Picture    string
	Created    time.Time
}

type Records []*Record

func (r *Record) ToDisplay() *Display {
	return &Display{
		Provider: r.Provider.Key,
		Email:    r.Email,
	}
}

func (r Records) FindByProvider(key string) Records {
	var ret Records
	for _, e := range r {
		if e.Provider.Key == key {
			ret = append(ret, e)
		}
	}
	return ret
}

func (r Records) Emails() []string {
	ret := make([]string, len(r))
	for i, e := range r {
		ret[i] = e.Email
	}
	return ret
}

type recordDTO struct {
	ID         uuid.UUID  `db:"id"`
	UserID     uuid.UUID  `db:"user_id"`
	Provider   string     `db:"provider"`
	ProviderID string     `db:"provider_id"`
	Expires    *time.Time `db:"expires"`
	Name       string     `db:"name"`
	Email      string     `db:"email"`
	Picture    string     `db:"picture"`
	Created    time.Time  `db:"created"`
}

func (dto *recordDTO) ToRecord() *Record {
	return &Record{
		ID:         dto.ID,
		UserID:     dto.UserID,
		Provider:   ProviderFromString(dto.Provider),
		ProviderID: dto.ProviderID,
		Expires:    dto.Expires,
		Name:       dto.Name,
		Email:      dto.Email,
		Picture:    dto.Picture,
		Created:    dto.Created,
	}
}

var ErrorAuthDisabled = errors.New("authorization is disabled")
