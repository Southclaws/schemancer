export type HttpMethod =
  | "GET"
  | "POST"
  | "PUT"
  | "DELETE"
  | "PATCH"
  | "HEAD"
  | "OPTIONS";

export interface ApiRequest {
  body?: string;
  method: HttpMethod;
  url: string;
}

export type Color =
  | "red"
  | "green"
  | "blue"
  | "yellow";

export type Priority =
  | "low"
  | "medium"
  | "high"
  | "critical";

export type Status =
  | "pending"
  | "in_progress"
  | "completed"
  | "failed"
  | "cancelled";

export interface Task {
  color?: Color;
  id: string;
  priority?: Priority;
  status: Status;
  title: string;
}
