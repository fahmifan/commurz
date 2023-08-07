package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fahmifan/ulids"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	_sessionUserKey             = "user_session"
	_sessionUserMaxAge          = 3600 * 24
	_sessionUserMaxAgeImmidiate = 180
)

type Session struct {
	store sessions.Store
}

func NewSession(
	store sessions.Store,
) *Session {
	return &Session{
		store: store,
	}
}

// AuthUser ..
type AuthUser struct {
	UserID ulids.ULID
	// ExpiredAt at empty means use the provided maxAge
	ExpiredAt time.Time
}

func (sa *Session) SaveUser(c echo.Context, maxAge int, userID ulids.ULID) error {
	sess, err := sa.store.Get(c.Request(), _sessionUserKey)
	if err != nil {
		return nil
	}

	authUser := AuthUser{
		UserID:    userID,
		ExpiredAt: time.Now().Add(time.Second * _sessionUserMaxAge),
	}

	buf, err := json.Marshal(authUser)
	if err != nil {
		return fmt.Errorf("marshal user session: %w", err)
	}

	sess.Values[_sessionUserKey] = string(buf)

	sess.Options.MaxAge = _sessionUserMaxAge
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("unable to save session for user %s: %w", authUser.UserID, err)
	}

	return nil
}

func (sa *Session) GetUser(c echo.Context) *AuthUser {
	return sa.GetUserFromRequest(c.Request())
}

func (sa *Session) GetUserFromRequest(req *http.Request) *AuthUser {
	sess, err := sa.store.Get(req, _sessionUserKey)
	if err != nil {
		return nil
	}

	buf, ok := sess.Values[_sessionUserKey].(string)
	if !ok {
		return nil
	}

	authUser := AuthUser{}
	err = json.Unmarshal([]byte(buf), &authUser)
	if err != nil {
		return nil
	}

	return &authUser
}

// DeleteUser ..
func (sa *Session) DeleteUser(c echo.Context) error {
	sess, err := sa.store.Get(c.Request(), _sessionUserKey)
	if err != nil {
		return fmt.Errorf("unable to get session: %w", err)
	}

	sess.Options.MaxAge = -1
	return sess.Save(c.Request(), c.Response())
}
