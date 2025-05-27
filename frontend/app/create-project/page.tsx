'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import axios from 'axios'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'
import { Card, CardContent } from '@/components/ui/card'
import { Label } from '@/components/ui/label'

export default function CreateProjectPage() {
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')
    const router = useRouter()

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        const token = localStorage.getItem('token')
        if (!token) return router.push('/login')

        try {
            await axios.post('http://localhost:8080/api/projects', {
                title,
                description,
            }, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            })
            router.push('/')
        } catch (error) {
            console.error('Ошибка при создании проекта', error)
        }
    }

    return (
        <div className="max-w-xl mx-auto p-6">
            <Card>
                <CardContent className="p-6 space-y-4">
                    <h1 className="text-2xl font-bold">Создать проект</h1>
                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                            <Label>Название проекта</Label>
                            <Input value={title} onChange={(e) => setTitle(e.target.value)} required />
                        </div>
                        <div>
                            <Label>Описание</Label>
                            <Textarea value={description} onChange={(e) => setDescription(e.target.value)} />
                        </div>
                        <Button type="submit">Создать</Button>
                    </form>
                </CardContent>
            </Card>
        </div>
    )
}
