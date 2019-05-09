package storkctl

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func newCreateCommand(cmdFactory Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	createCommands := &cobra.Command{
		Use:   "create",
		Short: "Create stork resources",
	}

	createCommands.AddCommand(
		newCreateSnapshotCommand(cmdFactory, ioStreams),
		newRestoreSnapshotCommand(cmdFactory, ioStreams),
		newGetSnapshotRestoreCommand(cmdFactory, ioStreams),
		newListSnapshotRestoreCommand(cmdFactory, ioStreams),
		newDeleteSnapshotRestoreCommand(cmdFactory, ioStreams),
		newCreateMigrationCommand(cmdFactory, ioStreams),
		newCreateMigrationScheduleCommand(cmdFactory, ioStreams),
		newCreatePVCCommand(cmdFactory, ioStreams),
		newCreateSnapshotScheduleCommand(cmdFactory, ioStreams),
		newCreateGroupSnapshotCommand(cmdFactory, ioStreams),
	)

	return createCommands
}
