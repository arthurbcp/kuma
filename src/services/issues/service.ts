import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { Issue } from "../../dto/issue"


// Operations related to issues in a repository.
export class IssuesService {
  private http: IHttpProvider;

  constructor(http: IHttpProvider) {
    this.http = http;
  }
  
//Lists all issues for the specified repository.
    async listRepoIssues(data?: {params?: { 
        owner:string,
        repo:string,},},   config?: AxiosRequestConfig):Promise<AxiosResponse<Issue[]>> {
      return await this.http.request<Issue[]>("get","/repos/{owner}/{repo}/issues", data, config);
    }

  
  
}