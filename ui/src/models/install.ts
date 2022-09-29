export interface Install {
    createdAt: Date,
    updatedAt: Date,
    appName: string,
    productName: string,
    releaseVersion: string,
    error: string
}

export interface InstallCommit {
    data: string,
    message: string
}