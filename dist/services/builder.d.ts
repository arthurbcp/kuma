import { IHttpProvider } from "../providers/http";
import { PetService, StoreService, UserService } from "./";
export declare const buildServices: (provider: IHttpProvider) => {
    pet: PetService;
    store: StoreService;
    user: UserService;
};
