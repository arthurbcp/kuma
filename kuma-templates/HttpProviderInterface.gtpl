import { AxiosRequestConfig, AxiosResponse } from "axios";
import { RequestData } from "./dto"

export interface IHttpProvider {
  request<T>(
    method: string,
    url: string,
    data?: RequestData,
    config?: AxiosRequestConfig<any>
  ): Promise<AxiosResponse<T>>;
}
