package dns

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type DNS struct {
	ManagedZones []ManagedZone
}

type ManagedZone struct {
	Metadata   defsecTypes.Metadata
	DNSSec     DNSSec
	Visibility defsecTypes.StringValue
}

func (m ManagedZone) IsPrivate() bool {
	return m.Visibility.EqualTo("private", defsecTypes.IgnoreCase)
}

type DNSSec struct {
	Metadata        defsecTypes.Metadata
	Enabled         defsecTypes.BoolValue
	DefaultKeySpecs KeySpecs
}

type KeySpecs struct {
	Metadata       defsecTypes.Metadata
	KeySigningKey  Key
	ZoneSigningKey Key
}

type Key struct {
	Metadata  defsecTypes.Metadata
	Algorithm defsecTypes.StringValue
}
