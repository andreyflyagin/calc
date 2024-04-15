package conversion

import "calc/internal/model"

func fromConversionToRepo(c *model.Conversion) Conversion {
	return Conversion{
		Base:   c.Base,
		Amount: c.Amount,
		Quote:  c.Quote,
		Value:  c.Value,
		Rate:   c.Rate,
	}
}
