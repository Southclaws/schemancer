from __future__ import annotations

from typing import Any, Dict, List
from enum import Enum
from pydantic import AnyUrl, BaseModel, ConfigDict, Field, RootModel
"""The sender or recipient of messages and data in a conversation."""
class Role(str, Enum):
    ASSISTANT = "assistant"
    USER = "user"


class Annotations(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Describes who the intended audience of this object or data is.
    
    It can include multiple entries to indicate content useful for multiple audiences (e.g., `["user", "assistant"]`).
    """
    audience: List[Role] | None = None
    """
    The moment the resource was last modified, as an ISO 8601 formatted string.
    
    Should be an ISO 8601 formatted string (e.g., "2025-01-12T15:00:58Z").
    
    Examples: last activity timestamp in an open file, timestamp when the resource
    was attached, etc.
    """
    last_modified: str | None = None
    """
    Describes how important this data is for operating the server.
    
    A value of 1 means "most important," and indicates that the data is
    effectively required, while 0 means "least important," and indicates that
    the data is entirely optional.
    """
    priority: float | None = Field(ge=0, le=1, default=None)

class AudioContent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    """The base64-encoded audio data."""
    data: bytes
    """The MIME type of the audio. Different providers may support different audio types."""
    mime_type: str
    type: str

class BaseMetadata(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None

class BlobResourceContents(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """A base64-encoded string representing the binary data of the item."""
    blob: bytes
    """The MIME type of this resource, if known."""
    mime_type: str | None = None
    """The URI of this resource."""
    uri: AnyUrl

class BooleanSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    default: bool | None = None
    description: str | None = None
    title: str | None = None
    type: str
"""A progress token, used to associate progress notifications with the original request."""
class ProgressToken(RootModel[Any]):
    pass


class CallToolRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class TaskMetadata(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Requested duration in milliseconds to retain task from creation."""
    ttl: int | None = None

class CallToolRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: CallToolRequestParamsMeta | None = None
    """Arguments to use for the tool call."""
    arguments: Dict[str, Any] | None = None
    """The name of the tool."""
    name: str
    """
    If specified, the caller is requesting task-augmented execution for this request.
    The request will return a CreateTaskResult immediately, and the actual result can be
    retrieved later via tasks/result.
    
    Task augmentation is subject to capability negotiation - receivers MUST declare support
    for task augmentation of specific request types in their capabilities.
    """
    task: TaskMetadata | None = None
"""A uniquely identifying ID for a request in JSON-RPC."""
class RequestId(RootModel[Any]):
    pass


class CallToolRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: CallToolRequestParams

class EmbeddedResource(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    resource: Any
    type: str

class ImageContent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    """The base64-encoded image data."""
    data: bytes
    """The MIME type of the image. Different providers may support different image types."""
    mime_type: str
    type: str

class Icon(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Optional MIME type override if the source MIME type is missing or generic.
    For example: `"image/png"`, `"image/jpeg"`, or `"image/svg+xml"`.
    """
    mime_type: str | None = None
    """
    Optional array of strings that specify sizes at which the icon can be used.
    Each string should be in WxH format (e.g., `"48x48"`, `"96x96"`) or `"any"` for scalable formats like SVG.
    
    If not provided, the client should assume that the icon can be used at any size.
    """
    sizes: List[str] | None = None
    """
    A standard URI pointing to an icon resource. May be an HTTP/HTTPS URL or a
    `data:` URI with Base64-encoded image data.
    
    Consumers SHOULD takes steps to ensure URLs serving icons are from the
    same domain as the client/server or a trusted domain.
    
    Consumers SHOULD take appropriate precautions when consuming SVGs as they can contain
    executable JavaScript.
    """
    src: AnyUrl
    """
    Optional specifier for the theme this icon is designed for. `light` indicates
    the icon is designed to be used with a light background, and `dark` indicates
    the icon is designed to be used with a dark background.
    
    If not provided, the client should assume the icon can be used with any theme.
    """
    theme: str | None = None

class ResourceLink(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    """
    A description of what this resource represents.
    
    This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
    """
    description: str | None = None
    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None
    """The MIME type of this resource, if known."""
    mime_type: str | None = None
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
    
    This can be used by Hosts to display file sizes and estimate context window usage.
    """
    size: int | None = None
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None
    type: str
    """The URI of this resource."""
    uri: AnyUrl

class TextContent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    """The text content of the message."""
    text: str
    type: str
ContentBlock = Union[TextContent, ImageContent, AudioContent, ResourceLink, EmbeddedResource]


class CallToolResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """A list of content objects that represent the unstructured result of the tool call."""
    content: List[ContentBlock]
    """
    Whether the tool call ended in an error.
    
    If not set, this is assumed to be false (the call was successful).
    
    Any errors that originate from the tool SHOULD be reported inside the result
    object, with `isError` set to true, _not_ as an MCP protocol-level error
    response. Otherwise, the LLM would not be able to see that an error occurred
    and self-correct.
    
    However, any errors in _finding_ the tool, an error indicating that the
    server does not support tool calls, or any other exceptional conditions,
    should be reported as an MCP error response.
    """
    is_error: bool | None = None
    """An optional JSON object that represents the structured result of the tool call."""
    structured_content: Dict[str, Any] | None = None

class CancelTaskRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The task identifier to cancel."""
    task_id: str

class CancelTaskRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: CancelTaskRequestParams
"""The status of a task."""
class TaskStatus(str, Enum):
    CANCELLED = "cancelled"
    COMPLETED = "completed"
    FAILED = "failed"
    INPUT_REQUIRED = "input_required"
    WORKING = "working"


class CancelTaskResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """ISO 8601 timestamp when the task was created."""
    created_at: str
    """ISO 8601 timestamp when the task was last updated."""
    last_updated_at: str
    """Suggested polling interval in milliseconds."""
    poll_interval: int | None = None
    """Current task state."""
    status: TaskStatus
    """
    Optional human-readable message describing the current task state.
    This can provide context for any status, including:
    - Reasons for "cancelled" status
    - Summaries for "completed" status
    - Diagnostic information for "failed" status (e.g., error details, what went wrong)
    """
    status_message: str | None = None
    """The task identifier."""
    task_id: str
    """Actual retention duration from creation in milliseconds, null for unlimited."""
    ttl: int

class CancelledNotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """An optional string describing the reason for the cancellation. This MAY be logged or presented to the user."""
    reason: str | None = None
    """
    The ID of the request to cancel.
    
    This MUST correspond to the ID of a request previously issued in the same direction.
    This MUST be provided for cancelling non-task requests.
    This MUST NOT be used for cancelling tasks (use the `tasks/cancel` request instead).
    """
    request_id: RequestId | None = None

class CancelledNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: CancelledNotificationParams

class ClientCapabilitiesElicitationForm(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesElicitationURL(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesElicitation(BaseModel):
    model_config = ConfigDict(extra="forbid")

    form: ClientCapabilitiesElicitationForm | None = None
    url: ClientCapabilitiesElicitationURL | None = None

class ClientCapabilitiesRoots(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether the client supports notifications for changes to the roots list."""
    list_changed: bool | None = None

class ClientCapabilitiesSamplingContext(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesSamplingTools(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesSampling(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Whether the client supports context inclusion via includeContext parameter.
    If not declared, servers SHOULD only use `includeContext: "none"` (or omit it).
    """
    context: ClientCapabilitiesSamplingContext | None = None
    """Whether the client supports tool use via tools and toolChoice parameters."""
    tools: ClientCapabilitiesSamplingTools | None = None

class ClientCapabilitiesTasksCancel(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesTasksList(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesTasksRequestsElicitationCreate(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesTasksRequestsElicitation(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether the client supports task-augmented elicitation/create requests."""
    create: ClientCapabilitiesTasksRequestsElicitationCreate | None = None

class ClientCapabilitiesTasksRequestsSamplingCreateMessage(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ClientCapabilitiesTasksRequestsSampling(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether the client supports task-augmented sampling/createMessage requests."""
    create_message: ClientCapabilitiesTasksRequestsSamplingCreateMessage | None = None

class ClientCapabilitiesTasksRequests(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Task support for elicitation-related requests."""
    elicitation: ClientCapabilitiesTasksRequestsElicitation | None = None
    """Task support for sampling-related requests."""
    sampling: ClientCapabilitiesTasksRequestsSampling | None = None

class ClientCapabilitiesTasks(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether this client supports tasks/cancel."""
    cancel: ClientCapabilitiesTasksCancel | None = None
    """Whether this client supports tasks/list."""
    list: ClientCapabilitiesTasksList | None = None
    """Specifies which request types can be augmented with tasks."""
    requests: ClientCapabilitiesTasksRequests | None = None

class ClientCapabilities(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Present if the client supports elicitation from the server."""
    elicitation: ClientCapabilitiesElicitation | None = None
    """Experimental, non-standard capabilities that the client supports."""
    experimental: Dict[str, Any] | None = None
    """Present if the client supports listing roots."""
    roots: ClientCapabilitiesRoots | None = None
    """Present if the client supports sampling from an LLM."""
    sampling: ClientCapabilitiesSampling | None = None
    """Present if the client supports task-augmented requests."""
    tasks: ClientCapabilitiesTasks | None = None

class NotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None

class InitializedNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: NotificationParams | None = None

class ProgressNotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """An optional message describing the current progress."""
    message: str | None = None
    """The progress thus far. This should increase every time progress is made, even if the total is unknown."""
    progress: float
    """The progress token which was given in the initial request, used to associate this notification with the request that is proceeding."""
    progress_token: ProgressToken
    """Total number of items to process (or total progress required), if known."""
    total: float | None = None

class ProgressNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: ProgressNotificationParams

class RootsListChangedNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: NotificationParams | None = None

class TaskStatusNotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """ISO 8601 timestamp when the task was created."""
    created_at: str
    """ISO 8601 timestamp when the task was last updated."""
    last_updated_at: str
    """Suggested polling interval in milliseconds."""
    poll_interval: int | None = None
    """Current task state."""
    status: TaskStatus
    """
    Optional human-readable message describing the current task state.
    This can provide context for any status, including:
    - Reasons for "cancelled" status
    - Summaries for "completed" status
    - Diagnostic information for "failed" status (e.g., error details, what went wrong)
    """
    status_message: str | None = None
    """The task identifier."""
    task_id: str
    """Actual retention duration from creation in milliseconds, null for unlimited."""
    ttl: int

class TaskStatusNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: TaskStatusNotificationParams
ClientNotification = Union[CancelledNotification, InitializedNotification, ProgressNotification, TaskStatusNotification, RootsListChangedNotification]


class CompleteRequestParamsArgument(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The name of the argument"""
    name: str
    """The value of the argument to use for completion matching."""
    value: str

class CompleteRequestParamsContext(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Previously-resolved variables in a URI template or prompt."""
    arguments: Dict[str, Any] | None = None

class CompleteRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class CompleteRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: CompleteRequestParamsMeta | None = None
    """The argument's information"""
    argument: CompleteRequestParamsArgument
    """Additional, optional context for completions"""
    context: CompleteRequestParamsContext | None = None
    ref: Any

class CompleteRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: CompleteRequestParams

class GetPromptRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class GetPromptRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: GetPromptRequestParamsMeta | None = None
    """Arguments to use for templating the prompt."""
    arguments: Dict[str, Any] | None = None
    """The name of the prompt or prompt template."""
    name: str

class GetPromptRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: GetPromptRequestParams

class GetTaskPayloadRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The task identifier to retrieve results for."""
    task_id: str

class GetTaskPayloadRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: GetTaskPayloadRequestParams

class GetTaskRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The task identifier to query."""
    task_id: str

class GetTaskRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: GetTaskRequestParams

class Implementation(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    An optional human-readable description of what this implementation does.
    
    This can be used by clients or servers to provide context about their purpose
    and capabilities. For example, a server might describe the types of resources
    or tools it provides, while a client might describe its intended use case.
    """
    description: str | None = None
    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None
    version: str
    """An optional URL of the website for this implementation."""
    website_url: AnyUrl | None = None

class InitializeRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class InitializeRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: InitializeRequestParamsMeta | None = None
    capabilities: ClientCapabilities
    client_info: Implementation
    """The latest version of the Model Context Protocol that the client supports. The client MAY decide to support older versions as well."""
    protocol_version: str

class InitializeRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: InitializeRequestParams

class PaginatedRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class PaginatedRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: PaginatedRequestParamsMeta | None = None
    """
    An opaque token representing the current pagination position.
    If provided, the server should return results starting after this cursor.
    """
    cursor: str | None = None

class ListPromptsRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: PaginatedRequestParams | None = None

class ListResourceTemplatesRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: PaginatedRequestParams | None = None

class ListResourcesRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: PaginatedRequestParams | None = None

class ListTasksRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: PaginatedRequestParams | None = None

class ListToolsRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: PaginatedRequestParams | None = None

class RequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class RequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: RequestParamsMeta | None = None

class PingRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: RequestParams | None = None

class ReadResourceRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class ReadResourceRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: ReadResourceRequestParamsMeta | None = None
    """The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it."""
    uri: AnyUrl

class ReadResourceRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: ReadResourceRequestParams
"""
The severity of a log message.

These map to syslog message severities, as specified in RFC-5424:
https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1
"""
class LoggingLevel(str, Enum):
    ALERT = "alert"
    CRITICAL = "critical"
    DEBUG = "debug"
    EMERGENCY = "emergency"
    ERROR = "error"
    INFO = "info"
    NOTICE = "notice"
    WARNING = "warning"


class SetLevelRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class SetLevelRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: SetLevelRequestParamsMeta | None = None
    """The level of logging that the client wants to receive from the server. The server should send all logs at this level and higher (i.e., more severe) to the client as notifications/message."""
    level: LoggingLevel

class SetLevelRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: SetLevelRequestParams

class SubscribeRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class SubscribeRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: SubscribeRequestParamsMeta | None = None
    """The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it."""
    uri: AnyUrl

class SubscribeRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: SubscribeRequestParams

class UnsubscribeRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class UnsubscribeRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: UnsubscribeRequestParamsMeta | None = None
    """The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it."""
    uri: AnyUrl

class UnsubscribeRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: UnsubscribeRequestParams
ClientRequest = Union[InitializeRequest, PingRequest, ListResourcesRequest, ListResourceTemplatesRequest, ReadResourceRequest, SubscribeRequest, UnsubscribeRequest, ListPromptsRequest, GetPromptRequest, ListToolsRequest, CallToolRequest, GetTaskRequest, GetTaskPayloadRequest, CancelTaskRequest, ListTasksRequest, SetLevelRequest, CompleteRequest]


class CreateMessageResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    content: Any
    """The name of the model that generated the message."""
    model: str
    role: Role
    """
    The reason why sampling stopped, if known.
    
    Standard values:
    - "endTurn": Natural end of the assistant's turn
    - "stopSequence": A stop sequence was encountered
    - "maxTokens": Maximum token limit was reached
    - "toolUse": The model wants to use one or more tools
    
    This field is an open string to allow for provider-specific stop reasons.
    """
    stop_reason: str | None = None

class ElicitResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    The user action in response to the elicitation.
    - "accept": User submitted the form/confirmed the action
    - "decline": User explicitly decline the action
    - "cancel": User dismissed without making an explicit choice
    """
    action: str
    """
    The submitted form data, only present when action is "accept" and mode was "form".
    Contains values matching the requested schema.
    Omitted for out-of-band mode responses.
    """
    content: Dict[str, Any] | None = None

class GetTaskPayloadResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None

class GetTaskResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """ISO 8601 timestamp when the task was created."""
    created_at: str
    """ISO 8601 timestamp when the task was last updated."""
    last_updated_at: str
    """Suggested polling interval in milliseconds."""
    poll_interval: int | None = None
    """Current task state."""
    status: TaskStatus
    """
    Optional human-readable message describing the current task state.
    This can provide context for any status, including:
    - Reasons for "cancelled" status
    - Summaries for "completed" status
    - Diagnostic information for "failed" status (e.g., error details, what went wrong)
    """
    status_message: str | None = None
    """The task identifier."""
    task_id: str
    """Actual retention duration from creation in milliseconds, null for unlimited."""
    ttl: int

class Root(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An optional name for the root. This can be used to provide a human-readable
    identifier for the root, which may be useful for display purposes or for
    referencing the root in other parts of the application.
    """
    name: str | None = None
    """
    The URI identifying the root. This *must* start with file:// for now.
    This restriction may be relaxed in future versions of the protocol to allow
    other URI schemes.
    """
    uri: AnyUrl

class ListRootsResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    roots: List[Root]

class Task(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """ISO 8601 timestamp when the task was created."""
    created_at: str
    """ISO 8601 timestamp when the task was last updated."""
    last_updated_at: str
    """Suggested polling interval in milliseconds."""
    poll_interval: int | None = None
    """Current task state."""
    status: TaskStatus
    """
    Optional human-readable message describing the current task state.
    This can provide context for any status, including:
    - Reasons for "cancelled" status
    - Summaries for "completed" status
    - Diagnostic information for "failed" status (e.g., error details, what went wrong)
    """
    status_message: str | None = None
    """The task identifier."""
    task_id: str
    """Actual retention duration from creation in milliseconds, null for unlimited."""
    ttl: int

class ListTasksResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An opaque token representing the pagination position after the last returned result.
    If present, there may be more results available.
    """
    next_cursor: str | None = None
    tasks: List[Task]

class Result(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
ClientResult = Union[Result, GetTaskResult, GetTaskPayloadResult, CancelTaskResult, ListTasksResult, CreateMessageResult, ListRootsResult, ElicitResult]


class CompleteResultCompletion(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Indicates whether there are additional completion options beyond those provided in the current response, even if the exact total is unknown."""
    has_more: bool | None = None
    """The total number of completion options available. This can exceed the number of values actually sent in the response."""
    total: int | None = None
    """An array of completion values. Must not exceed 100 items."""
    values: List[str]

class CompleteResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    completion: CompleteResultCompletion

class CreateMessageRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class CreateMessageRequestParamsMetadata(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ModelHint(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    A hint for a model name.
    
    The client SHOULD treat this as a substring of a model name; for example:
     - `claude-3-5-sonnet` should match `claude-3-5-sonnet-20241022`
     - `sonnet` should match `claude-3-5-sonnet-20241022`, `claude-3-sonnet-20240229`, etc.
     - `claude` should match any Claude model
    
    The client MAY also map the string to a different provider's model name or a different model family, as long as it fills a similar niche; for example:
     - `gemini-1.5-flash` could match `claude-3-haiku-20240307`
    """
    name: str | None = None

class ModelPreferences(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    How much to prioritize cost when selecting a model. A value of 0 means cost
    is not important, while a value of 1 means cost is the most important
    factor.
    """
    cost_priority: float | None = Field(ge=0, le=1, default=None)
    """
    Optional hints to use for model selection.
    
    If multiple hints are specified, the client MUST evaluate them in order
    (such that the first match is taken).
    
    The client SHOULD prioritize these hints over the numeric priorities, but
    MAY still use the priorities to select from ambiguous matches.
    """
    hints: List[ModelHint] | None = None
    """
    How much to prioritize intelligence and capabilities when selecting a
    model. A value of 0 means intelligence is not important, while a value of 1
    means intelligence is the most important factor.
    """
    intelligence_priority: float | None = Field(ge=0, le=1, default=None)
    """
    How much to prioritize sampling speed (latency) when selecting a model. A
    value of 0 means speed is not important, while a value of 1 means speed is
    the most important factor.
    """
    speed_priority: float | None = Field(ge=0, le=1, default=None)

class SamplingMessage(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    content: Any
    role: Role

class ToolAnnotations(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    If true, the tool may perform destructive updates to its environment.
    If false, the tool performs only additive updates.
    
    (This property is meaningful only when `readOnlyHint == false`)
    
    Default: true
    """
    destructive_hint: bool | None = None
    """
    If true, calling the tool repeatedly with the same arguments
    will have no additional effect on its environment.
    
    (This property is meaningful only when `readOnlyHint == false`)
    
    Default: false
    """
    idempotent_hint: bool | None = None
    """
    If true, this tool may interact with an "open world" of external
    entities. If false, the tool's domain of interaction is closed.
    For example, the world of a web search tool is open, whereas that
    of a memory tool is not.
    
    Default: true
    """
    open_world_hint: bool | None = None
    """
    If true, the tool does not modify its environment.
    
    Default: false
    """
    read_only_hint: bool | None = None
    """A human-readable title for the tool."""
    title: str | None = None

class ToolExecution(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Indicates whether this tool supports task-augmented execution.
    This allows clients to handle long-running operations through polling
    the task system.
    
    - "forbidden": Tool does not support task-augmented execution (default when absent)
    - "optional": Tool may support task-augmented execution
    - "required": Tool requires task-augmented execution
    
    Default: "forbidden"
    """
    task_support: str | None = None

class ToolInputSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    schema: str | None = None
    properties: Dict[str, Any] | None = None
    required: List[str] | None = None
    type: str

class ToolOutputSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    schema: str | None = None
    properties: Dict[str, Any] | None = None
    required: List[str] | None = None
    type: str

class Tool(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    Optional additional tool information.
    
    Display name precedence order is: title, annotations.title, then name.
    """
    annotations: ToolAnnotations | None = None
    """
    A human-readable description of the tool.
    
    This can be used by clients to improve the LLM's understanding of available tools. It can be thought of like a "hint" to the model.
    """
    description: str | None = None
    """Execution-related properties for this tool."""
    execution: ToolExecution | None = None
    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None
    """A JSON Schema object defining the expected parameters for the tool."""
    input_schema: ToolInputSchema
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    An optional JSON Schema object defining the structure of the tool's output returned in
    the structuredContent field of a CallToolResult.
    
    Defaults to JSON Schema 2020-12 when no explicit $schema is provided.
    Currently restricted to type: "object" at the root level.
    """
    output_schema: ToolOutputSchema | None = None
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None

class ToolChoice(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Controls the tool use ability of the model:
    - "auto": Model decides whether to use tools (default)
    - "required": Model MUST use at least one tool before completing
    - "none": Model MUST NOT use any tools
    """
    mode: str | None = None

class CreateMessageRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: CreateMessageRequestParamsMeta | None = None
    """
    A request to include context from one or more MCP servers (including the caller), to be attached to the prompt.
    The client MAY ignore this request.
    
    Default is "none". Values "thisServer" and "allServers" are soft-deprecated. Servers SHOULD only use these values if the client
    declares ClientCapabilities.sampling.context. These values may be removed in future spec releases.
    """
    include_context: str | None = None
    """
    The requested maximum number of tokens to sample (to prevent runaway completions).
    
    The client MAY choose to sample fewer tokens than the requested maximum.
    """
    max_tokens: int
    messages: List[SamplingMessage]
    """Optional metadata to pass through to the LLM provider. The format of this metadata is provider-specific."""
    metadata: CreateMessageRequestParamsMetadata | None = None
    """The server's preferences for which model to select. The client MAY ignore these preferences."""
    model_preferences: ModelPreferences | None = None
    stop_sequences: List[str] | None = None
    """An optional system prompt the server wants to use for sampling. The client MAY modify or omit this prompt."""
    system_prompt: str | None = None
    """
    If specified, the caller is requesting task-augmented execution for this request.
    The request will return a CreateTaskResult immediately, and the actual result can be
    retrieved later via tasks/result.
    
    Task augmentation is subject to capability negotiation - receivers MUST declare support
    for task augmentation of specific request types in their capabilities.
    """
    task: TaskMetadata | None = None
    temperature: float | None = None
    """
    Controls how the model uses tools.
    The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
    Default is `{ mode: "auto" }`.
    """
    tool_choice: ToolChoice | None = None
    """
    Tools that the model may use during generation.
    The client MUST return an error if this field is provided but ClientCapabilities.sampling.tools is not declared.
    """
    tools: List[Tool] | None = None

class CreateMessageRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: CreateMessageRequestParams

class CreateTaskResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    task: Task
"""An opaque token used to represent a cursor for pagination."""
class Cursor(RootModel[str]):
    pass


class ElicitRequestFormParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class ElicitRequestFormParamsRequestedSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    schema: str | None = None
    properties: Dict[str, Any]
    required: List[str] | None = None
    type: str

class ElicitRequestFormParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: ElicitRequestFormParamsMeta | None = None
    """The message to present to the user describing what information is being requested."""
    message: str
    """The elicitation mode."""
    mode: str | None = None
    """
    A restricted subset of JSON Schema.
    Only top-level properties are allowed, without nesting.
    """
    requested_schema: ElicitRequestFormParamsRequestedSchema
    """
    If specified, the caller is requesting task-augmented execution for this request.
    The request will return a CreateTaskResult immediately, and the actual result can be
    retrieved later via tasks/result.
    
    Task augmentation is subject to capability negotiation - receivers MUST declare support
    for task augmentation of specific request types in their capabilities.
    """
    task: TaskMetadata | None = None

class ElicitRequestURLParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class ElicitRequestURLParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: ElicitRequestURLParamsMeta | None = None
    """
    The ID of the elicitation, which must be unique within the context of the server.
    The client MUST treat this ID as an opaque value.
    """
    elicitation_id: str
    """The message to present to the user explaining why the interaction is needed."""
    message: str
    """The elicitation mode."""
    mode: str
    """
    If specified, the caller is requesting task-augmented execution for this request.
    The request will return a CreateTaskResult immediately, and the actual result can be
    retrieved later via tasks/result.
    
    Task augmentation is subject to capability negotiation - receivers MUST declare support
    for task augmentation of specific request types in their capabilities.
    """
    task: TaskMetadata | None = None
    """The URL that the user should navigate to."""
    url: AnyUrl
"""The parameters for a request to elicit additional information from the user via the client."""
ElicitRequestParams = Union[ElicitRequestURLParams, ElicitRequestFormParams]


class ElicitRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: ElicitRequestParams

class ElicitationCompleteNotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The ID of the elicitation that completed."""
    elicitation_id: str

class ElicitationCompleteNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: ElicitationCompleteNotificationParams
class EmptyResult(RootModel[Any]):
    pass


class LegacyTitledEnumSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    default: str | None = None
    description: str | None = None
    enum: List[str]
    """
    (Legacy) Display names for enum values.
    Non-standard according to JSON schema 2020-12.
    """
    enum_names: List[str] | None = None
    title: str | None = None
    type: str

class TitledMultiSelectEnumSchemaItemsAnyOfItem(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The constant enum value."""
    const: str
    """Display title for this option."""
    title: str

class TitledMultiSelectEnumSchemaItems(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Array of enum options with values and display labels."""
    any_of: List[TitledMultiSelectEnumSchemaItemsAnyOfItem]

class TitledMultiSelectEnumSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Optional default value."""
    default: List[str] | None = None
    """Optional description for the enum field."""
    description: str | None = None
    """Schema for array items with enum options and display labels."""
    items: TitledMultiSelectEnumSchemaItems
    """Maximum number of items to select."""
    max_items: int | None = None
    """Minimum number of items to select."""
    min_items: int | None = None
    """Optional title for the enum field."""
    title: str | None = None
    type: str

class TitledSingleSelectEnumSchemaOneOfItem(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The enum value."""
    const: str
    """Display label for this option."""
    title: str

class TitledSingleSelectEnumSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Optional default value."""
    default: str | None = None
    """Optional description for the enum field."""
    description: str | None = None
    """Array of enum options with values and display labels."""
    one_of: List[TitledSingleSelectEnumSchemaOneOfItem]
    """Optional title for the enum field."""
    title: str | None = None
    type: str

class UntitledMultiSelectEnumSchemaItems(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Array of enum values to choose from."""
    enum: List[str]
    type: str

class UntitledMultiSelectEnumSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Optional default value."""
    default: List[str] | None = None
    """Optional description for the enum field."""
    description: str | None = None
    """Schema for the array items."""
    items: UntitledMultiSelectEnumSchemaItems
    """Maximum number of items to select."""
    max_items: int | None = None
    """Minimum number of items to select."""
    min_items: int | None = None
    """Optional title for the enum field."""
    title: str | None = None
    type: str

class UntitledSingleSelectEnumSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Optional default value."""
    default: str | None = None
    """Optional description for the enum field."""
    description: str | None = None
    """Array of enum values to choose from."""
    enum: List[str]
    """Optional title for the enum field."""
    title: str | None = None
    type: str
EnumSchema = Union[UntitledSingleSelectEnumSchema, TitledSingleSelectEnumSchema, UntitledMultiSelectEnumSchema, TitledMultiSelectEnumSchema, LegacyTitledEnumSchema]


class Error(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The error type that occurred."""
    code: int
    """Additional information about the error. The value of this member is defined by the sender (e.g. detailed error information, nested errors etc.)."""
    data: Any | None = None
    """A short description of the error. The message SHOULD be limited to a concise single sentence."""
    message: str

class PromptMessage(BaseModel):
    model_config = ConfigDict(extra="forbid")

    content: ContentBlock
    role: Role

class GetPromptResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """An optional description for the prompt."""
    description: str | None = None
    messages: List[PromptMessage]

class Icons(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None

class ServerCapabilitiesCompletions(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ServerCapabilitiesLogging(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ServerCapabilitiesPrompts(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether this server supports notifications for changes to the prompt list."""
    list_changed: bool | None = None

class ServerCapabilitiesResources(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether this server supports notifications for changes to the resource list."""
    list_changed: bool | None = None
    """Whether this server supports subscribing to resource updates."""
    subscribe: bool | None = None

class ServerCapabilitiesTasksCancel(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ServerCapabilitiesTasksList(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ServerCapabilitiesTasksRequestsToolsCall(BaseModel):
    model_config = ConfigDict(extra="forbid")

    pass

class ServerCapabilitiesTasksRequestsTools(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether the server supports task-augmented tools/call requests."""
    call: ServerCapabilitiesTasksRequestsToolsCall | None = None

class ServerCapabilitiesTasksRequests(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Task support for tool-related requests."""
    tools: ServerCapabilitiesTasksRequestsTools | None = None

class ServerCapabilitiesTasks(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether this server supports tasks/cancel."""
    cancel: ServerCapabilitiesTasksCancel | None = None
    """Whether this server supports tasks/list."""
    list: ServerCapabilitiesTasksList | None = None
    """Specifies which request types can be augmented with tasks."""
    requests: ServerCapabilitiesTasksRequests | None = None

class ServerCapabilitiesTools(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Whether this server supports notifications for changes to the tool list."""
    list_changed: bool | None = None

class ServerCapabilities(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Present if the server supports argument autocompletion suggestions."""
    completions: ServerCapabilitiesCompletions | None = None
    """Experimental, non-standard capabilities that the server supports."""
    experimental: Dict[str, Any] | None = None
    """Present if the server supports sending log messages to the client."""
    logging: ServerCapabilitiesLogging | None = None
    """Present if the server offers any prompt templates."""
    prompts: ServerCapabilitiesPrompts | None = None
    """Present if the server offers any resources to read."""
    resources: ServerCapabilitiesResources | None = None
    """Present if the server supports task-augmented requests."""
    tasks: ServerCapabilitiesTasks | None = None
    """Present if the server offers any tools to call."""
    tools: ServerCapabilitiesTools | None = None

class InitializeResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    capabilities: ServerCapabilities
    """
    Instructions describing how to use the server and its features.
    
    This can be used by clients to improve the LLM's understanding of available tools, resources, etc. It can be thought of like a "hint" to the model. For example, this information MAY be added to the system prompt.
    """
    instructions: str | None = None
    """The version of the Model Context Protocol that the server wants to use. This may not match the version that the client requested. If the client cannot support this version, it MUST disconnect."""
    protocol_version: str
    server_info: Implementation

class JSONRPCErrorResponse(BaseModel):
    model_config = ConfigDict(extra="forbid")

    error: Error
    id: RequestId | None = None
    jsonrpc: str

class JSONRPCNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: Dict[str, Any] | None = None

class JSONRPCRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: Dict[str, Any] | None = None

class JSONRPCResultResponse(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    result: Result
"""Refers to any valid JSON-RPC object that can be decoded off the wire, or encoded to be sent."""
JSONRPCMessage = Union[JSONRPCRequest, JSONRPCNotification, JSONRPCResultResponse, JSONRPCErrorResponse]

"""A response to a request, containing either the result or error."""
JSONRPCResponse = Union[JSONRPCResultResponse, JSONRPCErrorResponse]


class PromptArgument(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """A human-readable description of the argument."""
    description: str | None = None
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """Whether this argument must be provided."""
    required: bool | None = None
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None

class Prompt(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """A list of arguments to use for templating the prompt."""
    arguments: List[PromptArgument] | None = None
    """An optional description of what this prompt provides"""
    description: str | None = None
    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None

class ListPromptsResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An opaque token representing the pagination position after the last returned result.
    If present, there may be more results available.
    """
    next_cursor: str | None = None
    prompts: List[Prompt]

class ResourceTemplate(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    """
    A description of what this template is for.
    
    This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
    """
    description: str | None = None
    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None
    """The MIME type for all resources that match this template. This should only be included if all resources matching this template have the same type."""
    mime_type: str | None = None
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None
    """A URI template (according to RFC 6570) that can be used to construct resource URIs."""
    uritemplate: str

class ListResourceTemplatesResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An opaque token representing the pagination position after the last returned result.
    If present, there may be more results available.
    """
    next_cursor: str | None = None
    resource_templates: List[ResourceTemplate]

class Resource(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """Optional annotations for the client."""
    annotations: Annotations | None = None
    """
    A description of what this resource represents.
    
    This can be used by clients to improve the LLM's understanding of available resources. It can be thought of like a "hint" to the model.
    """
    description: str | None = None
    """
    Optional set of sized icons that the client can display in a user interface.
    
    Clients that support rendering icons MUST support at least the following MIME types:
    - `image/png` - PNG images (safe, universal compatibility)
    - `image/jpeg` (and `image/jpg`) - JPEG images (safe, universal compatibility)
    
    Clients that support rendering icons SHOULD also support:
    - `image/svg+xml` - SVG images (scalable but requires security precautions)
    - `image/webp` - WebP images (modern, efficient format)
    """
    icons: List[Icon] | None = None
    """The MIME type of this resource, if known."""
    mime_type: str | None = None
    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
    
    This can be used by Hosts to display file sizes and estimate context window usage.
    """
    size: int | None = None
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None
    """The URI of this resource."""
    uri: AnyUrl

class ListResourcesResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An opaque token representing the pagination position after the last returned result.
    If present, there may be more results available.
    """
    next_cursor: str | None = None
    resources: List[Resource]

class ListRootsRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: RequestParams | None = None

class ListToolsResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An opaque token representing the pagination position after the last returned result.
    If present, there may be more results available.
    """
    next_cursor: str | None = None
    tools: List[Tool]

class LoggingMessageNotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """The data to be logged, such as a string message or an object. Any JSON serializable type is allowed here."""
    data: Any
    """The severity of this log message."""
    level: LoggingLevel
    """An optional name of the logger issuing this message."""
    logger: str | None = None

class LoggingMessageNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: LoggingMessageNotificationParams
MultiSelectEnumSchema = Union[UntitledMultiSelectEnumSchema, TitledMultiSelectEnumSchema]


class Notification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    method: str
    params: Dict[str, Any] | None = None

class NumberSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    default: int | None = None
    description: str | None = None
    maximum: int | None = None
    minimum: int | None = None
    title: str | None = None
    type: str

class PaginatedRequest(BaseModel):
    model_config = ConfigDict(extra="forbid")

    id: RequestId
    jsonrpc: str
    method: str
    params: PaginatedRequestParams | None = None

class PaginatedResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """
    An opaque token representing the pagination position after the last returned result.
    If present, there may be more results available.
    """
    next_cursor: str | None = None

class StringSchema(BaseModel):
    model_config = ConfigDict(extra="forbid")

    default: str | None = None
    description: str | None = None
    format: str | None = None
    max_length: int | None = None
    min_length: int | None = None
    title: str | None = None
    type: str
"""
Restricted schema definitions that only allow primitive types
without nested objects or arrays.
"""
PrimitiveSchemaDefinition = Union[StringSchema, NumberSchema, BooleanSchema, UntitledSingleSelectEnumSchema, TitledSingleSelectEnumSchema, UntitledMultiSelectEnumSchema, TitledMultiSelectEnumSchema, LegacyTitledEnumSchema]


class PromptListChangedNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: NotificationParams | None = None

class PromptReference(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn't present)."""
    name: str
    """
    Intended for UI and end-user contexts — optimized to be human-readable and easily understood,
    even by those unfamiliar with domain-specific terminology.
    
    If not provided, the name should be used for display (except for Tool,
    where `annotations.title` should be given precedence over using `name`,
    if present).
    """
    title: str | None = None
    type: str

class ReadResourceResult(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    contents: List[Any]

class RelatedTaskMetadata(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """The task identifier this message is associated with."""
    task_id: str

class Request(BaseModel):
    model_config = ConfigDict(extra="forbid")

    method: str
    params: Dict[str, Any] | None = None

class ResourceContents(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """The MIME type of this resource, if known."""
    mime_type: str | None = None
    """The URI of this resource."""
    uri: AnyUrl

class ResourceListChangedNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: NotificationParams | None = None

class ResourceRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class ResourceRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: ResourceRequestParamsMeta | None = None
    """The URI of the resource. The URI can use any protocol; it is up to the server how to interpret it."""
    uri: AnyUrl

class ResourceTemplateReference(BaseModel):
    model_config = ConfigDict(extra="forbid")

    type: str
    """The URI or URI template of the resource."""
    uri: str

class ResourceUpdatedNotificationParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """The URI of the resource that has been updated. This might be a sub-resource of the one that the client actually subscribed to."""
    uri: AnyUrl

class ResourceUpdatedNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: ResourceUpdatedNotificationParams

class ToolResultContent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Optional metadata about the tool result. Clients SHOULD preserve this field when
    including tool results in subsequent sampling requests to enable caching optimizations.
    
    See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
    """
    meta: Dict[str, Any] | None = None
    """
    The unstructured result content of the tool use.
    
    This has the same format as CallToolResult.content and can include text, images,
    audio, resource links, and embedded resources.
    """
    content: List[ContentBlock]
    """
    Whether the tool use resulted in an error.
    
    If true, the content typically describes the error that occurred.
    Default: false
    """
    is_error: bool | None = None
    """
    An optional structured result object.
    
    If the tool defined an outputSchema, this SHOULD conform to that schema.
    """
    structured_content: Dict[str, Any] | None = None
    """
    The ID of the tool use this result corresponds to.
    
    This MUST match the ID from a previous ToolUseContent.
    """
    tool_use_id: str
    type: str

class ToolUseContent(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """
    Optional metadata about the tool use. Clients SHOULD preserve this field when
    including tool uses in subsequent sampling requests to enable caching optimizations.
    
    See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage.
    """
    meta: Dict[str, Any] | None = None
    """
    A unique identifier for this tool use.
    
    This ID is used to match tool results to their corresponding tool uses.
    """
    id: str
    """The arguments to pass to the tool, conforming to the tool's input schema."""
    input: Dict[str, Any]
    """The name of the tool to call."""
    name: str
    type: str
SamplingMessageContentBlock = Union[TextContent, ImageContent, AudioContent, ToolUseContent, ToolResultContent]


class ToolListChangedNotification(BaseModel):
    model_config = ConfigDict(extra="forbid")

    jsonrpc: str
    method: str
    params: NotificationParams | None = None
ServerNotification = Union[CancelledNotification, ProgressNotification, ResourceListChangedNotification, ResourceUpdatedNotification, PromptListChangedNotification, ToolListChangedNotification, TaskStatusNotification, LoggingMessageNotification, ElicitationCompleteNotification]

ServerRequest = Union[PingRequest, GetTaskRequest, GetTaskPayloadRequest, CancelTaskRequest, ListTasksRequest, CreateMessageRequest, ListRootsRequest, ElicitRequest]

ServerResult = Union[Result, InitializeResult, ListResourcesResult, ListResourceTemplatesResult, ReadResourceResult, ListPromptsResult, GetPromptResult, ListToolsResult, CallToolResult, GetTaskResult, GetTaskPayloadResult, CancelTaskResult, ListTasksResult, CompleteResult]

SingleSelectEnumSchema = Union[UntitledSingleSelectEnumSchema, TitledSingleSelectEnumSchema]


class TaskAugmentedRequestParamsMeta(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """If specified, the caller is requesting out-of-band progress notifications for this request (as represented by notifications/progress). The value of this parameter is an opaque token that will be attached to any subsequent notifications. The receiver is not obligated to provide these notifications."""
    progress_token: ProgressToken | None = None

class TaskAugmentedRequestParams(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: TaskAugmentedRequestParamsMeta | None = None
    """
    If specified, the caller is requesting task-augmented execution for this request.
    The request will return a CreateTaskResult immediately, and the actual result can be
    retrieved later via tasks/result.
    
    Task augmentation is subject to capability negotiation - receivers MUST declare support
    for task augmentation of specific request types in their capabilities.
    """
    task: TaskMetadata | None = None

class TextResourceContents(BaseModel):
    model_config = ConfigDict(extra="forbid")

    """See [General fields: `_meta`](/specification/2025-11-25/basic/index#meta) for notes on `_meta` usage."""
    meta: Dict[str, Any] | None = None
    """The MIME type of this resource, if known."""
    mime_type: str | None = None
    """The text of the item. This must only be set if the item can actually be represented as text (not binary data)."""
    text: str
    """The URI of this resource."""
    uri: AnyUrl

class URLElicitationRequiredError(BaseModel):
    model_config = ConfigDict(extra="forbid")

    error: Any
    id: RequestId | None = None
    jsonrpc: str
