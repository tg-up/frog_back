package responses

import "icecreambash/tgup_backend/internal/models"

type ServiceFieldsResponse struct {
	ID           uint         `json:"id"`
	ServiceID    uint         `json:"service_id"`
	FieldName    string       `json:"field_name"`
	FieldSlug    string       `json:"field_slug"`
	FieldType    string       `json:"field_type"`
	DefaultValue string       `json:"default_value"`
	MinValue     uint8        `json:"min_value"`
	MaxValue     uint8        `json:"max_value"`
	CostType     string       `json:"cost_type"`
	Options      models.JSONB `json:"options"`
}

func GetFieldsResponse(values []models.ServiceField) []ServiceFieldsResponse {
	var fields = make([]ServiceFieldsResponse, 0)

	for _, model := range values {
		fields = append(fields, ServiceFieldsResponse{
			ID:           model.ID,
			ServiceID:    model.ServiceID,
			FieldName:    model.FieldName,
			FieldSlug:    model.FieldSlug,
			FieldType:    model.FieldType,
			DefaultValue: model.DefaultValue,
			MinValue:     model.MinValue,
			CostType:     model.CostType,
			Options:      model.Options,
		})
	}

	return fields
}
