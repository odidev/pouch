// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PidsStats PidsStats contains the stats of a container's pids
// swagger:model PidsStats

type PidsStats struct {

	// Current is the number of pids in the cgroup
	Current uint64 `json:"current,omitempty"`

	// Limit is the hard limit on the number of pids in the cgroup.
	// A "Limit" of 0 means that there is no limit.
	//
	Limit uint64 `json:"limit,omitempty"`
}

/* polymorph PidsStats current false */

/* polymorph PidsStats limit false */

// Validate validates this pids stats
func (m *PidsStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PidsStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PidsStats) UnmarshalBinary(b []byte) error {
	var res PidsStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
