export interface LoginBody {
    email: string,
    password: string
}

export interface RegisterBody {
    email: string | undefined,
    password: string | undefined,
    name: string | undefined
}