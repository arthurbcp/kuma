

export type UpdateRepository = {
    // New description of the repository
    description?:string
    // True to enable issues for this repository
    has_issues?:boolean
    // True to enable projects for this repository
    has_projects?:boolean
    // True to enable the wiki for this repository
    has_wiki?:boolean
    // New homepage URL
    homepage?:string
    // New name of the repository
    name?:string
    // True to make the repository private
    private?:boolean
}