import axios, { AxiosResponse } from "axios";
import { V1Pod } from "@kubernetes/client-node"

export class KubernetesService {

    static async getPod(namespace = "", name = ""): Promise<AxiosResponse<V1Pod>> {
        return axios.get(`http://localhost:5000/api/v1/kubernetes/pods/${namespace}/${name}`)
    }

}