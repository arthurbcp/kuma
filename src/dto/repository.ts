
import { User } from './user'

export type Repository = {
    created_at?:string
    description?:string
    fork?:boolean
    full_name?:string
    html_url?:string
    id?:number
    name?:string
    owner?:User
    private?:boolean
    pushed_at?:string
    updated_at?:string
    url?:string
}