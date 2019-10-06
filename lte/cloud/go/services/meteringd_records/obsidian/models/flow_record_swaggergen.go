// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FlowRecord flow record
// swagger:model flow_record
type FlowRecord struct {

	// bytes rx
	// Required: true
	BytesRx *uint64 `json:"bytes_rx"`

	// bytes tx
	// Required: true
	BytesTx *uint64 `json:"bytes_tx"`

	// pkts rx
	// Required: true
	PktsRx *uint64 `json:"pkts_rx"`

	// pkts tx
	// Required: true
	PktsTx *uint64 `json:"pkts_tx"`

	// subscriber id
	// Required: true
	SubscriberID string `json:"subscriber_id"`
}

// Validate validates this flow record
func (m *FlowRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesRx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBytesTx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePktsRx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePktsTx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowRecord) validateBytesRx(formats strfmt.Registry) error {

	if err := validate.Required("bytes_rx", "body", m.BytesRx); err != nil {
		return err
	}

	return nil
}

func (m *FlowRecord) validateBytesTx(formats strfmt.Registry) error {

	if err := validate.Required("bytes_tx", "body", m.BytesTx); err != nil {
		return err
	}

	return nil
}

func (m *FlowRecord) validatePktsRx(formats strfmt.Registry) error {

	if err := validate.Required("pkts_rx", "body", m.PktsRx); err != nil {
		return err
	}

	return nil
}

func (m *FlowRecord) validatePktsTx(formats strfmt.Registry) error {

	if err := validate.Required("pkts_tx", "body", m.PktsTx); err != nil {
		return err
	}

	return nil
}

func (m *FlowRecord) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.RequiredString("subscriber_id", "body", string(m.SubscriberID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowRecord) UnmarshalBinary(b []byte) error {
	var res FlowRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
