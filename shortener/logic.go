package shortener

import (
	"time"
	"github.com/teris-io/shortid"
	"errors"
	validate "gopkg.in/dealancer/validate.v2"
	errs "github.com/pkg/errors"
)

var(
	ErrRedirectNotFound = errors.New("Redirect not found")
	ErrRedirectInvalid = errors.New("Redirect invalid")
)

type RedirectSerivce struct {
	redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectSerivce {
	return RedirectSerivce{redirectRepo: redirectRepo}
}

func (r *RedirectSerivce) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *RedirectSerivce) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		errs.Wrap(ErrRedirectInvalid, "service.Redirect")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().Unix()
	return r.redirectRepo.Store(redirect)
}