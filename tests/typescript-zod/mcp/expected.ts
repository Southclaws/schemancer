import { z } from "zod";


// The sender or recipient of messages and data in a conversation.
export const RoleSchema = z.enum(["assistant", "user"]);
export type Role = z.infer<typeof RoleSchema>;


// Optional annotations for the client. The client can use annotations to inform how objects are used or displayed
export const AnnotationsSchema = z.object({
  audience: z.array(RoleSchema).optional(),
  lastModified: z.string().optional(),
  priority: z.number().min(0).max(1).optional(),
});
export type Annotations = z.infer<typeof AnnotationsSchema>;


// Audio provided to or from an LLM.
export const AudioContentSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  data: z.string(),
  mimeType: z.string(),
  type: z.string(),
});
export type AudioContent = z.infer<typeof AudioContentSchema>;


// Base interface for metadata with name (identifier) and title (display name) properties.
export const BaseMetadataSchema = z.object({
  name: z.string(),
  title: z.string().optional(),
});
export type BaseMetadata = z.infer<typeof BaseMetadataSchema>;

export const BlobResourceContentsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  blob: z.string(),
  mimeType: z.string().optional(),
  uri: z.string().url(),
});
export type BlobResourceContents = z.infer<typeof BlobResourceContentsSchema>;

export const BooleanSchemaSchema = z.object({
  default: z.boolean().optional(),
  description: z.string().optional(),
  title: z.string().optional(),
  type: z.string(),
});
export type BooleanSchema = z.infer<typeof BooleanSchemaSchema>;


// A progress token, used to associate progress notifications with the original request.
export const ProgressTokenSchema = z.unknown();
export type ProgressToken = z.infer<typeof ProgressTokenSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const CallToolRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type CallToolRequestParamsMeta = z.infer<typeof CallToolRequestParamsMetaSchema>;


// Metadata for augmenting a request with task execution.
// Include this in the `task` field of the request parameters.
export const TaskMetadataSchema = z.object({
  ttl: z.number().int().optional(),
});
export type TaskMetadata = z.infer<typeof TaskMetadataSchema>;


// Parameters for a `tools/call` request.
export const CallToolRequestParamsSchema = z.object({
  _meta: CallToolRequestParamsMetaSchema.optional(),
  arguments: z.record(z.string(), z.unknown()).optional(),
  name: z.string(),
  task: TaskMetadataSchema.optional(),
});
export type CallToolRequestParams = z.infer<typeof CallToolRequestParamsSchema>;


// A uniquely identifying ID for a request in JSON-RPC.
export const RequestIdSchema = z.unknown();
export type RequestId = z.infer<typeof RequestIdSchema>;


// Used by the client to invoke a tool provided by the server.
export const CallToolRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: CallToolRequestParamsSchema,
});
export type CallToolRequest = z.infer<typeof CallToolRequestSchema>;


// The contents of a resource, embedded into a prompt or tool call result.
// 
// It is up to the client how best to render embedded resources for the benefit
// of the LLM and/or the user.
export const EmbeddedResourceSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  resource: z.unknown(),
  type: z.string(),
});
export type EmbeddedResource = z.infer<typeof EmbeddedResourceSchema>;


// An image provided to or from an LLM.
export const ImageContentSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  data: z.string(),
  mimeType: z.string(),
  type: z.string(),
});
export type ImageContent = z.infer<typeof ImageContentSchema>;


// An optionally-sized icon that can be displayed in a user interface.
export const IconSchema = z.object({
  mimeType: z.string().optional(),
  sizes: z.array(z.string()).optional(),
  src: z.string().url(),
  theme: z.string().optional(),
});
export type Icon = z.infer<typeof IconSchema>;


// A resource that the server is capable of reading, included in a prompt or tool call result.
// 
// Note: resource links returned by tools are not guaranteed to appear in the results of `resources/list` requests.
export const ResourceLinkSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  description: z.string().optional(),
  icons: z.array(IconSchema).optional(),
  mimeType: z.string().optional(),
  name: z.string(),
  size: z.number().int().optional(),
  title: z.string().optional(),
  type: z.string(),
  uri: z.string().url(),
});
export type ResourceLink = z.infer<typeof ResourceLinkSchema>;


// Text provided to or from an LLM.
export const TextContentSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  text: z.string(),
  type: z.string(),
});
export type TextContent = z.infer<typeof TextContentSchema>;

export const ContentBlockSchema = z.union([
  TextContentSchema,
  ImageContentSchema,
  AudioContentSchema,
  ResourceLinkSchema,
  EmbeddedResourceSchema,
]);
export type ContentBlock = z.infer<typeof ContentBlockSchema>;


// The server's response to a tool call.
export const CallToolResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  content: z.array(ContentBlockSchema),
  isError: z.boolean().optional(),
  structuredContent: z.record(z.string(), z.unknown()).optional(),
});
export type CallToolResult = z.infer<typeof CallToolResultSchema>;

export const CancelTaskRequestParamsSchema = z.object({
  taskId: z.string(),
});
export type CancelTaskRequestParams = z.infer<typeof CancelTaskRequestParamsSchema>;


// A request to cancel a task.
export const CancelTaskRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: CancelTaskRequestParamsSchema,
});
export type CancelTaskRequest = z.infer<typeof CancelTaskRequestSchema>;

export const ResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
});
export type Result = z.infer<typeof ResultSchema>;


// The status of a task.
export const TaskStatusSchema = z.enum(["cancelled", "completed", "failed", "input_required", "working"]);
export type TaskStatus = z.infer<typeof TaskStatusSchema>;


// The response to a tasks/cancel request.
export const CancelTaskResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  createdAt: z.string(),
  lastUpdatedAt: z.string(),
  pollInterval: z.number().int().optional(),
  status: TaskStatusSchema,
  statusMessage: z.string().optional(),
  taskId: z.string(),
  ttl: z.number().int(),
});
export type CancelTaskResult = z.infer<typeof CancelTaskResultSchema>;


// Parameters for a `notifications/cancelled` notification.
export const CancelledNotificationParamsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  reason: z.string().optional(),
  requestId: RequestIdSchema.optional(),
});
export type CancelledNotificationParams = z.infer<typeof CancelledNotificationParamsSchema>;


// This notification can be sent by either side to indicate that it is cancelling a previously-issued request.
// 
// The request SHOULD still be in-flight, but due to communication latency, it is always possible that this notification MAY arrive after the request has already finished.
// 
// This notification indicates that the result will be unused, so any associated processing SHOULD cease.
// 
// A client MUST NOT attempt to cancel its `initialize` request.
// 
// For task cancellation, use the `tasks/cancel` request instead of this notification.
export const CancelledNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: CancelledNotificationParamsSchema,
});
export type CancelledNotification = z.infer<typeof CancelledNotificationSchema>;

export const ClientCapabilitiesElicitationFormSchema = z.object({
});
export type ClientCapabilitiesElicitationForm = z.infer<typeof ClientCapabilitiesElicitationFormSchema>;

export const ClientCapabilitiesElicitationURLSchema = z.object({
});
export type ClientCapabilitiesElicitationURL = z.infer<typeof ClientCapabilitiesElicitationURLSchema>;


// Present if the client supports elicitation from the server.
export const ClientCapabilitiesElicitationSchema = z.object({
  form: ClientCapabilitiesElicitationFormSchema.optional(),
  url: ClientCapabilitiesElicitationURLSchema.optional(),
});
export type ClientCapabilitiesElicitation = z.infer<typeof ClientCapabilitiesElicitationSchema>;

export const ClientCapabilitiesExperimentalValueSchema = z.object({
});
export type ClientCapabilitiesExperimentalValue = z.infer<typeof ClientCapabilitiesExperimentalValueSchema>;


// Present if the client supports listing roots.
export const ClientCapabilitiesRootsSchema = z.object({
  listChanged: z.boolean().optional(),
});
export type ClientCapabilitiesRoots = z.infer<typeof ClientCapabilitiesRootsSchema>;


// Whether the client supports context inclusion via includeContext parameter.
// If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
export const ClientCapabilitiesSamplingContextSchema = z.object({
});
export type ClientCapabilitiesSamplingContext = z.infer<typeof ClientCapabilitiesSamplingContextSchema>;


// Whether the client supports tool use via tools and toolChoice parameters.
export const ClientCapabilitiesSamplingToolsSchema = z.object({
});
export type ClientCapabilitiesSamplingTools = z.infer<typeof ClientCapabilitiesSamplingToolsSchema>;


// Present if the client supports sampling from an LLM.
export const ClientCapabilitiesSamplingSchema = z.object({
  context: ClientCapabilitiesSamplingContextSchema.optional(),
  tools: ClientCapabilitiesSamplingToolsSchema.optional(),
});
export type ClientCapabilitiesSampling = z.infer<typeof ClientCapabilitiesSamplingSchema>;


// Whether this client supports tasks/cancel.
export const ClientCapabilitiesTasksCancelSchema = z.object({
});
export type ClientCapabilitiesTasksCancel = z.infer<typeof ClientCapabilitiesTasksCancelSchema>;


// Whether this client supports tasks/list.
export const ClientCapabilitiesTasksListSchema = z.object({
});
export type ClientCapabilitiesTasksList = z.infer<typeof ClientCapabilitiesTasksListSchema>;


// Whether the client supports task-augmented elicitation/create requests.
export const ClientCapabilitiesTasksRequestsElicitationCreateSchema = z.object({
});
export type ClientCapabilitiesTasksRequestsElicitationCreate = z.infer<typeof ClientCapabilitiesTasksRequestsElicitationCreateSchema>;


// Task support for elicitation-related requests.
export const ClientCapabilitiesTasksRequestsElicitationSchema = z.object({
  create: ClientCapabilitiesTasksRequestsElicitationCreateSchema.optional(),
});
export type ClientCapabilitiesTasksRequestsElicitation = z.infer<typeof ClientCapabilitiesTasksRequestsElicitationSchema>;


// Whether the client supports task-augmented sampling/createMessage requests.
export const ClientCapabilitiesTasksRequestsSamplingCreateMessageSchema = z.object({
});
export type ClientCapabilitiesTasksRequestsSamplingCreateMessage = z.infer<typeof ClientCapabilitiesTasksRequestsSamplingCreateMessageSchema>;


// Task support for sampling-related requests.
export const ClientCapabilitiesTasksRequestsSamplingSchema = z.object({
  createMessage: ClientCapabilitiesTasksRequestsSamplingCreateMessageSchema.optional(),
});
export type ClientCapabilitiesTasksRequestsSampling = z.infer<typeof ClientCapabilitiesTasksRequestsSamplingSchema>;


// Specifies which request types can be augmented with tasks.
export const ClientCapabilitiesTasksRequestsSchema = z.object({
  elicitation: ClientCapabilitiesTasksRequestsElicitationSchema.optional(),
  sampling: ClientCapabilitiesTasksRequestsSamplingSchema.optional(),
});
export type ClientCapabilitiesTasksRequests = z.infer<typeof ClientCapabilitiesTasksRequestsSchema>;


// Present if the client supports task-augmented requests.
export const ClientCapabilitiesTasksSchema = z.object({
  cancel: ClientCapabilitiesTasksCancelSchema.optional(),
  list: ClientCapabilitiesTasksListSchema.optional(),
  requests: ClientCapabilitiesTasksRequestsSchema.optional(),
});
export type ClientCapabilitiesTasks = z.infer<typeof ClientCapabilitiesTasksSchema>;


// Capabilities a client may support. Known capabilities are defined here, in this schema, but this is not a closed set: any client can define its own, additional capabilities.
export const ClientCapabilitiesSchema = z.object({
  elicitation: ClientCapabilitiesElicitationSchema.optional(),
  experimental: z.record(z.string(), ClientCapabilitiesExperimentalValueSchema).optional(),
  roots: ClientCapabilitiesRootsSchema.optional(),
  sampling: ClientCapabilitiesSamplingSchema.optional(),
  tasks: ClientCapabilitiesTasksSchema.optional(),
});
export type ClientCapabilities = z.infer<typeof ClientCapabilitiesSchema>;

export const NotificationParamsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
});
export type NotificationParams = z.infer<typeof NotificationParamsSchema>;


// This notification is sent from the client to the server after initialization has finished.
export const InitializedNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: NotificationParamsSchema.optional(),
});
export type InitializedNotification = z.infer<typeof InitializedNotificationSchema>;


// Parameters for a `notifications/progress` notification.
export const ProgressNotificationParamsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  message: z.string().optional(),
  progress: z.number(),
  progressToken: ProgressTokenSchema,
  total: z.number().optional(),
});
export type ProgressNotificationParams = z.infer<typeof ProgressNotificationParamsSchema>;


// An out-of-band notification used to inform the receiver of a progress update for a long-running request.
export const ProgressNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: ProgressNotificationParamsSchema,
});
export type ProgressNotification = z.infer<typeof ProgressNotificationSchema>;


// A notification from the client to the server, informing it that the list of roots has changed.
// This notification should be sent whenever the client adds, removes, or modifies any root.
// The server should then request an updated list of roots using the ListRootsRequest.
export const RootsListChangedNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: NotificationParamsSchema.optional(),
});
export type RootsListChangedNotification = z.infer<typeof RootsListChangedNotificationSchema>;


// Parameters for a `notifications/tasks/status` notification.
export const TaskStatusNotificationParamsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  createdAt: z.string(),
  lastUpdatedAt: z.string(),
  pollInterval: z.number().int().optional(),
  status: TaskStatusSchema,
  statusMessage: z.string().optional(),
  taskId: z.string(),
  ttl: z.number().int(),
});
export type TaskStatusNotificationParams = z.infer<typeof TaskStatusNotificationParamsSchema>;


// An optional notification from the receiver to the requestor, informing them that a task's status has changed. Receivers are not required to send these notifications.
export const TaskStatusNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: TaskStatusNotificationParamsSchema,
});
export type TaskStatusNotification = z.infer<typeof TaskStatusNotificationSchema>;

export const ClientNotificationSchema = z.union([
  CancelledNotificationSchema,
  InitializedNotificationSchema,
  ProgressNotificationSchema,
  TaskStatusNotificationSchema,
  RootsListChangedNotificationSchema,
]);
export type ClientNotification = z.infer<typeof ClientNotificationSchema>;


// The argument's information
export const CompleteRequestParamsArgumentSchema = z.object({
  name: z.string(),
  value: z.string(),
});
export type CompleteRequestParamsArgument = z.infer<typeof CompleteRequestParamsArgumentSchema>;


// Additional, optional context for completions
export const CompleteRequestParamsContextSchema = z.object({
  arguments: z.record(z.string(), z.string()).optional(),
});
export type CompleteRequestParamsContext = z.infer<typeof CompleteRequestParamsContextSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const CompleteRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type CompleteRequestParamsMeta = z.infer<typeof CompleteRequestParamsMetaSchema>;


// Parameters for a `completion/complete` request.
export const CompleteRequestParamsSchema = z.object({
  _meta: CompleteRequestParamsMetaSchema.optional(),
  argument: CompleteRequestParamsArgumentSchema,
  context: CompleteRequestParamsContextSchema.optional(),
  ref: z.unknown(),
});
export type CompleteRequestParams = z.infer<typeof CompleteRequestParamsSchema>;


// A request from the client to the server, to ask for completion options.
export const CompleteRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: CompleteRequestParamsSchema,
});
export type CompleteRequest = z.infer<typeof CompleteRequestSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const GetPromptRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type GetPromptRequestParamsMeta = z.infer<typeof GetPromptRequestParamsMetaSchema>;


// Parameters for a `prompts/get` request.
export const GetPromptRequestParamsSchema = z.object({
  _meta: GetPromptRequestParamsMetaSchema.optional(),
  arguments: z.record(z.string(), z.string()).optional(),
  name: z.string(),
});
export type GetPromptRequestParams = z.infer<typeof GetPromptRequestParamsSchema>;


// Used by the client to get a prompt provided by the server.
export const GetPromptRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: GetPromptRequestParamsSchema,
});
export type GetPromptRequest = z.infer<typeof GetPromptRequestSchema>;

export const GetTaskPayloadRequestParamsSchema = z.object({
  taskId: z.string(),
});
export type GetTaskPayloadRequestParams = z.infer<typeof GetTaskPayloadRequestParamsSchema>;


// A request to retrieve the result of a completed task.
export const GetTaskPayloadRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: GetTaskPayloadRequestParamsSchema,
});
export type GetTaskPayloadRequest = z.infer<typeof GetTaskPayloadRequestSchema>;

export const GetTaskRequestParamsSchema = z.object({
  taskId: z.string(),
});
export type GetTaskRequestParams = z.infer<typeof GetTaskRequestParamsSchema>;


// A request to retrieve the state of a task.
export const GetTaskRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: GetTaskRequestParamsSchema,
});
export type GetTaskRequest = z.infer<typeof GetTaskRequestSchema>;


// Describes the MCP implementation.
export const ImplementationSchema = z.object({
  description: z.string().optional(),
  icons: z.array(IconSchema).optional(),
  name: z.string(),
  title: z.string().optional(),
  version: z.string(),
  websiteUrl: z.string().url().optional(),
});
export type Implementation = z.infer<typeof ImplementationSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const InitializeRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type InitializeRequestParamsMeta = z.infer<typeof InitializeRequestParamsMetaSchema>;


// Parameters for an `initialize` request.
export const InitializeRequestParamsSchema = z.object({
  _meta: InitializeRequestParamsMetaSchema.optional(),
  capabilities: ClientCapabilitiesSchema,
  clientInfo: ImplementationSchema,
  protocolVersion: z.string(),
});
export type InitializeRequestParams = z.infer<typeof InitializeRequestParamsSchema>;


// This request is sent from the client to the server when it first connects, asking it to begin initialization.
export const InitializeRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: InitializeRequestParamsSchema,
});
export type InitializeRequest = z.infer<typeof InitializeRequestSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const PaginatedRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type PaginatedRequestParamsMeta = z.infer<typeof PaginatedRequestParamsMetaSchema>;


// Common parameters for paginated requests.
export const PaginatedRequestParamsSchema = z.object({
  _meta: PaginatedRequestParamsMetaSchema.optional(),
  cursor: z.string().optional(),
});
export type PaginatedRequestParams = z.infer<typeof PaginatedRequestParamsSchema>;


// Sent from the client to request a list of prompts and prompt templates the server has.
export const ListPromptsRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: PaginatedRequestParamsSchema.optional(),
});
export type ListPromptsRequest = z.infer<typeof ListPromptsRequestSchema>;


// Sent from the client to request a list of resource templates the server has.
export const ListResourceTemplatesRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: PaginatedRequestParamsSchema.optional(),
});
export type ListResourceTemplatesRequest = z.infer<typeof ListResourceTemplatesRequestSchema>;


// Sent from the client to request a list of resources the server has.
export const ListResourcesRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: PaginatedRequestParamsSchema.optional(),
});
export type ListResourcesRequest = z.infer<typeof ListResourcesRequestSchema>;


// A request to retrieve a list of tasks.
export const ListTasksRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: PaginatedRequestParamsSchema.optional(),
});
export type ListTasksRequest = z.infer<typeof ListTasksRequestSchema>;


// Sent from the client to request a list of tools the server has.
export const ListToolsRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: PaginatedRequestParamsSchema.optional(),
});
export type ListToolsRequest = z.infer<typeof ListToolsRequestSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const RequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type RequestParamsMeta = z.infer<typeof RequestParamsMetaSchema>;


// Common params for any request.
export const RequestParamsSchema = z.object({
  _meta: RequestParamsMetaSchema.optional(),
});
export type RequestParams = z.infer<typeof RequestParamsSchema>;


// A ping, issued by either the server or the client, to check that the other party is still alive. The receiver must promptly respond, or else may be disconnected.
export const PingRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: RequestParamsSchema.optional(),
});
export type PingRequest = z.infer<typeof PingRequestSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const ReadResourceRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type ReadResourceRequestParamsMeta = z.infer<typeof ReadResourceRequestParamsMetaSchema>;


// Parameters for a `resources/read` request.
export const ReadResourceRequestParamsSchema = z.object({
  _meta: ReadResourceRequestParamsMetaSchema.optional(),
  uri: z.string().url(),
});
export type ReadResourceRequestParams = z.infer<typeof ReadResourceRequestParamsSchema>;


// Sent from the client to the server, to read a specific resource URI.
export const ReadResourceRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: ReadResourceRequestParamsSchema,
});
export type ReadResourceRequest = z.infer<typeof ReadResourceRequestSchema>;


// The severity of a log message.
// 
// These map to syslog message severities, as specified in RFC-5424:
// https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1
export const LoggingLevelSchema = z.enum(["alert", "critical", "debug", "emergency", "error", "info", "notice", "warning"]);
export type LoggingLevel = z.infer<typeof LoggingLevelSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const SetLevelRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type SetLevelRequestParamsMeta = z.infer<typeof SetLevelRequestParamsMetaSchema>;


// Parameters for a `logging/setLevel` request.
export const SetLevelRequestParamsSchema = z.object({
  _meta: SetLevelRequestParamsMetaSchema.optional(),
  level: LoggingLevelSchema,
});
export type SetLevelRequestParams = z.infer<typeof SetLevelRequestParamsSchema>;


// A request from the client to the server, to enable or adjust logging.
export const SetLevelRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: SetLevelRequestParamsSchema,
});
export type SetLevelRequest = z.infer<typeof SetLevelRequestSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const SubscribeRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type SubscribeRequestParamsMeta = z.infer<typeof SubscribeRequestParamsMetaSchema>;


// Parameters for a `resources/subscribe` request.
export const SubscribeRequestParamsSchema = z.object({
  _meta: SubscribeRequestParamsMetaSchema.optional(),
  uri: z.string().url(),
});
export type SubscribeRequestParams = z.infer<typeof SubscribeRequestParamsSchema>;


// Sent from the client to request resources/updated notifications from the server whenever a particular resource changes.
export const SubscribeRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: SubscribeRequestParamsSchema,
});
export type SubscribeRequest = z.infer<typeof SubscribeRequestSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const UnsubscribeRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type UnsubscribeRequestParamsMeta = z.infer<typeof UnsubscribeRequestParamsMetaSchema>;


// Parameters for a `resources/unsubscribe` request.
export const UnsubscribeRequestParamsSchema = z.object({
  _meta: UnsubscribeRequestParamsMetaSchema.optional(),
  uri: z.string().url(),
});
export type UnsubscribeRequestParams = z.infer<typeof UnsubscribeRequestParamsSchema>;


// Sent from the client to request cancellation of resources/updated notifications from the server. This should follow a previous resources/subscribe request.
export const UnsubscribeRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: UnsubscribeRequestParamsSchema,
});
export type UnsubscribeRequest = z.infer<typeof UnsubscribeRequestSchema>;

export const ClientRequestSchema = z.union([
  InitializeRequestSchema,
  PingRequestSchema,
  ListResourcesRequestSchema,
  ListResourceTemplatesRequestSchema,
  ReadResourceRequestSchema,
  SubscribeRequestSchema,
  UnsubscribeRequestSchema,
  ListPromptsRequestSchema,
  GetPromptRequestSchema,
  ListToolsRequestSchema,
  CallToolRequestSchema,
  GetTaskRequestSchema,
  GetTaskPayloadRequestSchema,
  CancelTaskRequestSchema,
  ListTasksRequestSchema,
  SetLevelRequestSchema,
  CompleteRequestSchema,
]);
export type ClientRequest = z.infer<typeof ClientRequestSchema>;


// The client's response to a sampling/createMessage request from the server.
// The client should inform the user before returning the sampled message, to allow them
// to inspect the response (human in the loop) and decide whether to allow the server to see it.
export const CreateMessageResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  content: z.unknown(),
  model: z.string(),
  role: RoleSchema,
  stopReason: z.string().optional(),
});
export type CreateMessageResult = z.infer<typeof CreateMessageResultSchema>;


// The client's response to an elicitation request.
export const ElicitResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  action: z.string(),
  content: z.record(z.string(), z.unknown()).optional(),
});
export type ElicitResult = z.infer<typeof ElicitResultSchema>;


// The response to a tasks/result request.
// The structure matches the result type of the original request.
// For example, a tools/call task would return the CallToolResult structure.
export const GetTaskPayloadResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
});
export type GetTaskPayloadResult = z.infer<typeof GetTaskPayloadResultSchema>;


// The response to a tasks/get request.
export const GetTaskResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  createdAt: z.string(),
  lastUpdatedAt: z.string(),
  pollInterval: z.number().int().optional(),
  status: TaskStatusSchema,
  statusMessage: z.string().optional(),
  taskId: z.string(),
  ttl: z.number().int(),
});
export type GetTaskResult = z.infer<typeof GetTaskResultSchema>;


// Represents a root directory or file that the server can operate on.
export const RootSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  name: z.string().optional(),
  uri: z.string().url(),
});
export type Root = z.infer<typeof RootSchema>;


// The client's response to a roots/list request from the server.
// This result contains an array of Root objects, each representing a root directory
// or file that the server can operate on.
export const ListRootsResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  roots: z.array(RootSchema),
});
export type ListRootsResult = z.infer<typeof ListRootsResultSchema>;


// Data associated with a task.
export const TaskSchema = z.object({
  createdAt: z.string(),
  lastUpdatedAt: z.string(),
  pollInterval: z.number().int().optional(),
  status: TaskStatusSchema,
  statusMessage: z.string().optional(),
  taskId: z.string(),
  ttl: z.number().int(),
});
export type Task = z.infer<typeof TaskSchema>;


// The response to a tasks/list request.
export const ListTasksResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  nextCursor: z.string().optional(),
  tasks: z.array(TaskSchema),
});
export type ListTasksResult = z.infer<typeof ListTasksResultSchema>;

export const ClientResultSchema = z.union([
  ResultSchema,
  GetTaskResultSchema,
  GetTaskPayloadResultSchema,
  CancelTaskResultSchema,
  ListTasksResultSchema,
  CreateMessageResultSchema,
  ListRootsResultSchema,
  ElicitResultSchema,
]);
export type ClientResult = z.infer<typeof ClientResultSchema>;

export const CompleteResultCompletionSchema = z.object({
  hasMore: z.boolean().optional(),
  total: z.number().int().optional(),
  values: z.array(z.string()),
});
export type CompleteResultCompletion = z.infer<typeof CompleteResultCompletionSchema>;


// The server's response to a completion/complete request
export const CompleteResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  completion: CompleteResultCompletionSchema,
});
export type CompleteResult = z.infer<typeof CompleteResultSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const CreateMessageRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type CreateMessageRequestParamsMeta = z.infer<typeof CreateMessageRequestParamsMetaSchema>;


// Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific.
export const CreateMessageRequestParamsMetadataSchema = z.object({
});
export type CreateMessageRequestParamsMetadata = z.infer<typeof CreateMessageRequestParamsMetadataSchema>;


// Hints to use for model selection.
// 
// Keys not declared here are currently left unspecified by the spec and are up
// to the client to interpret.
export const ModelHintSchema = z.object({
  name: z.string().optional(),
});
export type ModelHint = z.infer<typeof ModelHintSchema>;


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
export const ModelPreferencesSchema = z.object({
  costPriority: z.number().min(0).max(1).optional(),
  hints: z.array(ModelHintSchema).optional(),
  intelligencePriority: z.number().min(0).max(1).optional(),
  speedPriority: z.number().min(0).max(1).optional(),
});
export type ModelPreferences = z.infer<typeof ModelPreferencesSchema>;


// Describes a message issued to or received from an LLM API.
export const SamplingMessageSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  content: z.unknown(),
  role: RoleSchema,
});
export type SamplingMessage = z.infer<typeof SamplingMessageSchema>;


// Additional properties describing a Tool to clients.
// 
// NOTE: all properties in ToolAnnotations are **hints**.
// They are not guaranteed to provide a faithful description of
// tool behavior (including descriptive properties like `title`).
// 
// Clients should never make tool use decisions based on ToolAnnotations
// received from untrusted servers.
export const ToolAnnotationsSchema = z.object({
  destructiveHint: z.boolean().optional(),
  idempotentHint: z.boolean().optional(),
  openWorldHint: z.boolean().optional(),
  readOnlyHint: z.boolean().optional(),
  title: z.string().optional(),
});
export type ToolAnnotations = z.infer<typeof ToolAnnotationsSchema>;


// Execution-related properties for a tool.
export const ToolExecutionSchema = z.object({
  taskSupport: z.string().optional(),
});
export type ToolExecution = z.infer<typeof ToolExecutionSchema>;

export const ToolInputSchemaPropertiesValueSchema = z.object({
});
export type ToolInputSchemaPropertiesValue = z.infer<typeof ToolInputSchemaPropertiesValueSchema>;


// A JSON Schema object defining the expected parameters for the tool.
export const ToolInputSchemaSchema = z.object({
  $schema: z.string().optional(),
  properties: z.record(z.string(), ToolInputSchemaPropertiesValueSchema).optional(),
  required: z.array(z.string()).optional(),
  type: z.string(),
});
export type ToolInputSchema = z.infer<typeof ToolInputSchemaSchema>;

export const ToolOutputSchemaPropertiesValueSchema = z.object({
});
export type ToolOutputSchemaPropertiesValue = z.infer<typeof ToolOutputSchemaPropertiesValueSchema>;


// An optional JSON Schema object defining the structure of the tool's output returned in
// the structuredContent field of a CallToolResult.
// 
// Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
// Currently restricted to type: "object" at the root level.
export const ToolOutputSchemaSchema = z.object({
  $schema: z.string().optional(),
  properties: z.record(z.string(), ToolOutputSchemaPropertiesValueSchema).optional(),
  required: z.array(z.string()).optional(),
  type: z.string(),
});
export type ToolOutputSchema = z.infer<typeof ToolOutputSchemaSchema>;


// Definition for a tool the client can call.
export const ToolSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: ToolAnnotationsSchema.optional(),
  description: z.string().optional(),
  execution: ToolExecutionSchema.optional(),
  icons: z.array(IconSchema).optional(),
  inputSchema: ToolInputSchemaSchema,
  name: z.string(),
  outputSchema: ToolOutputSchemaSchema.optional(),
  title: z.string().optional(),
});
export type Tool = z.infer<typeof ToolSchema>;


// Controls tool selection behavior for sampling requests.
export const ToolChoiceSchema = z.object({
  mode: z.string().optional(),
});
export type ToolChoice = z.infer<typeof ToolChoiceSchema>;


// Parameters for a `sampling/createMessage` request.
export const CreateMessageRequestParamsSchema = z.object({
  _meta: CreateMessageRequestParamsMetaSchema.optional(),
  includeContext: z.string().optional(),
  maxTokens: z.number().int(),
  messages: z.array(SamplingMessageSchema),
  metadata: CreateMessageRequestParamsMetadataSchema.optional(),
  modelPreferences: ModelPreferencesSchema.optional(),
  stopSequences: z.array(z.string()).optional(),
  systemPrompt: z.string().optional(),
  task: TaskMetadataSchema.optional(),
  temperature: z.number().optional(),
  toolChoice: ToolChoiceSchema.optional(),
  tools: z.array(ToolSchema).optional(),
});
export type CreateMessageRequestParams = z.infer<typeof CreateMessageRequestParamsSchema>;


// A request from the server to sample an LLM via the client. The client has full discretion over which model to select. The client should also inform the user before beginning sampling, to allow them to inspect the request (human in the loop) and decide whether to approve it.
export const CreateMessageRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: CreateMessageRequestParamsSchema,
});
export type CreateMessageRequest = z.infer<typeof CreateMessageRequestSchema>;


// A response to a task-augmented request.
export const CreateTaskResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  task: TaskSchema,
});
export type CreateTaskResult = z.infer<typeof CreateTaskResultSchema>;


// An opaque token used to represent a cursor for pagination.
export const CursorSchema = z.string();
export type Cursor = z.infer<typeof CursorSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const ElicitRequestFormParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type ElicitRequestFormParamsMeta = z.infer<typeof ElicitRequestFormParamsMetaSchema>;


// Use TitledSingleSelectEnumSchema instead.
// This interface will be removed in a future version.
export const LegacyTitledEnumSchemaSchema = z.object({
  default: z.string().optional(),
  description: z.string().optional(),
  enum: z.array(z.string()),
  enumNames: z.array(z.string()).optional(),
  title: z.string().optional(),
  type: z.string(),
});
export type LegacyTitledEnumSchema = z.infer<typeof LegacyTitledEnumSchemaSchema>;

export const NumberSchemaSchema = z.object({
  default: z.number().int().optional(),
  description: z.string().optional(),
  maximum: z.number().int().optional(),
  minimum: z.number().int().optional(),
  title: z.string().optional(),
  type: z.string(),
});
export type NumberSchema = z.infer<typeof NumberSchemaSchema>;

export const StringSchemaSchema = z.object({
  default: z.string().optional(),
  description: z.string().optional(),
  format: z.string().optional(),
  maxLength: z.number().int().optional(),
  minLength: z.number().int().optional(),
  title: z.string().optional(),
  type: z.string(),
});
export type StringSchema = z.infer<typeof StringSchemaSchema>;

export const TitledMultiSelectEnumSchemaItemsAnyOfItemSchema = z.object({
  const: z.string(),
  title: z.string(),
});
export type TitledMultiSelectEnumSchemaItemsAnyOfItem = z.infer<typeof TitledMultiSelectEnumSchemaItemsAnyOfItemSchema>;


// Schema for array items with enum options and display labels.
export const TitledMultiSelectEnumSchemaItemsSchema = z.object({
  anyOf: z.array(TitledMultiSelectEnumSchemaItemsAnyOfItemSchema),
});
export type TitledMultiSelectEnumSchemaItems = z.infer<typeof TitledMultiSelectEnumSchemaItemsSchema>;


// Schema for multiple-selection enumeration with display titles for each option.
export const TitledMultiSelectEnumSchemaSchema = z.object({
  default: z.array(z.string()).optional(),
  description: z.string().optional(),
  items: TitledMultiSelectEnumSchemaItemsSchema,
  maxItems: z.number().int().optional(),
  minItems: z.number().int().optional(),
  title: z.string().optional(),
  type: z.string(),
});
export type TitledMultiSelectEnumSchema = z.infer<typeof TitledMultiSelectEnumSchemaSchema>;

export const TitledSingleSelectEnumSchemaOneOfItemSchema = z.object({
  const: z.string(),
  title: z.string(),
});
export type TitledSingleSelectEnumSchemaOneOfItem = z.infer<typeof TitledSingleSelectEnumSchemaOneOfItemSchema>;


// Schema for single-selection enumeration with display titles for each option.
export const TitledSingleSelectEnumSchemaSchema = z.object({
  default: z.string().optional(),
  description: z.string().optional(),
  oneOf: z.array(TitledSingleSelectEnumSchemaOneOfItemSchema),
  title: z.string().optional(),
  type: z.string(),
});
export type TitledSingleSelectEnumSchema = z.infer<typeof TitledSingleSelectEnumSchemaSchema>;


// Schema for the array items.
export const UntitledMultiSelectEnumSchemaItemsSchema = z.object({
  enum: z.array(z.string()),
  type: z.string(),
});
export type UntitledMultiSelectEnumSchemaItems = z.infer<typeof UntitledMultiSelectEnumSchemaItemsSchema>;


// Schema for multiple-selection enumeration without display titles for options.
export const UntitledMultiSelectEnumSchemaSchema = z.object({
  default: z.array(z.string()).optional(),
  description: z.string().optional(),
  items: UntitledMultiSelectEnumSchemaItemsSchema,
  maxItems: z.number().int().optional(),
  minItems: z.number().int().optional(),
  title: z.string().optional(),
  type: z.string(),
});
export type UntitledMultiSelectEnumSchema = z.infer<typeof UntitledMultiSelectEnumSchemaSchema>;


// Schema for single-selection enumeration without display titles for options.
export const UntitledSingleSelectEnumSchemaSchema = z.object({
  default: z.string().optional(),
  description: z.string().optional(),
  enum: z.array(z.string()),
  title: z.string().optional(),
  type: z.string(),
});
export type UntitledSingleSelectEnumSchema = z.infer<typeof UntitledSingleSelectEnumSchemaSchema>;


// Restricted schema definitions that only allow primitive types
// without nested objects or arrays.
export const PrimitiveSchemaDefinitionSchema = z.union([
  StringSchemaSchema,
  NumberSchemaSchema,
  BooleanSchemaSchema,
  UntitledSingleSelectEnumSchemaSchema,
  TitledSingleSelectEnumSchemaSchema,
  UntitledMultiSelectEnumSchemaSchema,
  TitledMultiSelectEnumSchemaSchema,
  LegacyTitledEnumSchemaSchema,
]);
export type PrimitiveSchemaDefinition = z.infer<typeof PrimitiveSchemaDefinitionSchema>;


// A restricted subset of JSON Schema.
// Only top-level properties are allowed, without nesting.
export const ElicitRequestFormParamsRequestedSchemaSchema = z.object({
  $schema: z.string().optional(),
  properties: z.record(z.string(), PrimitiveSchemaDefinitionSchema),
  required: z.array(z.string()).optional(),
  type: z.string(),
});
export type ElicitRequestFormParamsRequestedSchema = z.infer<typeof ElicitRequestFormParamsRequestedSchemaSchema>;


// The parameters for a request to elicit non-sensitive information from the user via a form in the client.
export const ElicitRequestFormParamsSchema = z.object({
  _meta: ElicitRequestFormParamsMetaSchema.optional(),
  message: z.string(),
  mode: z.string().optional(),
  requestedSchema: ElicitRequestFormParamsRequestedSchemaSchema,
  task: TaskMetadataSchema.optional(),
});
export type ElicitRequestFormParams = z.infer<typeof ElicitRequestFormParamsSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const ElicitRequestURLParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type ElicitRequestURLParamsMeta = z.infer<typeof ElicitRequestURLParamsMetaSchema>;


// The parameters for a request to elicit information from the user via a URL in the client.
export const ElicitRequestURLParamsSchema = z.object({
  _meta: ElicitRequestURLParamsMetaSchema.optional(),
  elicitationId: z.string(),
  message: z.string(),
  mode: z.string(),
  task: TaskMetadataSchema.optional(),
  url: z.string().url(),
});
export type ElicitRequestURLParams = z.infer<typeof ElicitRequestURLParamsSchema>;


// The parameters for a request to elicit additional information from the user via the client.
export const ElicitRequestParamsSchema = z.union([
  ElicitRequestURLParamsSchema,
  ElicitRequestFormParamsSchema,
]);
export type ElicitRequestParams = z.infer<typeof ElicitRequestParamsSchema>;


// A request from the server to elicit additional information from the user via the client.
export const ElicitRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: ElicitRequestParamsSchema,
});
export type ElicitRequest = z.infer<typeof ElicitRequestSchema>;

export const ElicitationCompleteNotificationParamsSchema = z.object({
  elicitationId: z.string(),
});
export type ElicitationCompleteNotificationParams = z.infer<typeof ElicitationCompleteNotificationParamsSchema>;


// An optional notification from the server to the client, informing it of a completion of a out-of-band elicitation request.
export const ElicitationCompleteNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: ElicitationCompleteNotificationParamsSchema,
});
export type ElicitationCompleteNotification = z.infer<typeof ElicitationCompleteNotificationSchema>;

export const EmptyResultSchema = z.unknown();
export type EmptyResult = z.infer<typeof EmptyResultSchema>;

export const EnumSchemaSchema = z.union([
  UntitledSingleSelectEnumSchemaSchema,
  TitledSingleSelectEnumSchemaSchema,
  UntitledMultiSelectEnumSchemaSchema,
  TitledMultiSelectEnumSchemaSchema,
  LegacyTitledEnumSchemaSchema,
]);
export type EnumSchema = z.infer<typeof EnumSchemaSchema>;

export const ErrorSchema = z.object({
  code: z.number().int(),
  data: z.unknown().optional(),
  message: z.string(),
});
export type Error = z.infer<typeof ErrorSchema>;


// Describes a message returned as part of a prompt.
// 
// This is similar to `SamplingMessage`, but also supports the embedding of
// resources from the MCP server.
export const PromptMessageSchema = z.object({
  content: ContentBlockSchema,
  role: RoleSchema,
});
export type PromptMessage = z.infer<typeof PromptMessageSchema>;


// The server's response to a prompts/get request from the client.
export const GetPromptResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  description: z.string().optional(),
  messages: z.array(PromptMessageSchema),
});
export type GetPromptResult = z.infer<typeof GetPromptResultSchema>;


// Base interface to add `icons` property.
export const IconsSchema = z.object({
  icons: z.array(IconSchema).optional(),
});
export type Icons = z.infer<typeof IconsSchema>;


// Present if the server supports argument autocompletion suggestions.
export const ServerCapabilitiesCompletionsSchema = z.object({
});
export type ServerCapabilitiesCompletions = z.infer<typeof ServerCapabilitiesCompletionsSchema>;

export const ServerCapabilitiesExperimentalValueSchema = z.object({
});
export type ServerCapabilitiesExperimentalValue = z.infer<typeof ServerCapabilitiesExperimentalValueSchema>;


// Present if the server supports sending log messages to the client.
export const ServerCapabilitiesLoggingSchema = z.object({
});
export type ServerCapabilitiesLogging = z.infer<typeof ServerCapabilitiesLoggingSchema>;


// Present if the server offers any prompt templates.
export const ServerCapabilitiesPromptsSchema = z.object({
  listChanged: z.boolean().optional(),
});
export type ServerCapabilitiesPrompts = z.infer<typeof ServerCapabilitiesPromptsSchema>;


// Present if the server offers any resources to read.
export const ServerCapabilitiesResourcesSchema = z.object({
  listChanged: z.boolean().optional(),
  subscribe: z.boolean().optional(),
});
export type ServerCapabilitiesResources = z.infer<typeof ServerCapabilitiesResourcesSchema>;


// Whether this server supports tasks/cancel.
export const ServerCapabilitiesTasksCancelSchema = z.object({
});
export type ServerCapabilitiesTasksCancel = z.infer<typeof ServerCapabilitiesTasksCancelSchema>;


// Whether this server supports tasks/list.
export const ServerCapabilitiesTasksListSchema = z.object({
});
export type ServerCapabilitiesTasksList = z.infer<typeof ServerCapabilitiesTasksListSchema>;


// Whether the server supports task-augmented tools/call requests.
export const ServerCapabilitiesTasksRequestsToolsCallSchema = z.object({
});
export type ServerCapabilitiesTasksRequestsToolsCall = z.infer<typeof ServerCapabilitiesTasksRequestsToolsCallSchema>;


// Task support for tool-related requests.
export const ServerCapabilitiesTasksRequestsToolsSchema = z.object({
  call: ServerCapabilitiesTasksRequestsToolsCallSchema.optional(),
});
export type ServerCapabilitiesTasksRequestsTools = z.infer<typeof ServerCapabilitiesTasksRequestsToolsSchema>;


// Specifies which request types can be augmented with tasks.
export const ServerCapabilitiesTasksRequestsSchema = z.object({
  tools: ServerCapabilitiesTasksRequestsToolsSchema.optional(),
});
export type ServerCapabilitiesTasksRequests = z.infer<typeof ServerCapabilitiesTasksRequestsSchema>;


// Present if the server supports task-augmented requests.
export const ServerCapabilitiesTasksSchema = z.object({
  cancel: ServerCapabilitiesTasksCancelSchema.optional(),
  list: ServerCapabilitiesTasksListSchema.optional(),
  requests: ServerCapabilitiesTasksRequestsSchema.optional(),
});
export type ServerCapabilitiesTasks = z.infer<typeof ServerCapabilitiesTasksSchema>;


// Present if the server offers any tools to call.
export const ServerCapabilitiesToolsSchema = z.object({
  listChanged: z.boolean().optional(),
});
export type ServerCapabilitiesTools = z.infer<typeof ServerCapabilitiesToolsSchema>;


// Capabilities that a server may support. Known capabilities are defined here, in this schema, but this is not a closed set: any server can define its own, additional capabilities.
export const ServerCapabilitiesSchema = z.object({
  completions: ServerCapabilitiesCompletionsSchema.optional(),
  experimental: z.record(z.string(), ServerCapabilitiesExperimentalValueSchema).optional(),
  logging: ServerCapabilitiesLoggingSchema.optional(),
  prompts: ServerCapabilitiesPromptsSchema.optional(),
  resources: ServerCapabilitiesResourcesSchema.optional(),
  tasks: ServerCapabilitiesTasksSchema.optional(),
  tools: ServerCapabilitiesToolsSchema.optional(),
});
export type ServerCapabilities = z.infer<typeof ServerCapabilitiesSchema>;


// After receiving an initialize request from the client, the server sends this response.
export const InitializeResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  capabilities: ServerCapabilitiesSchema,
  instructions: z.string().optional(),
  protocolVersion: z.string(),
  serverInfo: ImplementationSchema,
});
export type InitializeResult = z.infer<typeof InitializeResultSchema>;


// A response to a request that indicates an error occurred.
export const JSONRPCErrorResponseSchema = z.object({
  error: ErrorSchema,
  id: RequestIdSchema.optional(),
  jsonrpc: z.string(),
});
export type JSONRPCErrorResponse = z.infer<typeof JSONRPCErrorResponseSchema>;


// A notification which does not expect a response.
export const JSONRPCNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: z.record(z.string(), z.unknown()).optional(),
});
export type JSONRPCNotification = z.infer<typeof JSONRPCNotificationSchema>;


// A request that expects a response.
export const JSONRPCRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: z.record(z.string(), z.unknown()).optional(),
});
export type JSONRPCRequest = z.infer<typeof JSONRPCRequestSchema>;


// A successful (non-error) response to a request.
export const JSONRPCResultResponseSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  result: ResultSchema,
});
export type JSONRPCResultResponse = z.infer<typeof JSONRPCResultResponseSchema>;


// Refers to any valid JSON-RPC object that can be decoded off the wire, or encoded to be sent.
export const JSONRPCMessageSchema = z.union([
  JSONRPCRequestSchema,
  JSONRPCNotificationSchema,
  JSONRPCResultResponseSchema,
  JSONRPCErrorResponseSchema,
]);
export type JSONRPCMessage = z.infer<typeof JSONRPCMessageSchema>;


// A response to a request, containing either the result or error.
export const JSONRPCResponseSchema = z.union([
  JSONRPCResultResponseSchema,
  JSONRPCErrorResponseSchema,
]);
export type JSONRPCResponse = z.infer<typeof JSONRPCResponseSchema>;


// Describes an argument that a prompt can accept.
export const PromptArgumentSchema = z.object({
  description: z.string().optional(),
  name: z.string(),
  required: z.boolean().optional(),
  title: z.string().optional(),
});
export type PromptArgument = z.infer<typeof PromptArgumentSchema>;


// A prompt or prompt template that the server offers.
export const PromptSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  arguments: z.array(PromptArgumentSchema).optional(),
  description: z.string().optional(),
  icons: z.array(IconSchema).optional(),
  name: z.string(),
  title: z.string().optional(),
});
export type Prompt = z.infer<typeof PromptSchema>;


// The server's response to a prompts/list request from the client.
export const ListPromptsResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  nextCursor: z.string().optional(),
  prompts: z.array(PromptSchema),
});
export type ListPromptsResult = z.infer<typeof ListPromptsResultSchema>;


// A template description for resources available on the server.
export const ResourceTemplateSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  description: z.string().optional(),
  icons: z.array(IconSchema).optional(),
  mimeType: z.string().optional(),
  name: z.string(),
  title: z.string().optional(),
  uriTemplate: z.string(),
});
export type ResourceTemplate = z.infer<typeof ResourceTemplateSchema>;


// The server's response to a resources/templates/list request from the client.
export const ListResourceTemplatesResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  nextCursor: z.string().optional(),
  resourceTemplates: z.array(ResourceTemplateSchema),
});
export type ListResourceTemplatesResult = z.infer<typeof ListResourceTemplatesResultSchema>;


// A known resource that the server is capable of reading.
export const ResourceSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  annotations: AnnotationsSchema.optional(),
  description: z.string().optional(),
  icons: z.array(IconSchema).optional(),
  mimeType: z.string().optional(),
  name: z.string(),
  size: z.number().int().optional(),
  title: z.string().optional(),
  uri: z.string().url(),
});
export type Resource = z.infer<typeof ResourceSchema>;


// The server's response to a resources/list request from the client.
export const ListResourcesResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  nextCursor: z.string().optional(),
  resources: z.array(ResourceSchema),
});
export type ListResourcesResult = z.infer<typeof ListResourcesResultSchema>;


// Sent from the server to request a list of root URIs from the client. Roots allow
// servers to ask for specific directories or files to operate on. A common example
// for roots is providing a set of repositories or directories a server should operate
// on.
// 
// This request is typically used when the server needs to understand the file system
// structure or access specific locations that the client has permission to read from.
export const ListRootsRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: RequestParamsSchema.optional(),
});
export type ListRootsRequest = z.infer<typeof ListRootsRequestSchema>;


// The server's response to a tools/list request from the client.
export const ListToolsResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  nextCursor: z.string().optional(),
  tools: z.array(ToolSchema),
});
export type ListToolsResult = z.infer<typeof ListToolsResultSchema>;


// Parameters for a `notifications/message` notification.
export const LoggingMessageNotificationParamsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  data: z.unknown(),
  level: LoggingLevelSchema,
  logger: z.string().optional(),
});
export type LoggingMessageNotificationParams = z.infer<typeof LoggingMessageNotificationParamsSchema>;


// JSONRPCNotification of a log message passed from server to client. If no logging/setLevel request has been sent from the client, the server MAY decide which messages to send automatically.
export const LoggingMessageNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: LoggingMessageNotificationParamsSchema,
});
export type LoggingMessageNotification = z.infer<typeof LoggingMessageNotificationSchema>;

export const MultiSelectEnumSchemaSchema = z.union([
  UntitledMultiSelectEnumSchemaSchema,
  TitledMultiSelectEnumSchemaSchema,
]);
export type MultiSelectEnumSchema = z.infer<typeof MultiSelectEnumSchemaSchema>;

export const NotificationSchema = z.object({
  method: z.string(),
  params: z.record(z.string(), z.unknown()).optional(),
});
export type Notification = z.infer<typeof NotificationSchema>;

export const PaginatedRequestSchema = z.object({
  id: RequestIdSchema,
  jsonrpc: z.string(),
  method: z.string(),
  params: PaginatedRequestParamsSchema.optional(),
});
export type PaginatedRequest = z.infer<typeof PaginatedRequestSchema>;

export const PaginatedResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  nextCursor: z.string().optional(),
});
export type PaginatedResult = z.infer<typeof PaginatedResultSchema>;


// An optional notification from the server to the client, informing it that the list of prompts it offers has changed. This may be issued by servers without any previous subscription from the client.
export const PromptListChangedNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: NotificationParamsSchema.optional(),
});
export type PromptListChangedNotification = z.infer<typeof PromptListChangedNotificationSchema>;


// Identifies a prompt.
export const PromptReferenceSchema = z.object({
  name: z.string(),
  title: z.string().optional(),
  type: z.string(),
});
export type PromptReference = z.infer<typeof PromptReferenceSchema>;


// The server's response to a resources/read request from the client.
export const ReadResourceResultSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  contents: z.array(z.unknown()),
});
export type ReadResourceResult = z.infer<typeof ReadResourceResultSchema>;


// Metadata for associating messages with a task.
// Include this in the `_meta` field under the key `io.modelcontextprotocol/related-task`.
export const RelatedTaskMetadataSchema = z.object({
  taskId: z.string(),
});
export type RelatedTaskMetadata = z.infer<typeof RelatedTaskMetadataSchema>;

export const RequestSchema = z.object({
  method: z.string(),
  params: z.record(z.string(), z.unknown()).optional(),
});
export type Request = z.infer<typeof RequestSchema>;


// The contents of a specific resource or sub-resource.
export const ResourceContentsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  mimeType: z.string().optional(),
  uri: z.string().url(),
});
export type ResourceContents = z.infer<typeof ResourceContentsSchema>;


// An optional notification from the server to the client, informing it that the list of resources it can read from has changed. This may be issued by servers without any previous subscription from the client.
export const ResourceListChangedNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: NotificationParamsSchema.optional(),
});
export type ResourceListChangedNotification = z.infer<typeof ResourceListChangedNotificationSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const ResourceRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type ResourceRequestParamsMeta = z.infer<typeof ResourceRequestParamsMetaSchema>;


// Common parameters when working with resources.
export const ResourceRequestParamsSchema = z.object({
  _meta: ResourceRequestParamsMetaSchema.optional(),
  uri: z.string().url(),
});
export type ResourceRequestParams = z.infer<typeof ResourceRequestParamsSchema>;


// A reference to a resource or resource template definition.
export const ResourceTemplateReferenceSchema = z.object({
  type: z.string(),
  uri: z.string(),
});
export type ResourceTemplateReference = z.infer<typeof ResourceTemplateReferenceSchema>;


// Parameters for a `notifications/resources/updated` notification.
export const ResourceUpdatedNotificationParamsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  uri: z.string().url(),
});
export type ResourceUpdatedNotificationParams = z.infer<typeof ResourceUpdatedNotificationParamsSchema>;


// A notification from the server to the client, informing it that a resource has changed and may need to be read again. This should only be sent if the client previously sent a resources/subscribe request.
export const ResourceUpdatedNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: ResourceUpdatedNotificationParamsSchema,
});
export type ResourceUpdatedNotification = z.infer<typeof ResourceUpdatedNotificationSchema>;


// The result of a tool use, provided by the user back to the assistant.
export const ToolResultContentSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  content: z.array(ContentBlockSchema),
  isError: z.boolean().optional(),
  structuredContent: z.record(z.string(), z.unknown()).optional(),
  toolUseId: z.string(),
  type: z.string(),
});
export type ToolResultContent = z.infer<typeof ToolResultContentSchema>;


// A request from the assistant to call a tool.
export const ToolUseContentSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  id: z.string(),
  input: z.record(z.string(), z.unknown()),
  name: z.string(),
  type: z.string(),
});
export type ToolUseContent = z.infer<typeof ToolUseContentSchema>;

export const SamplingMessageContentBlockSchema = z.union([
  TextContentSchema,
  ImageContentSchema,
  AudioContentSchema,
  ToolUseContentSchema,
  ToolResultContentSchema,
]);
export type SamplingMessageContentBlock = z.infer<typeof SamplingMessageContentBlockSchema>;


// An optional notification from the server to the client, informing it that the list of tools it offers has changed. This may be issued by servers without any previous subscription from the client.
export const ToolListChangedNotificationSchema = z.object({
  jsonrpc: z.string(),
  method: z.string(),
  params: NotificationParamsSchema.optional(),
});
export type ToolListChangedNotification = z.infer<typeof ToolListChangedNotificationSchema>;

export const ServerNotificationSchema = z.union([
  CancelledNotificationSchema,
  ProgressNotificationSchema,
  ResourceListChangedNotificationSchema,
  ResourceUpdatedNotificationSchema,
  PromptListChangedNotificationSchema,
  ToolListChangedNotificationSchema,
  TaskStatusNotificationSchema,
  LoggingMessageNotificationSchema,
  ElicitationCompleteNotificationSchema,
]);
export type ServerNotification = z.infer<typeof ServerNotificationSchema>;

export const ServerRequestSchema = z.union([
  PingRequestSchema,
  GetTaskRequestSchema,
  GetTaskPayloadRequestSchema,
  CancelTaskRequestSchema,
  ListTasksRequestSchema,
  CreateMessageRequestSchema,
  ListRootsRequestSchema,
  ElicitRequestSchema,
]);
export type ServerRequest = z.infer<typeof ServerRequestSchema>;

export const ServerResultSchema = z.union([
  ResultSchema,
  InitializeResultSchema,
  ListResourcesResultSchema,
  ListResourceTemplatesResultSchema,
  ReadResourceResultSchema,
  ListPromptsResultSchema,
  GetPromptResultSchema,
  ListToolsResultSchema,
  CallToolResultSchema,
  GetTaskResultSchema,
  GetTaskPayloadResultSchema,
  CancelTaskResultSchema,
  ListTasksResultSchema,
  CompleteResultSchema,
]);
export type ServerResult = z.infer<typeof ServerResultSchema>;

export const SingleSelectEnumSchemaSchema = z.union([
  UntitledSingleSelectEnumSchemaSchema,
  TitledSingleSelectEnumSchemaSchema,
]);
export type SingleSelectEnumSchema = z.infer<typeof SingleSelectEnumSchemaSchema>;


// See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
export const TaskAugmentedRequestParamsMetaSchema = z.object({
  progressToken: ProgressTokenSchema.optional(),
});
export type TaskAugmentedRequestParamsMeta = z.infer<typeof TaskAugmentedRequestParamsMetaSchema>;


// Common params for any task-augmented request.
export const TaskAugmentedRequestParamsSchema = z.object({
  _meta: TaskAugmentedRequestParamsMetaSchema.optional(),
  task: TaskMetadataSchema.optional(),
});
export type TaskAugmentedRequestParams = z.infer<typeof TaskAugmentedRequestParamsSchema>;

export const TextResourceContentsSchema = z.object({
  _meta: z.record(z.string(), z.unknown()).optional(),
  mimeType: z.string().optional(),
  text: z.string(),
  uri: z.string().url(),
});
export type TextResourceContents = z.infer<typeof TextResourceContentsSchema>;


// An error response that indicates that the server requires the client to provide additional information via an elicitation request.
export const URLElicitationRequiredErrorSchema = z.object({
  error: z.unknown(),
  id: RequestIdSchema.optional(),
  jsonrpc: z.string(),
});
export type URLElicitationRequiredError = z.infer<typeof URLElicitationRequiredErrorSchema>;
