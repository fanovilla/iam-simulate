package arn

import (
    "errors"
    "fmt"
    "strings"

    "github.com/cloud-copilot/iam-simulate-go/internal/util"
)

// ARN represents a parsed AWS ARN: arn:partition:service:region:account-id:resource
type ARN struct {
    Raw       string
    Partition string
    Service   string
    Region    string
    AccountID string
    Resource  string
}

var ErrInvalidARN = errors.New("invalid ARN")

// Parse parses an ARN string into an ARN struct.
func Parse(s string) (ARN, error) {
    a := ARN{Raw: s}
    if !strings.HasPrefix(s, "arn:") {
        return a, fmt.Errorf("%w: missing prefix", ErrInvalidARN)
    }
    parts := strings.SplitN(s, ":", 6)
    if len(parts) != 6 {
        return a, fmt.Errorf("%w: wrong number of segments", ErrInvalidARN)
    }
    a.Partition = parts[1]
    a.Service = parts[2]
    a.Region = parts[3]
    a.AccountID = parts[4]
    a.Resource = parts[5]
    return a, nil
}

// Matches reports whether arn a matches the pattern p (which may contain wildcards).
// This is a generic matcher applying wildcard matching per whole-field basis.
// Service-specific resource nuances will be layered later.
func (a ARN) Matches(p ARN) bool {
    return util.WildcardMatch(p.Partition, a.Partition) &&
        util.WildcardMatch(p.Service, a.Service) &&
        util.WildcardMatch(p.Region, a.Region) &&
        util.WildcardMatch(p.AccountID, a.AccountID) &&
        util.WildcardMatch(p.Resource, a.Resource)
}
