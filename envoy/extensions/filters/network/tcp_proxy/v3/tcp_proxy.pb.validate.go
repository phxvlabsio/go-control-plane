// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/tcp_proxy/v3/tcp_proxy.proto

package tcp_proxyv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TcpProxy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TcpProxy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpProxy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TcpProxyMultiError, or nil
// if none found.
func (m *TcpProxy) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpProxy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := TcpProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOnDemand()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "OnDemand",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "OnDemand",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnDemand()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				field:  "OnDemand",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadataMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetIdleTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDownstreamIdleTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "DownstreamIdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "DownstreamIdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDownstreamIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				field:  "DownstreamIdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpstreamIdleTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "UpstreamIdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "UpstreamIdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpstreamIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				field:  "UpstreamIdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TcpProxyValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TcpProxyValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxyValidationError{
					field:  fmt.Sprintf("AccessLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if wrapper := m.GetMaxConnectAttempts(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			err := TcpProxyValidationError{
				field:  "MaxConnectAttempts",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(m.GetHashPolicy()) > 1 {
		err := TcpProxyValidationError{
			field:  "HashPolicy",
			reason: "value must contain no more than 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetHashPolicy() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TcpProxyValidationError{
						field:  fmt.Sprintf("HashPolicy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TcpProxyValidationError{
						field:  fmt.Sprintf("HashPolicy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxyValidationError{
					field:  fmt.Sprintf("HashPolicy[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetTunnelingConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "TunnelingConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxyValidationError{
					field:  "TunnelingConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTunnelingConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxyValidationError{
				field:  "TunnelingConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetMaxDownstreamConnectionDuration(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = TcpProxyValidationError{
				field:  "MaxDownstreamConnectionDuration",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

			if dur < gte {
				err := TcpProxyValidationError{
					field:  "MaxDownstreamConnectionDuration",
					reason: "value must be greater than or equal to 1ms",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if d := m.GetAccessLogFlushInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = TcpProxyValidationError{
				field:  "AccessLogFlushInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

			if dur < gte {
				err := TcpProxyValidationError{
					field:  "AccessLogFlushInterval",
					reason: "value must be greater than or equal to 1ms",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	switch m.ClusterSpecifier.(type) {

	case *TcpProxy_Cluster:
		// no validation rules for Cluster

	case *TcpProxy_WeightedClusters:

		if all {
			switch v := interface{}(m.GetWeightedClusters()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TcpProxyValidationError{
						field:  "WeightedClusters",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TcpProxyValidationError{
						field:  "WeightedClusters",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetWeightedClusters()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxyValidationError{
					field:  "WeightedClusters",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := TcpProxyValidationError{
			field:  "ClusterSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return TcpProxyMultiError(errors)
	}

	return nil
}

// TcpProxyMultiError is an error wrapping multiple validation errors returned
// by TcpProxy.ValidateAll() if the designated constraints aren't met.
type TcpProxyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpProxyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpProxyMultiError) AllErrors() []error { return m }

// TcpProxyValidationError is the validation error returned by
// TcpProxy.Validate if the designated constraints aren't met.
type TcpProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpProxyValidationError) ErrorName() string { return "TcpProxyValidationError" }

// Error satisfies the builtin error interface
func (e TcpProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpProxyValidationError{}

// Validate checks the field values on TcpProxy_WeightedCluster with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TcpProxy_WeightedCluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpProxy_WeightedCluster with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TcpProxy_WeightedClusterMultiError, or nil if none found.
func (m *TcpProxy_WeightedCluster) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpProxy_WeightedCluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetClusters()) < 1 {
		err := TcpProxy_WeightedClusterValidationError{
			field:  "Clusters",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetClusters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TcpProxy_WeightedClusterValidationError{
						field:  fmt.Sprintf("Clusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TcpProxy_WeightedClusterValidationError{
						field:  fmt.Sprintf("Clusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxy_WeightedClusterValidationError{
					field:  fmt.Sprintf("Clusters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TcpProxy_WeightedClusterMultiError(errors)
	}

	return nil
}

// TcpProxy_WeightedClusterMultiError is an error wrapping multiple validation
// errors returned by TcpProxy_WeightedCluster.ValidateAll() if the designated
// constraints aren't met.
type TcpProxy_WeightedClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpProxy_WeightedClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpProxy_WeightedClusterMultiError) AllErrors() []error { return m }

// TcpProxy_WeightedClusterValidationError is the validation error returned by
// TcpProxy_WeightedCluster.Validate if the designated constraints aren't met.
type TcpProxy_WeightedClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpProxy_WeightedClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpProxy_WeightedClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpProxy_WeightedClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpProxy_WeightedClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpProxy_WeightedClusterValidationError) ErrorName() string {
	return "TcpProxy_WeightedClusterValidationError"
}

// Error satisfies the builtin error interface
func (e TcpProxy_WeightedClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_WeightedCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpProxy_WeightedClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpProxy_WeightedClusterValidationError{}

// Validate checks the field values on TcpProxy_TunnelingConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TcpProxy_TunnelingConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpProxy_TunnelingConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TcpProxy_TunnelingConfigMultiError, or nil if none found.
func (m *TcpProxy_TunnelingConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpProxy_TunnelingConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetHostname()) < 1 {
		err := TcpProxy_TunnelingConfigValidationError{
			field:  "Hostname",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for UsePost

	if len(m.GetHeadersToAdd()) > 1000 {
		err := TcpProxy_TunnelingConfigValidationError{
			field:  "HeadersToAdd",
			reason: "value must contain no more than 1000 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetHeadersToAdd() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TcpProxy_TunnelingConfigValidationError{
						field:  fmt.Sprintf("HeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TcpProxy_TunnelingConfigValidationError{
						field:  fmt.Sprintf("HeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpProxy_TunnelingConfigValidationError{
					field:  fmt.Sprintf("HeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PropagateResponseHeaders

	if len(errors) > 0 {
		return TcpProxy_TunnelingConfigMultiError(errors)
	}

	return nil
}

// TcpProxy_TunnelingConfigMultiError is an error wrapping multiple validation
// errors returned by TcpProxy_TunnelingConfig.ValidateAll() if the designated
// constraints aren't met.
type TcpProxy_TunnelingConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpProxy_TunnelingConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpProxy_TunnelingConfigMultiError) AllErrors() []error { return m }

// TcpProxy_TunnelingConfigValidationError is the validation error returned by
// TcpProxy_TunnelingConfig.Validate if the designated constraints aren't met.
type TcpProxy_TunnelingConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpProxy_TunnelingConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpProxy_TunnelingConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpProxy_TunnelingConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpProxy_TunnelingConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpProxy_TunnelingConfigValidationError) ErrorName() string {
	return "TcpProxy_TunnelingConfigValidationError"
}

// Error satisfies the builtin error interface
func (e TcpProxy_TunnelingConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_TunnelingConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpProxy_TunnelingConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpProxy_TunnelingConfigValidationError{}

// Validate checks the field values on TcpProxy_OnDemand with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TcpProxy_OnDemand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpProxy_OnDemand with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TcpProxy_OnDemandMultiError, or nil if none found.
func (m *TcpProxy_OnDemand) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpProxy_OnDemand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOdcdsConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxy_OnDemandValidationError{
					field:  "OdcdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxy_OnDemandValidationError{
					field:  "OdcdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOdcdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxy_OnDemandValidationError{
				field:  "OdcdsConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ResourcesLocator

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxy_OnDemandValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxy_OnDemandValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxy_OnDemandValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TcpProxy_OnDemandMultiError(errors)
	}

	return nil
}

// TcpProxy_OnDemandMultiError is an error wrapping multiple validation errors
// returned by TcpProxy_OnDemand.ValidateAll() if the designated constraints
// aren't met.
type TcpProxy_OnDemandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpProxy_OnDemandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpProxy_OnDemandMultiError) AllErrors() []error { return m }

// TcpProxy_OnDemandValidationError is the validation error returned by
// TcpProxy_OnDemand.Validate if the designated constraints aren't met.
type TcpProxy_OnDemandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpProxy_OnDemandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpProxy_OnDemandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpProxy_OnDemandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpProxy_OnDemandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpProxy_OnDemandValidationError) ErrorName() string {
	return "TcpProxy_OnDemandValidationError"
}

// Error satisfies the builtin error interface
func (e TcpProxy_OnDemandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_OnDemand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpProxy_OnDemandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpProxy_OnDemandValidationError{}

// Validate checks the field values on TcpProxy_WeightedCluster_ClusterWeight
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *TcpProxy_WeightedCluster_ClusterWeight) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// TcpProxy_WeightedCluster_ClusterWeight with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// TcpProxy_WeightedCluster_ClusterWeightMultiError, or nil if none found.
func (m *TcpProxy_WeightedCluster_ClusterWeight) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := TcpProxy_WeightedCluster_ClusterWeightValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetWeight() < 1 {
		err := TcpProxy_WeightedCluster_ClusterWeightValidationError{
			field:  "Weight",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetadataMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpProxy_WeightedCluster_ClusterWeightValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpProxy_WeightedCluster_ClusterWeightValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpProxy_WeightedCluster_ClusterWeightValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TcpProxy_WeightedCluster_ClusterWeightMultiError(errors)
	}

	return nil
}

// TcpProxy_WeightedCluster_ClusterWeightMultiError is an error wrapping
// multiple validation errors returned by
// TcpProxy_WeightedCluster_ClusterWeight.ValidateAll() if the designated
// constraints aren't met.
type TcpProxy_WeightedCluster_ClusterWeightMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpProxy_WeightedCluster_ClusterWeightMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpProxy_WeightedCluster_ClusterWeightMultiError) AllErrors() []error { return m }

// TcpProxy_WeightedCluster_ClusterWeightValidationError is the validation
// error returned by TcpProxy_WeightedCluster_ClusterWeight.Validate if the
// designated constraints aren't met.
type TcpProxy_WeightedCluster_ClusterWeightValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) ErrorName() string {
	return "TcpProxy_WeightedCluster_ClusterWeightValidationError"
}

// Error satisfies the builtin error interface
func (e TcpProxy_WeightedCluster_ClusterWeightValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProxy_WeightedCluster_ClusterWeight.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpProxy_WeightedCluster_ClusterWeightValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpProxy_WeightedCluster_ClusterWeightValidationError{}
