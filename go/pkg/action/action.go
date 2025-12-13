package action

import (
	"fmt"
	"strings"

	"github.com/fanovilla/iam-simulate-go/internal/util"
)

// Name represents an AWS action in the form service:Operation.
type Name struct {
	Raw       string
	Service   string
	Operation string
}

// Parse parses an action string like "s3:PutObject" or with wildcards like "s3:*" or "*".
func Parse(s string) (Name, error) {
	n := Name{Raw: s}
	parts := strings.SplitN(s, ":", 2)
	if len(parts) == 1 {
		// Treat single token as service wildcard or global '*'
		n.Service = parts[0]
		n.Operation = "*"
		return n, nil
	}
	if parts[0] == "" || parts[1] == "" {
		return n, fmt.Errorf("invalid action: %q", s)
	}
	n.Service = strings.ToLower(parts[0])
	n.Operation = parts[1]
	return n, nil
}

// Matches reports whether action a matches pattern p (which may include wildcards).
func (a Name) Matches(p Name) bool {
	// Service is case-insensitive in AWS semantics; canonicalized to lower for parsed values.
	return util.WildcardMatch(p.Service, a.Service) && util.WildcardMatch(p.Operation, a.Operation)
}
