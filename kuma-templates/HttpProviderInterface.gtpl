import { AxiosRequestConfig, AxiosResponse } from "axios";

export interface IHttpProvider {
  request<T>(
    method: string,
    url: string,
    data?: RequestData,
    config?: AxiosRequestConfig<any>
  ): Promise<AxiosResponse<T>>;
}
