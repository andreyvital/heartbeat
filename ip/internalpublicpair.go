package ip

func NewInternalPublicPair(internal, public string) InternalPublicPair {
	return InternalPublicPair{
		Internal: internal,
		Public:   public,
	}
}

type InternalPublicPair struct {
	Internal string `json:"internal"`
	Public   string `json:"public,omitempty"`
}
