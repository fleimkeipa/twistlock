package model

import (
	"net/http"
	"time"
)

type Scanner struct {
	BaseURL string
	Token   string
	Client  *http.Client
	User    User
	Report  ScanReport
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authenticate struct {
	Token string `json:"token"`
}

type ScanReport []struct {
	ID       string    `json:"_id"`
	Type     string    `json:"type"`
	Hostname string    `json:"hostname"`
	ScanTime time.Time `json:"scanTime"`
	Binaries []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		Md5        string `json:"md5"`
		CveCount   int    `json:"cveCount"`
		Version    string `json:"version,omitempty"`
		MissingPkg bool   `json:"missingPkg,omitempty"`
	} `json:"binaries"`
	Secrets         []interface{} `json:"Secrets"`
	StartupBinaries []struct {
		Name     string `json:"name"`
		Path     string `json:"path"`
		Md5      string `json:"md5"`
		CveCount int    `json:"cveCount"`
	} `json:"startupBinaries"`
	OsDistro        string `json:"osDistro"`
	OsDistroVersion string `json:"osDistroVersion"`
	OsDistroRelease string `json:"osDistroRelease"`
	Distro          string `json:"distro"`
	Packages        []struct {
		PkgsType string `json:"pkgsType"`
		Pkgs     []struct {
			Version    string   `json:"version"`
			Name       string   `json:"name"`
			CveCount   int      `json:"cveCount"`
			License    string   `json:"license"`
			LayerTime  int      `json:"layerTime"`
			BinaryIdx  []int    `json:"binaryIdx,omitempty"`
			BinaryPkgs []string `json:"binaryPkgs,omitempty"`
		} `json:"pkgs"`
	} `json:"packages"`
	Files          []interface{} `json:"files"`
	PackageManager bool          `json:"packageManager"`
	Applications   []struct {
		Name                 string `json:"name"`
		Version              string `json:"version"`
		Path                 string `json:"path"`
		LayerTime            int    `json:"layerTime"`
		KnownVulnerabilities int    `json:"knownVulnerabilities"`
	} `json:"applications,omitempty"`
	Image struct {
		Created time.Time `json:"created"`
	} `json:"image"`
	History []struct {
		Created     int    `json:"created"`
		Instruction string `json:"instruction"`
		EmptyLayer  bool   `json:"emptyLayer"`
		ID          string `json:"id,omitempty"`
	} `json:"history"`
	ComplianceIssues interface{} `json:"complianceIssues"`
	AllCompliance    struct {
	} `json:"allCompliance"`
	Vulnerabilities []struct {
		Text        string  `json:"text"`
		ID          int     `json:"id"`
		Severity    string  `json:"severity"`
		Cvss        float64 `json:"cvss"`
		Status      string  `json:"status"`
		Cve         string  `json:"cve"`
		Cause       string  `json:"cause"`
		Description string  `json:"description"`
		Title       string  `json:"title"`
		VecStr      string  `json:"vecStr"`
		Exploit     string  `json:"exploit"`
		RiskFactors struct {
			AttackComplexityLow struct {
			} `json:"Attack complexity: low"`
			AttackVectorNetwork struct {
			} `json:"Attack vector: network"`
			PackageInUse struct {
			} `json:"Package in use"`
		} `json:"riskFactors"`
		Link            string      `json:"link"`
		Type            string      `json:"type"`
		PackageName     string      `json:"packageName"`
		PackageVersion  string      `json:"packageVersion"`
		LayerTime       int         `json:"layerTime"`
		Templates       interface{} `json:"templates"`
		Twistlock       bool        `json:"twistlock"`
		Cri             bool        `json:"cri"`
		Published       int         `json:"published"`
		FixDate         int         `json:"fixDate"`
		ApplicableRules []string    `json:"applicableRules"`
		Discovered      time.Time   `json:"discovered"`
		FunctionLayer   string      `json:"functionLayer"`
		BinaryPkgs      []string    `json:"binaryPkgs,omitempty"`
	} `json:"vulnerabilities"`
	RepoTag struct {
		Registry string `json:"registry"`
		Repo     string `json:"repo"`
		Tag      string `json:"tag"`
	} `json:"repoTag"`
	Tags []struct {
		Registry string `json:"registry"`
		Repo     string `json:"repo"`
		Tag      string `json:"tag"`
	} `json:"tags"`
	RepoDigests               []string  `json:"repoDigests"`
	CreationTime              time.Time `json:"creationTime"`
	VulnerabilitiesCount      int       `json:"vulnerabilitiesCount"`
	ComplianceIssuesCount     int       `json:"complianceIssuesCount"`
	VulnerabilityDistribution struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Medium   int `json:"medium"`
		Low      int `json:"low"`
		Total    int `json:"total"`
	} `json:"vulnerabilityDistribution"`
	ComplianceDistribution struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Medium   int `json:"medium"`
		Low      int `json:"low"`
		Total    int `json:"total"`
	} `json:"complianceDistribution"`
	VulnerabilityRiskScore int      `json:"vulnerabilityRiskScore"`
	ComplianceRiskScore    int      `json:"complianceRiskScore"`
	Layers                 []string `json:"layers"`
	TopLayer               string   `json:"topLayer"`
	RiskFactors            struct {
		AttackComplexityLow struct {
		} `json:"Attack complexity: low"`
		AttackVectorNetwork struct {
		} `json:"Attack vector: network"`
		CriticalSeverity struct {
		} `json:"Critical severity"`
		DoS struct {
		} `json:"DoS"`
		HasFix struct {
		} `json:"Has fix"`
		HighSeverity struct {
		} `json:"High severity"`
		MediumSeverity struct {
		} `json:"Medium severity"`
		PackageInUse struct {
		} `json:"Package in use"`
		RecentVulnerability struct {
		} `json:"Recent vulnerability"`
		RemoteExecution struct {
		} `json:"Remote execution"`
	} `json:"riskFactors"`
	Labels            []string `json:"labels,omitempty"`
	InstalledProducts struct {
		OsDistro          string `json:"osDistro"`
		HasPackageManager bool   `json:"hasPackageManager"`
	} `json:"installedProducts"`
	ScanVersion   string    `json:"scanVersion"`
	FirstScanTime time.Time `json:"firstScanTime"`
	CloudMetadata struct {
		AccountID string `json:"accountID"`
	} `json:"cloudMetadata"`
	Namespaces []string `json:"namespaces"`
	Clusters   []string `json:"clusters"`
	Instances  []struct {
		Image    string    `json:"image"`
		Host     string    `json:"host"`
		Modified time.Time `json:"modified"`
		Tag      string    `json:"tag"`
		Repo     string    `json:"repo"`
		Registry string    `json:"registry"`
	} `json:"instances"`
	Hosts struct {
		K8SPPlatformP13PocMarsWorker1 struct {
			Modified time.Time `json:"modified"`
		} `json:"k8s-p-platform-p1-3poc-mars-worker-1"`
		K8SPPlatformP13PocMarsWorker2 struct {
			Modified time.Time `json:"modified"`
		} `json:"k8s-p-platform-p1-3poc-mars-worker-2"`
		K8SPPlatformP13PocMarsWorker3 struct {
			Modified time.Time `json:"modified"`
		} `json:"k8s-p-platform-p1-3poc-mars-worker-3"`
		K8SPPlatformP13PocMarsWorker4 struct {
			Modified time.Time `json:"modified"`
		} `json:"k8s-p-platform-p1-3poc-mars-worker-4"`
		K8SPPlatformP13PocMarsWorker5 struct {
			Modified time.Time `json:"modified"`
		} `json:"k8s-p-platform-p1-3poc-mars-worker-5"`
	} `json:"hosts"`
	Err                string   `json:"err"`
	Collections        []string `json:"collections"`
	ScanID             int      `json:"scanID"`
	TrustStatus        string   `json:"trustStatus"`
	FirewallProtection struct {
		Enabled   bool `json:"enabled"`
		Supported bool `json:"supported"`
	} `json:"firewallProtection"`
	AppEmbedded    bool        `json:"appEmbedded"`
	WildFireUsage  interface{} `json:"wildFireUsage"`
	ExternalLabels []struct {
		SourceType string    `json:"sourceType"`
		SourceName string    `json:"sourceName"`
		Timestamp  time.Time `json:"timestamp"`
		Key        string    `json:"key"`
		Value      string    `json:"value"`
	} `json:"externalLabels,omitempty"`
	RhelRepos []string `json:"rhelRepos,omitempty"`
}
