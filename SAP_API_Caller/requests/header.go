package requests

type Header struct {
	OrderID                   *string `json:"OrderID"`
	Sequence                  *string `json:"Sequence"`
	OrderType                 *string `json:"OrderType"`
	ConfirmationText          *string `json:"ConfirmationText"`
	ConfirmationYieldQuantity *string `json:"ConfirmationYieldQuantity"`
	ConfirmationScrapQuantity *string `json:"ConfirmationScrapQuantity"`
	ConfirmationGroup         *string `json:"ConfirmationGroup"`
	Language                  *string `json:"Language"`
	Material                  *string `json:"Material"`
	Plant                     *string `json:"Plant"`
	WorkCenter                *string `json:"WorkCenter"`
}
