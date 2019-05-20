/*
Copyright 2018 Openstorage.org

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/libopenstorage/stork/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ApplicationBackups returns a ApplicationBackupInformer.
	ApplicationBackups() ApplicationBackupInformer
	// ApplicationClones returns a ApplicationCloneInformer.
	ApplicationClones() ApplicationCloneInformer
	// ApplicationRestores returns a ApplicationRestoreInformer.
	ApplicationRestores() ApplicationRestoreInformer
	// BackupLocations returns a BackupLocationInformer.
	BackupLocations() BackupLocationInformer
	// ClusterDomainUpdates returns a ClusterDomainUpdateInformer.
	ClusterDomainUpdates() ClusterDomainUpdateInformer
	// ClusterDomainsStatuses returns a ClusterDomainsStatusInformer.
	ClusterDomainsStatuses() ClusterDomainsStatusInformer
	// ClusterPairs returns a ClusterPairInformer.
	ClusterPairs() ClusterPairInformer
	// GroupVolumeSnapshots returns a GroupVolumeSnapshotInformer.
	GroupVolumeSnapshots() GroupVolumeSnapshotInformer
	// Migrations returns a MigrationInformer.
	Migrations() MigrationInformer
	// MigrationSchedules returns a MigrationScheduleInformer.
	MigrationSchedules() MigrationScheduleInformer
	// Rules returns a RuleInformer.
	Rules() RuleInformer
	// SchedulePolicies returns a SchedulePolicyInformer.
	SchedulePolicies() SchedulePolicyInformer
	// StorageClusters returns a StorageClusterInformer.
	StorageClusters() StorageClusterInformer
	// VolumeSnapshotRestores returns a VolumeSnapshotRestoreInformer.
	VolumeSnapshotRestores() VolumeSnapshotRestoreInformer
	// VolumeSnapshotSchedules returns a VolumeSnapshotScheduleInformer.
	VolumeSnapshotSchedules() VolumeSnapshotScheduleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ApplicationBackups returns a ApplicationBackupInformer.
func (v *version) ApplicationBackups() ApplicationBackupInformer {
	return &applicationBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationClones returns a ApplicationCloneInformer.
func (v *version) ApplicationClones() ApplicationCloneInformer {
	return &applicationCloneInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationRestores returns a ApplicationRestoreInformer.
func (v *version) ApplicationRestores() ApplicationRestoreInformer {
	return &applicationRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupLocations returns a BackupLocationInformer.
func (v *version) BackupLocations() BackupLocationInformer {
	return &backupLocationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterDomainUpdates returns a ClusterDomainUpdateInformer.
func (v *version) ClusterDomainUpdates() ClusterDomainUpdateInformer {
	return &clusterDomainUpdateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterDomainsStatuses returns a ClusterDomainsStatusInformer.
func (v *version) ClusterDomainsStatuses() ClusterDomainsStatusInformer {
	return &clusterDomainsStatusInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterPairs returns a ClusterPairInformer.
func (v *version) ClusterPairs() ClusterPairInformer {
	return &clusterPairInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GroupVolumeSnapshots returns a GroupVolumeSnapshotInformer.
func (v *version) GroupVolumeSnapshots() GroupVolumeSnapshotInformer {
	return &groupVolumeSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Migrations returns a MigrationInformer.
func (v *version) Migrations() MigrationInformer {
	return &migrationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MigrationSchedules returns a MigrationScheduleInformer.
func (v *version) MigrationSchedules() MigrationScheduleInformer {
	return &migrationScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Rules returns a RuleInformer.
func (v *version) Rules() RuleInformer {
	return &ruleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SchedulePolicies returns a SchedulePolicyInformer.
func (v *version) SchedulePolicies() SchedulePolicyInformer {
	return &schedulePolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StorageClusters returns a StorageClusterInformer.
func (v *version) StorageClusters() StorageClusterInformer {
	return &storageClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeSnapshotRestores returns a VolumeSnapshotRestoreInformer.
func (v *version) VolumeSnapshotRestores() VolumeSnapshotRestoreInformer {
	return &volumeSnapshotRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeSnapshotSchedules returns a VolumeSnapshotScheduleInformer.
func (v *version) VolumeSnapshotSchedules() VolumeSnapshotScheduleInformer {
	return &volumeSnapshotScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
