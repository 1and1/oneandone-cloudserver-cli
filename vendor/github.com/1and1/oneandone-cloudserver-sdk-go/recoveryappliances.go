package oneandone

import "net/http"

type RecoveryAppliance struct {
	Identity
	Os                   Os       `json:"os,omitempty"`
	AvailableDatacenters []string `json:"available_datacenters,omitempty"`
	ApiPtr
}

type SingleRecoveryAppliance struct {
	Identity
	Os                   string   `json:"os,omitempty"`
	OsFamily             string   `json:"os_family,omitempty"`
	OsVersion            string   `json:"os_version,omitempty"`
	AvailableDatacenters []string `json:"available_datacenters,omitempty"`
	ApiPtr
}

type Os struct {
	Architecture int    `json:"architecture,omitempty"`
	Family       string `json:"family,omitempty"`
	SubFamily    string `json:"subfamily,omitempty"`
	Name         string `json:"name,omitempty"`
}

// GET /recovery_appliances
func (api *API) ListRecoveryAppliances(args ...interface{}) ([]RecoveryAppliance, error) {
	url, err := processQueryParams(createUrl(api, recoveryAppliancePathSegment), args...)
	if err != nil {
		return nil, err
	}
	res := []RecoveryAppliance{}
	err = api.Client.Get(url, &res, http.StatusOK)
	if err != nil {
		return nil, err
	}
	for index, _ := range res {
		res[index].api = api
	}
	return res, nil
}

// GET /server_appliances/{id}
func (api *API) GetRecoveryAppliance(ra_id string) (*SingleRecoveryAppliance, error) {
	res := new(SingleRecoveryAppliance)
	url := createUrl(api, recoveryAppliancePathSegment, ra_id)
	err := api.Client.Get(url, &res, http.StatusOK)
	if err != nil {
		return nil, err
	}
	//	res.api = api
	return res, nil
}
