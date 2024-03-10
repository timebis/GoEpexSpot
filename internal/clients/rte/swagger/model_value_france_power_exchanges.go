/*
 * Wholesale Market
 *
 * Data related to the whole electricity market within EPEX Spot and Nord Pool hub Stock exchange operations.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type ValueFrancePowerExchanges struct {
	// Start time interval (YYYY-MM-DDThh:mm:sszzzzzz)
	StartDate time.Time `json:"start_date,omitempty"`
	// End time interval (YYYY-MM-DDThh:mm:sszzzzzz)
	EndDate time.Time `json:"end_date,omitempty"`
	// Volume of the electricity market (in MW)
	Value float32 `json:"value,omitempty"`
	// Price (in €/MWh)
	Price float32 `json:"price,omitempty"`
}
