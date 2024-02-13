package helmvalues

var K3sDefault = K3s{
	BaseHelm: BaseHelm{
		Pro:      false,
		Headless: false,
		Monitoring: MonitoringValues{
			ServiceMonitor: ServiceMonitor{Enabled: false},
		},

		Etcd: EtcdValues{
			Embedded: EmbeddedEtcdValues{Enabled: false},
			Replicas: 1,
		},
		Storage: Storage{
			Persistence:                      true,
			Size:                             "5Gi",
			AutoDeletePersistentVolumeClaims: false,
		},
		Api: APIServerValues{
			ControlPlaneCommonValues: ControlPlaneCommonValues{
				Image: "rancher/k3s:v1.19.0-k3s1",
			},
			Command: []string{"/binaries/k3s"},
			BaseArgs: []string{
				"server",
				"--write-kubeconfig=/data/k3s-config/kube-config.yaml",
				"--data-dir=/data",
				"--disable=traefik,servicelb,metrics-server,local-storage,coredns",
				"--disable-network-policy",
				"--disable-agent",
				"--disable-cloud-controller",
				"--egress-selector-mode=disabled",
				"--flannel-backend=none",
				"--kube-apiserver-arg=bind-address=127.0.0.1",
			},
		},
		Syncer: SyncerValues{
			Replicas:              1,
			LivenessProbe:         EnabledSwitch{Enabled: false},
			ReadinessProbe:        EnabledSwitch{Enabled: false},
			VolumeMounts:          []map[string]interface{}{{"data": "/data"}},
			KubeConfigContextName: "my-vcluster",
		},
		Rbac: RBACValues{Role: RBACRoleValues{Create: true}},
		Service: ServiceValues{
			Type: "ClusterIP",
		},
		Ingress: IngressValues{
			Enabled: false,
			Host:    "vcluster.local",
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/backend-protocol": "HTTPS",
				"nginx.ingress.kubernetes.io/ssl-passthrough":  "true",
				"nginx.ingress.kubernetes.io/ssl-redirect":     "true",
			},
		},
		SecurityContext: map[string]interface{}{
			"allowPrivilegeEscalation": false,
			"runAsUser":                0,
			"runAsGroup":               0,
		},
		Coredns: CoreDNSValues{
			Enabled:    true,
			Integrated: false,
			Fallback:   "8.8.8.8",
			Plugin:     CoreDNSPluginValues{Enabled: false},
			Replicas:   1,
			Service:    CoreDNSServiceValues{Type: "ClusterIP"},
		},
		Telemetry: TelemetryValues{Disabled: false, InstanceCreator: "helm"},
		Sync: SyncValues{
			Services:               EnabledSwitch{true},
			Configmaps:             SyncConfigMaps{Enabled: true, All: false},
			Secrets:                SyncSecrets{Enabled: true, All: false},
			Endpoints:              EnabledSwitch{true},
			Pods:                   SyncPods{Enabled: true},
			Events:                 EnabledSwitch{true},
			PersistentVolumeClaims: EnabledSwitch{true},
			Ingresses:              EnabledSwitch{true},
			FakeNodes:              EnabledSwitch{true},
			FakePersistentvolumes:  EnabledSwitch{true},
			Nodes:                  SyncNodes{Enabled: false, FakeKubeletIPs: true},
		},
	},
}
