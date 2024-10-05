import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { Repository } from "../../dto/repository";
import { NewRepository } from "../../dto/new_repository";

// Operations related to repositories of the authenticated user.
export class UserRepositoriesService {
  private http: IHttpProvider;

  constructor(http: IHttpProvider) {
    this.http = http;
  }

  //Lists all repositories for the authenticated user.
  async listUserRepos(
    data?: {
      query?: {
        type?: "all" | "owner" | "public" | "private" | "member";
      };
    },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<Repository[]>> {
    return await this.http.request<Repository[]>(
      "get",
      "/user/repos",
      data,
      config
    );
  }

  //Creates a new repository for the authenticated user.
  async createUserRepo(
    data?: { body?: NewRepository },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<Repository>> {
    return await this.http.request<Repository>(
      "post",
      "/user/repos",
      data,
      config
    );
  }
}
