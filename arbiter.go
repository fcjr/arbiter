package arbiter

type FirewallParams struct{}

type Firewall struct{}

func NewFirewall(params *FirewallParams) (*Firewall, error) {
	return &Firewall{}, nil
}
