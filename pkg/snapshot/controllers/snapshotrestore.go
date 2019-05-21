package controllers

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/libopenstorage/stork/drivers/volume"
	"github.com/libopenstorage/stork/pkg/apis/stork"
	stork_api "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"github.com/libopenstorage/stork/pkg/controller"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/portworx/sched-ops/k8s"
	v1 "k8s.io/api/core/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"
)

// SnapshotRestoreController controller to watch over In-Place snap restore CRD's
type SnapshotRestoreController struct {
	Driver   volume.Driver
	Recorder record.EventRecorder
}

// Init initialize the cluster pair controller
func (c *SnapshotRestoreController) Init() error {
	err := c.createCRD()
	if err != nil {
		return err
	}

	return controller.Register(
		&schema.GroupVersionKind{
			Group:   stork.GroupName,
			Version: stork_api.SchemeGroupVersion.Version,
			Kind:    reflect.TypeOf(stork_api.VolumeSnapshotRestore{}).Name(),
		},
		"",
		1*time.Minute,
		c)
}

// Handle updates for SnapshotRestore objects
func (c *SnapshotRestoreController) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *stork_api.VolumeSnapshotRestore:
		snapRestore := o
		if snapRestore.Spec.SourceName == "" || snapRestore.Spec.SourceType == "" {
			return fmt.Errorf("empty Snapshot name/type or namespace")
		}

		if event.Deleted {
			return nil
		}

		if snapRestore.Status.Status == stork_api.SnapshotRestoreStatusReady {
			return nil
		}

		err := c.Driver.VolumeSnapshotRestore(snapRestore)
		if err != nil {
			snapRestore.Status.Status = stork_api.SnapshotRestoreStatusError
			return fmt.Errorf("failed to restore pvc details %v", err)
		}

		snapRestore.Status.Status = stork_api.SnapshotRestoreStatusReady
		c.Recorder.Event(snapRestore,
			v1.EventTypeNormal,
			string(snapRestore.Status.Status),
			"Snapshot In-Place  Restore completed(dummy)")

		err = sdk.Update(snapRestore)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SnapshotRestoreController) createCRD() error {
	resource := k8s.CustomResource{
		Name:    stork_api.SnapshotRestoreResourceName,
		Plural:  stork_api.SnapshotRestoreResourcePlural,
		Group:   stork.GroupName,
		Version: stork_api.SchemeGroupVersion.Version,
		Scope:   apiextensionsv1beta1.NamespaceScoped,
		Kind:    reflect.TypeOf(stork_api.VolumeSnapshotRestore{}).Name(),
	}
	err := k8s.Instance().CreateCRD(resource)
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}

	return k8s.Instance().ValidateCRD(resource, validateCRDTimeout, validateCRDInterval)
}
