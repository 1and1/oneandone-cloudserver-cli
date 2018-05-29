package oneandone

import "net/http"

type Pricing struct {
	Currency string       `json:"currency,omitempty"`
	Plan     *pricingPlan `json:"pricing_plans,omitempty"`
}

type pricingPlan struct {
	Image            *pricingItem   `json:"image,omitempty"`
	PublicIPs        []pricingItem  `json:"public_ips,omitempty"`
	Servers          *serverPricing `json:"servers,omitempty"`
	SharedStorage    *pricingItem   `json:"shared_storage,omitempty"`
	SoftwareLicenses []pricingItem  `json:"software_licences,omitempty"`
}

type serverPricing struct {
	FixedServers []pricingItem `json:"fixed_servers,omitempty"`
	FlexServers  []pricingItem `json:"flexible_server,omitempty"`
}

type pricingItem struct {
	Name       string  `json:"name,omitempty"`
	GrossPrice float64 `json:"price_gross"`
	NetPrice   float64 `json:"price_net"`
	Unit       string  `json:"unit,omitempty"`
}

// GET /pricing
func (api *API) GetPricing() (*Pricing, error) {
	result := new(Pricing)
	url := createUrl(api, pricingPathSegment)
	err := api.Client.Get(url, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return result, nil
}
