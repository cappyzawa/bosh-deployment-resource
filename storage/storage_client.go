package storage

import (
	"github.com/cloudfoundry/bosh-deployment-resource/concourse"
	"encoding/json"
	"github.com/cloudfoundry/bosh-deployment-resource/gcp"
)

type GCSConfig struct {
	FileName string `json:"file_name"`
	Bucket   string `json:"bucket"`
	JSONKey  string `json:"json_key"`
}

type StorageClient interface {
	Download(filePath string) error
	Upload(filePath string)   error
}

func NewStorageClient(source concourse.Source) (StorageClient, error) {
	if source.VarsStore.Provider == "gcs" {
		gcsConfig := GCSConfig{}
		json.Unmarshal(source.VarsStore.Config, &gcsConfig)
		return gcp.NewStorage(gcsConfig.JSONKey, gcsConfig.Bucket, gcsConfig.FileName)
	}

	return nil, nil
}