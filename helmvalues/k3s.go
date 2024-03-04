package helmvalues

import "github.com/loft-sh/vcluster-values/config"

type K3s struct {
	BaseHelm
	AutoDeletePersistentVolumeClaims bool               `json:"autoDeletePersistentVolumeClaims,omitempty"`
	K3sToken                         string             `json:"k3sToken,omitempty"`
	EmbeddedEtcd                     EmbeddedEtcdValues `json:"embeddedEtcd,omitempty"`
	Etcd                             K3SEtcdValues      `json:"etcd,omitempty"`
	VCluster                         VClusterValues     `json:"vcluster,omitempty"`
	Syncer                           SyncerValues       `json:"syncer,omitempty"`
}

func (v *K3s) Upgrade() *config.Values {
	return &config.Values{
		ExportKubeConfig: toExportKubeConfig(v.Syncer.ExtraArgs),
	}
}

func (v *K3s) FromString(values string) error {
	return nil
}

type K3SEtcdValues struct {
	Enabled bool `json:"enabled,omitempty"`
	Migrate bool `json:"migrate,omitempty"`

	CommonValues
	SyncerExORCommonValues
	ControlPlaneCommonValues
	Storage            Storage                `json:"storage,omitempty"`
	SecurityContext    map[string]interface{} `json:"securityContext,omitempty"`
	ServiceAnnotations map[string]string      `json:"serviceAnnotations,omitempty"`
}
