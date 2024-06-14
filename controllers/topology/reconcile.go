package topology

import (
	"context"

	clabernetesapisv1alpha1 "github.com/srl-labs/clabernetes/apis/v1alpha1"
	apimachineryerrors "k8s.io/apimachinery/pkg/api/errors"
	ctrlruntime "sigs.k8s.io/controller-runtime"
)

// Reconcile handles reconciliation for this controller.
func (c *Controller) Reconcile(
	ctx context.Context,
	req ctrlruntime.Request,
) (ctrlruntime.Result, error) {
	c.BaseController.LogReconcileStart(req)

	topology, err := c.getTopologyFromReq(ctx, req)
	if err != nil {
		if apimachineryerrors.IsNotFound(err) {
			// was deleted, nothing to do
			c.BaseController.LogReconcileCompleteObjectNotExist(req)

			return ctrlruntime.Result{}, nil
		}

		c.BaseController.LogReconcileFailedGettingObject(req, err)

		return ctrlruntime.Result{}, err
	}

	if topology.DeletionTimestamp != nil {
		// deleting nothing to do, we have no finalizers or anything at this point
		return ctrlruntime.Result{}, nil
	}

	if c.BaseController.ShouldIgnoreReconcile(topology) {
		return ctrlruntime.Result{}, nil
	}

	// we always reconcile the "namespace" resources first -- meaning the resources that exist in
	// the namespace that are not 1:1 to a Topology -- for example: service account and role
	// binding. These resources are created for the namespace on creation of the first Topology in
	// the namespace, and are removed when the last Topology is removed from the namespace.
	err = c.TopologyReconciler.ReconcileNamespaceResources(ctx, topology)
	if err != nil {
		return ctrlruntime.Result{}, err
	}

	reconcileData, err := NewReconcileData(topology)
	if err != nil {
		c.BaseController.Log.Criticalf(
			"failed processing previously stored containerlab resource, error: %s", err,
		)

		return ctrlruntime.Result{}, err
	}

	// reconcile the naming -- we *must* do this to ensure that our status field is set!
	c.TopologyReconciler.ReconcileNaming(topology, reconcileData)

	err = c.processDefinition(topology, reconcileData)
	if err != nil {
		c.BaseController.Log.Criticalf("failed processing topology definition, error: %s", err)

		return ctrlruntime.Result{}, err
	}

	err = c.reconcileResources(ctx, topology, reconcileData)
	if err != nil {
		return ctrlruntime.Result{}, err
	}

	if reconcileData.ShouldUpdateResource {
		// we should update because config hash or something changed, so snag the updated status
		// data out of the reconcile data, put it in the resource, and push the update
		err = reconcileData.SetStatus(topology.Status)
		if err != nil {
			c.BaseController.Log.Criticalf(
				"failed setting object '%s/%s' status, error: %s",
				topology.Namespace,
				topology.Name,
				err,
			)

			return ctrlruntime.Result{}, err
		}

		err = c.BaseController.Client.Update(ctx, topology)
		if err != nil {
			c.BaseController.Log.Criticalf(
				"failed updating object '%s/%s' error: %s",
				topology.Namespace,
				topology.Name,
				err,
			)

			return ctrlruntime.Result{}, err
		}
	}

	c.BaseController.LogReconcileCompleteSuccess(req)

	return ctrlruntime.Result{}, nil
}

func (c *Controller) reconcileResources(
	ctx context.Context,
	topology *clabernetesapisv1alpha1.Topology,
	reconcileData *ReconcileData,
) error {
	err := c.TopologyReconciler.ReconcileConfigMap(
		ctx,
		topology,
		reconcileData,
	)
	if err != nil {
		c.BaseController.Log.Criticalf(
			"failed reconciling clabernetes config map, error: %s",
			err,
		)

		return err
	}

	err = c.TopologyReconciler.ReconcileConnectivity(
		ctx,
		topology,
		reconcileData,
	)
	if err != nil {
		c.BaseController.Log.Criticalf(
			"failed reconciling clabernetes connectivity resource, error: %s",
			err,
		)

		return err
	}

	err = c.TopologyReconciler.ReconcileServices(
		ctx,
		topology,
		reconcileData,
	)
	if err != nil {
		// error already logged
		return err
	}

	err = c.TopologyReconciler.ReconcilePersistentVolumeClaim(
		ctx,
		topology,
		reconcileData,
	)
	if err != nil {
		c.BaseController.Log.Criticalf("failed reconciling clabernetes pvcs, error: %s", err)

		return err
	}

	err = c.TopologyReconciler.ReconcileDeployments(
		ctx,
		topology,
		reconcileData,
	)
	if err != nil {
		c.BaseController.Log.Criticalf("failed reconciling clabernetes deployments, error: %s", err)

		return err
	}

	return nil
}
