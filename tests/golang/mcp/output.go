package mcp

import (
	"net/url"
)

// The sender or recipient of messages and data in a conversation.
type Role string

const (
	RoleAssistant Role = "assistant"
	RoleUser      Role = "user"
)

// Optional annotations for the client. The client can use annotations to inform how objects are used or displayed
type Annotations struct {
	// Describes who the intended audience of this object or data is.
	//
	// It can include multiple entries to indicate content useful for multiple audiences (e.g., `["user", "assistant"]`).
	Audience []Role `json:"audience,omitempty"`
	// The moment the resource was last modified, as an ISO 8601 formatted string.
	//
	// Should be an ISO 8601 formatted string (e.g., "2025-01-12T15:00:58Z").
	//
	// Examples: last activity timestamp in an open file, timestamp when the resource
	// was attached, etc.
	LastModified *string `json:"lastModified,omitempty"`
	// Describes how important this data is for operating the server.
	//
	// A value of 1 means "most important," and indicates that the data is
	// effectively required, while 0 means "least important," and indicates that
	// the data is entirely optional.
	Priority *float64 `json:"priority,omitempty"`
}

// Audio provided to or from an LLM.
type AudioContent struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	// The base64-encoded audio data.
	Data []byte `json:"data"`
	// The MIME type of the audio. Different providers may support different audio types.
	MimeType string `json:"mimeType"`
	Type     string `json:"type"`
}

// Base interface for metadata with name (identifier) and title (display name) properties.
type BaseMetadata struct {
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
}

type BlobResourceContents struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// A base64-encoded string representing the binary data of the item.
	Blob []byte `json:"blob"`
	// The MIME type of this resource, if known.
	MimeType *string `json:"mimeType,omitempty"`
	// The URI of this resource.
	URI url.URL `json:"uri"`
}

type BooleanSchema struct {
	Default     *bool   `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	Title       *string `json:"title,omitempty"`
	Type        string  `json:"type"`
}

// A progress token, used to associate progress notifications with the original request.
type ProgressToken = interface{}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type CallToolRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Metadata for augmenting a request with task execution.
// Include this in the `task` field of the request parameters.
type TaskMetadata struct {
	// Requested duration in milliseconds to retain task from creation.
	Ttl *int `json:"ttl,omitempty"`
}

// Parameters for a `tools/call` request.
type CallToolRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *CallToolRequestParamsMeta `json:"_meta,omitempty"`
	// Arguments to use for the tool call.
	Arguments map[string]interface{} `json:"arguments,omitempty"`
	// The name of the tool.
	Name string `json:"name"`
	// If specified, the caller is requesting task-augmented execution for this request.
	// The request will return a CreateTaskResult immediately, and the actual result can be
	// retrieved later via tasks/result.
	//
	// Task augmentation is subject to capability negotiation - receivers MUST declare support
	// for task augmentation of specific request types in their capabilities.
	Task *TaskMetadata `json:"task,omitempty"`
}

// A uniquely identifying ID for a request in JSON-RPC.
type RequestId = interface{}

// Used by the client to invoke a tool provided by the server.
type CallToolRequest struct {
	ID      RequestId             `json:"id"`
	Jsonrpc string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  CallToolRequestParams `json:"params"`
}

// The contents of a resource, embedded into a prompt or tool call result.
//
// It is up to the client how best to render embedded resources for the benefit
// of the LLM and/or the user.
type EmbeddedResource struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	Resource    interface{}  `json:"resource"`
	Type        string       `json:"type"`
}

// An image provided to or from an LLM.
type ImageContent struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	// The base64-encoded image data.
	Data []byte `json:"data"`
	// The MIME type of the image. Different providers may support different image types.
	MimeType string `json:"mimeType"`
	Type     string `json:"type"`
}

// An optionally-sized icon that can be displayed in a user interface.
type Icon struct {
	// Optional MIME type override if the source MIME type is missing or generic.
	// For example: `"image/png"`, `"image/jpeg"`, or `"image/svg+xml"`.
	MimeType *string `json:"mimeType,omitempty"`
	// Optional array of strings that specify sizes at which the icon can be used.
	// Each string should be in WxH format (e.g., `"48x48"`, `"96x96"`) or `"any"` for scalable formats like SVG.
	//
	// If not provided, the client should assume that the icon can be used at any size.
	Sizes []string `json:"sizes,omitempty"`
	// A standard URI pointing to an icon resource. May be an HTTP/HTTPS URL or a
	// `data:` URI with Base64-encoded image data.
	//
	// Consumers SHOULD takes steps to ensure URLs serving icons are from the
	// same domain as the client/server or a trusted domain.
	//
	// Consumers SHOULD take appropriate precautions when consuming SVGs as they can contain
	// executable JavaScript.
	Src url.URL `json:"src"`
	// Optional specifier for the theme this icon is designed for. `light` indicates
	// the icon is designed to be used with a light background, and `dark` indicates
	// the icon is designed to be used with a dark background.
	//
	// If not provided, the client should assume the icon can be used with any theme.
	Theme *string `json:"theme,omitempty"`
}

// A resource that the server is capable of reading, included in a prompt or tool call result.
//
// Note: resource links returned by tools are not guaranteed to appear in the results of `resources/list` requests.
type ResourceLink struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	// A description of what this resource represents.
	//
	// This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
	Description *string `json:"description,omitempty"`
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
	// The MIME type of this resource, if known.
	MimeType *string `json:"mimeType,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
	//
	// This can be used by Hosts to display file sizes and estimate context window usage.
	Size *int `json:"size,omitempty"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
	Type  string  `json:"type"`
	// The URI of this resource.
	URI url.URL `json:"uri"`
}

// Text provided to or from an LLM.
type TextContent struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	// The text content of the message.
	Text string `json:"text"`
	Type string `json:"type"`
}

type ContentBlock = interface{}

// The server's response to a tool call.
type CallToolResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// A list of content objects that represent the unstructured result of the tool call.
	Content []ContentBlock `json:"content"`
	// Whether the tool call ended in an error.
	//
	// If not set, this is assumed to be false (the call was successful).
	//
	// Any errors that originate from the tool SHOULD be reported inside the result
	// object, with `isError` set to true, _not_ as an MCP protocol-level error
	// response. Otherwise, the LLM would not be able to see that an error occurred
	// and self-correct.
	//
	// However, any errors in _finding_ the tool, an error indicating that the
	// server does not support tool calls, or any other exceptional conditions,
	// should be reported as an MCP error response.
	IsError *bool `json:"isError,omitempty"`
	// An optional JSON object that represents the structured result of the tool call.
	StructuredContent map[string]interface{} `json:"structuredContent,omitempty"`
}

type CancelTaskRequestParams struct {
	// The task identifier to cancel.
	TaskID string `json:"taskId"`
}

// A request to cancel a task.
type CancelTaskRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  CancelTaskRequestParams `json:"params"`
}

type Result struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
}

// The status of a task.
type TaskStatus string

const (
	TaskStatusCancelled     TaskStatus = "cancelled"
	TaskStatusCompleted     TaskStatus = "completed"
	TaskStatusFailed        TaskStatus = "failed"
	TaskStatusInputRequired TaskStatus = "input_required"
	TaskStatusWorking       TaskStatus = "working"
)

// The response to a tasks/cancel request.
type CancelTaskResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// ISO 8601 timestamp when the task was created.
	CreatedAt string `json:"createdAt"`
	// ISO 8601 timestamp when the task was last updated.
	LastUpdatedAt string `json:"lastUpdatedAt"`
	// Suggested polling interval in milliseconds.
	PollInterval *int `json:"pollInterval,omitempty"`
	// Current task state.
	Status TaskStatus `json:"status"`
	// Optional human-readable message describing the current task state.
	// This can provide context for any status, including:
	// - Reasons for "cancelled" status
	// - Summaries for "completed" status
	// - Diagnostic information for "failed" status (e.g., error details, what went wrong)
	StatusMessage *string `json:"statusMessage,omitempty"`
	// The task identifier.
	TaskID string `json:"taskId"`
	// Actual retention duration from creation in milliseconds, null for unlimited.
	Ttl int `json:"ttl"`
}

// Parameters for a `notifications/cancelled` notification.
type CancelledNotificationParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An optional string describing the reason for the cancellation. This MAY be logged or presented to the user.
	Reason *string `json:"reason,omitempty"`
	// The ID of the request to cancel.
	//
	// This MUST correspond to the ID of a request previously issued in the same direction.
	// This MUST be provided for cancelling non-task requests.
	// This MUST NOT be used for cancelling tasks (use the `tasks/cancel` request instead).
	RequestID *RequestId `json:"requestId,omitempty"`
}

// This notification can be sent by either side to indicate that it is cancelling a previously-issued request.
//
// The request SHOULD still be in-flight, but due to communication latency, it is always possible that this notification MAY arrive after the request has already finished.
//
// This notification indicates that the result will be unused, so any associated processing SHOULD cease.
//
// A client MUST NOT attempt to cancel its `initialize` request.
//
// For task cancellation, use the `tasks/cancel` request instead of this notification.
type CancelledNotification struct {
	Jsonrpc string                      `json:"jsonrpc"`
	Method  string                      `json:"method"`
	Params  CancelledNotificationParams `json:"params"`
}

type ClientCapabilitiesElicitationForm struct {
}

type ClientCapabilitiesElicitationURL struct {
}

// Present if the client supports elicitation from the server.
type ClientCapabilitiesElicitation struct {
	Form *ClientCapabilitiesElicitationForm `json:"form,omitempty"`
	URL  *ClientCapabilitiesElicitationURL  `json:"url,omitempty"`
}

type ClientCapabilitiesExperimentalValue struct {
}

// Present if the client supports listing roots.
type ClientCapabilitiesRoots struct {
	// Whether the client supports notifications for changes to the roots list.
	ListChanged *bool `json:"listChanged,omitempty"`
}

// Whether the client supports context inclusion via includeContext parameter.
// If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
type ClientCapabilitiesSamplingContext struct {
}

// Whether the client supports tool use via tools and toolChoice parameters.
type ClientCapabilitiesSamplingTools struct {
}

// Present if the client supports sampling from an LLM.
type ClientCapabilitiesSampling struct {
	// Whether the client supports context inclusion via includeContext parameter.
	// If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
	Context *ClientCapabilitiesSamplingContext `json:"context,omitempty"`
	// Whether the client supports tool use via tools and toolChoice parameters.
	Tools *ClientCapabilitiesSamplingTools `json:"tools,omitempty"`
}

// Whether this client supports tasks/cancel.
type ClientCapabilitiesTasksCancel struct {
}

// Whether this client supports tasks/list.
type ClientCapabilitiesTasksList struct {
}

// Whether the client supports task-augmented elicitation/create requests.
type ClientCapabilitiesTasksRequestsElicitationCreate struct {
}

// Task support for elicitation-related requests.
type ClientCapabilitiesTasksRequestsElicitation struct {
	// Whether the client supports task-augmented elicitation/create requests.
	Create *ClientCapabilitiesTasksRequestsElicitationCreate `json:"create,omitempty"`
}

// Whether the client supports task-augmented sampling/createMessage requests.
type ClientCapabilitiesTasksRequestsSamplingCreateMessage struct {
}

// Task support for sampling-related requests.
type ClientCapabilitiesTasksRequestsSampling struct {
	// Whether the client supports task-augmented sampling/createMessage requests.
	CreateMessage *ClientCapabilitiesTasksRequestsSamplingCreateMessage `json:"createMessage,omitempty"`
}

// Specifies which request types can be augmented with tasks.
type ClientCapabilitiesTasksRequests struct {
	// Task support for elicitation-related requests.
	Elicitation *ClientCapabilitiesTasksRequestsElicitation `json:"elicitation,omitempty"`
	// Task support for sampling-related requests.
	Sampling *ClientCapabilitiesTasksRequestsSampling `json:"sampling,omitempty"`
}

// Present if the client supports task-augmented requests.
type ClientCapabilitiesTasks struct {
	// Whether this client supports tasks/cancel.
	Cancel *ClientCapabilitiesTasksCancel `json:"cancel,omitempty"`
	// Whether this client supports tasks/list.
	List *ClientCapabilitiesTasksList `json:"list,omitempty"`
	// Specifies which request types can be augmented with tasks.
	Requests *ClientCapabilitiesTasksRequests `json:"requests,omitempty"`
}

// Capabilities a client may support. Known capabilities are defined here, in this schema, but this is not a closed set: any client can define its own, additional capabilities.
type ClientCapabilities struct {
	// Present if the client supports elicitation from the server.
	Elicitation *ClientCapabilitiesElicitation `json:"elicitation,omitempty"`
	// Experimental, non-standard capabilities that the client supports.
	Experimental map[string]ClientCapabilitiesExperimentalValue `json:"experimental,omitempty"`
	// Present if the client supports listing roots.
	Roots *ClientCapabilitiesRoots `json:"roots,omitempty"`
	// Present if the client supports sampling from an LLM.
	Sampling *ClientCapabilitiesSampling `json:"sampling,omitempty"`
	// Present if the client supports task-augmented requests.
	Tasks *ClientCapabilitiesTasks `json:"tasks,omitempty"`
}

type NotificationParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
}

// This notification is sent from the client to the server after initialization has finished.
type InitializedNotification struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  *NotificationParams `json:"params,omitempty"`
}

// Parameters for a `notifications/progress` notification.
type ProgressNotificationParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An optional message describing the current progress.
	Message *string `json:"message,omitempty"`
	// The progress thus far. This should increase every time progress is made, even if the total is unknown.
	Progress float64 `json:"progress"`
	// The progress token which was given in the initial request, used to associate this notification with the request that is proceeding.
	ProgressToken ProgressToken `json:"progressToken"`
	// Total number of items to process (or total progress required), if known.
	Total *float64 `json:"total,omitempty"`
}

// An out-of-band notification used to inform the receiver of a progress update for a long-running request.
type ProgressNotification struct {
	Jsonrpc string                     `json:"jsonrpc"`
	Method  string                     `json:"method"`
	Params  ProgressNotificationParams `json:"params"`
}

// A notification from the client to the server, informing it that the list of roots has changed.
// This notification should be sent whenever the client adds, removes, or modifies any root.
// The server should then request an updated list of roots using the ListRootsRequest.
type RootsListChangedNotification struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  *NotificationParams `json:"params,omitempty"`
}

// Parameters for a `notifications/tasks/status` notification.
type TaskStatusNotificationParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// ISO 8601 timestamp when the task was created.
	CreatedAt string `json:"createdAt"`
	// ISO 8601 timestamp when the task was last updated.
	LastUpdatedAt string `json:"lastUpdatedAt"`
	// Suggested polling interval in milliseconds.
	PollInterval *int `json:"pollInterval,omitempty"`
	// Current task state.
	Status TaskStatus `json:"status"`
	// Optional human-readable message describing the current task state.
	// This can provide context for any status, including:
	// - Reasons for "cancelled" status
	// - Summaries for "completed" status
	// - Diagnostic information for "failed" status (e.g., error details, what went wrong)
	StatusMessage *string `json:"statusMessage,omitempty"`
	// The task identifier.
	TaskID string `json:"taskId"`
	// Actual retention duration from creation in milliseconds, null for unlimited.
	Ttl int `json:"ttl"`
}

// An optional notification from the receiver to the requestor, informing them that a task's status has changed. Receivers are not required to send these notifications.
type TaskStatusNotification struct {
	Jsonrpc string                       `json:"jsonrpc"`
	Method  string                       `json:"method"`
	Params  TaskStatusNotificationParams `json:"params"`
}

type ClientNotification = interface{}

// The argument's information
type CompleteRequestParamsArgument struct {
	// The name of the argument
	Name string `json:"name"`
	// The value of the argument to use for completion matching.
	Value string `json:"value"`
}

// Additional, optional context for completions
type CompleteRequestParamsContext struct {
	// Previously-resolved variables in a URI template or prompt.
	Arguments map[string]string `json:"arguments,omitempty"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type CompleteRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for a `completion/complete` request.
type CompleteRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *CompleteRequestParamsMeta `json:"_meta,omitempty"`
	// The argument's information
	Argument CompleteRequestParamsArgument `json:"argument"`
	// Additional, optional context for completions
	Context *CompleteRequestParamsContext `json:"context,omitempty"`
	Ref     interface{}                   `json:"ref"`
}

// A request from the client to the server, to ask for completion options.
type CompleteRequest struct {
	ID      RequestId             `json:"id"`
	Jsonrpc string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  CompleteRequestParams `json:"params"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type GetPromptRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for a `prompts/get` request.
type GetPromptRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *GetPromptRequestParamsMeta `json:"_meta,omitempty"`
	// Arguments to use for templating the prompt.
	Arguments map[string]string `json:"arguments,omitempty"`
	// The name of the prompt or prompt template.
	Name string `json:"name"`
}

// Used by the client to get a prompt provided by the server.
type GetPromptRequest struct {
	ID      RequestId              `json:"id"`
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  GetPromptRequestParams `json:"params"`
}

type GetTaskPayloadRequestParams struct {
	// The task identifier to retrieve results for.
	TaskID string `json:"taskId"`
}

// A request to retrieve the result of a completed task.
type GetTaskPayloadRequest struct {
	ID      RequestId                   `json:"id"`
	Jsonrpc string                      `json:"jsonrpc"`
	Method  string                      `json:"method"`
	Params  GetTaskPayloadRequestParams `json:"params"`
}

type GetTaskRequestParams struct {
	// The task identifier to query.
	TaskID string `json:"taskId"`
}

// A request to retrieve the state of a task.
type GetTaskRequest struct {
	ID      RequestId            `json:"id"`
	Jsonrpc string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  GetTaskRequestParams `json:"params"`
}

// Describes the MCP implementation.
type Implementation struct {
	// An optional human-readable description of what this implementation does.
	//
	// This can be used by clients or servers to provide context about their purpose
	// and capabilities. For example, a server might describe the types of resources
	// or tools it provides, while a client might describe its intended use case.
	Description *string `json:"description,omitempty"`
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title   *string `json:"title,omitempty"`
	Version string  `json:"version"`
	// An optional URL of the website for this implementation.
	WebsiteURL *url.URL `json:"websiteUrl,omitempty"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type InitializeRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for an `initialize` request.
type InitializeRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta         *InitializeRequestParamsMeta `json:"_meta,omitempty"`
	Capabilities ClientCapabilities           `json:"capabilities"`
	ClientInfo   Implementation               `json:"clientInfo"`
	// The latest version of the Model Context Protocol that the client supports. The client MAY decide to support older versions as well.
	ProtocolVersion string `json:"protocolVersion"`
}

// This request is sent from the client to the server when it first connects, asking it to begin initialization.
type InitializeRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  InitializeRequestParams `json:"params"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type PaginatedRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Common parameters for paginated requests.
type PaginatedRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *PaginatedRequestParamsMeta `json:"_meta,omitempty"`
	// An opaque token representing the current pagination position.
	// If provided, the server should return results starting after this cursor.
	Cursor *string `json:"cursor,omitempty"`
}

// Sent from the client to request a list of prompts and prompt templates the server has.
type ListPromptsRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  *PaginatedRequestParams `json:"params,omitempty"`
}

// Sent from the client to request a list of resource templates the server has.
type ListResourceTemplatesRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  *PaginatedRequestParams `json:"params,omitempty"`
}

// Sent from the client to request a list of resources the server has.
type ListResourcesRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  *PaginatedRequestParams `json:"params,omitempty"`
}

// A request to retrieve a list of tasks.
type ListTasksRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  *PaginatedRequestParams `json:"params,omitempty"`
}

// Sent from the client to request a list of tools the server has.
type ListToolsRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  *PaginatedRequestParams `json:"params,omitempty"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type RequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Common params for any request.
type RequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *RequestParamsMeta `json:"_meta,omitempty"`
}

// A ping, issued by either the server or the client, to check that the other party is still alive. The receiver must promptly respond, or else may be disconnected.
type PingRequest struct {
	ID      RequestId      `json:"id"`
	Jsonrpc string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  *RequestParams `json:"params,omitempty"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type ReadResourceRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for a `resources/read` request.
type ReadResourceRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *ReadResourceRequestParamsMeta `json:"_meta,omitempty"`
	// The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
	URI url.URL `json:"uri"`
}

// Sent from the client to the server, to read a specific resource URI.
type ReadResourceRequest struct {
	ID      RequestId                 `json:"id"`
	Jsonrpc string                    `json:"jsonrpc"`
	Method  string                    `json:"method"`
	Params  ReadResourceRequestParams `json:"params"`
}

// The severity of a log message.
//
// These map to syslog message severities, as specified in RFC-5424:
// https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1
type LoggingLevel string

const (
	LoggingLevelAlert     LoggingLevel = "alert"
	LoggingLevelCritical  LoggingLevel = "critical"
	LoggingLevelDebug     LoggingLevel = "debug"
	LoggingLevelEmergency LoggingLevel = "emergency"
	LoggingLevelError     LoggingLevel = "error"
	LoggingLevelInfo      LoggingLevel = "info"
	LoggingLevelNotice    LoggingLevel = "notice"
	LoggingLevelWarning   LoggingLevel = "warning"
)

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type SetLevelRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for a `logging/setLevel` request.
type SetLevelRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *SetLevelRequestParamsMeta `json:"_meta,omitempty"`
	// The level of logging that the client wants to receive from the server. The server should send all logs at this level and higher (i.e., more severe) to the client as notifications/message.
	Level LoggingLevel `json:"level"`
}

// A request from the client to the server, to enable or adjust logging.
type SetLevelRequest struct {
	ID      RequestId             `json:"id"`
	Jsonrpc string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  SetLevelRequestParams `json:"params"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type SubscribeRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for a `resources/subscribe` request.
type SubscribeRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *SubscribeRequestParamsMeta `json:"_meta,omitempty"`
	// The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
	URI url.URL `json:"uri"`
}

// Sent from the client to request resources/updated notifications from the server whenever a particular resource changes.
type SubscribeRequest struct {
	ID      RequestId              `json:"id"`
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  SubscribeRequestParams `json:"params"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type UnsubscribeRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Parameters for a `resources/unsubscribe` request.
type UnsubscribeRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *UnsubscribeRequestParamsMeta `json:"_meta,omitempty"`
	// The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
	URI url.URL `json:"uri"`
}

// Sent from the client to request cancellation of resources/updated notifications from the server. This should follow a previous resources/subscribe request.
type UnsubscribeRequest struct {
	ID      RequestId                `json:"id"`
	Jsonrpc string                   `json:"jsonrpc"`
	Method  string                   `json:"method"`
	Params  UnsubscribeRequestParams `json:"params"`
}

type ClientRequest = interface{}

// The client's response to a sampling/createMessage request from the server.
// The client should inform the user before returning the sampled message, to allow them
// to inspect the response (human in the loop) and decide whether to allow the server to see it.
type CreateMessageResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta    map[string]interface{} `json:"_meta,omitempty"`
	Content interface{}            `json:"content"`
	// The name of the model that generated the message.
	Model string `json:"model"`
	Role  Role   `json:"role"`
	// The reason why sampling stopped, if known.
	//
	// Standard values:
	// - "endTurn": Natural end of the assistant's turn
	// - "stopSequence": A stop sequence was encountered
	// - "maxTokens": Maximum token limit was reached
	// - "toolUse": The model wants to use one or more tools
	//
	// This field is an open string to allow for provider-specific stop reasons.
	StopReason *string `json:"stopReason,omitempty"`
}

// The client's response to an elicitation request.
type ElicitResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// The user action in response to the elicitation.
	// - "accept": User submitted the form/confirmed the action
	// - "decline": User explicitly decline the action
	// - "cancel": User dismissed without making an explicit choice
	Action string `json:"action"`
	// The submitted form data, only present when action is "accept" and mode was "form".
	// Contains values matching the requested schema.
	// Omitted for out-of-band mode responses.
	Content map[string]interface{} `json:"content,omitempty"`
}

// The response to a tasks/result request.
// The structure matches the result type of the original request.
// For example, a tools/call task would return the CallToolResult structure.
type GetTaskPayloadResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
}

// The response to a tasks/get request.
type GetTaskResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// ISO 8601 timestamp when the task was created.
	CreatedAt string `json:"createdAt"`
	// ISO 8601 timestamp when the task was last updated.
	LastUpdatedAt string `json:"lastUpdatedAt"`
	// Suggested polling interval in milliseconds.
	PollInterval *int `json:"pollInterval,omitempty"`
	// Current task state.
	Status TaskStatus `json:"status"`
	// Optional human-readable message describing the current task state.
	// This can provide context for any status, including:
	// - Reasons for "cancelled" status
	// - Summaries for "completed" status
	// - Diagnostic information for "failed" status (e.g., error details, what went wrong)
	StatusMessage *string `json:"statusMessage,omitempty"`
	// The task identifier.
	TaskID string `json:"taskId"`
	// Actual retention duration from creation in milliseconds, null for unlimited.
	Ttl int `json:"ttl"`
}

// Represents a root directory or file that the server can operate on.
type Root struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An optional name for the root. This can be used to provide a human-readable
	// identifier for the root, which may be useful for display purposes or for
	// referencing the root in other parts of the application.
	Name *string `json:"name,omitempty"`
	// The URI identifying the root. This *must* start with file:// for now.
	// This restriction may be relaxed in future versions of the protocol to allow
	// other URI schemes.
	URI url.URL `json:"uri"`
}

// The client's response to a roots/list request from the server.
// This result contains an array of Root objects, each representing a root directory
// or file that the server can operate on.
type ListRootsResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta  map[string]interface{} `json:"_meta,omitempty"`
	Roots []Root                 `json:"roots"`
}

// Data associated with a task.
type Task struct {
	// ISO 8601 timestamp when the task was created.
	CreatedAt string `json:"createdAt"`
	// ISO 8601 timestamp when the task was last updated.
	LastUpdatedAt string `json:"lastUpdatedAt"`
	// Suggested polling interval in milliseconds.
	PollInterval *int `json:"pollInterval,omitempty"`
	// Current task state.
	Status TaskStatus `json:"status"`
	// Optional human-readable message describing the current task state.
	// This can provide context for any status, including:
	// - Reasons for "cancelled" status
	// - Summaries for "completed" status
	// - Diagnostic information for "failed" status (e.g., error details, what went wrong)
	StatusMessage *string `json:"statusMessage,omitempty"`
	// The task identifier.
	TaskID string `json:"taskId"`
	// Actual retention duration from creation in milliseconds, null for unlimited.
	Ttl int `json:"ttl"`
}

// The response to a tasks/list request.
type ListTasksResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An opaque token representing the pagination position after the last returned result.
	// If present, there may be more results available.
	NextCursor *string `json:"nextCursor,omitempty"`
	Tasks      []Task  `json:"tasks"`
}

type ClientResult = interface{}

type CompleteResultCompletion struct {
	// Indicates whether there are additional completion options beyond those provided in the current response, even if the exact total is unknown.
	HasMore *bool `json:"hasMore,omitempty"`
	// The total number of completion options available. This can exceed the number of values actually sent in the response.
	Total *int `json:"total,omitempty"`
	// An array of completion values. Must not exceed 100 items.
	Values []string `json:"values"`
}

// The server's response to a completion/complete request
type CompleteResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta       map[string]interface{}   `json:"_meta,omitempty"`
	Completion CompleteResultCompletion `json:"completion"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type CreateMessageRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific.
type CreateMessageRequestParamsMetadata struct {
}

// Hints to use for model selection.
//
// Keys not declared here are currently left unspecified by the spec and are up
// to the client to interpret.
type ModelHint struct {
	// A hint for a model name.
	//
	// The client SHOULD treat this as a substring of a model name; for example:
	//  - `claude-3-5-sonnet` should match `claude-3-5-sonnet-20241022`
	//  - `sonnet` should match `claude-3-5-sonnet-20241022`, `claude-3-sonnet-20240229`, etc.
	//  - `claude` should match any Claude model
	//
	// The client MAY also map the string to a different provider's model name or a different model family, as long as it fills a similar niche; for example:
	//  - `gemini-1.5-flash` could match `claude-3-haiku-20240307`
	Name *string `json:"name,omitempty"`
}

// The server's preferences for model selection, requested of the client during sampling.
//
// Because LLMs can vary along multiple dimensions, choosing the "best" model is
// rarely straightforward.  Different models excel in different areas—some are
// faster but less capable, others are more capable but more expensive, and so
// on. This interface allows servers to express their priorities across multiple
// dimensions to help clients make an appropriate selection for their use case.
//
// These preferences are always advisory. The client MAY ignore them. It is also
// up to the client to decide how to interpret these preferences and how to
// balance them against other considerations.
type ModelPreferences struct {
	// How much to prioritize cost when selecting a model. A value of 0 means cost
	// is not important, while a value of 1 means cost is the most important
	// factor.
	CostPriority *float64 `json:"costPriority,omitempty"`
	// Optional hints to use for model selection.
	//
	// If multiple hints are specified, the client MUST evaluate them in order
	// (such that the first match is taken).
	//
	// The client SHOULD prioritize these hints over the numeric priorities, but
	// MAY still use the priorities to select from ambiguous matches.
	Hints []ModelHint `json:"hints,omitempty"`
	// How much to prioritize intelligence and capabilities when selecting a
	// model. A value of 0 means intelligence is not important, while a value of 1
	// means intelligence is the most important factor.
	IntelligencePriority *float64 `json:"intelligencePriority,omitempty"`
	// How much to prioritize sampling speed (latency) when selecting a model. A
	// value of 0 means speed is not important, while a value of 1 means speed is
	// the most important factor.
	SpeedPriority *float64 `json:"speedPriority,omitempty"`
}

// Describes a message issued to or received from an LLM API.
type SamplingMessage struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta    map[string]interface{} `json:"_meta,omitempty"`
	Content interface{}            `json:"content"`
	Role    Role                   `json:"role"`
}

// Additional properties describing a Tool to clients.
//
// NOTE: all properties in ToolAnnotations are **hints**.
// They are not guaranteed to provide a faithful description of
// tool behavior (including descriptive properties like `title`).
//
// Clients should never make tool use decisions based on ToolAnnotations
// received from untrusted servers.
type ToolAnnotations struct {
	// If true, the tool may perform destructive updates to its environment.
	// If false, the tool performs only additive updates.
	//
	// (This property is meaningful only when `readOnlyHint == false`)
	//
	// Default: true
	DestructiveHint *bool `json:"destructiveHint,omitempty"`
	// If true, calling the tool repeatedly with the same arguments
	// will have no additional effect on its environment.
	//
	// (This property is meaningful only when `readOnlyHint == false`)
	//
	// Default: false
	IdempotentHint *bool `json:"idempotentHint,omitempty"`
	// If true, this tool may interact with an "open world" of external
	// entities. If false, the tool's domain of interaction is closed.
	// For example, the world of a web search tool is open, whereas that
	// of a memory tool is not.
	//
	// Default: true
	OpenWorldHint *bool `json:"openWorldHint,omitempty"`
	// If true, the tool does not modify its environment.
	//
	// Default: false
	ReadOnlyHint *bool `json:"readOnlyHint,omitempty"`
	// A human-readable title for the tool.
	Title *string `json:"title,omitempty"`
}

// Execution-related properties for a tool.
type ToolExecution struct {
	// Indicates whether this tool supports task-augmented execution.
	// This allows clients to handle long-running operations through polling
	// the task system.
	//
	// - "forbidden": Tool does not support task-augmented execution (default when absent)
	// - "optional": Tool may support task-augmented execution
	// - "required": Tool requires task-augmented execution
	//
	// Default: "forbidden"
	TaskSupport *string `json:"taskSupport,omitempty"`
}

type ToolInputSchemaPropertiesValue struct {
}

// A JSON Schema object defining the expected parameters for the tool.
type ToolInputSchema struct {
	Schema     *string                                   `json:"$schema,omitempty"`
	Properties map[string]ToolInputSchemaPropertiesValue `json:"properties,omitempty"`
	Required   []string                                  `json:"required,omitempty"`
	Type       string                                    `json:"type"`
}

type ToolOutputSchemaPropertiesValue struct {
}

// An optional JSON Schema object defining the structure of the tool's output returned in
// the structuredContent field of a CallToolResult.
//
// Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
// Currently restricted to type: "object" at the root level.
type ToolOutputSchema struct {
	Schema     *string                                    `json:"$schema,omitempty"`
	Properties map[string]ToolOutputSchemaPropertiesValue `json:"properties,omitempty"`
	Required   []string                                   `json:"required,omitempty"`
	Type       string                                     `json:"type"`
}

// Definition for a tool the client can call.
type Tool struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional additional tool information.
	//
	// Display name precedence order is: title, annotations.title, then name.
	Annotations *ToolAnnotations `json:"annotations,omitempty"`
	// A human-readable description of the tool.
	//
	// This can be used by clients to improve the LLM's understanding of available tools. It can be thought of like a "hint" to the model.
	Description *string `json:"description,omitempty"`
	// Execution-related properties for this tool.
	Execution *ToolExecution `json:"execution,omitempty"`
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
	// A JSON Schema object defining the expected parameters for the tool.
	InputSchema ToolInputSchema `json:"inputSchema"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// An optional JSON Schema object defining the structure of the tool's output returned in
	// the structuredContent field of a CallToolResult.
	//
	// Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
	// Currently restricted to type: "object" at the root level.
	OutputSchema *ToolOutputSchema `json:"outputSchema,omitempty"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
}

// Controls tool selection behavior for sampling requests.
type ToolChoice struct {
	// Controls the tool use ability of the model:
	// - "auto": Model decides whether to use tools (default)
	// - "required": Model MUST use at least one tool before completing
	// - "none": Model MUST NOT use any tools
	Mode *string `json:"mode,omitempty"`
}

// Parameters for a `sampling/createMessage` request.
type CreateMessageRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *CreateMessageRequestParamsMeta `json:"_meta,omitempty"`
	// A request to include context from one or more MCP servers (including the caller), to be attached to the prompt.
	// The client MAY ignore this request.
	//
	// Default is "none". Values "thisServer" and "allServers" are soft-deprecated. Servers SHOULD only use these values if the client
	// declares ClientCapabilities.sampling.context. These values may be removed in future spec releases.
	IncludeContext *string `json:"includeContext,omitempty"`
	// The requested maximum number of tokens to sample (to prevent runaway completions).
	//
	// The client MAY choose to sample fewer tokens than the requested maximum.
	MaxTokens int               `json:"maxTokens"`
	Messages  []SamplingMessage `json:"messages"`
	// Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific.
	Metadata *CreateMessageRequestParamsMetadata `json:"metadata,omitempty"`
	// The server's preferences for which model to select. The client MAY ignore these preferences.
	ModelPreferences *ModelPreferences `json:"modelPreferences,omitempty"`
	StopSequences    []string          `json:"stopSequences,omitempty"`
	// An optional system prompt the server wants to use for sampling. The client MAY modify or omit this prompt.
	SystemPrompt *string `json:"systemPrompt,omitempty"`
	// If specified, the caller is requesting task-augmented execution for this request.
	// The request will return a CreateTaskResult immediately, and the actual result can be
	// retrieved later via tasks/result.
	//
	// Task augmentation is subject to capability negotiation - receivers MUST declare support
	// for task augmentation of specific request types in their capabilities.
	Task        *TaskMetadata `json:"task,omitempty"`
	Temperature *float64      `json:"temperature,omitempty"`
	// Controls how the model uses tools.
	// The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
	// Default is `{ mode: "auto" }`.
	ToolChoice *ToolChoice `json:"toolChoice,omitempty"`
	// Tools that the model may use during generation.
	// The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
	Tools []Tool `json:"tools,omitempty"`
}

// A request from the server to sample an LLM via the client. The client has full discretion over which model to select. The client should also inform the user before beginning sampling, to allow them to inspect the request (human in the loop) and decide whether to approve it.
type CreateMessageRequest struct {
	ID      RequestId                  `json:"id"`
	Jsonrpc string                     `json:"jsonrpc"`
	Method  string                     `json:"method"`
	Params  CreateMessageRequestParams `json:"params"`
}

// A response to a task-augmented request.
type CreateTaskResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	Task Task                   `json:"task"`
}

// An opaque token used to represent a cursor for pagination.
type Cursor = string

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type ElicitRequestFormParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Use TitledSingleSelectEnumSchema instead.
// This interface will be removed in a future version.
type LegacyTitledEnumSchema struct {
	Default     *string  `json:"default,omitempty"`
	Description *string  `json:"description,omitempty"`
	Enum        []string `json:"enum"`
	// (Legacy) Display names for enum values.
	// Non-standard according to JSON schema 2020-12.
	EnumNames []string `json:"enumNames,omitempty"`
	Title     *string  `json:"title,omitempty"`
	Type      string   `json:"type"`
}

type NumberSchema struct {
	Default     *int    `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	Maximum     *int    `json:"maximum,omitempty"`
	Minimum     *int    `json:"minimum,omitempty"`
	Title       *string `json:"title,omitempty"`
	Type        string  `json:"type"`
}

type StringSchema struct {
	Default     *string `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	Format      *string `json:"format,omitempty"`
	MaxLength   *int    `json:"maxLength,omitempty"`
	MinLength   *int    `json:"minLength,omitempty"`
	Title       *string `json:"title,omitempty"`
	Type        string  `json:"type"`
}

type TitledMultiSelectEnumSchemaItemsAnyOfItem struct {
	// The constant enum value.
	Const string `json:"const"`
	// Display title for this option.
	Title string `json:"title"`
}

// Schema for array items with enum options and display labels.
type TitledMultiSelectEnumSchemaItems struct {
	// Array of enum options with values and display labels.
	AnyOf []TitledMultiSelectEnumSchemaItemsAnyOfItem `json:"anyOf"`
}

// Schema for multiple-selection enumeration with display titles for each option.
type TitledMultiSelectEnumSchema struct {
	// Optional default value.
	Default []string `json:"default,omitempty"`
	// Optional description for the enum field.
	Description *string `json:"description,omitempty"`
	// Schema for array items with enum options and display labels.
	Items TitledMultiSelectEnumSchemaItems `json:"items"`
	// Maximum number of items to select.
	MaxItems *int `json:"maxItems,omitempty"`
	// Minimum number of items to select.
	MinItems *int `json:"minItems,omitempty"`
	// Optional title for the enum field.
	Title *string `json:"title,omitempty"`
	Type  string  `json:"type"`
}

type TitledSingleSelectEnumSchemaOneOfItem struct {
	// The enum value.
	Const string `json:"const"`
	// Display label for this option.
	Title string `json:"title"`
}

// Schema for single-selection enumeration with display titles for each option.
type TitledSingleSelectEnumSchema struct {
	// Optional default value.
	Default *string `json:"default,omitempty"`
	// Optional description for the enum field.
	Description *string `json:"description,omitempty"`
	// Array of enum options with values and display labels.
	OneOf []TitledSingleSelectEnumSchemaOneOfItem `json:"oneOf"`
	// Optional title for the enum field.
	Title *string `json:"title,omitempty"`
	Type  string  `json:"type"`
}

// Schema for the array items.
type UntitledMultiSelectEnumSchemaItems struct {
	// Array of enum values to choose from.
	Enum []string `json:"enum"`
	Type string   `json:"type"`
}

// Schema for multiple-selection enumeration without display titles for options.
type UntitledMultiSelectEnumSchema struct {
	// Optional default value.
	Default []string `json:"default,omitempty"`
	// Optional description for the enum field.
	Description *string `json:"description,omitempty"`
	// Schema for the array items.
	Items UntitledMultiSelectEnumSchemaItems `json:"items"`
	// Maximum number of items to select.
	MaxItems *int `json:"maxItems,omitempty"`
	// Minimum number of items to select.
	MinItems *int `json:"minItems,omitempty"`
	// Optional title for the enum field.
	Title *string `json:"title,omitempty"`
	Type  string  `json:"type"`
}

// Schema for single-selection enumeration without display titles for options.
type UntitledSingleSelectEnumSchema struct {
	// Optional default value.
	Default *string `json:"default,omitempty"`
	// Optional description for the enum field.
	Description *string `json:"description,omitempty"`
	// Array of enum values to choose from.
	Enum []string `json:"enum"`
	// Optional title for the enum field.
	Title *string `json:"title,omitempty"`
	Type  string  `json:"type"`
}

// Restricted schema definitions that only allow primitive types
// without nested objects or arrays.
type PrimitiveSchemaDefinition = interface{}

// A restricted subset of JSON Schema.
// Only top-level properties are allowed, without nesting.
type ElicitRequestFormParamsRequestedSchema struct {
	Schema     *string                              `json:"$schema,omitempty"`
	Properties map[string]PrimitiveSchemaDefinition `json:"properties"`
	Required   []string                             `json:"required,omitempty"`
	Type       string                               `json:"type"`
}

// The parameters for a request to elicit non-sensitive information from the user via a form in the client.
type ElicitRequestFormParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *ElicitRequestFormParamsMeta `json:"_meta,omitempty"`
	// The message to present to the user describing what information is being requested.
	Message string `json:"message"`
	// The elicitation mode.
	Mode *string `json:"mode,omitempty"`
	// A restricted subset of JSON Schema.
	// Only top-level properties are allowed, without nesting.
	RequestedSchema ElicitRequestFormParamsRequestedSchema `json:"requestedSchema"`
	// If specified, the caller is requesting task-augmented execution for this request.
	// The request will return a CreateTaskResult immediately, and the actual result can be
	// retrieved later via tasks/result.
	//
	// Task augmentation is subject to capability negotiation - receivers MUST declare support
	// for task augmentation of specific request types in their capabilities.
	Task *TaskMetadata `json:"task,omitempty"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type ElicitRequestURLParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// The parameters for a request to elicit information from the user via a URL in the client.
type ElicitRequestURLParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *ElicitRequestURLParamsMeta `json:"_meta,omitempty"`
	// The ID of the elicitation, which must be unique within the context of the server.
	// The client MUST treat this ID as an opaque value.
	ElicitationID string `json:"elicitationId"`
	// The message to present to the user explaining why the interaction is needed.
	Message string `json:"message"`
	// The elicitation mode.
	Mode string `json:"mode"`
	// If specified, the caller is requesting task-augmented execution for this request.
	// The request will return a CreateTaskResult immediately, and the actual result can be
	// retrieved later via tasks/result.
	//
	// Task augmentation is subject to capability negotiation - receivers MUST declare support
	// for task augmentation of specific request types in their capabilities.
	Task *TaskMetadata `json:"task,omitempty"`
	// The URL that the user should navigate to.
	URL url.URL `json:"url"`
}

// The parameters for a request to elicit additional information from the user via the client.
type ElicitRequestParams = interface{}

// A request from the server to elicit additional information from the user via the client.
type ElicitRequest struct {
	ID      RequestId           `json:"id"`
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  ElicitRequestParams `json:"params"`
}

type ElicitationCompleteNotificationParams struct {
	// The ID of the elicitation that completed.
	ElicitationID string `json:"elicitationId"`
}

// An optional notification from the server to the client, informing it of a completion of a out-of-band elicitation request.
type ElicitationCompleteNotification struct {
	Jsonrpc string                                `json:"jsonrpc"`
	Method  string                                `json:"method"`
	Params  ElicitationCompleteNotificationParams `json:"params"`
}

type EmptyResult = interface{}

type EnumSchema = interface{}

type Error struct {
	// The error type that occurred.
	Code int `json:"code"`
	// Additional information about the error. The value of this member is defined by the sender (e.g. detailed error information, nested errors etc.).
	Data interface{} `json:"data,omitempty"`
	// A short description of the error. The message SHOULD be limited to a concise single sentence.
	Message string `json:"message"`
}

// Describes a message returned as part of a prompt.
//
// This is similar to `SamplingMessage`, but also supports the embedding of
// resources from the MCP server.
type PromptMessage struct {
	Content ContentBlock `json:"content"`
	Role    Role         `json:"role"`
}

// The server's response to a prompts/get request from the client.
type GetPromptResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An optional description for the prompt.
	Description *string         `json:"description,omitempty"`
	Messages    []PromptMessage `json:"messages"`
}

// Base interface to add `icons` property.
type Icons struct {
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
}

// Present if the server supports argument autocompletion suggestions.
type ServerCapabilitiesCompletions struct {
}

type ServerCapabilitiesExperimentalValue struct {
}

// Present if the server supports sending log messages to the client.
type ServerCapabilitiesLogging struct {
}

// Present if the server offers any prompt templates.
type ServerCapabilitiesPrompts struct {
	// Whether this server supports notifications for changes to the prompt list.
	ListChanged *bool `json:"listChanged,omitempty"`
}

// Present if the server offers any resources to read.
type ServerCapabilitiesResources struct {
	// Whether this server supports notifications for changes to the resource list.
	ListChanged *bool `json:"listChanged,omitempty"`
	// Whether this server supports subscribing to resource updates.
	Subscribe *bool `json:"subscribe,omitempty"`
}

// Whether this server supports tasks/cancel.
type ServerCapabilitiesTasksCancel struct {
}

// Whether this server supports tasks/list.
type ServerCapabilitiesTasksList struct {
}

// Whether the server supports task-augmented tools/call requests.
type ServerCapabilitiesTasksRequestsToolsCall struct {
}

// Task support for tool-related requests.
type ServerCapabilitiesTasksRequestsTools struct {
	// Whether the server supports task-augmented tools/call requests.
	Call *ServerCapabilitiesTasksRequestsToolsCall `json:"call,omitempty"`
}

// Specifies which request types can be augmented with tasks.
type ServerCapabilitiesTasksRequests struct {
	// Task support for tool-related requests.
	Tools *ServerCapabilitiesTasksRequestsTools `json:"tools,omitempty"`
}

// Present if the server supports task-augmented requests.
type ServerCapabilitiesTasks struct {
	// Whether this server supports tasks/cancel.
	Cancel *ServerCapabilitiesTasksCancel `json:"cancel,omitempty"`
	// Whether this server supports tasks/list.
	List *ServerCapabilitiesTasksList `json:"list,omitempty"`
	// Specifies which request types can be augmented with tasks.
	Requests *ServerCapabilitiesTasksRequests `json:"requests,omitempty"`
}

// Present if the server offers any tools to call.
type ServerCapabilitiesTools struct {
	// Whether this server supports notifications for changes to the tool list.
	ListChanged *bool `json:"listChanged,omitempty"`
}

// Capabilities that a server may support. Known capabilities are defined here, in this schema, but this is not a closed set: any server can define its own, additional capabilities.
type ServerCapabilities struct {
	// Present if the server supports argument autocompletion suggestions.
	Completions *ServerCapabilitiesCompletions `json:"completions,omitempty"`
	// Experimental, non-standard capabilities that the server supports.
	Experimental map[string]ServerCapabilitiesExperimentalValue `json:"experimental,omitempty"`
	// Present if the server supports sending log messages to the client.
	Logging *ServerCapabilitiesLogging `json:"logging,omitempty"`
	// Present if the server offers any prompt templates.
	Prompts *ServerCapabilitiesPrompts `json:"prompts,omitempty"`
	// Present if the server offers any resources to read.
	Resources *ServerCapabilitiesResources `json:"resources,omitempty"`
	// Present if the server supports task-augmented requests.
	Tasks *ServerCapabilitiesTasks `json:"tasks,omitempty"`
	// Present if the server offers any tools to call.
	Tools *ServerCapabilitiesTools `json:"tools,omitempty"`
}

// After receiving an initialize request from the client, the server sends this response.
type InitializeResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta         map[string]interface{} `json:"_meta,omitempty"`
	Capabilities ServerCapabilities     `json:"capabilities"`
	// Instructions describing how to use the server and its features.
	//
	// This can be used by clients to improve the LLM's understanding of available tools, resources, etc. It can be thought of like a "hint" to the model. For example, this information MAY be added to the system prompt.
	Instructions *string `json:"instructions,omitempty"`
	// The version of the Model Context Protocol that the server wants to use. This may not match the version that the client requested. If the client cannot support this version, it MUST disconnect.
	ProtocolVersion string         `json:"protocolVersion"`
	ServerInfo      Implementation `json:"serverInfo"`
}

// A response to a request that indicates an error occurred.
type JSONRPCErrorResponse struct {
	Error   Error      `json:"error"`
	ID      *RequestId `json:"id,omitempty"`
	Jsonrpc string     `json:"jsonrpc"`
}

// A notification which does not expect a response.
type JSONRPCNotification struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

// A request that expects a response.
type JSONRPCRequest struct {
	ID      RequestId              `json:"id"`
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

// A successful (non-error) response to a request.
type JSONRPCResultResponse struct {
	ID      RequestId `json:"id"`
	Jsonrpc string    `json:"jsonrpc"`
	Result  Result    `json:"result"`
}

// Refers to any valid JSON-RPC object that can be decoded off the wire, or encoded to be sent.
type JSONRPCMessage = interface{}

// A response to a request, containing either the result or error.
type JSONRPCResponse = interface{}

// Describes an argument that a prompt can accept.
type PromptArgument struct {
	// A human-readable description of the argument.
	Description *string `json:"description,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Whether this argument must be provided.
	Required *bool `json:"required,omitempty"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
}

// A prompt or prompt template that the server offers.
type Prompt struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// A list of arguments to use for templating the prompt.
	Arguments []PromptArgument `json:"arguments,omitempty"`
	// An optional description of what this prompt provides
	Description *string `json:"description,omitempty"`
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
}

// The server's response to a prompts/list request from the client.
type ListPromptsResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An opaque token representing the pagination position after the last returned result.
	// If present, there may be more results available.
	NextCursor *string  `json:"nextCursor,omitempty"`
	Prompts    []Prompt `json:"prompts"`
}

// A template description for resources available on the server.
type ResourceTemplate struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	// A description of what this template is for.
	//
	// This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
	Description *string `json:"description,omitempty"`
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
	// The MIME type for all resources that match this template. This should only be included if all resources matching this template have the same type.
	MimeType *string `json:"mimeType,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
	// A URI template (according to RFC 6570) that can be used to construct resource URIs.
	URITemplate string `json:"uriTemplate"`
}

// The server's response to a resources/templates/list request from the client.
type ListResourceTemplatesResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An opaque token representing the pagination position after the last returned result.
	// If present, there may be more results available.
	NextCursor        *string            `json:"nextCursor,omitempty"`
	ResourceTemplates []ResourceTemplate `json:"resourceTemplates"`
}

// A known resource that the server is capable of reading.
type Resource struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// Optional annotations for the client.
	Annotations *Annotations `json:"annotations,omitempty"`
	// A description of what this resource represents.
	//
	// This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
	Description *string `json:"description,omitempty"`
	// Optional set of sized icons that the client can display in a user interface.
	//
	// Clients that support rendering icons MUST support at least the following MIME types:
	// - `image/png` - PNG images (safe, universal compatibility)
	// - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
	//
	// Clients that support rendering icons SHOULD also support:
	// - `image/svg+xml` - SVG images (scalable but requires security precautions)
	// - `image/webp` - WebP images (modern, efficient format)
	Icons []Icon `json:"icons,omitempty"`
	// The MIME type of this resource, if known.
	MimeType *string `json:"mimeType,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
	//
	// This can be used by Hosts to display file sizes and estimate context window usage.
	Size *int `json:"size,omitempty"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
	// The URI of this resource.
	URI url.URL `json:"uri"`
}

// The server's response to a resources/list request from the client.
type ListResourcesResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An opaque token representing the pagination position after the last returned result.
	// If present, there may be more results available.
	NextCursor *string    `json:"nextCursor,omitempty"`
	Resources  []Resource `json:"resources"`
}

// Sent from the server to request a list of root URIs from the client. Roots allow
// servers to ask for specific directories or files to operate on. A common example
// for roots is providing a set of repositories or directories a server should operate
// on.
//
// This request is typically used when the server needs to understand the file system
// structure or access specific locations that the client has permission to read from.
type ListRootsRequest struct {
	ID      RequestId      `json:"id"`
	Jsonrpc string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  *RequestParams `json:"params,omitempty"`
}

// The server's response to a tools/list request from the client.
type ListToolsResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An opaque token representing the pagination position after the last returned result.
	// If present, there may be more results available.
	NextCursor *string `json:"nextCursor,omitempty"`
	Tools      []Tool  `json:"tools"`
}

// Parameters for a `notifications/message` notification.
type LoggingMessageNotificationParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// The data to be logged, such as a string message or an object. Any JSON serializable type is allowed here.
	Data interface{} `json:"data"`
	// The severity of this log message.
	Level LoggingLevel `json:"level"`
	// An optional name of the logger issuing this message.
	Logger *string `json:"logger,omitempty"`
}

// JSONRPCNotification of a log message passed from server to client. If no logging/setLevel request has been sent from the client, the server MAY decide which messages to send automatically.
type LoggingMessageNotification struct {
	Jsonrpc string                           `json:"jsonrpc"`
	Method  string                           `json:"method"`
	Params  LoggingMessageNotificationParams `json:"params"`
}

type MultiSelectEnumSchema = interface{}

type Notification struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type PaginatedRequest struct {
	ID      RequestId               `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  *PaginatedRequestParams `json:"params,omitempty"`
}

type PaginatedResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// An opaque token representing the pagination position after the last returned result.
	// If present, there may be more results available.
	NextCursor *string `json:"nextCursor,omitempty"`
}

// An optional notification from the server to the client, informing it that the list of prompts it offers has changed. This may be issued by servers without any previous subscription from the client.
type PromptListChangedNotification struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  *NotificationParams `json:"params,omitempty"`
}

// Identifies a prompt.
type PromptReference struct {
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
	// even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool,
	// where `annotations.title` should be given precedence over using `name`,
	// if present).
	Title *string `json:"title,omitempty"`
	Type  string  `json:"type"`
}

// The server's response to a resources/read request from the client.
type ReadResourceResult struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta     map[string]interface{} `json:"_meta,omitempty"`
	Contents []interface{}          `json:"contents"`
}

// Metadata for associating messages with a task.
// Include this in the `_meta` field under the key `io.modelcontextprotocol/related-task`.
type RelatedTaskMetadata struct {
	// The task identifier this message is associated with.
	TaskID string `json:"taskId"`
}

type Request struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

// The contents of a specific resource or sub-resource.
type ResourceContents struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// The MIME type of this resource, if known.
	MimeType *string `json:"mimeType,omitempty"`
	// The URI of this resource.
	URI url.URL `json:"uri"`
}

// An optional notification from the server to the client, informing it that the list of resources it can read from has changed. This may be issued by servers without any previous subscription from the client.
type ResourceListChangedNotification struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  *NotificationParams `json:"params,omitempty"`
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type ResourceRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Common parameters when working with resources.
type ResourceRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *ResourceRequestParamsMeta `json:"_meta,omitempty"`
	// The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
	URI url.URL `json:"uri"`
}

// A reference to a resource or resource template definition.
type ResourceTemplateReference struct {
	Type string `json:"type"`
	// The URI or URI template of the resource.
	URI string `json:"uri"`
}

// Parameters for a `notifications/resources/updated` notification.
type ResourceUpdatedNotificationParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// The URI of the resource that has been updated. This might be a sub-resource of the one that the client actually subscribed to.
	URI url.URL `json:"uri"`
}

// A notification from the server to the client, informing it that a resource has changed and may need to be read again. This should only be sent if the client previously sent a resources/subscribe request.
type ResourceUpdatedNotification struct {
	Jsonrpc string                            `json:"jsonrpc"`
	Method  string                            `json:"method"`
	Params  ResourceUpdatedNotificationParams `json:"params"`
}

// The result of a tool use, provided by the user back to the assistant.
type ToolResultContent struct {
	// Optional metadata about the tool result. Clients SHOULD preserve this field when
	// including tool results in subsequent sampling requests to enable caching optimizations.
	//
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// The unstructured result content of the tool use.
	//
	// This has the same format as CallToolResult.content and can include text, images,
	// audio, resource links, and embedded resources.
	Content []ContentBlock `json:"content"`
	// Whether the tool use resulted in an error.
	//
	// If true, the content typically describes the error that occurred.
	// Default: false
	IsError *bool `json:"isError,omitempty"`
	// An optional structured result object.
	//
	// If the tool defined an outputSchema, this SHOULD conform to that schema.
	StructuredContent map[string]interface{} `json:"structuredContent,omitempty"`
	// The ID of the tool use this result corresponds to.
	//
	// This MUST match the ID from a previous ToolUseContent.
	ToolUseID string `json:"toolUseId"`
	Type      string `json:"type"`
}

// A request from the assistant to call a tool.
type ToolUseContent struct {
	// Optional metadata about the tool use. Clients SHOULD preserve this field when
	// including tool uses in subsequent sampling requests to enable caching optimizations.
	//
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// A unique identifier for this tool use.
	//
	// This ID is used to match tool results to their corresponding tool uses.
	ID string `json:"id"`
	// The arguments to pass to the tool, conforming to the tool's input schema.
	Input map[string]interface{} `json:"input"`
	// The name of the tool to call.
	Name string `json:"name"`
	Type string `json:"type"`
}

type SamplingMessageContentBlock = interface{}

// An optional notification from the server to the client, informing it that the list of tools it offers has changed. This may be issued by servers without any previous subscription from the client.
type ToolListChangedNotification struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  *NotificationParams `json:"params,omitempty"`
}

type ServerNotification = interface{}

type ServerRequest = interface{}

type ServerResult = interface{}

type SingleSelectEnumSchema = interface{}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
type TaskAugmentedRequestParamsMeta struct {
	// If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
	ProgressToken *ProgressToken `json:"progressToken,omitempty"`
}

// Common params for any task-augmented request.
type TaskAugmentedRequestParams struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta *TaskAugmentedRequestParamsMeta `json:"_meta,omitempty"`
	// If specified, the caller is requesting task-augmented execution for this request.
	// The request will return a CreateTaskResult immediately, and the actual result can be
	// retrieved later via tasks/result.
	//
	// Task augmentation is subject to capability negotiation - receivers MUST declare support
	// for task augmentation of specific request types in their capabilities.
	Task *TaskMetadata `json:"task,omitempty"`
}

type TextResourceContents struct {
	// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
	Meta map[string]interface{} `json:"_meta,omitempty"`
	// The MIME type of this resource, if known.
	MimeType *string `json:"mimeType,omitempty"`
	// The text of the item. This must only be set if the item can actually be represented as text (not binary data).
	Text string `json:"text"`
	// The URI of this resource.
	URI url.URL `json:"uri"`
}

// An error response that indicates that the server requires the client to provide additional information via an elicitation request.
type URLElicitationRequiredError struct {
	Error   interface{} `json:"error"`
	ID      *RequestId  `json:"id,omitempty"`
	Jsonrpc string      `json:"jsonrpc"`
}
