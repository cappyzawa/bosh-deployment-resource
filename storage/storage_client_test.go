package storage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry/bosh-deployment-resource/concourse"
	"github.com/cloudfoundry/bosh-deployment-resource/storage"
	"github.com/cloudfoundry/bosh-deployment-resource/gcp"
)

var _ = Describe("StorageClient", func() {
	Describe("NewStorageClient", func() {
		Context("when asking for a GCS client", func() {
			It("returns a GCS client", func() {
				source := concourse.Source{
					VarsStore: concourse.VarsStore{
						Provider: "gcs",
						Config: []byte("{\"json_key\": \"{\\\"type\\\": \\\"service_account\\\"}\", \"file_name\": \"bar\", \"bucket\": \"baz\"}"),
					},
				}

				storageClient, err := storage.NewStorageClient(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(storageClient).To(BeAssignableToTypeOf(gcp.Storage{}))
			})
		})

		Context("otherwise", func() {
			It("returns nil", func() {
				source := concourse.Source{
					VarsStore: concourse.VarsStore{
						Provider: "unknown",
						Config: []byte{},
					},
				}

				storageClient, err := storage.NewStorageClient(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(storageClient).To(BeNil())
			})
		})
	})
})