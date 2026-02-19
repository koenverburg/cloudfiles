// Dagger module for validating Ansible playbooks locally.
//
// This module provides functions to test each machine configuration
// (control-plane, worker-node, ai-node, and shared) in isolated Ubuntu containers.

package main

import (
	"context"
	"dagger/cloudfiles/internal/dagger"
	"fmt"
)

type Cloudfiles struct{}

// baseContainer returns a base Ubuntu container with Ansible installed
func (m *Cloudfiles) baseContainer() *dagger.Container {
	return dag.Container().
		From("ubuntu:22.04").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "python3", "python3-pip", "sudo", "curl", "git", "build-essential", "unzip"}).
		WithExec([]string{"pip3", "install", "ansible"})
}

// Shared runs the shared setup playbook in an Ubuntu container
func (m *Cloudfiles) Shared(
	ctx context.Context,
	// Source directory containing the playbooks (defaults to current directory)
	// +optional
	source *dagger.Directory,
) (string, error) {
	if source == nil {
		source = dag.CurrentModule().Source().Directory("..")
	}

	return m.baseContainer().
		WithDirectory("/workspace", source).
		WithWorkdir("/workspace").
		WithExec([]string{
			"ansible-playbook",
			"-i", "localhost,",
			"-c", "local",
			"machines/shared/setup.yml",
		}).
		Stdout(ctx)
}

// ControlPlane runs the control-plane setup playbook in an Ubuntu container
func (m *Cloudfiles) ControlPlane(
	ctx context.Context,
	// Source directory containing the playbooks (defaults to current directory)
	// +optional
	source *dagger.Directory,
) (string, error) {
	if source == nil {
		source = dag.CurrentModule().Source().Directory("..")
	}

	return m.baseContainer().
		WithDirectory("/workspace", source).
		WithWorkdir("/workspace").
		WithExec([]string{
			"ansible-playbook",
			"-i", "localhost,",
			"-c", "local",
			"machines/control-plane/setup.yml",
		}).
		Stdout(ctx)
}

// WorkerNode runs the worker-node setup playbook in an Ubuntu container
func (m *Cloudfiles) WorkerNode(
	ctx context.Context,
	// Source directory containing the playbooks (defaults to current directory)
	// +optional
	source *dagger.Directory,
	// Consul server address
	// +optional
	// +default="127.0.0.1"
	consulServer string,
) (string, error) {
	if source == nil {
		source = dag.CurrentModule().Source().Directory("..")
	}

	if consulServer == "" {
		consulServer = "127.0.0.1"
	}

	return m.baseContainer().
		WithDirectory("/workspace", source).
		WithWorkdir("/workspace").
		WithExec([]string{
			"ansible-playbook",
			"-i", "localhost,",
			"-c", "local",
			"-e", fmt.Sprintf("consul_server=%s", consulServer),
			"machines/worker-node/setup.yml",
		}).
		Stdout(ctx)
}

// AiNode runs the ai-node setup playbook in an Ubuntu container
func (m *Cloudfiles) AiNode(
	ctx context.Context,
	// Source directory containing the playbooks (defaults to current directory)
	// +optional
	source *dagger.Directory,
	// Consul server address
	// +optional
	// +default="127.0.0.1"
	consulServer string,
) (string, error) {
	if source == nil {
		source = dag.CurrentModule().Source().Directory("..")
	}

	if consulServer == "" {
		consulServer = "127.0.0.1"
	}

	return m.baseContainer().
		WithDirectory("/workspace", source).
		WithWorkdir("/workspace").
		WithExec([]string{
			"ansible-playbook",
			"-i", "localhost,",
			"-c", "local",
			"-e", fmt.Sprintf("consul_server=%s", consulServer),
			"machines/ai-node/setup.yml",
		}).
		Stdout(ctx)
}
