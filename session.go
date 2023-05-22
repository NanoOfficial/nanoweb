//
//
// @filename: session.go
// COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package nanoweb

type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	SessionID() string
	isExpired() bool
}
