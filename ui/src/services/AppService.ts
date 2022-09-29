import axios, { AxiosResponse } from "axios";
import { App, CreateAppRequest } from "../models/app";
import { QueuedChange, StoredChange } from "../models/change";
import { Install, InstallCommit } from "../models/install";
import { Pagination } from "../models/pagination";
import { ResourceList } from "../models/resources";

export class AppService {

    static async create(req: CreateAppRequest): Promise<AxiosResponse<App>> {
        return axios.post("http://localhost:5000/api/v1/apps", req)
    }

    static async addProduct(appName: string, name: string): Promise<AxiosResponse<App>> {
        return axios.post(`http://localhost:5000/api/v1/apps/${appName}/products`, { name })
    }

    static async listApps(limit = 10, offset = 0, name = ""): Promise<AxiosResponse<Pagination<App>>> {
        return axios.get("http://localhost:5000/api/v1/apps", { params: { limit, offset, name } })
    }

    static async get(name = ""): Promise<AxiosResponse<App>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}`)
    }

    static async getInstall(name = "", productName = ""): Promise<AxiosResponse<Install>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/installs/${productName}`)
    }

    static async getInstallResources(name = "", productName = ""): Promise<AxiosResponse<ResourceList>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/installs/${productName}/resources`)
    }

    static async getInstallConfig(name = "", productName = ""): Promise<AxiosResponse<{data: string}>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/installs/${productName}/config`)
    }

    static async updateInstallConfig(name = "", productName = "", data: InstallCommit): Promise<AxiosResponse<{data: string}>> {
        return axios.put(`http://localhost:5000/api/v1/apps/${name}/installs/${productName}/config`, data)
    }

    static async getStoredChanges(name = "", limit = 10, offset = 0): Promise<AxiosResponse<Pagination<StoredChange>>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/changes/stored`, { params: { limit, offset } })
    }

    static async getQueuedChanges(name = "", limit = 10, offset = 0): Promise<AxiosResponse<Pagination<QueuedChange>>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/changes/queued`, { params: { limit, offset } })
    }

}