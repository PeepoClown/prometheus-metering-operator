package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrometheusMeteringSpec defines the desired state of PrometheusMetering
type PrometheusMeteringSpec struct {
	Name             string                    `json:"name,omitempty"`
	Namespace        string                    `json:"namespace,omitempty"`
	ReplicasCount    int32                     `json:"replicasCount,omitempty"`
	PrometheusConfig *PrometheusMeteringConfig `json:"prometheusConfig,omitempty"`
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type PrometheusMeteringConfig struct {
	GlobalConfig *GlobalPrometheusConfig `json:"globalConfig,omitempty" yaml:"global"`
	// +optional
	RemoteWrite []RemoteStorageUrl `json:"remoteWrite,omitempty" yaml:"remote_write"`
	// +optional
	RemoteRead []RemoteStorageUrl `json:"remoteRead,omitempty" yaml:"remote_read"`
	// +optional
	ScrapeJobs []ScrapeJob `json:"scrapeJobs,omitempty" yaml:"scrape_configs"`
}

type GlobalPrometheusConfig struct {
	ScrapeInterval string `json:"scrapeInterval,omitempty" yaml:"scrape_interval"`
	ScrapeTimeout  string `json:"scrapeTimeout,omitempty" yaml:"scrape_timeout"`
	// +optional
	EvaluationInterval string `json:"evaluationInterval,omitempty" yaml:"evaluation_interval"`
}

type RemoteStorageUrl struct {
	Url string `json:"url,omitempty" yaml:"url"`
}

type ScrapeJob struct {
	JobName       string         `json:"jobName,omitempty" yaml:"job_name"`
	MetricsPath   string         `json:"metricsPath,omitempty" yaml:"metrics_path"`
	StaticConfigs []StaticConfig `json:"staticConfigs,omitempty" yaml:"static_configs"`
	// +optional
	ScrapeInterval string `json:"scrapeInterval,omitempty" yaml:"scrape_interval"`
	// +optional
	ScrapeTimeout string `json:"scrapeTimeout,omitempty" yaml:"scrape_timeout"`
}

type StaticConfig struct {
	Targets []string `json:"targets" yaml:"targets"`
}

// PrometheusMeteringStatus defines the observed state of PrometheusMetering
type PrometheusMeteringStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PrometheusMetering is the Schema for the prometheusmeterings API
type PrometheusMetering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrometheusMeteringSpec   `json:"spec,omitempty"`
	Status PrometheusMeteringStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PrometheusMeteringList contains a list of PrometheusMetering
type PrometheusMeteringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrometheusMetering `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrometheusMetering{}, &PrometheusMeteringList{})
}
