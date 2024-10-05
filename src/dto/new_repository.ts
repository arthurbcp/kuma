

export type NewRepository = {
    // A short description of the repository
    description?:string
    // True to enable issues for this repository
    has_issues?:boolean
    // True to enable projects for this repository
    has_projects?:boolean
    // True to enable the wiki for this repository
    has_wiki?:boolean
    // Repository homepage URL
    homepage?:string
    // Name of the repository
    name?:string
    // True to create a private repository
    private?:boolean
}