package oneandone

import (
	"net/http"
	"time"
)

type BlockStorageRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Size           *int   `json:"size"`
	ServerId       string `json:"server,omitempty"`
	DatacenterId   string `json:"datacenter_id,omitempty"`
	ExecutionGroup string `json:"execution_group,omitempty"`
}

type BlockStorage struct {
	Identity
	descField
	Size  int    `json:"size"`
	State string `json:"state,omitempty"`
	// Name         string              `json:"name,omitempty"`
	CreationDate time.Time           `json:"creation_date,omitempty"`
	Datacenter   *Datacenter         `json:"datacenter,omitempty"`
	Server       *BlockStorageServer `json:"server,omitempty"`
	DiskID       string              `json:"disk_id,omitemtpy"`
	UUID         string              `json:"uuid,omitemtpy"`
	ApiPtr
}

type BlockStorageServer struct {
	Id       string `json:"id,omitempty"`
	ServerId string `json:"server,omitempty"`
	Name     string `json:"name,omitempty"`
}

type UpdateBlockStorageRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (api *API) ListBlockStorages(args ...interface{}) ([]BlockStorage, error) {
	url, err := processQueryParams(createUrl(api, blockStoragePathSegment), args...)
	if err != nil {
		return nil, err
	}
	result := []BlockStorage{}
	err = api.Client.Get(url, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	for index := range result {
		result[index].api = api
	}
	return result, nil
}

func (api *API) GetBlockStorage(id string) (*BlockStorage, error) {
	result := new(BlockStorage)
	url := createUrl(api, blockStoragePathSegment, id)
	err := api.Client.Get(url, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

func (api *API) GetBlockStorageServer(id string) (*BlockStorageServer, error) {
	result := new(BlockStorageServer)
	url := createUrl(api, blockStoragePathSegment, id, "server")
	err := api.Client.Get(url, &result, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api *API) AddBlockStorageServer(blockStorageId string, serverId string) (*BlockStorage, error) {
	result := new(BlockStorage)
	req := BlockStorageServer{ServerId: serverId}
	url := createUrl(api, blockStoragePathSegment, blockStorageId, "server")
	err := api.Client.Post(url, &req, &result, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

func (api *API) RemoveBlockStorageServer(blockStorageId string, serverId string) (*BlockStorage, error) {
	result := new(BlockStorage)
	blockStorage, err := api.GetBlockStorage(blockStorageId)

	if err != nil {
		return nil, err
	}

	req := BlockStorageServer{ServerId: serverId}
	url := createUrl(api, blockStoragePathSegment, blockStorage.Id, "server")

	err = api.Client.Delete(url, &req, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api

	return result, nil
}

func (api *API) CreateBlockStorage(request *BlockStorageRequest) (string, *BlockStorage, error) {
	result := new(BlockStorage)
	url := createUrl(api, blockStoragePathSegment)
	err := api.Client.Post(url, request, &result, http.StatusCreated)

	if err != nil {
		return "", nil, err
	}
	result.api = api
	return result.Id, result, nil
}

func (bs *BlockStorage) GetState() (string, error) {
	in, err := bs.api.GetBlockStorage(bs.Id)
	if in == nil {
		return "", err
	}
	return in.State, err
}

func (api *API) DeleteBlockStorage(id string) (*BlockStorage, error) {
	result := new(BlockStorage)
	url := createUrl(api, blockStoragePathSegment, id)
	err := api.Client.Delete(url, nil, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

func (api *API) UpdateBlockStorage(id string, request *UpdateBlockStorageRequest) (*BlockStorage, error) {
	result := new(BlockStorage)
	url := createUrl(api, blockStoragePathSegment, id)
	err := api.Client.Put(url, &request, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}
