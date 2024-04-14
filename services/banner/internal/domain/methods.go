package domain

import (
	"avito-go/pkg/xstringset"
	"github.com/google/uuid"
)

func NewActor(ID uuid.UUID, roles []string) Actor {
	a := Actor{ID: ID,
		roles: xstringset.New()}
	a.initRoles(roles)
	return a
}

func (c *Actor) HasRole(role string) bool {
	return c.roles.Contains(role)
}

func (c *Actor) HasOneOfRoles(roles ...string) bool {
	if roles == nil || len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if c.roles.Contains(role) {
			return true
		}
	}
	return false
}

func (c *Actor) HasAllRoles(roles ...string) bool {
	if roles == nil || len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if !c.roles.Contains(role) {
			return false
		}
	}
	return true
}

func (c *Actor) initRoles(roles []string) {
	c.roles.AddItems(roles)
}

func (c *Actor) GetRoles() []string {
	// TODO - does it necessary to return make([]string, 0)
	if c == nil || c.roles == nil {
		return make([]string, 0)
	}
	copyRoles := make([]string, 0, c.roles.Size())
	for r, _ := range c.roles {
		copyRoles = append(copyRoles, r)
	}
	return copyRoles
}

func (c *BannerFilter) Validate() error {
	//xapp.NewError(http.StatusBadRequest, " ", " ", nil)
	return nil
}
