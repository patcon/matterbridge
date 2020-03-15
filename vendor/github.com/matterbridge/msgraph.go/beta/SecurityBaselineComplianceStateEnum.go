// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecurityBaselineComplianceState undocumented
type SecurityBaselineComplianceState int

const (
	// SecurityBaselineComplianceStateVUnknown undocumented
	SecurityBaselineComplianceStateVUnknown SecurityBaselineComplianceState = 0
	// SecurityBaselineComplianceStateVSecure undocumented
	SecurityBaselineComplianceStateVSecure SecurityBaselineComplianceState = 1
	// SecurityBaselineComplianceStateVNotApplicable undocumented
	SecurityBaselineComplianceStateVNotApplicable SecurityBaselineComplianceState = 2
	// SecurityBaselineComplianceStateVNotSecure undocumented
	SecurityBaselineComplianceStateVNotSecure SecurityBaselineComplianceState = 3
	// SecurityBaselineComplianceStateVError undocumented
	SecurityBaselineComplianceStateVError SecurityBaselineComplianceState = 4
	// SecurityBaselineComplianceStateVConflict undocumented
	SecurityBaselineComplianceStateVConflict SecurityBaselineComplianceState = 5
)

// SecurityBaselineComplianceStatePUnknown returns a pointer to SecurityBaselineComplianceStateVUnknown
func SecurityBaselineComplianceStatePUnknown() *SecurityBaselineComplianceState {
	v := SecurityBaselineComplianceStateVUnknown
	return &v
}

// SecurityBaselineComplianceStatePSecure returns a pointer to SecurityBaselineComplianceStateVSecure
func SecurityBaselineComplianceStatePSecure() *SecurityBaselineComplianceState {
	v := SecurityBaselineComplianceStateVSecure
	return &v
}

// SecurityBaselineComplianceStatePNotApplicable returns a pointer to SecurityBaselineComplianceStateVNotApplicable
func SecurityBaselineComplianceStatePNotApplicable() *SecurityBaselineComplianceState {
	v := SecurityBaselineComplianceStateVNotApplicable
	return &v
}

// SecurityBaselineComplianceStatePNotSecure returns a pointer to SecurityBaselineComplianceStateVNotSecure
func SecurityBaselineComplianceStatePNotSecure() *SecurityBaselineComplianceState {
	v := SecurityBaselineComplianceStateVNotSecure
	return &v
}

// SecurityBaselineComplianceStatePError returns a pointer to SecurityBaselineComplianceStateVError
func SecurityBaselineComplianceStatePError() *SecurityBaselineComplianceState {
	v := SecurityBaselineComplianceStateVError
	return &v
}

// SecurityBaselineComplianceStatePConflict returns a pointer to SecurityBaselineComplianceStateVConflict
func SecurityBaselineComplianceStatePConflict() *SecurityBaselineComplianceState {
	v := SecurityBaselineComplianceStateVConflict
	return &v
}