/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeployments implements DeploymentInterface
type FakeDeployments struct {
	Fake *FakeAppsV1
	ns   string
	te   string
}

var deploymentsResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

var deploymentsKind = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *FakeDeployments) Get(name string, options v1.GetOptions) (result *appsv1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(deploymentsResource, c.ns, name, c.te), &appsv1.Deployment{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *FakeDeployments) List(opts v1.ListOptions) (result *appsv1.DeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(deploymentsResource, deploymentsKind, c.ns, opts, c.te), &appsv1.DeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.DeploymentList{ListMeta: obj.(*appsv1.DeploymentList).ListMeta}
	for _, item := range obj.(*appsv1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *FakeDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(deploymentsResource, c.ns, opts, c.te))

}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Create(deployment *appsv1.Deployment) (result *appsv1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(deploymentsResource, c.ns, deployment, c.te), &appsv1.Deployment{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.Deployment), err
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Update(deployment *appsv1.Deployment) (result *appsv1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(deploymentsResource, c.ns, deployment, c.te), &appsv1.Deployment{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.Deployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeployments) UpdateStatus(deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(deploymentsResource, "status", c.ns, deployment, c.te), &appsv1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *FakeDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(deploymentsResource, c.ns, name, c.te), &appsv1.Deployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(deploymentsResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &appsv1.DeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deployment.
func (c *FakeDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *appsv1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(deploymentsResource, c.te, c.ns, name, pt, data, subresources...), &appsv1.Deployment{})

	if obj == nil {
		return nil, err
	}

	return obj.(*appsv1.Deployment), err
}

// GetScale takes name of the deployment, and returns the corresponding scale object, and an error if there is any.
func (c *FakeDeployments) GetScale(deploymentName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithMultiTenancy(deploymentsResource, c.ns, "scale", deploymentName, c.te), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}

	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeDeployments) UpdateScale(deploymentName string, scale *autoscalingv1.Scale) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(deploymentsResource, "scale", c.ns, scale, c.te), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}

	return obj.(*autoscalingv1.Scale), err
}