
import { User } from './user'

export type Issue = {
    closed_at?:string
    comments?:number
    created_at?:string
    id?:number
    locked?:boolean
    number?:number
    state?:"open" | "closed"
    title?:string
    updated_at?:string
    user?:User
}