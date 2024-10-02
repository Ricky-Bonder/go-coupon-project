package entity

type Basket struct {
	Value                 int  `json:"value"`
	AppliedDiscount       int  `json:"applied_discount"`
	ApplicationSuccessful bool `json:"application_successful"`
}
