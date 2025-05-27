export const getToken = (): string | null => {
    if (typeof window !== "undefined") {
        return localStorage.getItem("token")
    }
    return null
}

export const saveToken = (token: string) => {
    localStorage.setItem("token", token)
}

export const removeToken = () => {
    localStorage.removeItem("token")
}
