package helmvalues

import "github.com/loft-sh/vcluster-values/config"

type K0s struct {
	BaseHelm
	AutoDeletePersistentVolumeClaims bool               `json:"autoDeletePersistentVolumeClaims,omitempty"`
	VCluster                         VClusterValues     `json:"vcluster,omitempty"`
	Syncer                           SyncerValues       `json:"syncer,omitempty"`
	EmbeddedEtcd                     EmbeddedEtcdValues `json:"embeddedEtcd,omitempty"`
}

func (v *K0s) Upgrade() *config.Values {
	return &config.Values{
		ExportKubeConfig: toExportKubeConfig(v.Syncer.ExtraArgs),
	}
}

func (v *K0s) FromString(values string) error {
	return nil
}
