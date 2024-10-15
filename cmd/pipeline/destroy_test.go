// Copyright 2023 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipeline

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/okteto/okteto/internal/test/client"
	"github.com/okteto/okteto/pkg/okteto"
	"github.com/okteto/okteto/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestDestroyPipelineSuccessful(t *testing.T) {
	ctx := context.Background()
	response := &client.FakePipelineResponses{
		DestroyResponse: &types.GitDeployResponse{
			Action: &types.Action{
				ID:   "test",
				Name: "test",
			},
		},
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
		},
		okCtx: &okteto.Context{},
	}
	opts := &DestroyOptions{
		Name: "test",
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.NoError(t, err)
}

func TestDestroyPipelineSuccessfulWithWait(t *testing.T) {
	ctx := context.Background()
	response := &client.FakePipelineResponses{
		DestroyResponse: &types.GitDeployResponse{
			Action: &types.Action{
				ID:   "test",
				Name: "test",
			},
		},
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
			StreamClient:   client.NewFakeStreamClient(&client.FakeStreamResponse{}),
		},
		okCtx: &okteto.Context{},
	}
	opts := &DestroyOptions{
		Name: "test",
		Wait: true,
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.NoError(t, err)
}

func TestDestroyPipelineSuccessfulWithWaitStreamErr(t *testing.T) {
	ctx := context.Background()
	response := &client.FakePipelineResponses{
		DestroyResponse: &types.GitDeployResponse{
			Action: &types.Action{
				ID:   "test",
				Name: "test",
			},
		},
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
			StreamClient:   client.NewFakeStreamClient(&client.FakeStreamResponse{StreamErr: errors.New("error")}),
		},
		okCtx: &okteto.Context{},
	}
	opts := &DestroyOptions{
		Name: "test",
		Wait: true,
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.NoError(t, err)
}

func TestDestroyNonExistentPipeline(t *testing.T) {
	ctx := context.Background()
	response := &client.FakePipelineResponses{
		DestroyErr: fmt.Errorf("not found"),
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
		},
		okCtx: &okteto.Context{},
	}
	opts := &DestroyOptions{
		Name: "no exists",
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.NoError(t, err)
}

func TestDestroyNonExistentPipelineWithWait(t *testing.T) {
	ctx := context.Background()
	response := &client.FakePipelineResponses{
		DestroyErr: fmt.Errorf("not found"),
		DestroyResponse: &types.GitDeployResponse{
			Action: &types.Action{},
		},
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
			StreamClient:   client.NewFakeStreamClient(&client.FakeStreamResponse{}),
		},
		okCtx: &okteto.Context{},
	}

	opts := &DestroyOptions{
		Name: "no exists",
		Wait: true,
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.NoError(t, err)
}

func TestDestroyExistentPipelineWithError(t *testing.T) {
	ctx := context.Background()
	pipelineErr := fmt.Errorf("test error")
	response := &client.FakePipelineResponses{
		DestroyErr: pipelineErr,
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
		},
		okCtx: &okteto.Context{},
	}

	opts := &DestroyOptions{
		Name: "no exists",
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.ErrorIs(t, err, pipelineErr)
}

func TestDestroyExistentPipelineTimeoutError(t *testing.T) {
	ctx := context.Background()
	pipelineErr := fmt.Errorf("test error")
	response := &client.FakePipelineResponses{
		DestroyResponse: &types.GitDeployResponse{
			Action: &types.Action{},
		},
		WaitErr: pipelineErr,
	}
	pc := &Command{
		okClient: &client.FakeOktetoClient{
			PipelineClient: client.NewFakePipelineClient(response),
			StreamClient:   client.NewFakeStreamClient(&client.FakeStreamResponse{}),
		},
		okCtx: &okteto.Context{},
	}

	opts := &DestroyOptions{
		Name:    "test",
		Wait:    true,
		Timeout: 0 * time.Second,
	}
	err := pc.ExecuteDestroyPipeline(ctx, opts)
	assert.ErrorIs(t, err, pipelineErr)
}
