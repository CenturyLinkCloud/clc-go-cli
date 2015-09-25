package billing

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type GetInvoiceData struct {
	Year                string `valid:"required" URIParam:"yes"`
	Month               string `valid:"required" URIParam:"yes"`
	PricingAccountAlias string `URIParam:"yes"`
}

func (g *GetInvoiceData) Validate() error {
	match, err := regexp.Match("^[[:digit:]]{4}$", []byte(g.Year))
	if err != nil || !match {
		return fmt.Errorf("Year must be specified in YYYY format")
	}

	yetValid := govalidator.IsInt(g.Month)
	if yetValid {
		i, err := govalidator.ToInt(g.Month)
		if err != nil || i < 1 || i > 12 {
			yetValid = false
		}
	}
	if !yetValid {
		return fmt.Errorf("Month must be an integer between 1 and 12")
	}
	return nil
}
