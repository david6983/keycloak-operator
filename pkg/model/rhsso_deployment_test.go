package model

import (
	"testing"
)

func TestRHSSODeployment_testExperimentalEnvs(t *testing.T) {
	testExperimentalEnvs(t, RHSSODeployment)
}

func TestRHSSODeployment_testExperimentalArgs(t *testing.T) {
	testExperimentalArgs(t, RHSSODeployment)
}

func TestRHSSODeployment_testExperimentalCommand(t *testing.T) {
	testExperimentalCommand(t, RHSSODeployment)
}

func TestRHSSODeployment_testExperimentalVolumesWithConfigMaps(t *testing.T) {
	testExperimentalVolumesWithConfigMaps(t, RHSSODeployment)
}

func TestRHSSODeployment_testAffinityDefaultMultiAZ(t *testing.T) {
	testAffinityDefaultMultiAZ(t, RHSSODeployment)
}

func TestRHSSODeployment_testTolerationsExperimental(t *testing.T) {
	testTolerationsExperimentalTolerationsSet(t, RHSSODeployment)
}

func TestRHSSODeployment_testNodeSelectorExperimental(t *testing.T) {
	testNodeSelectorExperimentalNodeSelectorSet(t, RHSSODeployment)
}

func TestRHSSODeployment_testAffinityExperimental(t *testing.T) {
	testAffinityExperimentalAffinitySet(t, RHSSODeployment)
}

func TestRHSSODeployment_testDeploymentSpecImagePolicy(t *testing.T) {
	testDeploymentSpecImagePolicy(t, RHSSODeployment)
}

func TestRHSSODeploymentReconciled_testDisableRHSSOReplicasSyncingFalse(t *testing.T) {
	testDisableDeploymentReplicasSyncingFalse(t, RHSSODeployment, RHSSODeploymentReconciled)
}

func TestRHSSODeploymentReconciled_testDisableRHSSOReplicasSyncingTrue(t *testing.T) {
	testDisableDeploymentReplicasSyncingTrue(t, RHSSODeployment, RHSSODeploymentReconciled)
}

func TestRHSSODeployment_testServiceAccountSetExperimental(t *testing.T) {
	testServiceAccountSet(t, RHSSODeployment)
}

func TestRHSSODeployment_testServiceAccountReconciledSetExperimental(t *testing.T) {
	testServiceAccountReconciledSet(t, RHSSODeployment, RHSSODeploymentReconciled)
}
