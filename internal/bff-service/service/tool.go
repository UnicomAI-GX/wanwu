// @Author wangxm 10/24/星期五 14:46:00
package service

import (
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	errs "github.com/UnicomAI/wanwu/api/proto/err-code"
	mcp_service "github.com/UnicomAI/wanwu/api/proto/mcp-service"
	"github.com/UnicomAI/wanwu/internal/bff-service/model/request"
	"github.com/UnicomAI/wanwu/internal/bff-service/model/response"
	"github.com/UnicomAI/wanwu/internal/bff-service/pkg/util"
	"github.com/UnicomAI/wanwu/pkg/constant"
	grpc_util "github.com/UnicomAI/wanwu/pkg/grpc-util"
	openapi3_util "github.com/UnicomAI/wanwu/pkg/openapi3-util"
	"github.com/gin-gonic/gin"
)

func GetToolSelect(ctx *gin.Context, userID, orgID string, name string) (*response.ListResult, error) {
	resp, err := mcp.GetToolSelect(ctx.Request.Context(), &mcp_service.GetToolSelectReq{
		Name: name,
		Identity: &mcp_service.Identity{
			UserId: userID,
			OrgId:  orgID,
		},
	})
	if err != nil {
		return nil, err
	}

	var list []response.ToolSelect
	for _, item := range resp.List {
		list = append(list, response.ToolSelect{
			UniqueId: util.ConcatAssistantToolUniqueId("tool-", item.ToolId),
			ToolInfo: response.ToolInfo{
				ToolId:          item.ToolId,
				ToolName:        item.ToolName,
				ToolType:        item.ToolType,
				Desc:            item.Desc,
				APIKey:          item.ApiKey,
				NeedApiKeyInput: item.NeedApiKeyInput,
			},
		})
	}
	return &response.ListResult{
		List:  list,
		Total: int64(len(list)),
	}, nil
}

func GetToolActionList(ctx *gin.Context, userID, orgID string, req request.ToolActionListReq) (*response.ToolActionList, error) {
	schema, _, _, err := getToolSchema(ctx, userID, orgID, req.ToolId, req.ToolType)
	if err != nil {
		return nil, err
	}

	// 加载schema
	var actions []*protocol.Tool
	doc, err := openapi3_util.LoadFromData(ctx.Request.Context(), []byte(schema))
	if err != nil {
		return nil, grpc_util.ErrorStatus(errs.Code_BFFInvalidArg, err.Error())
	}
	if err := openapi3_util.ValidateDoc(ctx.Request.Context(), doc); err != nil {
		return nil, grpc_util.ErrorStatus(errs.Code_BFFInvalidArg, err.Error())
	}

	// Operation2ProtocolTool
	for _, pathItem := range doc.Paths.Map() {
		for _, operation := range pathItem.Operations() {
			action := openapi3_util.Operation2ProtocolTool(operation)
			actions = append(actions, action)
		}
	}

	return &response.ToolActionList{
		Actions: actions,
	}, nil
}

func GetToolActionDetail(ctx *gin.Context, userID, orgID string, req request.ToolActionReq) (*response.ToolActionDetail, error) {
	schema, needApiKeyInput, apikey, err := getToolSchema(ctx, userID, orgID, req.ToolId, req.ToolType)
	if err != nil {
		return nil, err
	}

	// Schema2ProtocolTool
	action, err := openapi3_util.Schema2ProtocolTool(ctx.Request.Context(), []byte(schema), req.ActionName)
	if err != nil {
		return nil, grpc_util.ErrorStatus(errs.Code_BFFInvalidArg, err.Error())
	}

	// apikey
	return &response.ToolActionDetail{
		NeedApiKeyInput: needApiKeyInput,
		APIKey:          apikey,
		Action:          action,
	}, nil
}

// --- internal ---
func getToolSchema(ctx *gin.Context, userID, orgID, toolID string, toolType string) (string, bool, string, error) {
	if toolType == constant.ToolTypeBuiltIn {
		// 获取内置工具详情
		resp, err := mcp.GetSquareTool(ctx.Request.Context(), &mcp_service.GetSquareToolReq{
			ToolSquareId: toolID,
			Identity: &mcp_service.Identity{
				UserId: userID,
				OrgId:  orgID,
			},
		})
		if err != nil {
			return "", false, "", err
		}

		return resp.Schema, resp.BuiltInTools.NeedApiKeyInput, resp.BuiltInTools.ApiKey, nil
	} else if toolType == constant.ToolTypeCustom {
		// 获取自定义工具详情
		resp, err := mcp.GetCustomToolInfo(ctx.Request.Context(), &mcp_service.GetCustomToolInfoReq{
			CustomToolId: toolID,
			Identity: &mcp_service.Identity{
				UserId: userID,
				OrgId:  orgID,
			},
		})
		if err != nil {
			return "", false, "", err
		}
		return resp.Schema, false, "", nil
	}
	return "", false, "", nil
}
