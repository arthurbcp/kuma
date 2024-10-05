import { AxiosResponse } from "axios";
import { IHttpProvider } from "./http_provider_interface";
import { RequestData } from "./dto";
export declare class HttpProvider implements IHttpProvider {
    private client;
    constructor(baseURL: string);
    request<T>(method: "get" | "post" | "put" | "delete" | "patch" | "options" | "head" | "trace", url: string, data?: RequestData, config?: any): Promise<AxiosResponse<T>>;
}
