import axios, { AxiosResponse } from "axios";
import { App } from "../models/app";
import { QueuedChange, StoredChange } from "../models/change";
import { Install } from "../models/install";
import { Pagination } from "../models/pagination";
import { ResourceList } from "../models/resources";

export class AppService {

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

    static async getStoredChanges(name = "", limit = 10, offset = 0): Promise<AxiosResponse<Pagination<StoredChange>>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/changes/stored`, { params: { limit, offset } })
    }

    static async getQueuedChanges(name = "", limit = 10, offset = 0): Promise<AxiosResponse<Pagination<QueuedChange>>> {
        return axios.get(`http://localhost:5000/api/v1/apps/${name}/changes/queued`, { params: { limit, offset } })
    }

}