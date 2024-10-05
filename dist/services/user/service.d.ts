import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { User } from "../../dto/user";
export declare class UserService {
    private http;
    constructor(http: IHttpProvider);
    createUser(data?: {
        body?: User;
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    createUsersWithArrayInput(data?: {
        body?: User[];
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    createUsersWithListInput(data?: {
        body?: User[];
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    loginUser(data?: {
        query?: {
            username: string;
            password: string;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<string>>;
    logoutUser(data?: {}, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    deleteUser(data?: {
        params?: {
            username: string;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    getUserByName(data?: {
        params?: {
            username: string;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<User>>;
    updateUser(data?: {
        params?: {
            username: string;
        };
        body?: User;
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
}
