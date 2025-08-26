package storage

type Url struct {
	ID string `json:"id"`
	OldUrl string `json:"old_url"`
	NewUrl string `json:"new_url"`
	CreatedAt string `json:"created_at"`
	ExpiredAt string `json:"expired_at"`
}