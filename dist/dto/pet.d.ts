import { Category } from './category';
import { Tag } from './tag';
export type Pet = {
    category?: Category;
    id?: number;
    name?: string;
    photoUrls?: string[];
    status?: "available" | "pending" | "sold";
    tags?: Tag[];
};
