import axios, { AxiosInstance, AxiosResponse } from "axios";
import { IHttpProvider } from "./http_provider_interface";
import { RequestData } from "./dto"

export class HttpProvider implements IHttpProvider {
  private client: AxiosInstance;
  constructor(baseURL: string) {
    this.client = axios.create({ baseURL });
  }

  async request<T>(
    method:
      | "get"
      | "post"
      | "put"
      | "delete"
      | "patch"
      | "options"
      | "head"
      | "trace",
    url: string,
    data?: RequestData,
    config?: any
  ): Promise<AxiosResponse<T>> {
    const response = await this.client.request<T>({
      method,
      url,
      data,
      ...config,
    });
    return response;
  }
}
