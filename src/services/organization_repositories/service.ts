import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { Repository } from "../../dto/repository";

// Operations related to repositories of an organization.
export class OrganizationRepositoriesService {
  private http: IHttpProvider;

  constructor(http: IHttpProvider) {
    this.http = http;
  }

  //Lists repositories for the specified organization.
  async listOrgRepos(
    data?: {
      params?: {
        org: string;
      };
    },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<Repository[]>> {
    return await this.http.request<Repository[]>(
      "get",
      "/orgs/{org}/repos",
      data,
      config
    );
  }
}
