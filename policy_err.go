package archp

import "fmt"

type PolicyError struct {
	TargetAnalized string
	Policy         string
	TriggerErr     string
}

func (pe *PolicyError) Error() string {
	return fmt.Sprintf("target analized %v failed on %v by %v", pe.TargetAnalized, pe.Policy, pe.TriggerErr)
}

func newPolicyError(target, policy, trigger string) *PolicyError {
	return &PolicyError{
		TargetAnalized: target,
		Policy:         policy,
		TriggerErr:     trigger,
	}
}
