"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import axios from "axios"

export default function RegisterPage() {
    const router = useRouter()
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState("")
    const [success, setSuccess] = useState(false)

    const handleRegister = async () => {
        try {
            await axios.post("http://localhost:8080/api/register", {
                username,
                email,
                password,
            })

            setSuccess(true)
            setError("")
            router.push("/login")
        } catch (e) {
            setError("Ошибка регистрации. Проверьте введённые данные.")
        }
    }

    const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => setUsername(e.target.value)
    const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)
    const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)

    return (
        <div className="h-screen flex justify-center items-center">
            <div className="space-y-4 w-full max-w-sm">
                <h1 className="text-2xl font-bold">Регистрация</h1>
                <div>
                    <Label>Имя пользователя</Label>
                    <Input value={username} onChange={handleUsernameChange} />
                </div>
                <div>
                    <Label>Email</Label>
                    <Input value={email} onChange={handleEmailChange} />
                </div>
                <div>
                    <Label>Пароль</Label>
                    <Input type="password" value={password} onChange={handlePasswordChange} />
                </div>
                {error && <p className="text-sm text-red-500">{error}</p>}
                {success && <p className="text-sm text-green-500">Регистрация прошла успешно!</p>}
                <Button className="w-full" onClick={handleRegister}>
                    Зарегистрироваться
                </Button>
            </div>
        </div>
    )
}
