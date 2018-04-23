// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package job

import (
	clientutil "openpitrix.io/openpitrix/pkg/client"
	clusterclient "openpitrix.io/openpitrix/pkg/client/cluster"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/plugins"
)

type Processor struct {
	Job *models.Job
}

func NewProcessor(job *models.Job) *Processor {
	return &Processor{
		Job: job,
	}
}

// Pre process when job is start
func (j *Processor) Pre() error {
	var err error
	ctx := clientutil.GetSystemUserContext()
	client, err := clusterclient.NewClusterManagerClient(ctx)
	if err != nil {
		logger.Errorf("Executing job [%s] pre processor failed: %+v", j.Job.JobId, err)
		return err
	}
	switch j.Job.JobAction {
	case constants.ActionCreateCluster:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusCreating,
			}),
		})
	case constants.ActionUpgradeCluster:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusUpgrading,
			}),
		})
	case constants.ActionRollbackCluster:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusRollbacking,
			}),
		})
	case constants.ActionResizeCluster:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusResizing,
			}),
		})
	case constants.ActionAddClusterNodes:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusScaling,
			}),
		})
	case constants.ActionDeleteClusterNodes:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusScaling,
			}),
		})
	case constants.ActionStopClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{

			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusStopping,
			}),
		})
	case constants.ActionStartClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusStarting,
			}),
		})
	case constants.ActionDeleteClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusDeleting,
			}),
		})
	case constants.ActionRecoverClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusRecovering,
			}),
		})
	case constants.ActionCeaseClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusCeasing,
			}),
		})
	case constants.ActionUpdateClusterEnv:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId:        j.Job.ClusterId,
				TransitionStatus: constants.StatusUpdating,
			}),
		})
	default:
		logger.Errorf("Unknown job action [%s]", j.Job.JobAction)
	}

	if err != nil {
		logger.Panicf("Executing job [%s] pre processor failed: %+v", j.Job.JobId, err)
	}
	return err
}

// Post process when job is done
func (j *Processor) Post() error {
	var err error
	ctx := clientutil.GetSystemUserContext()
	client, err := clusterclient.NewClusterManagerClient(ctx)
	if err != nil {
		logger.Errorf("Executing job [%s] post processor failed: %+v", j.Job.JobId, err)
		return err
	}
	switch j.Job.JobAction {
	case constants.ActionCreateCluster:
		providerInterface, err := plugins.GetProviderPlugin(j.Job.Provider)
		if err != nil {
			logger.Errorf("No such provider [%s]. ", j.Job.Provider)
			return err
		}
		err = providerInterface.UpdateClusterStatus(j.Job)
		if err != nil {
			logger.Errorf("Executing job [%s] post processor failed: %+v", j.Job.JobId, err)
			return err
		}

		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionUpgradeCluster:
		providerInterface, err := plugins.GetProviderPlugin(j.Job.Provider)
		if err != nil {
			logger.Errorf("No such provider [%s]. ", j.Job.Provider)
			return err
		}
		err = providerInterface.UpdateClusterStatus(j.Job)
		if err != nil {
			logger.Errorf("Executing job [%s] post processor failed: %+v", j.Job.JobId, err)
			return err
		}

		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionRollbackCluster:
		providerInterface, err := plugins.GetProviderPlugin(j.Job.Provider)
		if err != nil {
			logger.Errorf("No such provider [%s]. ", j.Job.Provider)
			return err
		}
		err = providerInterface.UpdateClusterStatus(j.Job)
		if err != nil {
			logger.Errorf("Executing job [%s] post processor failed: %+v", j.Job.JobId, err)
			return err
		}

		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionResizeCluster:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionAddClusterNodes:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionDeleteClusterNodes:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionStopClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusStopped,
			}),
		})
	case constants.ActionStartClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionDeleteClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusDeleted,
			}),
		})
	case constants.ActionRecoverClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	case constants.ActionCeaseClusters:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusCeased,
			}),
		})
	case constants.ActionUpdateClusterEnv:
		_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
			Cluster: models.ClusterToPb(&models.Cluster{
				ClusterId: j.Job.ClusterId,
				Status:    constants.StatusActive,
			}),
		})
	default:
		logger.Errorf("Unknown job action [%s]", j.Job.JobAction)
	}

	if err != nil {
		logger.Errorf("Executing job [%s] post processor failed: %+v", j.Job.JobId, err)
	}
	return err
}

func (j *Processor) Final() {
	ctx := clientutil.GetSystemUserContext()
	client, err := clusterclient.NewClusterManagerClient(ctx)
	if err != nil {
		logger.Errorf("Executing job [%s] final processor failed: %+v", j.Job.JobId, err)
		return
	}
	_, err = client.ModifyCluster(ctx, &pb.ModifyClusterRequest{
		Cluster: models.ClusterToPb(&models.Cluster{
			ClusterId:        j.Job.ClusterId,
			TransitionStatus: "",
		}),
	})
	if err != nil {
		logger.Errorf("Executing job [%s] final processor failed: %+v", j.Job.JobId, err)
	}
}
