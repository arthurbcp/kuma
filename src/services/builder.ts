import { IHttpProvider } from "../providers/http";
import {UserRepositoriesService,RepositoriesService,OrganizationRepositoriesService,IssuesService } from "./";

export const buildServices = (provider: IHttpProvider) => {
  return {
    userRepositories: new UserRepositoriesService(provider),
    repositories: new RepositoriesService(provider),
    organizationRepositories: new OrganizationRepositoriesService(provider),
    issues: new IssuesService(provider),};
};
