package vault

import (
	"strings"
)

type VaultHoldings struct {
	Holdings []VaultItem
}

type VaultItem struct {
	Name         string
	Type         string
	Public       bool
	Project      string
	RelativePath string
	FileName     string
	ParentDir    string
	FileExt      string
	TypeURI      string
	Age          float64
	DateCreated  string
}

// Prjs returns all unique projets in
// VaultHoldings as an slice of strings
func (v *VaultHoldings) Prjs() []string {
	m := make(map[string]bool)
	var sa []string

	for _, item := range v.Holdings {
		p := item.Project
		if _, ok := m[p]; !ok {
			m[p] = true
		}
	}

	// WHY DO THIS?  just return the map?
	for k := range m {
		sa = append(sa, k)
	}

	return sa
}

// PrjFiles returns an array string of all files
// for a given project
// thought:  return [][]string and get the items to
// loop on now?   or return a slice of the stuct
// and return a VaultHoldings ??
func (v *VaultHoldings) PrjFiles(t string) VaultHoldings {
	var pi []VaultItem

	for _, item := range v.Holdings {
		p := item.Project
		// This line is an ERROR....     if strings.Contains(p, t) { // I could also skip dot files here too..  rather than in main..
		if strings.Compare(p, t) == 0 { // I could also skip dot files here too..  rather than in main..
			pi = append(pi, item)
		}
	}

	return VaultHoldings{pi}
}
