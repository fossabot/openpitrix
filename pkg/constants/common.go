// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package constants

import "time"

const (
	prefix             = "openpitrix-"
	ApiGatewayHost     = prefix + "api-gateway"
	RepoManagerHost    = prefix + "repo-manager"
	AppManagerHost     = prefix + "app-manager"
	RuntimeManagerHost = prefix + "runtime-manager"
	ClusterManagerHost = prefix + "cluster-manager"
	JobManagerHost     = prefix + "job-manager"
	TaskManagerHost    = prefix + "task-manager"
	PilotManagerHost   = prefix + "pilot-manager"
	RepoIndexerHost    = prefix + "repo-indexer"
)

const (
	ApiGatewayPort       = 9100 // 91 is similar as Pi, Open[Pi]trix
	RepoManagerPort      = 9101
	AppManagerPort       = 9102
	RuntimeManagerPort   = 9103
	ClusterManagerPort   = 9104
	JobManagerPort       = 9106
	TaskManagerPort      = 9107
	RepoIndexerPort      = 9108
	PilotManagerPort     = 9110
	FrontgateServicePort = 9111
	DroneServicePort     = 9112
)

const (
	StatusActive      = "active"
	StatusCreating    = "creating"
	StatusDeleted     = "deleted"
	StatusDeleting    = "deleting"
	StatusUpgrading   = "upgrading"
	StatusUpdating    = "updating"
	StatusRollbacking = "rollbacking"
	StatusStopped     = "stopped"
	StatusStopping    = "stopping"
	StatusStarting    = "starting"
	StatusRecovering  = "recovering"
	StatusCeased      = "ceased"
	StatusCeasing     = "ceasing"
	StatusResizing    = "resizing"
	StatusScaling     = "scaling"
	StatusWorking     = "working"
	StatusPending     = "pending"
	StatusSuccessful  = "successful"
	StatusFailed      = "failed"
)

const (
	JobLength       = 20
	TaskLength      = 20
	RepoEventLength = 20
	InstanceSize    = 20

	DefaultMountPoint = "/data"

	Ext4FileSystem = "ext4"
	XfsFileSystem  = "xfs"

	DefaultExt4MountOption = "defaults,noatime"
	DefaultXfsMountOption  = "rw,noatime,inode64,allocsize=16m"
)

const (
	WaitTaskTimeout  = 600 * time.Second
	WaitTaskInterval = 10 * time.Second

	TimeoutName           = "Timeout"
	DefaultServiceTimeout = 600
)

const (
	ActionCreateCluster      = "CreateCluster"
	ActionUpgradeCluster     = "UpgradeCluster"
	ActionRollbackCluster    = "RollbackCluster"
	ActionResizeCluster      = "ResizeCluster"
	ActionAddClusterNodes    = "AddClusterNodes"
	ActionDeleteClusterNodes = "DeleteClusterNodes"
	ActionStopClusters       = "StopClusters"
	ActionStartClusters      = "StartClusters"
	ActionDeleteClusters     = "DeleteClusters"
	ActionRecoverClusters    = "RecoverClusters"
	ActionCeaseClusters      = "CeaseClusters"
	ActionUpdateClusterEnv   = "UpdateClusterEnv"
)

const (
	ProviderQingCloud  = "qingcloud"
	ProviderKubernetes = "kubernetes"
	TargetPilot        = "pilot"
)

var VmBaseProviders = []string{ProviderQingCloud}

const (
	PlaceHolder       = "*"
	ReplicaRoleSuffix = "-replica"
)

const (
	NodesToExecuteOnName    = "nodes_to_execute_on"
	PostStartServiceName    = "post_start_service"
	PostStopServiceName     = "post_stop_service"
	AgentInstalledName      = "agent_installed"
	ServiceOrderName        = "order"
	ServiceTimeoutName      = "timeout"
	ServiceCmdName          = "cmd"
	ScalingPolicyParallel   = "parallel"
	ScalingPolicySequential = "sequential"

	NormalClusterType    = 0
	FrontgateClusterType = 1

	ServiceInit           = "init"
	ServiceStart          = "start"
	ServiceStop           = "stop"
	ServiceScaleIn        = "scale_in"
	ServiceScaleOut       = "scale_out"
	ServiceCustom         = "custom_service"
	ServiceRestart        = "restart"
	ServiceDestroy        = "destroy"
	ServiceBackup         = "backup"
	ServiceRestore        = "restore"
	ServiceDeleteSnapshot = "delete_snapshot"
	ServiceUpgrade        = "upgrade"
)

var ServiceNames = []string{
	ServiceInit, ServiceStart, ServiceStop, ServiceScaleIn, ServiceScaleOut, ServiceRestart,
	ServiceDestroy, ServiceBackup, ServiceRestore, ServiceDeleteSnapshot, ServiceUpgrade,
}
