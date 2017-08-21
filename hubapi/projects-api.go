package hubapi

const (
	ProjectVersionPhasePlanning    = "PLANNING"
	ProjectVersionPhaseDevelopment = "DEVELOPMENT"
	ProjectVersionPhaseReleased    = "RELEASED"
	ProjectVersionPhaseDeprecated  = "DEPRECATED"
	ProjectVersionPhaseArchived    = "ARCHIVED"
)

const (
	ProjectVersionDistributionExternal   = "EXTERNAL"
	ProjectVersionDistributionSaaS       = "SAAS"
	ProjectVersionDistributionInternal   = "INTERNAL"
	ProjectVersionDistributionOpenSource = "OPENSOURCE"
)

type ProjectList struct {
	TotalCount uint32    `json:"totalCount"`
	Items      []Project `json:"items"`
}

type Project struct {
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Source                  string `json:"source"`
	ProjectTier             uint32 `json:"projectTier"`
	ProjectLevelAdjustments bool   `json:"projectLevelAdjustments"`
	ProjectOwner            string `json:"projectOwner"`
	Meta                    Meta   `json:"_meta"`
}

type ProjectRequest struct {
	Name                    string                 `json:"name"`
	Description             *string                `json:"description"`
	ProjectTier             *int                   `json:"projectTier"`
	ProjectOwner            *string                `json:"projectOwner"`
	ProjectLevelAdjustments bool                   `json:"projectLevelAdjustments"`
	VersionRequest          *ProjectVersionRequest `json:"versionRequest"`
}

type ProjectVersionList struct {
	TotalCount uint32           `json:"totalCount"`
	Items      []ProjectVersion `json:"items"`
}

type ProjectVersion struct {
	VersionName     string `json:"versionName"`
	Nickname        string `json:"nickname"`
	ReleaseComments string `json:"releaseComments"`
	ReleasedOn      string `json:"releasedOn"` // TODO: change this to a date
	Phase           string `json:"phase"`
	Distribution    string `json:"distribution"`
	Meta            Meta   `json:"_meta"`
}

type ProjectVersionRequest struct {
	VersionName     string  `json:"versionName"`
	Nickname        string  `json:"nickname"`
	ReleaseComments string  `json:"releaseComments"`
	ReleasedOn      *string `json:"releasedOn"` // TODO: change this to a date
	Phase           string  `json:"phase"`
	Distribution    string  `json:"distribution"`
}

// TODO: This is horrible, we need to fix up this API
type ProjectVersionRiskProfile struct {
	Categories       map[string]map[string]int `json:"categories"`
	BomLastUpdatedAt string                    `json:"bomLastUpdatedAt"` // TODO: Should be a date/time
	Meta             Meta                      `json:"_meta"`
}

func (p *Project) GetProjectVersionsLink() (*ResourceLink, error) {
	return p.Meta.FindLinkByRel("versions")
}

func (v *ProjectVersion) GetProjectLink() (*ResourceLink, error) {
	return v.Meta.FindLinkByRel("project")
}

func (v *ProjectVersion) GetCodeLocationsLink() (*ResourceLink, error) {
	return v.Meta.FindLinkByRel("codelocations")
}

func (v *ProjectVersion) GetComponentsLink() (*ResourceLink, error) {
	return v.Meta.FindLinkByRel("components")
}

func (v *ProjectVersion) GetVulnerableComponentsLink() (*ResourceLink, error) {
	return v.Meta.FindLinkByRel("vulnerable-components")
}

func (v *ProjectVersion) GetProjectVersionRiskProfileLink() (*ResourceLink, error) {
	return v.Meta.FindLinkByRel("riskProfile")
}
