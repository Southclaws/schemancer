
// The sender or recipient of messages and data in a conversation.
export type Role =
  | "assistant"
  | "user";

// Optional annotations for the client. The client can use annotations to inform how objects are used or displayed
export interface Annotations {
  // Describes who the intended audience of this object or data is.
  // 
  // It can include multiple entries to indicate content useful for multiple audiences (e.g., `["user", "assistant"]`).
  audience?: Role[];
  // The moment the resource was last modified, as an ISO 8601 formatted string.
  // 
  // Should be an ISO 8601 formatted string (e.g., "2025-01-12T15:00:58Z").
  // 
  // Examples: last activity timestamp in an open file, timestamp when the resource
  // was attached, etc.
  lastModified?: string;
  // Describes how important this data is for operating the server.
  // 
  // A value of 1 means "most important," and indicates that the data is
  // effectively required, while 0 means "least important," and indicates that
  // the data is entirely optional.
  priority?: number;
}

// Audio provided to or from an LLM.
export interface AudioContent {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  // The base64-encoded audio data.
  data: string;
  // The MIME type of the audio. Different providers may support different audio types.
  mimeType: string;
  type: string;
}

// Base interface for metadata with name (identifier) and title (display name) properties.
export interface BaseMetadata {
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
}

export interface BlobResourceContents {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // A base64-encoded string representing the binary data of the item.
  blob: string;
  // The MIME type of this resource, if known.
  mimeType?: string;
  // The URI of this resource.
  uri: string;
}

export interface BooleanSchema {
  default?: boolean;
  description?: string;
  title?: string;
  type: string;
}

// A progress token, used to associate progress notifications with the original request.
export type ProgressToken = unknown;

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface CallToolRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Metadata for augmenting a request with task execution.
// Include this in the `task` field of the request parameters.
export interface TaskMetadata {
  // Requested duration in milliseconds to retain task from creation.
  ttl?: number;
}

// Parameters for a `tools/call` request.
export interface CallToolRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: CallToolRequestParamsMeta;
  // Arguments to use for the tool call.
  arguments?: Record<string, unknown>;
  // The name of the tool.
  name: string;
  // If specified, the caller is requesting task-augmented execution for this request.
  // The request will return a CreateTaskResult immediately, and the actual result can be
  // retrieved later via tasks/result.
  // 
  // Task augmentation is subject to capability negotiation - receivers MUST declare support
  // for task augmentation of specific request types in their capabilities.
  task?: TaskMetadata;
}

// A uniquely identifying ID for a request in JSON-RPC.
export type RequestId = unknown;

// Used by the client to invoke a tool provided by the server.
export interface CallToolRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: CallToolRequestParams;
}

// The contents of a resource, embedded into a prompt or tool call result.
// 
// It is up to the client how best to render embedded resources for the benefit
// of the LLM and/or the user.
export interface EmbeddedResource {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  resource: unknown;
  type: string;
}

// An image provided to or from an LLM.
export interface ImageContent {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  // The base64-encoded image data.
  data: string;
  // The MIME type of the image. Different providers may support different image types.
  mimeType: string;
  type: string;
}

// An optionally-sized icon that can be displayed in a user interface.
export interface Icon {
  // Optional MIME type override if the source MIME type is missing or generic.
  // For example: `"image/png"`, `"image/jpeg"`, or `"image/svg+xml"`.
  mimeType?: string;
  // Optional array of strings that specify sizes at which the icon can be used.
  // Each string should be in WxH format (e.g., `"48x48"`, `"96x96"`) or `"any"` for scalable formats like SVG.
  // 
  // If not provided, the client should assume that the icon can be used at any size.
  sizes?: string[];
  // A standard URI pointing to an icon resource. May be an HTTP/HTTPS URL or a
  // `data:` URI with Base64-encoded image data.
  // 
  // Consumers SHOULD takes steps to ensure URLs serving icons are from the
  // same domain as the client/server or a trusted domain.
  // 
  // Consumers SHOULD take appropriate precautions when consuming SVGs as they can contain
  // executable JavaScript.
  src: string;
  // Optional specifier for the theme this icon is designed for. `light` indicates
  // the icon is designed to be used with a light background, and `dark` indicates
  // the icon is designed to be used with a dark background.
  // 
  // If not provided, the client should assume the icon can be used with any theme.
  theme?: string;
}

// A resource that the server is capable of reading, included in a prompt or tool call result.
// 
// Note: resource links returned by tools are not guaranteed to appear in the results of `resources/list` requests.
export interface ResourceLink {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  // A description of what this resource represents.
  // 
  // This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
  description?: string;
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
  // The MIME type of this resource, if known.
  mimeType?: string;
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
  // 
  // This can be used by Hosts to display file sizes and estimate context window usage.
  size?: number;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
  type: string;
  // The URI of this resource.
  uri: string;
}

// Text provided to or from an LLM.
export interface TextContent {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  // The text content of the message.
  text: string;
  type: string;
}

export type ContentBlock = TextContent | ImageContent | AudioContent | ResourceLink | EmbeddedResource;

// The server's response to a tool call.
export interface CallToolResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // A list of content objects that represent the unstructured result of the tool call.
  content: ContentBlock[];
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
  isError?: boolean;
  // An optional JSON object that represents the structured result of the tool call.
  structuredContent?: Record<string, unknown>;
}

export interface CancelTaskRequestParams {
  // The task identifier to cancel.
  taskId: string;
}

// A request to cancel a task.
export interface CancelTaskRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: CancelTaskRequestParams;
}


// The status of a task.
export type TaskStatus =
  | "cancelled"
  | "completed"
  | "failed"
  | "input_required"
  | "working";

// The response to a tasks/cancel request.
export interface CancelTaskResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // ISO 8601 timestamp when the task was created.
  createdAt: string;
  // ISO 8601 timestamp when the task was last updated.
  lastUpdatedAt: string;
  // Suggested polling interval in milliseconds.
  pollInterval?: number;
  // Current task state.
  status: TaskStatus;
  // Optional human-readable message describing the current task state.
  // This can provide context for any status, including:
  // - Reasons for "cancelled" status
  // - Summaries for "completed" status
  // - Diagnostic information for "failed" status (e.g., error details, what went wrong)
  statusMessage?: string;
  // The task identifier.
  taskId: string;
  // Actual retention duration from creation in milliseconds, null for unlimited.
  ttl: number;
}

// Parameters for a `notifications/cancelled` notification.
export interface CancelledNotificationParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An optional string describing the reason for the cancellation. This MAY be logged or presented to the user.
  reason?: string;
  // The ID of the request to cancel.
  // 
  // This MUST correspond to the ID of a request previously issued in the same direction.
  // This MUST be provided for cancelling non-task requests.
  // This MUST NOT be used for cancelling tasks (use the `tasks/cancel` request instead).
  requestId?: RequestId;
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
export interface CancelledNotification {
  jsonrpc: string;
  method: string;
  params: CancelledNotificationParams;
}

export interface ClientCapabilitiesElicitationForm {
}

export interface ClientCapabilitiesElicitationURL {
}

// Present if the client supports elicitation from the server.
export interface ClientCapabilitiesElicitation {
  form?: ClientCapabilitiesElicitationForm;
  url?: ClientCapabilitiesElicitationURL;
}

export interface ClientCapabilitiesExperimentalValue {
}

// Present if the client supports listing roots.
export interface ClientCapabilitiesRoots {
  // Whether the client supports notifications for changes to the roots list.
  listChanged?: boolean;
}

// Whether the client supports context inclusion via includeContext parameter.
// If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
export interface ClientCapabilitiesSamplingContext {
}

// Whether the client supports tool use via tools and toolChoice parameters.
export interface ClientCapabilitiesSamplingTools {
}

// Present if the client supports sampling from an LLM.
export interface ClientCapabilitiesSampling {
  // Whether the client supports context inclusion via includeContext parameter.
  // If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
  context?: ClientCapabilitiesSamplingContext;
  // Whether the client supports tool use via tools and toolChoice parameters.
  tools?: ClientCapabilitiesSamplingTools;
}

// Whether this client supports tasks/cancel.
export interface ClientCapabilitiesTasksCancel {
}

// Whether this client supports tasks/list.
export interface ClientCapabilitiesTasksList {
}

// Whether the client supports task-augmented elicitation/create requests.
export interface ClientCapabilitiesTasksRequestsElicitationCreate {
}

// Task support for elicitation-related requests.
export interface ClientCapabilitiesTasksRequestsElicitation {
  // Whether the client supports task-augmented elicitation/create requests.
  create?: ClientCapabilitiesTasksRequestsElicitationCreate;
}

// Whether the client supports task-augmented sampling/createMessage requests.
export interface ClientCapabilitiesTasksRequestsSamplingCreateMessage {
}

// Task support for sampling-related requests.
export interface ClientCapabilitiesTasksRequestsSampling {
  // Whether the client supports task-augmented sampling/createMessage requests.
  createMessage?: ClientCapabilitiesTasksRequestsSamplingCreateMessage;
}

// Specifies which request types can be augmented with tasks.
export interface ClientCapabilitiesTasksRequests {
  // Task support for elicitation-related requests.
  elicitation?: ClientCapabilitiesTasksRequestsElicitation;
  // Task support for sampling-related requests.
  sampling?: ClientCapabilitiesTasksRequestsSampling;
}

// Present if the client supports task-augmented requests.
export interface ClientCapabilitiesTasks {
  // Whether this client supports tasks/cancel.
  cancel?: ClientCapabilitiesTasksCancel;
  // Whether this client supports tasks/list.
  list?: ClientCapabilitiesTasksList;
  // Specifies which request types can be augmented with tasks.
  requests?: ClientCapabilitiesTasksRequests;
}

// Capabilities a client may support. Known capabilities are defined here, in this schema, but this is not a closed set: any client can define its own, additional capabilities.
export interface ClientCapabilities {
  // Present if the client supports elicitation from the server.
  elicitation?: ClientCapabilitiesElicitation;
  // Experimental, non-standard capabilities that the client supports.
  experimental?: Record<string, ClientCapabilitiesExperimentalValue>;
  // Present if the client supports listing roots.
  roots?: ClientCapabilitiesRoots;
  // Present if the client supports sampling from an LLM.
  sampling?: ClientCapabilitiesSampling;
  // Present if the client supports task-augmented requests.
  tasks?: ClientCapabilitiesTasks;
}

export interface NotificationParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
}

// This notification is sent from the client to the server after initialization has finished.
export interface InitializedNotification {
  jsonrpc: string;
  method: string;
  params?: NotificationParams;
}

// Parameters for a `notifications/progress` notification.
export interface ProgressNotificationParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An optional message describing the current progress.
  message?: string;
  // The progress thus far. This should increase every time progress is made, even if the total is unknown.
  progress: number;
  // The progress token which was given in the initial request, used to associate this notification with the request that is proceeding.
  progressToken: ProgressToken;
  // Total number of items to process (or total progress required), if known.
  total?: number;
}

// An out-of-band notification used to inform the receiver of a progress update for a long-running request.
export interface ProgressNotification {
  jsonrpc: string;
  method: string;
  params: ProgressNotificationParams;
}

// A notification from the client to the server, informing it that the list of roots has changed.
// This notification should be sent whenever the client adds, removes, or modifies any root.
// The server should then request an updated list of roots using the ListRootsRequest.
export interface RootsListChangedNotification {
  jsonrpc: string;
  method: string;
  params?: NotificationParams;
}

// Parameters for a `notifications/tasks/status` notification.
export interface TaskStatusNotificationParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // ISO 8601 timestamp when the task was created.
  createdAt: string;
  // ISO 8601 timestamp when the task was last updated.
  lastUpdatedAt: string;
  // Suggested polling interval in milliseconds.
  pollInterval?: number;
  // Current task state.
  status: TaskStatus;
  // Optional human-readable message describing the current task state.
  // This can provide context for any status, including:
  // - Reasons for "cancelled" status
  // - Summaries for "completed" status
  // - Diagnostic information for "failed" status (e.g., error details, what went wrong)
  statusMessage?: string;
  // The task identifier.
  taskId: string;
  // Actual retention duration from creation in milliseconds, null for unlimited.
  ttl: number;
}

// An optional notification from the receiver to the requestor, informing them that a task's status has changed. Receivers are not required to send these notifications.
export interface TaskStatusNotification {
  jsonrpc: string;
  method: string;
  params: TaskStatusNotificationParams;
}

export type ClientNotification = CancelledNotification | InitializedNotification | ProgressNotification | TaskStatusNotification | RootsListChangedNotification;

// The argument's information
export interface CompleteRequestParamsArgument {
  // The name of the argument
  name: string;
  // The value of the argument to use for completion matching.
  value: string;
}

// Additional, optional context for completions
export interface CompleteRequestParamsContext {
  // Previously-resolved variables in a URI template or prompt.
  arguments?: Record<string, string>;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface CompleteRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for a `completion/complete` request.
export interface CompleteRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: CompleteRequestParamsMeta;
  // The argument's information
  argument: CompleteRequestParamsArgument;
  // Additional, optional context for completions
  context?: CompleteRequestParamsContext;
  ref: unknown;
}

// A request from the client to the server, to ask for completion options.
export interface CompleteRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: CompleteRequestParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface GetPromptRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for a `prompts/get` request.
export interface GetPromptRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: GetPromptRequestParamsMeta;
  // Arguments to use for templating the prompt.
  arguments?: Record<string, string>;
  // The name of the prompt or prompt template.
  name: string;
}

// Used by the client to get a prompt provided by the server.
export interface GetPromptRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: GetPromptRequestParams;
}

export interface GetTaskPayloadRequestParams {
  // The task identifier to retrieve results for.
  taskId: string;
}

// A request to retrieve the result of a completed task.
export interface GetTaskPayloadRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: GetTaskPayloadRequestParams;
}

export interface GetTaskRequestParams {
  // The task identifier to query.
  taskId: string;
}

// A request to retrieve the state of a task.
export interface GetTaskRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: GetTaskRequestParams;
}

// Describes the MCP implementation.
export interface Implementation {
  // An optional human-readable description of what this implementation does.
  // 
  // This can be used by clients or servers to provide context about their purpose
  // and capabilities. For example, a server might describe the types of resources
  // or tools it provides, while a client might describe its intended use case.
  description?: string;
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
  version: string;
  // An optional URL of the website for this implementation.
  websiteUrl?: string;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface InitializeRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for an `initialize` request.
export interface InitializeRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: InitializeRequestParamsMeta;
  capabilities: ClientCapabilities;
  clientInfo: Implementation;
  // The latest version of the Model Context Protocol that the client supports. The client MAY decide to support older versions as well.
  protocolVersion: string;
}

// This request is sent from the client to the server when it first connects, asking it to begin initialization.
export interface InitializeRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: InitializeRequestParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface PaginatedRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Common parameters for paginated requests.
export interface PaginatedRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: PaginatedRequestParamsMeta;
  // An opaque token representing the current pagination position.
  // If provided, the server should return results starting after this cursor.
  cursor?: string;
}

// Sent from the client to request a list of prompts and prompt templates the server has.
export interface ListPromptsRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: PaginatedRequestParams;
}

// Sent from the client to request a list of resource templates the server has.
export interface ListResourceTemplatesRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: PaginatedRequestParams;
}

// Sent from the client to request a list of resources the server has.
export interface ListResourcesRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: PaginatedRequestParams;
}

// A request to retrieve a list of tasks.
export interface ListTasksRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: PaginatedRequestParams;
}

// Sent from the client to request a list of tools the server has.
export interface ListToolsRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: PaginatedRequestParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface RequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Common params for any request.
export interface RequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: RequestParamsMeta;
}

// A ping, issued by either the server or the client, to check that the other party is still alive. The receiver must promptly respond, or else may be disconnected.
export interface PingRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: RequestParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface ReadResourceRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for a `resources/read` request.
export interface ReadResourceRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: ReadResourceRequestParamsMeta;
  // The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
  uri: string;
}

// Sent from the client to the server, to read a specific resource URI.
export interface ReadResourceRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: ReadResourceRequestParams;
}


// The severity of a log message.
// 
// These map to syslog message severities, as specified in RFC-5424:
// https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1
export type LoggingLevel =
  | "alert"
  | "critical"
  | "debug"
  | "emergency"
  | "error"
  | "info"
  | "notice"
  | "warning";

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface SetLevelRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for a `logging/setLevel` request.
export interface SetLevelRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: SetLevelRequestParamsMeta;
  // The level of logging that the client wants to receive from the server. The server should send all logs at this level and higher (i.e., more severe) to the client as notifications/message.
  level: LoggingLevel;
}

// A request from the client to the server, to enable or adjust logging.
export interface SetLevelRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: SetLevelRequestParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface SubscribeRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for a `resources/subscribe` request.
export interface SubscribeRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: SubscribeRequestParamsMeta;
  // The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
  uri: string;
}

// Sent from the client to request resources/updated notifications from the server whenever a particular resource changes.
export interface SubscribeRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: SubscribeRequestParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface UnsubscribeRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Parameters for a `resources/unsubscribe` request.
export interface UnsubscribeRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: UnsubscribeRequestParamsMeta;
  // The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
  uri: string;
}

// Sent from the client to request cancellation of resources/updated notifications from the server. This should follow a previous resources/subscribe request.
export interface UnsubscribeRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: UnsubscribeRequestParams;
}

export type ClientRequest = InitializeRequest | PingRequest | ListResourcesRequest | ListResourceTemplatesRequest | ReadResourceRequest | SubscribeRequest | UnsubscribeRequest | ListPromptsRequest | GetPromptRequest | ListToolsRequest | CallToolRequest | GetTaskRequest | GetTaskPayloadRequest | CancelTaskRequest | ListTasksRequest | SetLevelRequest | CompleteRequest;

// The client's response to a sampling/createMessage request from the server.
// The client should inform the user before returning the sampled message, to allow them
// to inspect the response (human in the loop) and decide whether to allow the server to see it.
export interface CreateMessageResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  content: unknown;
  // The name of the model that generated the message.
  model: string;
  role: Role;
  // The reason why sampling stopped, if known.
  // 
  // Standard values:
  // - "endTurn": Natural end of the assistant's turn
  // - "stopSequence": A stop sequence was encountered
  // - "maxTokens": Maximum token limit was reached
  // - "toolUse": The model wants to use one or more tools
  // 
  // This field is an open string to allow for provider-specific stop reasons.
  stopReason?: string;
}

// The client's response to an elicitation request.
export interface ElicitResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // The user action in response to the elicitation.
  // - "accept": User submitted the form/confirmed the action
  // - "decline": User explicitly decline the action
  // - "cancel": User dismissed without making an explicit choice
  action: string;
  // The submitted form data, only present when action is "accept" and mode was "form".
  // Contains values matching the requested schema.
  // Omitted for out-of-band mode responses.
  content?: Record<string, unknown>;
}

// The response to a tasks/result request.
// The structure matches the result type of the original request.
// For example, a tools/call task would return the CallToolResult structure.
export interface GetTaskPayloadResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
}

// The response to a tasks/get request.
export interface GetTaskResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // ISO 8601 timestamp when the task was created.
  createdAt: string;
  // ISO 8601 timestamp when the task was last updated.
  lastUpdatedAt: string;
  // Suggested polling interval in milliseconds.
  pollInterval?: number;
  // Current task state.
  status: TaskStatus;
  // Optional human-readable message describing the current task state.
  // This can provide context for any status, including:
  // - Reasons for "cancelled" status
  // - Summaries for "completed" status
  // - Diagnostic information for "failed" status (e.g., error details, what went wrong)
  statusMessage?: string;
  // The task identifier.
  taskId: string;
  // Actual retention duration from creation in milliseconds, null for unlimited.
  ttl: number;
}

// Represents a root directory or file that the server can operate on.
export interface Root {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An optional name for the root. This can be used to provide a human-readable
  // identifier for the root, which may be useful for display purposes or for
  // referencing the root in other parts of the application.
  name?: string;
  // The URI identifying the root. This *must* start with file:// for now.
  // This restriction may be relaxed in future versions of the protocol to allow
  // other URI schemes.
  uri: string;
}

// The client's response to a roots/list request from the server.
// This result contains an array of Root objects, each representing a root directory
// or file that the server can operate on.
export interface ListRootsResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  roots: Root[];
}

// Data associated with a task.
export interface Task {
  // ISO 8601 timestamp when the task was created.
  createdAt: string;
  // ISO 8601 timestamp when the task was last updated.
  lastUpdatedAt: string;
  // Suggested polling interval in milliseconds.
  pollInterval?: number;
  // Current task state.
  status: TaskStatus;
  // Optional human-readable message describing the current task state.
  // This can provide context for any status, including:
  // - Reasons for "cancelled" status
  // - Summaries for "completed" status
  // - Diagnostic information for "failed" status (e.g., error details, what went wrong)
  statusMessage?: string;
  // The task identifier.
  taskId: string;
  // Actual retention duration from creation in milliseconds, null for unlimited.
  ttl: number;
}

// The response to a tasks/list request.
export interface ListTasksResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An opaque token representing the pagination position after the last returned result.
  // If present, there may be more results available.
  nextCursor?: string;
  tasks: Task[];
}

export interface Result {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
}

export type ClientResult = Result | GetTaskResult | GetTaskPayloadResult | CancelTaskResult | ListTasksResult | CreateMessageResult | ListRootsResult | ElicitResult;

export interface CompleteResultCompletion {
  // Indicates whether there are additional completion options beyond those provided in the current response, even if the exact total is unknown.
  hasMore?: boolean;
  // The total number of completion options available. This can exceed the number of values actually sent in the response.
  total?: number;
  // An array of completion values. Must not exceed 100 items.
  values: string[];
}

// The server's response to a completion/complete request
export interface CompleteResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  completion: CompleteResultCompletion;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface CreateMessageRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific.
export interface CreateMessageRequestParamsMetadata {
}

// Hints to use for model selection.
// 
// Keys not declared here are currently left unspecified by the spec and are up
// to the client to interpret.
export interface ModelHint {
  // A hint for a model name.
  // 
  // The client SHOULD treat this as a substring of a model name; for example:
  //  - `claude-3-5-sonnet` should match `claude-3-5-sonnet-20241022`
  //  - `sonnet` should match `claude-3-5-sonnet-20241022`, `claude-3-sonnet-20240229`, etc.
  //  - `claude` should match any Claude model
  // 
  // The client MAY also map the string to a different provider's model name or a different model family, as long as it fills a similar niche; for example:
  //  - `gemini-1.5-flash` could match `claude-3-haiku-20240307`
  name?: string;
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
export interface ModelPreferences {
  // How much to prioritize cost when selecting a model. A value of 0 means cost
  // is not important, while a value of 1 means cost is the most important
  // factor.
  costPriority?: number;
  // Optional hints to use for model selection.
  // 
  // If multiple hints are specified, the client MUST evaluate them in order
  // (such that the first match is taken).
  // 
  // The client SHOULD prioritize these hints over the numeric priorities, but
  // MAY still use the priorities to select from ambiguous matches.
  hints?: ModelHint[];
  // How much to prioritize intelligence and capabilities when selecting a
  // model. A value of 0 means intelligence is not important, while a value of 1
  // means intelligence is the most important factor.
  intelligencePriority?: number;
  // How much to prioritize sampling speed (latency) when selecting a model. A
  // value of 0 means speed is not important, while a value of 1 means speed is
  // the most important factor.
  speedPriority?: number;
}

// Describes a message issued to or received from an LLM API.
export interface SamplingMessage {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  content: unknown;
  role: Role;
}

// Additional properties describing a Tool to clients.
// 
// NOTE: all properties in ToolAnnotations are **hints**.
// They are not guaranteed to provide a faithful description of
// tool behavior (including descriptive properties like `title`).
// 
// Clients should never make tool use decisions based on ToolAnnotations
// received from untrusted servers.
export interface ToolAnnotations {
  // If true, the tool may perform destructive updates to its environment.
  // If false, the tool performs only additive updates.
  // 
  // (This property is meaningful only when `readOnlyHint == false`)
  // 
  // Default: true
  destructiveHint?: boolean;
  // If true, calling the tool repeatedly with the same arguments
  // will have no additional effect on its environment.
  // 
  // (This property is meaningful only when `readOnlyHint == false`)
  // 
  // Default: false
  idempotentHint?: boolean;
  // If true, this tool may interact with an "open world" of external
  // entities. If false, the tool's domain of interaction is closed.
  // For example, the world of a web search tool is open, whereas that
  // of a memory tool is not.
  // 
  // Default: true
  openWorldHint?: boolean;
  // If true, the tool does not modify its environment.
  // 
  // Default: false
  readOnlyHint?: boolean;
  // A human-readable title for the tool.
  title?: string;
}

// Execution-related properties for a tool.
export interface ToolExecution {
  // Indicates whether this tool supports task-augmented execution.
  // This allows clients to handle long-running operations through polling
  // the task system.
  // 
  // - "forbidden": Tool does not support task-augmented execution (default when absent)
  // - "optional": Tool may support task-augmented execution
  // - "required": Tool requires task-augmented execution
  // 
  // Default: "forbidden"
  taskSupport?: string;
}

export interface ToolInputSchemaPropertiesValue {
}

// A JSON Schema object defining the expected parameters for the tool.
export interface ToolInputSchema {
  $schema?: string;
  properties?: Record<string, ToolInputSchemaPropertiesValue>;
  required?: string[];
  type: string;
}

export interface ToolOutputSchemaPropertiesValue {
}

// An optional JSON Schema object defining the structure of the tool's output returned in
// the structuredContent field of a CallToolResult.
// 
// Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
// Currently restricted to type: "object" at the root level.
export interface ToolOutputSchema {
  $schema?: string;
  properties?: Record<string, ToolOutputSchemaPropertiesValue>;
  required?: string[];
  type: string;
}

// Definition for a tool the client can call.
export interface Tool {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional additional tool information.
  // 
  // Display name precedence order is: title, annotations.title, then name.
  annotations?: ToolAnnotations;
  // A human-readable description of the tool.
  // 
  // This can be used by clients to improve the LLM's understanding of available tools. It can be thought of like a "hint" to the model.
  description?: string;
  // Execution-related properties for this tool.
  execution?: ToolExecution;
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
  // A JSON Schema object defining the expected parameters for the tool.
  inputSchema: ToolInputSchema;
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // An optional JSON Schema object defining the structure of the tool's output returned in
  // the structuredContent field of a CallToolResult.
  // 
  // Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
  // Currently restricted to type: "object" at the root level.
  outputSchema?: ToolOutputSchema;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
}

// Controls tool selection behavior for sampling requests.
export interface ToolChoice {
  // Controls the tool use ability of the model:
  // - "auto": Model decides whether to use tools (default)
  // - "required": Model MUST use at least one tool before completing
  // - "none": Model MUST NOT use any tools
  mode?: string;
}

// Parameters for a `sampling/createMessage` request.
export interface CreateMessageRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: CreateMessageRequestParamsMeta;
  // A request to include context from one or more MCP servers (including the caller), to be attached to the prompt.
  // The client MAY ignore this request.
  // 
  // Default is "none". Values "thisServer" and "allServers" are soft-deprecated. Servers SHOULD only use these values if the client
  // declares ClientCapabilities.sampling.context. These values may be removed in future spec releases.
  includeContext?: string;
  // The requested maximum number of tokens to sample (to prevent runaway completions).
  // 
  // The client MAY choose to sample fewer tokens than the requested maximum.
  maxTokens: number;
  messages: SamplingMessage[];
  // Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific.
  metadata?: CreateMessageRequestParamsMetadata;
  // The server's preferences for which model to select. The client MAY ignore these preferences.
  modelPreferences?: ModelPreferences;
  stopSequences?: string[];
  // An optional system prompt the server wants to use for sampling. The client MAY modify or omit this prompt.
  systemPrompt?: string;
  // If specified, the caller is requesting task-augmented execution for this request.
  // The request will return a CreateTaskResult immediately, and the actual result can be
  // retrieved later via tasks/result.
  // 
  // Task augmentation is subject to capability negotiation - receivers MUST declare support
  // for task augmentation of specific request types in their capabilities.
  task?: TaskMetadata;
  temperature?: number;
  // Controls how the model uses tools.
  // The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
  // Default is `{ mode: "auto" }`.
  toolChoice?: ToolChoice;
  // Tools that the model may use during generation.
  // The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
  tools?: Tool[];
}

// A request from the server to sample an LLM via the client. The client has full discretion over which model to select. The client should also inform the user before beginning sampling, to allow them to inspect the request (human in the loop) and decide whether to approve it.
export interface CreateMessageRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: CreateMessageRequestParams;
}

// A response to a task-augmented request.
export interface CreateTaskResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  task: Task;
}

// An opaque token used to represent a cursor for pagination.
export type Cursor = string;

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface ElicitRequestFormParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Use TitledSingleSelectEnumSchema instead.
// This interface will be removed in a future version.
export interface LegacyTitledEnumSchema {
  default?: string;
  description?: string;
  enum: string[];
  // (Legacy) Display names for enum values.
  // Non-standard according to JSON schema 2020-12.
  enumNames?: string[];
  title?: string;
  type: string;
}

export interface NumberSchema {
  default?: number;
  description?: string;
  maximum?: number;
  minimum?: number;
  title?: string;
  type: string;
}

export interface StringSchema {
  default?: string;
  description?: string;
  format?: string;
  maxLength?: number;
  minLength?: number;
  title?: string;
  type: string;
}

export interface TitledMultiSelectEnumSchemaItemsAnyOfItem {
  // The constant enum value.
  const: string;
  // Display title for this option.
  title: string;
}

// Schema for array items with enum options and display labels.
export interface TitledMultiSelectEnumSchemaItems {
  // Array of enum options with values and display labels.
  anyOf: TitledMultiSelectEnumSchemaItemsAnyOfItem[];
}

// Schema for multiple-selection enumeration with display titles for each option.
export interface TitledMultiSelectEnumSchema {
  // Optional default value.
  default?: string[];
  // Optional description for the enum field.
  description?: string;
  // Schema for array items with enum options and display labels.
  items: TitledMultiSelectEnumSchemaItems;
  // Maximum number of items to select.
  maxItems?: number;
  // Minimum number of items to select.
  minItems?: number;
  // Optional title for the enum field.
  title?: string;
  type: string;
}

export interface TitledSingleSelectEnumSchemaOneOfItem {
  // The enum value.
  const: string;
  // Display label for this option.
  title: string;
}

// Schema for single-selection enumeration with display titles for each option.
export interface TitledSingleSelectEnumSchema {
  // Optional default value.
  default?: string;
  // Optional description for the enum field.
  description?: string;
  // Array of enum options with values and display labels.
  oneOf: TitledSingleSelectEnumSchemaOneOfItem[];
  // Optional title for the enum field.
  title?: string;
  type: string;
}

// Schema for the array items.
export interface UntitledMultiSelectEnumSchemaItems {
  // Array of enum values to choose from.
  enum: string[];
  type: string;
}

// Schema for multiple-selection enumeration without display titles for options.
export interface UntitledMultiSelectEnumSchema {
  // Optional default value.
  default?: string[];
  // Optional description for the enum field.
  description?: string;
  // Schema for the array items.
  items: UntitledMultiSelectEnumSchemaItems;
  // Maximum number of items to select.
  maxItems?: number;
  // Minimum number of items to select.
  minItems?: number;
  // Optional title for the enum field.
  title?: string;
  type: string;
}

// Schema for single-selection enumeration without display titles for options.
export interface UntitledSingleSelectEnumSchema {
  // Optional default value.
  default?: string;
  // Optional description for the enum field.
  description?: string;
  // Array of enum values to choose from.
  enum: string[];
  // Optional title for the enum field.
  title?: string;
  type: string;
}


// Restricted schema definitions that only allow primitive types
// without nested objects or arrays.
export type PrimitiveSchemaDefinition = StringSchema | NumberSchema | BooleanSchema | UntitledSingleSelectEnumSchema | TitledSingleSelectEnumSchema | UntitledMultiSelectEnumSchema | TitledMultiSelectEnumSchema | LegacyTitledEnumSchema;

// A restricted subset of JSON Schema.
// Only top-level properties are allowed, without nesting.
export interface ElicitRequestFormParamsRequestedSchema {
  $schema?: string;
  properties: Record<string, PrimitiveSchemaDefinition>;
  required?: string[];
  type: string;
}

// The parameters for a request to elicit non-sensitive information from the user via a form in the client.
export interface ElicitRequestFormParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: ElicitRequestFormParamsMeta;
  // The message to present to the user describing what information is being requested.
  message: string;
  // The elicitation mode.
  mode?: string;
  // A restricted subset of JSON Schema.
  // Only top-level properties are allowed, without nesting.
  requestedSchema: ElicitRequestFormParamsRequestedSchema;
  // If specified, the caller is requesting task-augmented execution for this request.
  // The request will return a CreateTaskResult immediately, and the actual result can be
  // retrieved later via tasks/result.
  // 
  // Task augmentation is subject to capability negotiation - receivers MUST declare support
  // for task augmentation of specific request types in their capabilities.
  task?: TaskMetadata;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface ElicitRequestURLParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// The parameters for a request to elicit information from the user via a URL in the client.
export interface ElicitRequestURLParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: ElicitRequestURLParamsMeta;
  // The ID of the elicitation, which must be unique within the context of the server.
  // The client MUST treat this ID as an opaque value.
  elicitationId: string;
  // The message to present to the user explaining why the interaction is needed.
  message: string;
  // The elicitation mode.
  mode: string;
  // If specified, the caller is requesting task-augmented execution for this request.
  // The request will return a CreateTaskResult immediately, and the actual result can be
  // retrieved later via tasks/result.
  // 
  // Task augmentation is subject to capability negotiation - receivers MUST declare support
  // for task augmentation of specific request types in their capabilities.
  task?: TaskMetadata;
  // The URL that the user should navigate to.
  url: string;
}


// The parameters for a request to elicit additional information from the user via the client.
export type ElicitRequestParams = ElicitRequestURLParams | ElicitRequestFormParams;

// A request from the server to elicit additional information from the user via the client.
export interface ElicitRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params: ElicitRequestParams;
}

export interface ElicitationCompleteNotificationParams {
  // The ID of the elicitation that completed.
  elicitationId: string;
}

// An optional notification from the server to the client, informing it of a completion of a out-of-band elicitation request.
export interface ElicitationCompleteNotification {
  jsonrpc: string;
  method: string;
  params: ElicitationCompleteNotificationParams;
}

export type EmptyResult = unknown;

export type EnumSchema = UntitledSingleSelectEnumSchema | TitledSingleSelectEnumSchema | UntitledMultiSelectEnumSchema | TitledMultiSelectEnumSchema | LegacyTitledEnumSchema;

export interface Error {
  // The error type that occurred.
  code: number;
  // Additional information about the error. The value of this member is defined by the sender (e.g. detailed error information, nested errors etc.).
  data?: unknown;
  // A short description of the error. The message SHOULD be limited to a concise single sentence.
  message: string;
}

// Describes a message returned as part of a prompt.
// 
// This is similar to `SamplingMessage`, but also supports the embedding of
// resources from the MCP server.
export interface PromptMessage {
  content: ContentBlock;
  role: Role;
}

// The server's response to a prompts/get request from the client.
export interface GetPromptResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An optional description for the prompt.
  description?: string;
  messages: PromptMessage[];
}

// Base interface to add `icons` property.
export interface Icons {
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
}

// Present if the server supports argument autocompletion suggestions.
export interface ServerCapabilitiesCompletions {
}

export interface ServerCapabilitiesExperimentalValue {
}

// Present if the server supports sending log messages to the client.
export interface ServerCapabilitiesLogging {
}

// Present if the server offers any prompt templates.
export interface ServerCapabilitiesPrompts {
  // Whether this server supports notifications for changes to the prompt list.
  listChanged?: boolean;
}

// Present if the server offers any resources to read.
export interface ServerCapabilitiesResources {
  // Whether this server supports notifications for changes to the resource list.
  listChanged?: boolean;
  // Whether this server supports subscribing to resource updates.
  subscribe?: boolean;
}

// Whether this server supports tasks/cancel.
export interface ServerCapabilitiesTasksCancel {
}

// Whether this server supports tasks/list.
export interface ServerCapabilitiesTasksList {
}

// Whether the server supports task-augmented tools/call requests.
export interface ServerCapabilitiesTasksRequestsToolsCall {
}

// Task support for tool-related requests.
export interface ServerCapabilitiesTasksRequestsTools {
  // Whether the server supports task-augmented tools/call requests.
  call?: ServerCapabilitiesTasksRequestsToolsCall;
}

// Specifies which request types can be augmented with tasks.
export interface ServerCapabilitiesTasksRequests {
  // Task support for tool-related requests.
  tools?: ServerCapabilitiesTasksRequestsTools;
}

// Present if the server supports task-augmented requests.
export interface ServerCapabilitiesTasks {
  // Whether this server supports tasks/cancel.
  cancel?: ServerCapabilitiesTasksCancel;
  // Whether this server supports tasks/list.
  list?: ServerCapabilitiesTasksList;
  // Specifies which request types can be augmented with tasks.
  requests?: ServerCapabilitiesTasksRequests;
}

// Present if the server offers any tools to call.
export interface ServerCapabilitiesTools {
  // Whether this server supports notifications for changes to the tool list.
  listChanged?: boolean;
}

// Capabilities that a server may support. Known capabilities are defined here, in this schema, but this is not a closed set: any server can define its own, additional capabilities.
export interface ServerCapabilities {
  // Present if the server supports argument autocompletion suggestions.
  completions?: ServerCapabilitiesCompletions;
  // Experimental, non-standard capabilities that the server supports.
  experimental?: Record<string, ServerCapabilitiesExperimentalValue>;
  // Present if the server supports sending log messages to the client.
  logging?: ServerCapabilitiesLogging;
  // Present if the server offers any prompt templates.
  prompts?: ServerCapabilitiesPrompts;
  // Present if the server offers any resources to read.
  resources?: ServerCapabilitiesResources;
  // Present if the server supports task-augmented requests.
  tasks?: ServerCapabilitiesTasks;
  // Present if the server offers any tools to call.
  tools?: ServerCapabilitiesTools;
}

// After receiving an initialize request from the client, the server sends this response.
export interface InitializeResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  capabilities: ServerCapabilities;
  // Instructions describing how to use the server and its features.
  // 
  // This can be used by clients to improve the LLM's understanding of available tools, resources, etc. It can be thought of like a "hint" to the model. For example, this information MAY be added to the system prompt.
  instructions?: string;
  // The version of the Model Context Protocol that the server wants to use. This may not match the version that the client requested. If the client cannot support this version, it MUST disconnect.
  protocolVersion: string;
  serverInfo: Implementation;
}

// A response to a request that indicates an error occurred.
export interface JSONRPCErrorResponse {
  error: Error;
  id?: RequestId;
  jsonrpc: string;
}

// A notification which does not expect a response.
export interface JSONRPCNotification {
  jsonrpc: string;
  method: string;
  params?: Record<string, unknown>;
}

// A request that expects a response.
export interface JSONRPCRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: Record<string, unknown>;
}

// A successful (non-error) response to a request.
export interface JSONRPCResultResponse {
  id: RequestId;
  jsonrpc: string;
  result: Result;
}


// Refers to any valid JSON-RPC object that can be decoded off the wire, or encoded to be sent.
export type JSONRPCMessage = JSONRPCRequest | JSONRPCNotification | JSONRPCResultResponse | JSONRPCErrorResponse;


// A response to a request, containing either the result or error.
export type JSONRPCResponse = JSONRPCResultResponse | JSONRPCErrorResponse;

// Describes an argument that a prompt can accept.
export interface PromptArgument {
  // A human-readable description of the argument.
  description?: string;
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // Whether this argument must be provided.
  required?: boolean;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
}

// A prompt or prompt template that the server offers.
export interface Prompt {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // A list of arguments to use for templating the prompt.
  arguments?: PromptArgument[];
  // An optional description of what this prompt provides
  description?: string;
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
}

// The server's response to a prompts/list request from the client.
export interface ListPromptsResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An opaque token representing the pagination position after the last returned result.
  // If present, there may be more results available.
  nextCursor?: string;
  prompts: Prompt[];
}

// A template description for resources available on the server.
export interface ResourceTemplate {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  // A description of what this template is for.
  // 
  // This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
  description?: string;
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
  // The MIME type for all resources that match this template. This should only be included if all resources matching this template have the same type.
  mimeType?: string;
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
  // A URI template (according to RFC 6570) that can be used to construct resource URIs.
  uriTemplate: string;
}

// The server's response to a resources/templates/list request from the client.
export interface ListResourceTemplatesResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An opaque token representing the pagination position after the last returned result.
  // If present, there may be more results available.
  nextCursor?: string;
  resourceTemplates: ResourceTemplate[];
}

// A known resource that the server is capable of reading.
export interface Resource {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // Optional annotations for the client.
  annotations?: Annotations;
  // A description of what this resource represents.
  // 
  // This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
  description?: string;
  // Optional set of sized icons that the client can display in a user interface.
  // 
  // Clients that support rendering icons MUST support at least the following MIME types:
  // - `image/png` - PNG images (safe, universal compatibility)
  // - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
  // 
  // Clients that support rendering icons SHOULD also support:
  // - `image/svg+xml` - SVG images (scalable but requires security precautions)
  // - `image/webp` - WebP images (modern, efficient format)
  icons?: Icon[];
  // The MIME type of this resource, if known.
  mimeType?: string;
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
  // 
  // This can be used by Hosts to display file sizes and estimate context window usage.
  size?: number;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
  // The URI of this resource.
  uri: string;
}

// The server's response to a resources/list request from the client.
export interface ListResourcesResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An opaque token representing the pagination position after the last returned result.
  // If present, there may be more results available.
  nextCursor?: string;
  resources: Resource[];
}

// Sent from the server to request a list of root URIs from the client. Roots allow
// servers to ask for specific directories or files to operate on. A common example
// for roots is providing a set of repositories or directories a server should operate
// on.
// 
// This request is typically used when the server needs to understand the file system
// structure or access specific locations that the client has permission to read from.
export interface ListRootsRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: RequestParams;
}

// The server's response to a tools/list request from the client.
export interface ListToolsResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An opaque token representing the pagination position after the last returned result.
  // If present, there may be more results available.
  nextCursor?: string;
  tools: Tool[];
}

// Parameters for a `notifications/message` notification.
export interface LoggingMessageNotificationParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // The data to be logged, such as a string message or an object. Any JSON serializable type is allowed here.
  data: unknown;
  // The severity of this log message.
  level: LoggingLevel;
  // An optional name of the logger issuing this message.
  logger?: string;
}

// JSONRPCNotification of a log message passed from server to client. If no logging/setLevel request has been sent from the client, the server MAY decide which messages to send automatically.
export interface LoggingMessageNotification {
  jsonrpc: string;
  method: string;
  params: LoggingMessageNotificationParams;
}

export type MultiSelectEnumSchema = UntitledMultiSelectEnumSchema | TitledMultiSelectEnumSchema;

export interface Notification {
  method: string;
  params?: Record<string, unknown>;
}

export interface PaginatedRequest {
  id: RequestId;
  jsonrpc: string;
  method: string;
  params?: PaginatedRequestParams;
}

export interface PaginatedResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // An opaque token representing the pagination position after the last returned result.
  // If present, there may be more results available.
  nextCursor?: string;
}

// An optional notification from the server to the client, informing it that the list of prompts it offers has changed. This may be issued by servers without any previous subscription from the client.
export interface PromptListChangedNotification {
  jsonrpc: string;
  method: string;
  params?: NotificationParams;
}

// Identifies a prompt.
export interface PromptReference {
  // Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present).
  name: string;
  // Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
  // even by those unfamiliar with domain-specific terminology.
  // 
  // If not provided, the name should be used for display (except for Tool,
  // where `annotations.title` should be given precedence over using `name`,
  // if present).
  title?: string;
  type: string;
}

// The server's response to a resources/read request from the client.
export interface ReadResourceResult {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  contents: unknown[];
}

// Metadata for associating messages with a task.
// Include this in the `_meta` field under the key `io.modelcontextprotocol/related-task`.
export interface RelatedTaskMetadata {
  // The task identifier this message is associated with.
  taskId: string;
}

export interface Request {
  method: string;
  params?: Record<string, unknown>;
}

// The contents of a specific resource or sub-resource.
export interface ResourceContents {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // The MIME type of this resource, if known.
  mimeType?: string;
  // The URI of this resource.
  uri: string;
}

// An optional notification from the server to the client, informing it that the list of resources it can read from has changed. This may be issued by servers without any previous subscription from the client.
export interface ResourceListChangedNotification {
  jsonrpc: string;
  method: string;
  params?: NotificationParams;
}

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface ResourceRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Common parameters when working with resources.
export interface ResourceRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: ResourceRequestParamsMeta;
  // The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it.
  uri: string;
}

// A reference to a resource or resource template definition.
export interface ResourceTemplateReference {
  type: string;
  // The URI or URI template of the resource.
  uri: string;
}

// Parameters for a `notifications/resources/updated` notification.
export interface ResourceUpdatedNotificationParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // The URI of the resource that has been updated. This might be a sub-resource of the one that the client actually subscribed to.
  uri: string;
}

// A notification from the server to the client, informing it that a resource has changed and may need to be read again. This should only be sent if the client previously sent a resources/subscribe request.
export interface ResourceUpdatedNotification {
  jsonrpc: string;
  method: string;
  params: ResourceUpdatedNotificationParams;
}

// The result of a tool use, provided by the user back to the assistant.
export interface ToolResultContent {
  // Optional metadata about the tool result. Clients SHOULD preserve this field when
  // including tool results in subsequent sampling requests to enable caching optimizations.
  // 
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // The unstructured result content of the tool use.
  // 
  // This has the same format as CallToolResult.content and can include text, images,
  // audio, resource links, and embedded resources.
  content: ContentBlock[];
  // Whether the tool use resulted in an error.
  // 
  // If true, the content typically describes the error that occurred.
  // Default: false
  isError?: boolean;
  // An optional structured result object.
  // 
  // If the tool defined an outputSchema, this SHOULD conform to that schema.
  structuredContent?: Record<string, unknown>;
  // The ID of the tool use this result corresponds to.
  // 
  // This MUST match the ID from a previous ToolUseContent.
  toolUseId: string;
  type: string;
}

// A request from the assistant to call a tool.
export interface ToolUseContent {
  // Optional metadata about the tool use. Clients SHOULD preserve this field when
  // including tool uses in subsequent sampling requests to enable caching optimizations.
  // 
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // A unique identifier for this tool use.
  // 
  // This ID is used to match tool results to their corresponding tool uses.
  id: string;
  // The arguments to pass to the tool, conforming to the tool's input schema.
  input: Record<string, unknown>;
  // The name of the tool to call.
  name: string;
  type: string;
}

export type SamplingMessageContentBlock = TextContent | ImageContent | AudioContent | ToolUseContent | ToolResultContent;

// An optional notification from the server to the client, informing it that the list of tools it offers has changed. This may be issued by servers without any previous subscription from the client.
export interface ToolListChangedNotification {
  jsonrpc: string;
  method: string;
  params?: NotificationParams;
}

export type ServerNotification = CancelledNotification | ProgressNotification | ResourceListChangedNotification | ResourceUpdatedNotification | PromptListChangedNotification | ToolListChangedNotification | TaskStatusNotification | LoggingMessageNotification | ElicitationCompleteNotification;

export type ServerRequest = PingRequest | GetTaskRequest | GetTaskPayloadRequest | CancelTaskRequest | ListTasksRequest | CreateMessageRequest | ListRootsRequest | ElicitRequest;

export type ServerResult = Result | InitializeResult | ListResourcesResult | ListResourceTemplatesResult | ReadResourceResult | ListPromptsResult | GetPromptResult | ListToolsResult | CallToolResult | GetTaskResult | GetTaskPayloadResult | CancelTaskResult | ListTasksResult | CompleteResult;

export type SingleSelectEnumSchema = UntitledSingleSelectEnumSchema | TitledSingleSelectEnumSchema;

// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export interface TaskAugmentedRequestParamsMeta {
  // If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications.
  progressToken?: ProgressToken;
}

// Common params for any task-augmented request.
export interface TaskAugmentedRequestParams {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: TaskAugmentedRequestParamsMeta;
  // If specified, the caller is requesting task-augmented execution for this request.
  // The request will return a CreateTaskResult immediately, and the actual result can be
  // retrieved later via tasks/result.
  // 
  // Task augmentation is subject to capability negotiation - receivers MUST declare support
  // for task augmentation of specific request types in their capabilities.
  task?: TaskMetadata;
}

export interface TextResourceContents {
  // See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
  _meta?: Record<string, unknown>;
  // The MIME type of this resource, if known.
  mimeType?: string;
  // The text of the item. This must only be set if the item can actually be represented as text (not binary data).
  text: string;
  // The URI of this resource.
  uri: string;
}

// An error response that indicates that the server requires the client to provide additional information via an elicitation request.
export interface URLElicitationRequiredError {
  error: unknown;
  id?: RequestId;
  jsonrpc: string;
}
