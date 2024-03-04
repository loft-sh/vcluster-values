package helmvalues

import (
	"strings"

	"github.com/loft-sh/vcluster-values/config"
)

type Values interface {
	Upgrade() *config.Values
	FromString(values string) error
}

func toExportKubeConfig(args []string) config.ExportKubeConfig {
	parsedArgs := parseArgs(args)

	return config.ExportKubeConfig{
		Context: parsedArgs["--kube-config-context-name"],
		Server:  parsedArgs["--out-kube-config-server"],
		Secret: config.SecretReference{
			Name:      parsedArgs["--out-kube-config-secret"],
			Namespace: parsedArgs["--out-kube-config-secret-namespace"],
		},
	}
}

func toSync(values *BaseHelm, syncer *SyncerValues) config.Sync {
	if values == nil {
		return config.Sync{
			ToHost: config.SyncToHost{},
		}
	}

	parsedArgs := parseArgs(syncer.ExtraArgs)

	return config.Sync{
		ToHost: config.SyncToHost{
			Services:               toEnableSwitch(values.Sync.Services),
			Endpoints:              toEnableSwitch(values.Sync.Endpoints),
			Ingresses:              toEnableSwitch(values.Sync.Ingresses),
			PriorityClasses:        toEnableSwitch(values.Sync.Priorityclasses),
			NetworkPolicies:        toEnableSwitch(values.Sync.Networkpolicies),
			VolumeSnapshots:        toEnableSwitch(values.Sync.Volumesnapshots),
			PodDisruptionBudgets:   toEnableSwitch(values.Sync.Poddisruptionbudgets),
			ServiceAccounts:        toEnableSwitch(values.Sync.Serviceaccounts),
			StorageClasses:         toEnableSwitch(values.Sync.StorageClasses),
			PersistentVolumes:      toEnableSwitch(values.Sync.PersistentVolumes),
			PersistentVolumeClaims: toEnableSwitch(values.Sync.PersistentVolumeClaims),
			ConfigMaps: config.SyncAllResource{
				All:     values.Sync.Configmaps.All,
				Enabled: values.Sync.Configmaps.Enabled,
			},
			Secrets: config.SyncAllResource{
				All:     values.Sync.Secrets.All,
				Enabled: values.Sync.Secrets.Enabled,
			},
			Pods: config.SyncPods{
				EnableSwitch:   config.EnableSwitch{Enabled: values.Sync.Pods.Enabled},
				TranslateImage: parseTranslateImages(parsedArgs["--translate-image"]),
			},
		},
	}
}

func toEnableSwitch(enabledSwitch EnabledSwitch) config.EnableSwitch {
	return config.EnableSwitch{
		Enabled: enabledSwitch.Enabled,
	}
}

func parseArgs(args []string) map[string]string {
	parsedArgs := map[string]string{}
	for _, arg := range args {
		tokens := strings.Split(arg, "=")

		if len(tokens) == 0 {
			continue
		}

		argName := tokens[0]

		if len(tokens) == 1 {
			parsedArgs[argName] = "true"
			break
		}

		argValue := tokens[1]
		argValue = strings.TrimPrefix(argValue, `"`)
		argValue = strings.TrimSuffix(argValue, `"`)
		parsedArgs[argName] = argValue
	}
	return parsedArgs
}

func parseTranslateImages(translateImage string) map[string]string {
	return nil
}
