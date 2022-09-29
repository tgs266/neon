export interface CreateCredentialsRequest {
    name: string,
    username?: string,
    password?: string,
    token?: string
}

export interface CreateCredentialsResponse {
    name: string,
}

export interface Credentials {
    name: string,
    basicAuth: boolean,
    tokenAuth: boolean
}