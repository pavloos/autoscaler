/*
Copyright 2016 The Kubernetes Authors.

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

package providers

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/autoscaler/cluster-autoscaler/processors/nodegroupconfig"
)

// NewDefaultMaxNodeProvisionTimeProvider returns the default maxNodeProvisionTimeProvider which uses the NodeGroupConfigProcessor.
func NewDefaultMaxNodeProvisionTimeProvider(context *context.AutoscalingContext, nodeGroupConfigProcessor nodegroupconfig.NodeGroupConfigProcessor) *defultMaxNodeProvisionTimeProvider {
	return &defultMaxNodeProvisionTimeProvider{context: context, nodeGroupConfigProcessor: nodeGroupConfigProcessor}
}

type defultMaxNodeProvisionTimeProvider struct {
	context                  *context.AutoscalingContext
	nodeGroupConfigProcessor nodegroupconfig.NodeGroupConfigProcessor
}

// GetMaxNodeProvisionTime returns MaxNodeProvisionTime value that should be used for the given NodeGroup.
func (p *defultMaxNodeProvisionTimeProvider) GetMaxNodeProvisionTime(nodeGroup cloudprovider.NodeGroup) (time.Duration, error) {
	return p.nodeGroupConfigProcessor.GetMaxNodeProvisionTime(p.context, nodeGroup)
}
