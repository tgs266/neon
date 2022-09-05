export interface Release {
    createdAt: Date,
    updatedAt: Date,
    productName: string,
    productVersion: string,
    releaseChannel: number,
    helmChart: string,
    recalled: boolean,
    dependencies: any
}