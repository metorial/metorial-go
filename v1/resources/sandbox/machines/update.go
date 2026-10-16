package machines

import (
	"encoding/json"
	"time"
)

// SandboxMachinesUpdateOutputImage represents the sandbox machines update output image type.
type SandboxMachinesUpdateOutputImage struct {
	// Id - The image the machine boots from
	Id *string `json:"id,omitempty"`
	// Name - The image name
	Name *string `json:"name,omitempty"`
	// Identifier - The global image identifier
	Identifier *string `json:"identifier,omitempty"`
	// VersionId - The image version the machine boots from
	VersionId *string `json:"version_id,omitempty"`
}

// SandboxMachinesUpdateOutputSpec - The machine's size. Null once a terminated machine's details expired, about 30 days after termination.
type SandboxMachinesUpdateOutputSpec struct {
	// Id - The spec the machine was created with
	Id *string `json:"id,omitempty"`
	// Identifier - The spec's human-readable identifier
	Identifier string `json:"identifier"`
	// MemoryGib - The machine's memory in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The machine's CPU count
	Cpus float64 `json:"cpus"`
}

// SandboxMachinesUpdateOutputMaxBoost - The maximum resources the machine can boost to, based on four times its backing size. Null once the remote machine details expire.
type SandboxMachinesUpdateOutputMaxBoost struct {
	// MemoryGib - The maximum memory the machine can boost to, in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The maximum CPU count the machine can boost to
	Cpus float64 `json:"cpus"`
}

// SandboxMachinesUpdateOutputRootDisk - The machine's root filesystem disk
type SandboxMachinesUpdateOutputRootDisk struct {
	// Id - The root disk's identifier
	Id string `json:"id"`
	// SizeBytes - The root disk's size in bytes
	SizeBytes *float64 `json:"size_bytes,omitempty"`
}

// SandboxMachinesUpdateOutputMounts represents the sandbox machines update output mounts type.
type SandboxMachinesUpdateOutputMounts struct {
	// DiskId - The mounted disk
	DiskId string `json:"disk_id"`
	// MountPath - Where the disk is mounted in the machine
	MountPath string `json:"mount_path"`
	// State - Whether the disk is currently mounted
	State string `json:"state"`
}

// SandboxMachinesUpdateOutputLifecyclePolicy represents the sandbox machines update output lifecycle policy type.
type SandboxMachinesUpdateOutputLifecyclePolicy struct {
	// PauseAfterIdleSeconds - Pause the machine after this many idle seconds
	PauseAfterIdleSeconds *float64 `json:"pause_after_idle_seconds,omitempty"`
	// HibernateAfterPausedSeconds - Hibernate the machine after it was paused for this many seconds
	HibernateAfterPausedSeconds *float64 `json:"hibernate_after_paused_seconds,omitempty"`
	// ConnectionIdleSeconds - Consider a connection idle after this many seconds
	ConnectionIdleSeconds *float64 `json:"connection_idle_seconds,omitempty"`
}

// SandboxMachinesUpdateOutputRuntime represents the sandbox machines update output runtime type.
type SandboxMachinesUpdateOutputRuntime struct {
	// Setup - Commands run once before the entrypoint starts
	Setup *[]string `json:"setup,omitempty"`
	// Entrypoint - The entrypoint command
	Entrypoint *[]string `json:"entrypoint,omitempty"`
	// Cmd - Default arguments for the entrypoint
	Cmd *[]string `json:"cmd,omitempty"`
	// Env - Environment variables
	Env *map[string]string `json:"env,omitempty"`
	// Workdir - The working directory
	Workdir *string `json:"workdir,omitempty"`
	// User - The user the entrypoint runs as
	User *string `json:"user,omitempty"`
	// StopSignal - The signal sent to the entrypoint when the machine stops
	StopSignal *string `json:"stop_signal,omitempty"`
}

// SandboxMachinesUpdateOutputLifecycleHooks represents the sandbox machines update output lifecycle hooks type.
type SandboxMachinesUpdateOutputLifecycleHooks struct {
	// Type - How the hook is invoked
	Type string `json:"type"`
	// On - The lifecycle transition the hook runs on
	On string `json:"on"`
	// Port - The port requested by an http hook
	Port *float64 `json:"port,omitempty"`
	// Path - The URL path of an http hook or the file of a script hook
	Path string `json:"path"`
	// Arguments - The arguments passed to a script hook
	Arguments *[]string `json:"arguments,omitempty"`
	// Wait - Whether the transition waits for the hook to finish
	Wait *bool `json:"wait,omitempty"`
	// TimeoutSeconds - How long the hook may run before it is considered failed
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty"`
}

// SandboxMachinesUpdateOutputOperation - The operation currently in flight, if any
type SandboxMachinesUpdateOutputOperation struct {
	Id          string    `json:"id"`
	Kind        string    `json:"kind"`
	Trigger     string    `json:"trigger"`
	RequestedAt time.Time `json:"requested_at"`
}

// SandboxMachinesUpdateOutputCurrentInstance - The instance currently backing the machine, if any
type SandboxMachinesUpdateOutputCurrentInstance struct {
	Id         string  `json:"id"`
	Generation float64 `json:"generation"`
	Status     string  `json:"status"`
	Phase      string  `json:"phase"`
}

// SandboxMachinesUpdateOutputUpcomingEvents represents the sandbox machines update output upcoming events type.
type SandboxMachinesUpdateOutputUpcomingEvents struct {
	// Type - What will happen
	Type string `json:"type"`
	// OccursAt - When it will happen
	OccursAt time.Time `json:"occurs_at"`
}

// SandboxMachinesUpdateOutputFailure represents the sandbox machines update output failure type.
type SandboxMachinesUpdateOutputFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesUpdateOutput represents the sandbox machines update output type.
type SandboxMachinesUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Metadata - User-defined metadata stored in Metorial
	Metadata map[string]any `json:"metadata"`
	// Id - The machine's unique identifier
	Id string `json:"id"`
	// Status - The machine's current status
	Status string `json:"status"`
	// Name - The machine's name
	Name *string `json:"name,omitempty"`
	// Slug - The machine's URL-safe slug, used in access hostnames. Null once a terminated machine's details expired.
	Slug *string `json:"slug,omitempty"`
	// Region - The identifier of the region the machine runs in
	Region string                           `json:"region"`
	Image  SandboxMachinesUpdateOutputImage `json:"image"`
	// Spec - The machine's size. Null once a terminated machine's details expired, about 30 days after termination.
	Spec *SandboxMachinesUpdateOutputSpec `json:"spec,omitempty"`
	// MaxBoost - The maximum resources the machine can boost to, based on four times its backing size. Null once the remote machine details expire.
	MaxBoost *SandboxMachinesUpdateOutputMaxBoost `json:"max_boost,omitempty"`
	// RootDisk - The machine's root filesystem disk
	RootDisk *SandboxMachinesUpdateOutputRootDisk `json:"root_disk,omitempty"`
	// Mounts - The additional disks mounted into the machine
	Mounts []SandboxMachinesUpdateOutputMounts `json:"mounts"`
	// FirewallIds - The firewalls applied to the machine
	FirewallIds []string `json:"firewall_ids"`
	// PublicPorts - The ports reachable without an access token
	PublicPorts     []float64                                  `json:"public_ports"`
	LifecyclePolicy SandboxMachinesUpdateOutputLifecyclePolicy `json:"lifecycle_policy"`
	Runtime         *SandboxMachinesUpdateOutputRuntime        `json:"runtime,omitempty"`
	// LifecycleHooks - Hooks that run on lifecycle transitions
	LifecycleHooks []SandboxMachinesUpdateOutputLifecycleHooks `json:"lifecycle_hooks"`
	// IgnoreImageHooks - Whether the image's own lifecycle hooks are ignored
	IgnoreImageHooks bool `json:"ignore_image_hooks"`
	// EntrypointWaitSeconds - How long startup waits for the entrypoint
	EntrypointWaitSeconds *float64 `json:"entrypoint_wait_seconds,omitempty"`
	// Operation - The operation currently in flight, if any
	Operation *SandboxMachinesUpdateOutputOperation `json:"operation,omitempty"`
	// CurrentInstance - The instance currently backing the machine, if any
	CurrentInstance *SandboxMachinesUpdateOutputCurrentInstance `json:"current_instance,omitempty"`
	// RestartRequired - Whether the running instance still uses an older configuration than the machine was last updated to. Metorial replaces such an instance on its own, restarting a running machine and hibernating a paused one, and the update is live once the machine starts again.
	RestartRequired bool `json:"restart_required"`
	// LastActivityAt - When the machine was last active
	LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
	// PauseAt - When the machine will be paused next
	PauseAt *time.Time `json:"pause_at,omitempty"`
	// HibernatesAt - When the machine will be hibernated next
	HibernatesAt *time.Time `json:"hibernates_at,omitempty"`
	// UpcomingEvents - Scheduled events within the next hour, soonest first. Reprovision replaces the underlying instance.
	UpcomingEvents []SandboxMachinesUpdateOutputUpcomingEvents `json:"upcoming_events"`
	Failure        *SandboxMachinesUpdateOutputFailure         `json:"failure,omitempty"`
	// CreatedAt - The machine's creation date
	CreatedAt time.Time `json:"created_at"`
	// RunningAt - When the machine last started running
	RunningAt *time.Time `json:"running_at,omitempty"`
	// PausedAt - When the machine was last paused
	PausedAt *time.Time `json:"paused_at,omitempty"`
	// HibernatedAt - When the machine was last hibernated
	HibernatedAt *time.Time `json:"hibernated_at,omitempty"`
	// TerminatedAt - When the machine was terminated
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

// MapSandboxMachinesUpdateOutputFromJSON deserializes JSON data into a SandboxMachinesUpdateOutput.
func MapSandboxMachinesUpdateOutputFromJSON(data []byte) (*SandboxMachinesUpdateOutput, error) {
	var v SandboxMachinesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesUpdateOutputToJSON serializes a SandboxMachinesUpdateOutput to JSON.
func MapSandboxMachinesUpdateOutputToJSON(v *SandboxMachinesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesUpdateBodyLifecyclePolicy represents the sandbox machines update body lifecycle policy type.
type SandboxMachinesUpdateBodyLifecyclePolicy struct {
	// PauseAfterIdleSeconds - Pause the machine after this many idle seconds
	PauseAfterIdleSeconds *float64 `json:"pause_after_idle_seconds,omitempty"`
	// HibernateAfterPausedSeconds - Hibernate the machine after it was paused for this many seconds
	HibernateAfterPausedSeconds *float64 `json:"hibernate_after_paused_seconds,omitempty"`
	// ConnectionIdleSeconds - Consider a connection idle after this many seconds
	ConnectionIdleSeconds *float64 `json:"connection_idle_seconds,omitempty"`
}

// SandboxMachinesUpdateBodyRuntime represents the sandbox machines update body runtime type.
type SandboxMachinesUpdateBodyRuntime struct {
	// Setup - Commands run once before the entrypoint starts
	Setup *[]string `json:"setup,omitempty"`
	// Entrypoint - The entrypoint command
	Entrypoint *[]string `json:"entrypoint,omitempty"`
	// Cmd - Default arguments for the entrypoint
	Cmd *[]string `json:"cmd,omitempty"`
	// Env - Environment variables
	Env *map[string]string `json:"env,omitempty"`
	// Workdir - The working directory
	Workdir *string `json:"workdir,omitempty"`
	// User - The user the entrypoint runs as
	User *string `json:"user,omitempty"`
	// StopSignal - The signal sent to the entrypoint when the machine stops
	StopSignal *string `json:"stop_signal,omitempty"`
}

// SandboxMachinesUpdateBodyMounts represents the sandbox machines update body mounts type.
type SandboxMachinesUpdateBodyMounts struct {
	// DiskId - The disk to mount
	DiskId string `json:"disk_id"`
	// MountPath - Where to mount the disk inside the machine
	MountPath string `json:"mount_path"`
}

// SandboxMachinesUpdateBody represents the sandbox machines update body type.
type SandboxMachinesUpdateBody struct {
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the machine
	Name *string `json:"name,omitempty"`
	// SpecId - The ID of the spec that sets the size of the machine
	SpecId *string `json:"spec_id,omitempty"`
	// SpecIdentifier - The identifier of the spec that sets the size of the machine
	SpecIdentifier *string `json:"spec_identifier,omitempty"`
	// MemoryGib - The memory of the machine in GiB, in steps of 0.5. Use together with cpus instead of a spec ID.
	MemoryGib *float64 `json:"memory_gib,omitempty"`
	// Cpus - The CPUs of the machine, in steps of 0.5. Use together with memory_gib instead of a spec ID.
	Cpus            *float64                                  `json:"cpus,omitempty"`
	LifecyclePolicy *SandboxMachinesUpdateBodyLifecyclePolicy `json:"lifecycle_policy,omitempty"`
	Runtime         *SandboxMachinesUpdateBodyRuntime         `json:"runtime,omitempty"`
	LifecycleHooks  *[]any                                    `json:"lifecycle_hooks,omitempty"`
	// IgnoreImageHooks - Ignore the image's own lifecycle hooks
	IgnoreImageHooks *bool `json:"ignore_image_hooks,omitempty"`
	// EntrypointWaitSeconds - How long startup waits for the entrypoint
	EntrypointWaitSeconds *float64                           `json:"entrypoint_wait_seconds,omitempty"`
	Mounts                *[]SandboxMachinesUpdateBodyMounts `json:"mounts,omitempty"`
	// FirewallIds - Firewalls to apply to the machine
	FirewallIds *[]string `json:"firewall_ids,omitempty"`
	// PublicPorts - Ports reachable without an access token
	PublicPorts *[]float64 `json:"public_ports,omitempty"`
}

// MapSandboxMachinesUpdateBodyFromJSON deserializes JSON data into a SandboxMachinesUpdateBody.
func MapSandboxMachinesUpdateBodyFromJSON(data []byte) (*SandboxMachinesUpdateBody, error) {
	var v SandboxMachinesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesUpdateBodyToJSON serializes a SandboxMachinesUpdateBody to JSON.
func MapSandboxMachinesUpdateBodyToJSON(v *SandboxMachinesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
