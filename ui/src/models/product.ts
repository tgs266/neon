import { Install } from "./install";
import { Release } from "./release";

export interface Product {
    name: string
    releases?: Release[]
    installs?: Install[]
    createdAt: string
    updatedAt: string
}