package machines

import (
	"encoding/json"
	"time"
)

// SandboxMachinesListOutputItemsImage represents the sandbox machines list output items image type.
type SandboxMachinesListOutputItemsImage struct {
	// Id - The image the machine boots from
	Id *string `json:"id,omitempty"`
	// Name - The image name
	Name *string `json:"name,omitempty"`
	// Identifier - The global image identifier
	Identifier *string `json:"identifier,omitempty"`
	// VersionId - The image version the machine boots from
	VersionId *string `json:"version_id,omitempty"`
}

// SandboxMachinesListOutputItemsSpec - The machine's size. Null once a terminated machine's details expired, about 30 days after termination.
type SandboxMachinesListOutputItemsSpec struct {
	// Id - The spec the machine was created with
	Id *string `json:"id,omitempty"`
	// Identifier - The spec's human-readable identifier
	Identifier string `json:"identifier"`
	// MemoryGib - The machine's memory in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The machine's CPU count
	Cpus float64 `json:"cpus"`
}

// SandboxMachinesListOutputItemsMaxBoost - The maximum resources the machine can boost to, based on four times its backing size. Null once the remote machine details expire.
type SandboxMachinesListOutputItemsMaxBoost struct {
	// MemoryGib - The maximum memory the machine can boost to, in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The maximum CPU count the machine can boost to
	Cpus float64 `json:"cpus"`
}

// SandboxMachinesListOutputItemsRootDisk - The machine's root filesystem disk
type SandboxMachinesListOutputItemsRootDisk struct {
	// Id - The root disk's identifier
	Id string `json:"id"`
	// SizeBytes - The root disk's size in bytes
	SizeBytes *float64 `json:"size_bytes,omitempty"`
}

// SandboxMachinesListOutputItemsMounts represents the sandbox machines list output items mounts type.
type SandboxMachinesListOutputItemsMounts struct {
	// DiskId - The mounted disk
	DiskId string `json:"disk_id"`
	// MountPath - Where the disk is mounted in the machine
	MountPath string `json:"mount_path"`
	// State - Whether the disk is currently mounted
	State string `json:"state"`
}

// SandboxMachinesListOutputItemsLifecyclePolicy represents the sandbox machines list output items lifecycle policy type.
type SandboxMachinesListOutputItemsLifecyclePolicy struct {
	// PauseAfterIdleSeconds - Pause the machine after this many idle seconds
	PauseAfterIdleSeconds *float64 `json:"pause_after_idle_seconds,omitempty"`
	// HibernateAfterPausedSeconds - Hibernate the machine after it was paused for this many seconds
	HibernateAfterPausedSeconds *float64 `json:"hibernate_after_paused_seconds,omitempty"`
	// ConnectionIdleSeconds - Consider a connection idle after this many seconds
	ConnectionIdleSeconds *float64 `json:"connection_idle_seconds,omitempty"`
}

// SandboxMachinesListOutputItemsRuntime represents the sandbox machines list output items runtime type.
type SandboxMachinesListOutputItemsRuntime struct {
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

// SandboxMachinesListOutputItemsLifecycleHooks represents the sandbox machines list output items lifecycle hooks type.
type SandboxMachinesListOutputItemsLifecycleHooks struct {
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

// SandboxMachinesListOutputItemsOperation - The operation currently in flight, if any
type SandboxMachinesListOutputItemsOperation struct {
	Id          string    `json:"id"`
	Kind        string    `json:"kind"`
	Trigger     string    `json:"trigger"`
	RequestedAt time.Time `json:"requested_at"`
}

// SandboxMachinesListOutputItemsCurrentInstance - The instance currently backing the machine, if any
type SandboxMachinesListOutputItemsCurrentInstance struct {
	Id         string  `json:"id"`
	Generation float64 `json:"generation"`
	Status     string  `json:"status"`
	Phase      string  `json:"phase"`
}

// SandboxMachinesListOutputItemsUpcomingEvents represents the sandbox machines list output items upcoming events type.
type SandboxMachinesListOutputItemsUpcomingEvents struct {
	// Type - What will happen
	Type string `json:"type"`
	// OccursAt - When it will happen
	OccursAt time.Time `json:"occurs_at"`
}

// SandboxMachinesListOutputItemsFailure represents the sandbox machines list output items failure type.
type SandboxMachinesListOutputItemsFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesListOutputItems represents the sandbox machines list output items type.
type SandboxMachinesListOutputItems struct {
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
	Region string                              `json:"region"`
	Image  SandboxMachinesListOutputItemsImage `json:"image"`
	// Spec - The machine's size. Null once a terminated machine's details expired, about 30 days after termination.
	Spec *SandboxMachinesListOutputItemsSpec `json:"spec,omitempty"`
	// MaxBoost - The maximum resources the machine can boost to, based on four times its backing size. Null once the remote machine details expire.
	MaxBoost *SandboxMachinesListOutputItemsMaxBoost `json:"max_boost,omitempty"`
	// RootDisk - The machine's root filesystem disk
	RootDisk *SandboxMachinesListOutputItemsRootDisk `json:"root_disk,omitempty"`
	// Mounts - The additional disks mounted into the machine
	Mounts []SandboxMachinesListOutputItemsMounts `json:"mounts"`
	// FirewallIds - The firewalls applied to the machine
	FirewallIds []string `json:"firewall_ids"`
	// PublicPorts - The ports reachable without an access token
	PublicPorts     []float64                                     `json:"public_ports"`
	LifecyclePolicy SandboxMachinesListOutputItemsLifecyclePolicy `json:"lifecycle_policy"`
	Runtime         *SandboxMachinesListOutputItemsRuntime        `json:"runtime,omitempty"`
	// LifecycleHooks - Hooks that run on lifecycle transitions
	LifecycleHooks []SandboxMachinesListOutputItemsLifecycleHooks `json:"lifecycle_hooks"`
	// IgnoreImageHooks - Whether the image's own lifecycle hooks are ignored
	IgnoreImageHooks bool `json:"ignore_image_hooks"`
	// EntrypointWaitSeconds - How long startup waits for the entrypoint
	EntrypointWaitSeconds *float64 `json:"entrypoint_wait_seconds,omitempty"`
	// Operation - The operation currently in flight, if any
	Operation *SandboxMachinesListOutputItemsOperation `json:"operation,omitempty"`
	// CurrentInstance - The instance currently backing the machine, if any
	CurrentInstance *SandboxMachinesListOutputItemsCurrentInstance `json:"current_instance,omitempty"`
	// RestartRequired - Whether the running instance still uses an older configuration than the machine was last updated to. Metorial replaces such an instance on its own, restarting a running machine and hibernating a paused one, and the update is live once the machine starts again.
	RestartRequired bool `json:"restart_required"`
	// LastActivityAt - When the machine was last active
	LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
	// PauseAt - When the machine will be paused next
	PauseAt *time.Time `json:"pause_at,omitempty"`
	// HibernatesAt - When the machine will be hibernated next
	HibernatesAt *time.Time `json:"hibernates_at,omitempty"`
	// UpcomingEvents - Scheduled events within the next hour, soonest first. Reprovision replaces the underlying instance.
	UpcomingEvents []SandboxMachinesListOutputItemsUpcomingEvents `json:"upcoming_events"`
	Failure        *SandboxMachinesListOutputItemsFailure         `json:"failure,omitempty"`
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

// SandboxMachinesListOutputPagination represents the sandbox machines list output pagination type.
type SandboxMachinesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxMachinesListOutput represents the sandbox machines list output type.
type SandboxMachinesListOutput struct {
	Items      []SandboxMachinesListOutputItems    `json:"items"`
	Pagination SandboxMachinesListOutputPagination `json:"pagination"`
}

// MapSandboxMachinesListOutputFromJSON deserializes JSON data into a SandboxMachinesListOutput.
func MapSandboxMachinesListOutputFromJSON(data []byte) (*SandboxMachinesListOutput, error) {
	var v SandboxMachinesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesListOutputToJSON serializes a SandboxMachinesListOutput to JSON.
func MapSandboxMachinesListOutputToJSON(v *SandboxMachinesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesListQuery represents the sandbox machines list query type.
type SandboxMachinesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by resource IDs
	Id *any `json:"id,omitempty"`
	// Status - Filter by whether the machine still exists (active) or was terminated (deleted). Defaults to active.
	Status *any `json:"status,omitempty"`
	// State - Filter by the machine's current runtime status, the status field of the machine. Terminated machines are only listed together with status=deleted.
	State *any `json:"state,omitempty"`
	// ImageId - Filter by the image the machines boot from
	ImageId *any `json:"image_id,omitempty"`
	// Region - Filter by the identifiers of the regions the machines run in
	Region *any `json:"region,omitempty"`
	// Search - Search machines by name, slug, or ID
	Search *string `json:"search,omitempty"`
}

// MapSandboxMachinesListQueryFromJSON deserializes JSON data into a SandboxMachinesListQuery.
func MapSandboxMachinesListQueryFromJSON(data []byte) (*SandboxMachinesListQuery, error) {
	var v SandboxMachinesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesListQueryToJSON serializes a SandboxMachinesListQuery to JSON.
func MapSandboxMachinesListQueryToJSON(v *SandboxMachinesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
