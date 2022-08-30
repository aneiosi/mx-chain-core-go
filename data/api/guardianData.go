package api

// Guardian holds the relevant information for an account guardian
type Guardian struct {
	Address string `json:"address"`
	Epoch   uint32 `json:"epoch"`
}

// GuardianData holds data relating to the configured guardian(s) and frozen state of an account
type GuardianData struct {
	ActiveGuardian  *Guardian `json:"activeGuardian,omitempty"`
	PendingGuardian *Guardian `json:"pendingGuardian,omitempty"`
	Frozen          bool      `json:"frozen,omitempty"`
}
