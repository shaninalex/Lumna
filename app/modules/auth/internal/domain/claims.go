package domain

type AccessClaims struct {
	IdentityID int `json:"iid"`
}

func (c AccessClaims) Validate() error {
	if c.IdentityID <= 0 {
		return ErrAccessClaimsInvalid
	}
	return nil
}
