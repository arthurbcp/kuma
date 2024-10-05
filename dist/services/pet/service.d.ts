import { AxiosRequestConfig, AxiosResponse } from "axios";
import { IHttpProvider } from "../../providers/http/http_provider_interface";
import { Pet } from "../../dto/pet";
import { ApiResponse } from "../../dto/api_response";
export declare class PetService {
    private http;
    constructor(http: IHttpProvider);
    addPet(data?: {
        body?: Pet;
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    updatePet(data?: {
        body?: Pet;
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    findPetsByStatus(data?: {
        query?: {
            status: "available" | "pending" | "sold"[];
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<Pet[]>>;
    findPetsByTags(data?: {
        query?: {
            tags: string[];
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<Pet[]>>;
    deletePet(data?: {
        params?: {
            petId: number;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    getPetById(data?: {
        params?: {
            petId: number;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<Pet>>;
    updatePetWithForm(data?: {
        params?: {
            petId: number;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<undefined>>;
    uploadFile(data?: {
        params?: {
            petId: number;
        };
    }, config?: AxiosRequestConfig): Promise<AxiosResponse<ApiResponse>>;
}
