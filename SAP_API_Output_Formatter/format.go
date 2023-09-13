package sap_api_output_formatter

import (
	"encoding/json"
	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
	"sap-api-integrations-production-order-conf-creates-rmq-kube/SAP_API_Caller/responses"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}

	header := make([]Header, 0, 10)
	data := pm.D
	header = append(header, Header{
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
	})

	return header, nil
}
