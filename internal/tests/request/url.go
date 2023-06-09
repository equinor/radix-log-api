package request

import (
	"fmt"
	"net/url"
)

type UrlOption func(*url.URL)

func WithQueryParam(name, value string) UrlOption {
	return func(u *url.URL) {
		q := u.Query()
		q.Add(name, value)
		u.RawQuery = q.Encode()
	}
}

func ComponentInventoryUrl(appName, envName, componentName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s", appName, envName, componentName), options...)
}

func ComponentLogUrl(appName, envName, componentName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s/log", appName, envName, componentName), options...)
}

func ReplicaLogUrl(appName, envName, componentName, replicaName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s/replicas/%s/log", appName, envName, componentName, replicaName), options...)
}

func ContainerLogUrl(appName, envName, componentName, replicaName, containerID string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s/replicas/%s/containers/%s/log", appName, envName, componentName, replicaName, containerID), options...)
}

func URL(path string, options ...UrlOption) string {
	u := &url.URL{Path: path}

	for _, o := range options {
		o(u)
	}

	return u.String()
}
