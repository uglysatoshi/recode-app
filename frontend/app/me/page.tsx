'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import axios from 'axios'
import { Card, CardContent } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'

type User = {
    id: number
    username: string
    email: string
    role: string
}

export default function MePage() {
    const [user, setUser] = useState<User | null>(null)
    const [loading, setLoading] = useState(true)
    const router = useRouter()

    useEffect(() => {
        const token = localStorage.getItem('token')
        if (!token) {
            router.push('/login')
            return
        }

        axios
            .get('http://localhost:8080/api/me', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            })
            .then((res) => {
                setUser(res.data.data)
            })
            .catch((err) => {
                console.error('Ошибка при получении пользователя:', err)
                router.push('/login')
            })
            .finally(() => setLoading(false))
    }, [router])

    if (loading) {
        return (
            <div className="p-6">
                <Skeleton className="h-10 w-[200px] mb-4" />
                <Skeleton className="h-6 w-full" />
            </div>
        )
    }

    return (
        <div className="p-6">
            <h1 className="text-2xl font-bold mb-4">Информация о пользователе</h1>
            <Card>
                <CardContent className="p-4 space-y-2">
                    <p><strong>ID:</strong> {user?.id}</p>
                    <p><strong>Username:</strong> {user?.username}</p>
                    <p><strong>Email:</strong> {user?.email}</p>
                    <p><strong>Role:</strong> {user?.role}</p>
                </CardContent>
            </Card>
        </div>
    )
}
