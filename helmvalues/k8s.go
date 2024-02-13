package helmvalues

type K8s struct {
	BaseHelm
}

type APIServerValues struct {
	SyncerExORCommonValues
	ControlPlaneCommonValues
	SecurityContext    map[string]interface{} `json:"securityContext,omitempty"`
	ServiceAnnotations map[string]string      `json:"serviceAnnotations,omitempty"`
	Command            []string               `json:"command,omitempty"`
	BaseArgs           []string               `json:"baseArgs,omitempty"`
	ExtraArgs          []string               `json:"extraArgs,omitempty"`
}

type ControllerValues struct {
	SyncerExORCommonValues
	ControlPlaneCommonValues
}

type SchedulerValues struct {
	SyncerExORCommonValues
	ControlPlaneCommonValues
	Disabled bool `json:"disabled,omitempty"`
}
