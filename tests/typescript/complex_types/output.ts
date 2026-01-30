// Describes an error that occurred.
export interface Error {
  // Machine-readable error code.
  code: string;
  // The field that caused the error, if applicable.
  field?: string;
  // Human-readable error message.
  message: string;
}

// Represents an item in the system.
// Items can be of various types and contain metadata.
export interface Item {
  // When this item was created.
  createdAt: Date;
  // Unique identifier for the item.
  id: string;
  // Additional key-value metadata.
  metadata?: Record<string, unknown>;
  // Tags associated with this item.
  tags?: string[];
  // The type classification of this item.
  type: string;
  // When this item was last updated.
  updatedAt?: Date;
}

// Pagination and request metadata.
export interface Metadata {
  // Whether there are more pages available.
  hasMore?: boolean;
  // Current page number (1-indexed).
  page: number;
  // Number of items per page.
  perPage: number;
  // Total number of items available.
  total: number;
}

// A generic API response wrapper with pagination support.
export interface ApiResponse {
  // The array of result items.
  data: Item[];
  // Any errors that occurred during the request.
  errors?: Error[];
  meta: Metadata;
}
