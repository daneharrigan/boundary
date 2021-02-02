package main

import (
	"github.com/hashicorp/boundary/internal/types/resource"
)

type cmdInfo struct {
	// The type of the resource, e.g. "target"
	ResourceType string

	// The import path of the API package
	PkgPath string

	// Standard actions (with standard parameters) used by this resource
	StdActions []string

	// HasCustomActionFlags controls whether to generate code to add extra
	// flags, useful for subtype actions
	HasCustomActionFlags bool

	// HasExtraCommandVars controls whether to generate an embedded struct with
	// extra command variables
	HasExtraCommandVars bool

	// HasExtraSynopsisFunc controls whether to generate code to look for an
	// extra synopsis function
	HasExtraSynopsisFunc bool

	// HasExtraActions controls whether to generate code to populate extra
	// flags into the map and switch on those functions
	HasExtraActions bool

	// HasExtraFlagsFunc controls whether to insert code to add extra flags into
	// the map
	HasExtraFlagsFunc bool

	// SkipNormalHelp indicates skipping the normal help case for when it needs
	// to be only custom (mainly for subactions)
	SkipNormalHelp bool

	// HasExtraHelpFunc controls whether to include a default statement chaining
	// to an extra help function
	HasExtraHelpFunc bool

	// HasExampleCliOutput controls whether to generate code to look for a CLI
	// output env var and print
	HasExampleCliOutput bool

	// IsAbstractType triggers some behavior specialized for abstract types,
	// e.g. those that have subcommands for create/update
	IsAbstractType bool

	// HasExtraFlagHandlingFunc controls whether to call out to an external command
	// for extra flag handling
	HasExtraFlagHandlingFunc bool

	// HasId controls whether to add ID emptiness checking. Note that some
	// commands that allow name/scope id or name/scope name handling may skip
	// this in favor of custom logic.
	HasId bool

	// Container controls what to generate for a required container (scope ID,
	// auth method ID, etc.); see ContainerRequiredActions as well
	Container string

	// ContainerRequiredActions controls which actions require that the
	// container ID is not empty.
	ContainerRequiredActions []string

	// HasName controls whether to add name options
	HasName bool

	// HasDescription controls whether to add description options
	HasDescription bool

	// HasScopeName controls whether to add scope name options
	HasScopeName bool

	// HasRecursiveListing controls whether to add in options for recursive
	// listing
	HasRecursiveListing bool

	// VersionedActions controls which actions to add a case for version checking
	VersionedActions []string

	// HasExtraActionsOutput controls whether to generate code to call a
	// function for custom output
	HasExtraActionsOutput bool

	// SubActionPrefix specifies the prefix to use when generating sub-action
	// commands (e.g. "targets update tcp")
	SubActionPrefix string
}

var inputStructs = map[string][]*cmdInfo{
	"accounts": {
		{
			ResourceType:             resource.Account.String(),
			PkgPath:                  "github.com/hashicorp/boundary/api/accounts",
			StdActions:               []string{"read", "delete", "list"},
			HasCustomActionFlags:     true,
			HasExtraCommandVars:      true,
			HasExtraSynopsisFunc:     true,
			HasExtraActions:          true,
			HasExtraFlagsFunc:        true,
			HasExtraHelpFunc:         true,
			IsAbstractType:           true,
			HasExtraFlagHandlingFunc: true,
			Container:                "AuthMethod",
			HasId:                    true,
			HasName:                  true,
			HasDescription:           true,
			ContainerRequiredActions: []string{"list"},
			VersionedActions:         []string{"change-password", "set-password"},
		},
		{
			ResourceType:             resource.Account.String(),
			PkgPath:                  "github.com/hashicorp/boundary/api/accounts",
			StdActions:               []string{"create", "update"},
			HasCustomActionFlags:     true,
			SubActionPrefix:          "password",
			HasExtraCommandVars:      true,
			HasExtraSynopsisFunc:     true,
			SkipNormalHelp:           true,
			HasExtraHelpFunc:         true,
			HasExtraFlagsFunc:        true,
			HasExtraFlagHandlingFunc: true,
			HasId:                    true,
			HasName:                  true,
			Container:                "AuthMethod",
			ContainerRequiredActions: []string{"create"},
			HasDescription:           true,
			VersionedActions:         []string{"update"},
		},
	},
	"groups": {
		{
			ResourceType:             resource.Group.String(),
			PkgPath:                  "github.com/hashicorp/boundary/api/groups",
			StdActions:               []string{"create", "read", "update", "delete", "list"},
			HasCustomActionFlags:     true,
			HasExtraCommandVars:      true,
			HasExtraSynopsisFunc:     true,
			HasExtraActions:          true,
			HasExtraFlagsFunc:        true,
			HasExtraHelpFunc:         true,
			HasExtraFlagHandlingFunc: true,
			HasId:                    true,
			Container:                "Scope",
			ContainerRequiredActions: []string{"create", "list"},
			HasName:                  true,
			HasDescription:           true,
			HasRecursiveListing:      true,
			VersionedActions:         []string{"update", "add-members", "remove-members", "set-members"},
		},
	},
	"targets": {
		{
			ResourceType:             resource.Target.String(),
			PkgPath:                  "github.com/hashicorp/boundary/api/targets",
			StdActions:               []string{"read", "delete", "list"},
			HasCustomActionFlags:     true,
			HasExtraCommandVars:      true,
			HasExtraSynopsisFunc:     true,
			HasExtraActions:          true,
			HasExtraFlagsFunc:        true,
			HasExtraHelpFunc:         true,
			HasExampleCliOutput:      true,
			IsAbstractType:           true,
			HasExtraFlagHandlingFunc: true,
			HasName:                  true,
			HasDescription:           true,
			HasRecursiveListing:      true,
			Container:                "Scope",
			ContainerRequiredActions: []string{"list"},
			VersionedActions:         []string{"add-host-sets", "remove-host-sets", "set-host-sets"},
			HasExtraActionsOutput:    true,
		},
		{
			ResourceType:             resource.Target.String(),
			PkgPath:                  "github.com/hashicorp/boundary/api/targets",
			StdActions:               []string{"create", "update"},
			HasCustomActionFlags:     true,
			SubActionPrefix:          "tcp",
			HasExtraCommandVars:      true,
			HasExtraSynopsisFunc:     true,
			SkipNormalHelp:           true,
			HasExtraHelpFunc:         true,
			HasExtraFlagsFunc:        true,
			HasExtraFlagHandlingFunc: true,
			HasId:                    true,
			HasName:                  true,
			Container:                "Scope",
			ContainerRequiredActions: []string{"create"},
			HasDescription:           true,
			VersionedActions:         []string{"update"},
		},
	},
}
