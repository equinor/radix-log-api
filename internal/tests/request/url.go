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

func ComponentReplicaLogUrl(appName, envName, componentName, replicaName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s/replicas/%s/log", appName, envName, componentName, replicaName), options...)
}

func ComponentContainerLogUrl(appName, envName, componentName, replicaName, containerID string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s/replicas/%s/containers/%s/log", appName, envName, componentName, replicaName, containerID), options...)
}

func JobInventoryUrl(appName, envName, jobComponentName, jobName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/jobcomponents/%s/jobs/%s", appName, envName, jobComponentName, jobName), options...)
}

func JobLogUrl(appName, envName, jobComponentName, jobName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/jobcomponents/%s/jobs/%s/log", appName, envName, jobComponentName, jobName), options...)
}

func JobReplicaLogUrl(appName, envName, jobComponentName, jobName, replicaName string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/jobcomponents/%s/jobs/%s/replicas/%s/log", appName, envName, jobComponentName, jobName, replicaName), options...)
}

func JobContainerLogUrl(appName, envName, jobComponentName, jobName, replicaName, containerID string, options ...UrlOption) string {
	return URL(fmt.Sprintf("/api/v1/applications/%s/environments/%s/jobcomponents/%s/jobs/%s/replicas/%s/containers/%s/log", appName, envName, jobComponentName, jobName, replicaName, containerID), options...)
}

func URL(path string, options ...UrlOption) string {
	u := &url.URL{Path: path}

	for _, o := range options {
		o(u)
	}

	return u.String()
}
