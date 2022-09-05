import { Release } from "./release";

export interface Product {
    name: string
    releases?: Release[]
}