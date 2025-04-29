package plugsdk

type MessageCode string

const (
	BuildRequest         = "OPENMISO_BuildRequest"
	BuildArtifact        = "OPENMISO_BuildArtifact"
	ReleaseRequest       = "OPENMISO_ReleaseRequest"
	ReleaseArtifact      = "OPENMISO_ReleaseArtifact"
	DeployRequest        = "OPENMISO_DeployRequest"
	DeployArtifact       = "OPENMISO_DeployArtifact"
	LogMessage           = "OPENMISO_LogMessage"
	UnsupportedOperation = "OPENMISO_UnsupportedOperation"
)
