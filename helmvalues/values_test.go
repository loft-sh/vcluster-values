package helmvalues

import (
	"testing"

	"github.com/loft-sh/vcluster-values/config"
	"gotest.tools/v3/assert"
)

var (
	TRUE  = true
	FALSE = false
)

func TestParseArgs(t *testing.T) {
	assert.DeepEqual(
		t,
		parseArgs([]string{
			"--some-bool-flag",
		}),
		map[string]string{
			"--some-bool-flag": "true",
		},
	)
	assert.DeepEqual(
		t,
		parseArgs([]string{
			"--out-kube-config-secret=vc-secret",
		}),
		map[string]string{
			"--out-kube-config-secret": "vc-secret",
		},
	)
	assert.DeepEqual(
		t,
		parseArgs([]string{
			"--out-kube-config-secret=\"vc-secret\"",
		}),
		map[string]string{
			"--out-kube-config-secret": "vc-secret",
		},
	)
	assert.DeepEqual(
		t,
		parseArgs([]string{
			"--translate-image=coredns/coredns=mirror.io/coredns/coredns",
		}),
		map[string]string{
			"--translate-image": "coredns/coredns=mirror.io/coredns/coredns",
		},
	)
}

func TestToExportKubeConfig(t *testing.T) {
	assert.DeepEqual(t,
		toExportKubeConfig([]string{}),
		config.ExportKubeConfig{},
	)
	assert.DeepEqual(
		t,
		toExportKubeConfig([]string{
			"--out-kube-config-secret=vc-secret",
		}),
		config.ExportKubeConfig{
			Secret: config.SecretReference{
				Name: "vc-secret",
			},
		},
	)
	assert.DeepEqual(
		t,
		toExportKubeConfig([]string{
			"--out-kube-config-secret-namespace=vc-namespace",
		}),
		config.ExportKubeConfig{
			Secret: config.SecretReference{
				Namespace: "vc-namespace",
			},
		},
	)
	assert.DeepEqual(
		t,
		toExportKubeConfig([]string{
			"--kube-config-context-name=vc-context",
		}),
		config.ExportKubeConfig{
			Context: "vc-context",
		},
	)
	assert.DeepEqual(
		t,
		toExportKubeConfig([]string{
			"--out-kube-config-server=vc-server",
		}),
		config.ExportKubeConfig{
			Server: "vc-server",
		},
	)
}

func TestToSync(t *testing.T) {
	assert.DeepEqual(t,
		toSync(nil, nil),
		config.Sync{ToHost: config.SyncToHost{}},
	)

	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Services: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{Services: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Endpoints: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{Endpoints: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Ingresses: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{Ingresses: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Priorityclasses: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{PriorityClasses: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Networkpolicies: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{NetworkPolicies: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Volumesnapshots: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{VolumeSnapshots: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Poddisruptionbudgets: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{PodDisruptionBudgets: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Serviceaccounts: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{ServiceAccounts: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{StorageClasses: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{StorageClasses: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{PersistentVolumes: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{PersistentVolumes: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{PersistentVolumeClaims: EnabledSwitch{Enabled: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{PersistentVolumeClaims: config.EnableSwitch{Enabled: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Configmaps: SyncConfigMaps{Enabled: &TRUE, All: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{ConfigMaps: config.SyncAllResource{Enabled: &TRUE, All: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(&BaseHelm{Sync: SyncValues{Secrets: SyncSecrets{Enabled: &TRUE, All: &TRUE}}}, nil),
		config.Sync{ToHost: config.SyncToHost{Secrets: config.SyncAllResource{Enabled: &TRUE, All: &TRUE}}},
	)
	assert.DeepEqual(t,
		toSync(
			&BaseHelm{
				Sync: SyncValues{
					Pods: SyncPods{
						Enabled: &TRUE,
					},
				},
			},
			&SyncerValues{
				ExtraArgs: []string{
					"--translate-image=coredns/coredns=mirror.io/coredns/coredns",
				},
			},
		),
		config.Sync{ToHost: config.SyncToHost{
			Pods: config.SyncPods{
				EnableSwitch: config.EnableSwitch{Enabled: &TRUE},
			},
		}},
	)
}
