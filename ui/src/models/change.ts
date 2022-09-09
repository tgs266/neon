import { Release } from "./release";

export interface QueuedChange {
    id: string
    type: string,
    details: string,
    release: Release,
    lastChecked: string
}

export interface StoredChange {
    id: string
    type: string,
    details: string,
    release: Release,
    completedAt: string
}