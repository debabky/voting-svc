package cookies

import (
	"github.com/debabky/voting-svc/internal/config"
	"net/http"

	"github.com/debabky/voting-svc/internal/service/jwt"
)

type Cookies struct {
	Domain string
	Secure bool
}

func NewCookies(cfg *config.CookiesConfig) *Cookies {
	return &Cookies{
		Domain: cfg.Domain,
		Secure: cfg.Secure,
	}
}

func (c *Cookies) SetTokensCookies(w http.ResponseWriter, access, refresh string) {
	refreshCookie := &http.Cookie{
		Name:     jwt.RefreshTokenType.String(),
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		Secure:   c.Secure,
		SameSite: http.SameSiteLaxMode,
		Domain:   c.Domain,
	}

	http.SetCookie(w, refreshCookie)

	accessCookie := &http.Cookie{
		Name:     jwt.AccessTokenType.String(),
		Value:    access,
		Path:     "/",
		HttpOnly: true,
		Secure:   c.Secure,
		SameSite: http.SameSiteLaxMode,
		Domain:   c.Domain,
	}

	http.SetCookie(w, accessCookie)
}

func (c *Cookies) ClearTokensCookies(w http.ResponseWriter) {
	refreshCookie := &http.Cookie{
		Name:     jwt.RefreshTokenType.String(),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   c.Secure,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Domain:   c.Domain,
	}

	http.SetCookie(w, refreshCookie)

	accessCookie := &http.Cookie{
		Name:     jwt.AccessTokenType.String(),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   c.Secure,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Domain:   c.Domain,
	}

	http.SetCookie(w, accessCookie)
}
