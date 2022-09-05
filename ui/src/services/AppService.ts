import axios, { AxiosResponse } from "axios";
import { App } from "../models/app";
import { Pagination } from "../models/pagination";

export class AppService {

    static async listApps(limit = 10, offset = 0, name = ""): Promise<AxiosResponse<Pagination<App>>> {
        return axios.get("http://localhost:5000/api/v1/apps", { params: { limit, offset, name } })
    }

}