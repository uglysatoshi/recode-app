"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { saveToken } from "@/lib/auth"
import axios from "axios"

export default function LoginPage() {
    const router = useRouter()
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState("")

    const handleLogin = async () => {
        try {
            const res = await axios.post("http://localhost:8080/api/login", {
                email,
                password,
            })

            const token = res.data.token
            if (token) {
                saveToken(token)
                router.push("/")
            } else {
                setError("Ошибка авторизации")
            }
        } catch (e) {
            setError("Неверный email или пароль")
        }
    }

    return (
        <div className="h-screen flex justify-center items-center">
            <div className="space-y-4 w-full max-w-sm">
                <h1 className="text-2xl font-bold">Вход</h1>
                <div>
                    <Label>Email</Label>
                    <Input value={email} onChange={(e) => setEmail(e.target.value)} />
                </div>
                <div>
                    <Label>Пароль</Label>
                    <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </div>
                {error && <p className="text-sm text-red-500">{error}</p>}
                <Button className="w-full" onClick={handleLogin}>
                    Войти
                </Button>
            </div>
        </div>
    )
}
