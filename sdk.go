// Package potato_api_sdk provides client library for Potato API gRPC services
package potato_api_sdk

import (
	"github.com/sztu-ywy/potato-api-sdk/proto"
)

// Re-export main types from proto package
type (
	// Request/Response types
	PingRequest                  = proto.PingRequest
	PingResponse                 = proto.PingResponse
	AudioStreamRequest           = proto.AudioStreamRequest
	AudioRecognitionResponse     = proto.AudioRecognitionResponse
	DialogueRequest              = proto.DialogueRequest
	DialogueResponse             = proto.DialogueResponse
	UpdateMemoryRequest          = proto.UpdateMemoryRequest
	UpdateMemoryResponse         = proto.UpdateMemoryResponse
	ConversationSummaryRequest   = proto.ConversationSummaryRequest
	ConversationSummaryResponse  = proto.ConversationSummaryResponse
	SaveConversation4RAGRequest  = proto.SaveConversation4RAGRequest
	SaveConversation4RAGResponse = proto.SaveConversation4RAGResponse

	// Configuration types
	InitialConfig   = proto.InitialConfig
	Voiceprint      = proto.Voiceprint
	AudioParams     = proto.AudioParams
	DialogueContext = proto.DialogueContext
	DeivceState     = proto.DeivceState
	RoleList        = proto.RoleList
	AsrResult       = proto.AsrResult

	// Service interfaces
	PotatoChatServiceClient              = proto.PotatoChatServiceClient
	PotatoChatServiceServer              = proto.PotatoChatServiceServer
	UnimplementedPotatoChatServiceServer = proto.UnimplementedPotatoChatServiceServer
)

// Service constructor functions
var (
	NewPotatoChatServiceClient      = proto.NewPotatoChatServiceClient
	RegisterPotatoChatServiceServer = proto.RegisterPotatoChatServiceServer
)
