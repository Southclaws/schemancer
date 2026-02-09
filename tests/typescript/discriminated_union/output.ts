export interface BaseEvent {
  timestamp: Date;
  type: string;
}


export interface CreatedEvent {
  type: "created";
  id: string;
  name: string;
  timestamp: Date;
}

export interface UpdatedEvent {
  type: "updated";
  changes: Record<string, unknown>;
  id: string;
  timestamp: Date;
}

export interface DeletedEvent {
  type: "deleted";
  id: string;
  reason?: string;
  timestamp: Date;
}

export type Event =
  | CreatedEvent
  | UpdatedEvent
  | DeletedEvent;

export function isCreatedEvent(value: Event): value is CreatedEvent {
  return value.type === "created";
}

export function isUpdatedEvent(value: Event): value is UpdatedEvent {
  return value.type === "updated";
}

export function isDeletedEvent(value: Event): value is DeletedEvent {
  return value.type === "deleted";
}
