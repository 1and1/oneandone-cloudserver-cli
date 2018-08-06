package oneandone

import (
	"net/http"
)

type SSHKeyRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
}

type SSHKey struct {
	Identity
	descField
	State        string       `json:"state,omitempty"`
	Servers      *[]SSHServer `json:"servers,omitempty"`
	Md5          string       `json:"md5,omitempty"`
	PublicKey    string       `json:"public_key,omitempty"`
	CreationDate string       `json:"creation_date,omitempty"`
	ApiPtr
}

type SSHServer struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (api *API) ListSSHKeys(args ...interface{}) ([]SSHKey, error) {
	url, err := processQueryParams(createUrl(api, sshkeyPathSegment), args...)
	if err != nil {
		return nil, err
	}
	result := []SSHKey{}
	err = api.Client.Get(url, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	for index := range result {
		result[index].api = api
	}
	return result, nil
}

func (api *API) GetSSHKey(id string) (*SSHKey, error) {
	result := new(SSHKey)
	url := createUrl(api, sshkeyPathSegment, id)
	err := api.Client.Get(url, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

func (api *API) CreateSSHKey(request *SSHKeyRequest) (string, *SSHKey, error) {
	result := new(SSHKey)
	url := createUrl(api, sshkeyPathSegment)
	err := api.Client.Post(url, request, &result, http.StatusCreated)

	if err != nil {
		return "", nil, err
	}
	result.api = api
	return result.Id, result, nil
}

func (api *API) DeleteSSHKey(id string) (*SSHKey, error) {
	result := new(SSHKey)
	url := createUrl(api, sshkeyPathSegment, id)
	err := api.Client.Delete(url, nil, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

func (api *API) RenameSSHKey(id string, new_name string, new_desc string) (*SSHKey, error) {
	data := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{Name: new_name, Description: new_desc}
	result := new(SSHKey)
	url := createUrl(api, sshkeyPathSegment, id)
	err := api.Client.Put(url, &data, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}
