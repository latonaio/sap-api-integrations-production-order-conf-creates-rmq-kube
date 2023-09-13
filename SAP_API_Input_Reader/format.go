package sap_api_input_reader

import (
	"sap-api-integrations-production-order-conf-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header

	return &requests.Header{
		OrderID:                   data.OrderID,
		Sequence:                  data.Sequence,
		OrderType:                 data.OrderType,
		ConfirmationText:          data.ConfirmationText,
		ConfirmationYieldQuantity: data.ConfirmationYieldQuantity,
		ConfirmationScrapQuantity: data.ConfirmationScrapQuantity,
		ConfirmationGroup:         data.ConfirmationGroup,
		Language:                  data.Language,
		Material:                  data.Material,
		Plant:                     data.Plant,
		WorkCenter:                data.WorkCenter,
	}
}
