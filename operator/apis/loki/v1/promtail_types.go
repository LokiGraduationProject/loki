package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PromtailSpec defines the desired state of Promtail
type PromtailSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Config PromtailConfig `json:"promtail_config"`
}

// PromtailConfig defines the configuration of the promtail deployment
type PromtailConfig struct {
	// +optional
	Server ServerConfig `json:"server,omitempty"`

	// +optional
	Clients []ClientConfig `json:"clients,omitempty"`

	// +optional
	Positions PositionsConfig `json:"positions,omitempty"`

	// +optional
	TargetConfig TargetConfigConfig `json:"target_config,omitempty"`

	// +optional
	ScrapeConfigs []ScrapeConfigsConfig `json:"scrape_configs,omitempty"`
}

type ServerConfig struct {
	Disable                   bool   `json:"disable,omitempty"`
	ProfilingEnabled           bool   `json:"profiling_enabled,omitempty"`
	HttpListenAddress          string `json:"http_listen_address,omitempty"`
	HttpListenPort             int    `json:"http_listen_port,omitempty"`
	GrpcListenAddress          string `json:"grpc_listen_address,omitempty"`
	GrpcListenPort             int    `json:"grpc_listen_port,omitempty"`
	RegisterInstrumentation    bool   `json:"register_instrumentation,omitempty"`
	GracefulShutdownTimeout    string `json:"graceful_shutdown_timeout,omitempty"`
	HttpServerReadTimeout      string `json:"http_server_read_timeout,omitempty"`
	HttpServerWriteTimeout     string `json:"http_server_write_timeout,omitempty"`
	HttpServerIdleTimeout      string `json:"http_server_idle_timeout,omitempty"`
	GrpcServerMaxRecvMsgSize   int    `json:"grpc_server_max_recv_msg_size,omitempty"`
	GrpcServerMaxSendMsgSize   int    `json:"grpc_server_max_send_msg_size,omitempty"`
	GrpcServerMaxConcurrentStreams int `json:"grpc_server_max_concurrent_streams,omitempty"`
	LogLevel                   string `json:"log_level,omitempty"`
	HttpPathPrefix             string `json:"http_path_prefix,omitempty"`
	HealthCheckTarget          bool   `json:"health_check_target,omitempty"`
	EnableRuntimeReload        bool   `json:"enable_runtime_reload,omitempty"`
}

type ClientConfig struct {
	// Loki의 HTTP URL
	// +optional
	URL string `json:"url"`

	// HTTP 헤더 설정
	// +optional
	Headers map[string]string `json:"headers,omitempty"`

	// 테넌트 ID
	// +optional
	TenantID string `json:"tenant_id,omitempty"`

	// 최대 대기 시간
	// +optional
	BatchWait string `json:"batchwait,omitempty"`

	// 최대 배치 크기
	// +optional
	BatchSize int `json:"batchsize,omitempty"`

	// Basic Auth 설정
	// +optional
	BasicAuth *BasicAuthConfig `json:"basic_auth,omitempty"`

	// OAuth2 설정
	// +optional
	OAuth2 *OAuth2Config `json:"oauth2,omitempty"`

	// Bearer Token 설정
	// +optional
	BearerToken      string `json:"bearer_token,omitempty"`
	BearerTokenFile  string `json:"bearer_token_file,omitempty"`

	// 프록시 설정
	// +optional
	ProxyURL string `json:"proxy_url,omitempty"`

	// TLS 설정
	// +optional
	TLSConfig *TLSConfig `json:"tls_config,omitempty"`

	// 백오프 설정
	// +optional
	BackoffConfig *BackoffConfig `json:"backoff_config,omitempty"`

	// Rate Limit 관련 설정
	// +optional
	DropRateLimitedBatches bool `json:"drop_rate_limited_batches,omitempty"`

	// 외부 레이블 설정
	// +optional
	ExternalLabels map[string]string `json:"external_labels,omitempty"`

	// 요청 타임아웃
	// +optional
	Timeout string `json:"timeout,omitempty"`
}

type BasicAuthConfig struct {
	// +optional
	Username     string `json:"username,omitempty"`

	// +optional
	Password     string `json:"password,omitempty"`

	// +optional
	PasswordFile string `json:"password_file,omitempty"`
}

type OAuth2Config struct {
	// +optional
	ClientID         string            `json:"client_id,omitempty"`

	// +optional
	ClientSecret     string            `json:"client_secret,omitempty"`

	// +optional
	ClientSecretFile string            `json:"client_secret_file,omitempty"`

	// +optional
	Scopes           []string          `json:"scopes,omitempty"`

	// +optional
	TokenURL         string            `json:"token_url,omitempty"`

	// +optional
	EndpointParams   map[string]string `json:"endpoint_params,omitempty"`
}

type TLSConfig struct {
	// +optional
	CAFile             string `json:"ca_file,omitempty"`

	// +optional
	CertFile           string `json:"cert_file,omitempty"`

	// +optional
	KeyFile            string `json:"key_file,omitempty"`

	// +optional
	ServerName         string `json:"server_name,omitempty"`

	// +optional
	InsecureSkipVerify bool   `json:"insecure_skip_verify,omitempty"`
}

type BackoffConfig struct {
	// +optional
	MinPeriod   string `json:"min_period,omitempty"`

	// +optional
	MaxPeriod   string `json:"max_period,omitempty"`

	// +optional
	MaxRetries  int    `json:"max_retries,omitempty"`
}

type PositionsConfig struct {
	// +optional
	Filename          string `json:"filename,omitempty"`

	// +optional
	SyncPeriod        string `json:"sync_period,omitempty"`

	// +optional
	IgnoreInvalidYaml bool `json:"ignore_invalid_yaml,omitempty"`
}

type TargetConfigConfig struct {
	// +optional
	SyncPeriod string `json:"sync_period,omitempty"`
}

type ScrapeConfigsConfig struct {
	// +optional
	JobName             string               `json:"job_name"`

	KubernetesSDConfigs []KubernetesSDConfig `json:"kubernetes_sd_configs,omitempty"`

	PipelineStages      []PipelineStage      `json:"pipeline_stages,omitempty"` 

	RelabelConfigs      []RelabelConfigsConfig      `json:"relabel_configs,omitempty"`
}

type KubernetesSDConfig struct {
	// +optional
	Role string `json:"role"`
}

type PipelineStage struct {
	Docker bool `json:"docker,omitempty"`
	Cri    bool `json:"cri,omitempty"`
}

type RelabelConfigsConfig struct {

	// Their content is concatenated using the configured separator.
	SourceLabels []string `json:"source_labels,omitempty"` // 기본값 없음

	// Separator placed between concatenated source label values.
	Separator string `json:"separator,omitempty"` // default = ";"

	// TargetLabel is the label to which the resulting value is written in a replace action.
	TargetLabel string `json:"target_label,omitempty"` // 기본값 없음

	// Regex is the regular expression against which the extracted value is matched.
	Regex string `json:"regex,omitempty"` // default = "(.*)"

	// Modulus to take of the hash of the source label values.
	Modulus uint64 `json:"modulus,omitempty"` // 기본값 없음

	// Replacement is the value against which a regex replace is performed if the regular expression matches.
	Replacement string `json:"replacement,omitempty"` // default = "$1"

	// Action to perform based on regex matching.
	Action string `json:"action,omitempty"` // default = "replace"
}

// PromtailStatus defines the observed state of Promtail
type PromtailStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Promtail is the Schema for the promtails API
type Promtail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PromtailSpec   `json:"spec,omitempty"`
	Status PromtailStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PromtailList contains a list of Promtail
type PromtailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Promtail `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Promtail{}, &PromtailList{})
}
