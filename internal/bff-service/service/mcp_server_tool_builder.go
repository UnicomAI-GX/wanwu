package service

import (
	"context"

	mcp_service "github.com/UnicomAI/wanwu/api/proto/mcp-service"
	"github.com/UnicomAI/wanwu/internal/bff-service/model/request"
	mcp_util "github.com/UnicomAI/wanwu/internal/bff-service/pkg/mcp-util"
	"github.com/UnicomAI/wanwu/pkg/constant"
)

type mcpServerToolBuilder interface {
	MCPServerToolType() string
	AppID() string
	AppName() string
	GetOpenapiSchema(ctx context.Context) (string, *mcp_util.APIAuth, error)
}

// --- mcpServerCustomToolBuilder ---

type mcpServerCustomToolBuilder struct {
	customToolID   string
	customToolName string
}

func (mcpServerCustomToolBuilder) MCPServerToolType() string {
	return constant.MCPServerToolTypeCustomTool
}

func (builder *mcpServerCustomToolBuilder) AppID() string {
	return builder.customToolID
}

func (builder *mcpServerCustomToolBuilder) AppName() string {
	return builder.customToolName
}

func (builder *mcpServerCustomToolBuilder) GetOpenapiSchema(ctx context.Context) (string, *mcp_util.APIAuth, error) {
	customToolInfo, err := mcp.GetCustomToolInfo(ctx, &mcp_service.GetCustomToolInfoReq{
		CustomToolId: builder.customToolID,
	})
	if err != nil {
		return "", nil, err
	}
	builder.customToolName = customToolInfo.Name
	return customToolInfo.Schema, convertToolApiAuth(customToolInfo.ApiAuth), nil
}

// --- mcpServerOpenapiSchemaBuilder ---

type mcpServerOpenapiSchemaBuilder struct {
	schema string
	name   string
	auth   request.CustomToolApiAuthWebRequest
}

func (mcpServerOpenapiSchemaBuilder) MCPServerToolType() string {
	return constant.MCPServerToolTypeOpenAPI
}

func (builder *mcpServerOpenapiSchemaBuilder) AppID() string {
	return ""
}

func (builder *mcpServerOpenapiSchemaBuilder) AppName() string {
	return builder.name
}

func (builder *mcpServerOpenapiSchemaBuilder) GetOpenapiSchema(ctx context.Context) (string, *mcp_util.APIAuth, error) {
	auth := &mcp_util.APIAuth{}
	if builder.auth.Type != "" && builder.auth.Type != "None" {
		auth.Type = "API Key"
		auth.In = "header"
		if builder.auth.AuthType == "Custom" {
			if builder.auth.CustomHeaderName != "" {
				auth.Name = builder.auth.CustomHeaderName
				auth.Value = builder.auth.APIKey
			}
		} else {
			auth.Name = "Authorization"
			auth.Value = "Bearer " + builder.auth.APIKey
		}
	}
	return builder.schema, auth, nil
}
