import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { Order } from "../../dto/order";
export declare class StoreService {
    private http;
    constructor(http: IHttpProvider);
    getInventory(data?: {}, config?: AxiosRequestConfig): Promise<AxiosResponse<{
        [key: string]: any;
    }>>;
    placeOrder(data?: {
        body?: Order;
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<Order>>;
    deleteOrder(data?: {
        params?: {
            orderId: number;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    getOrderById(data?: {
        params?: {
            orderId: number;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<Order>>;
}
