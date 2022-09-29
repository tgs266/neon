import axios, { AxiosResponse } from "axios";
import { App } from "../models/app";
import { Pagination } from "../models/pagination";
import { CreateCredentialsRequest, CreateCredentialsResponse } from "../models/credentials";

export class CredentialsService {
    static async create(req: CreateCredentialsRequest): Promise<AxiosResponse<CreateCredentialsResponse>> {
        return axios.post("http://localhost:5000/api/v1/credentials", req)
    }
}