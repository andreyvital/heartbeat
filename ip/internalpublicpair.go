package ip

// NewInternalPublicPair creates a new pair of internal and
// public (external) IP addresses
func NewInternalPublicPair(internal, public string) InternalPublicPair {
	return InternalPublicPair{
		Internal: internal,
		Public:   public,
	}
}

// InternalPublicPair represents a pair of both internal and
// public (external) IP addresses
type InternalPublicPair struct {
	Internal string `json:"internal"`
	Public   string `json:"public,omitempty"`
}
