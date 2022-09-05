import axios, { AxiosResponse } from "axios";
import { Pagination } from "../models/pagination";
import { Product } from "../models/product";

export class ProductService {

    static async listProducts(limit = 10, offset = 0, name = ""): Promise<AxiosResponse<Pagination<Product>>> {
        return axios.get("http://localhost:5000/api/v1/products", { params: { limit, offset, name } })
    }

    static async get(name = ""): Promise<AxiosResponse<Product>> {
        return axios.get(`http://localhost:5000/api/v1/products/${name}`)
    }

}