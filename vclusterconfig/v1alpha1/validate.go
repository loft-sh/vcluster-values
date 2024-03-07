package v1alpha1

import (
	"errors"
	"regexp"
)

var validationFunctions = []func(Config) error{
	syncerImageFormat,
	k3sToken,
	controllerAndApi,
}

func Validate(b Config) []error {
	errors := make([]error, 0)
	for _, function := range validationFunctions {
		err := function(b)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func syncerImageFormat(b Config) error {
	r := regexp.MustCompile("[a-zA-Z]+.*$")
	if !r.Match([]byte(b.Syncer.Image)) {
		return errors.New("the syncer image doesn't have the right format")
	}
	return nil
}

func k3sToken(b Config) error {
	if b.Distro != "k3s" && b.K3sToken != "" {
		return errors.New("k3sToken is only valid when the distro is k3s")
	}
	return nil
}

func controllerAndApi(b Config) error {
	if b.Distro == "eks" || b.Distro == "k8s" {
		return nil
	}

	if b.Controller.Image != "" ||
		len(b.Etcd.ExtraArgs) > 0 ||
		len(b.Etcd.Resources) > 0 ||
		b.Controller.ImagePullPolicy != "" {
		return errors.New("controller field is only valid with k8s and eks")
	}
	return nil
}
