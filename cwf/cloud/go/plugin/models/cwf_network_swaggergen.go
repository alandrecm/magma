// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	models2 "magma/feg/cloud/go/plugin/models"
	models4 "magma/orc8r/cloud/go/models"
	models5 "magma/orc8r/cloud/go/pluginimpl/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CwfNetwork Carrier Wifi Network spec
// swagger:model cwf_network
type CwfNetwork struct {

	// carrier wifi
	// Required: true
	CarrierWifi *NetworkCarrierWifiConfigs `json:"carrier_wifi"`

	// description
	// Required: true
	Description models4.NetworkDescription `json:"description"`

	// dns
	// Required: true
	DNS *models5.NetworkDNSConfig `json:"dns"`

	// features
	Features *models5.NetworkFeatures `json:"features,omitempty"`

	// federation
	// Required: true
	Federation *models2.FederatedNetworkConfigs `json:"federation"`

	// id
	// Required: true
	ID models4.NetworkID `json:"id"`

	// name
	// Required: true
	Name models4.NetworkName `json:"name"`
}

// Validate validates this cwf network
func (m *CwfNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierWifi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CwfNetwork) validateCarrierWifi(formats strfmt.Registry) error {

	if err := validate.Required("carrier_wifi", "body", m.CarrierWifi); err != nil {
		return err
	}

	if m.CarrierWifi != nil {
		if err := m.CarrierWifi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("carrier_wifi")
			}
			return err
		}
	}

	return nil
}

func (m *CwfNetwork) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *CwfNetwork) validateDNS(formats strfmt.Registry) error {

	if err := validate.Required("dns", "body", m.DNS); err != nil {
		return err
	}

	if m.DNS != nil {
		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *CwfNetwork) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *CwfNetwork) validateFederation(formats strfmt.Registry) error {

	if err := validate.Required("federation", "body", m.Federation); err != nil {
		return err
	}

	if m.Federation != nil {
		if err := m.Federation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("federation")
			}
			return err
		}
	}

	return nil
}

func (m *CwfNetwork) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *CwfNetwork) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CwfNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CwfNetwork) UnmarshalBinary(b []byte) error {
	var res CwfNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
