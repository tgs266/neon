import { Install } from "./install";

export interface App {
    createdAt: string,
    updatedAt: string,
    name: string,
    products: string[],
    installs?: Install[],
    releaseChannel: number,
    error: string,
    installStatus: string
}