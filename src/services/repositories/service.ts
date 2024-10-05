import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { Repository } from "../../dto/repository";
import { UpdateRepository } from "../../dto/update_repository";

// Operations related to public repositories.
export class RepositoriesService {
  private http: IHttpProvider;

  constructor(http: IHttpProvider) {
    this.http = http;
  }

  //Deletes a repository.
  async deleteRepo(
    data?: {
      params?: {
        owner: string;
        repo: string;
      };
    },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<undefined>> {
    return await this.http.request<undefined>(
      "delete",
      "/repos/{owner}/{repo}",
      data,
      config
    );
  }

  //Retrieves a repository by owner and repo name.
  async getRepo(
    data?: {
      params?: {
        owner: string;
        repo: string;
      };
    },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<Repository>> {
    return await this.http.request<Repository>(
      "get",
      "/repos/{owner}/{repo}",
      data,
      config
    );
  }

  //Updates repository information.
  async updateRepo(
    data?: {
      params?: {
        owner: string;
        repo: string;
      };
      body?: UpdateRepository;
    },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<Repository>> {
    return await this.http.request<Repository>(
      "patch",
      "/repos/{owner}/{repo}",
      data,
      config
    );
  }

  //Lists public repositories for the specified user.
  async listUserReposByUsername(
    data?: {
      params?: {
        username: string;
      };
    },
    config?: AxiosRequestConfig
  ): Promise<AxiosResponse<Repository[]>> {
    return await this.http.request<Repository[]>(
      "get",
      "/users/{username}/repos",
      data,
      config
    );
  }
}
