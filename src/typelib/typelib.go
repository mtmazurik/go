package typelib

import "time"

// Repository data object for storage in MongoDB (Atlas-hosted solution)
type Repository struct {
	ID           int        `json:"_id"`
	KeyName      string     `json:"keyName"`
	KeyValue     string     `json:"keyValue"`
	Tags         [10]string `json:"tags"`
	CreatedDate  time.Time  `json:"createdDate"`
	CreatedBy    string     `json:"createdBy"`
	ModifiedDate time.Time  `json:"modifiedDate"`
	ModifiedBy   string     `json:"modifiedBy"`
	App          string     `json:"app"`
	Repository   string     `json:"repository"`
	Collection   string     `json:"collection"`
	SchemaName   string     `json:"schemaName"`
	SchemaURI    string     `json:"schemaURI"`
	Body         string     `json:"body"`
}

// CRUD interface for creating different datastore concrete implementations (MongoDB, Firebase)
type CRUD interface {
	Create(r Repository)
	Read(keyName string, keyValue string) (repos []Repository)
	Update(r Repository)
	Delete(keyName string, keyValue string)
}
