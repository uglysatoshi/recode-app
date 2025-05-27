'use client'

import { useEffect, useState } from 'react'
import axios from 'axios'
import { useRouter } from 'next/navigation'
import { Card, CardContent } from '@/components/ui/card'

type Project = {
    ID: number
    Title: string
    Description: string
}

type Task = {
    ID: number
    Title: string
    Description: string
    Status: string
    Priority: string
}

export default function Dashboard() {
    const router = useRouter()
    const [projects, setProjects] = useState<Project[]>([])
    const [tasks, setTasks] = useState<Task[]>([])
    const [loading, setLoading] = useState(true)

    const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null

    useEffect(() => {
        if (!token) {
            router.push('/login')
            return
        }

        const fetchData = async () => {
            try {
                const [projectsRes, tasksRes] = await Promise.all([
                    axios.get('http://localhost:8080/api/projects', {
                        headers: { Authorization: `Bearer ${token}` },
                    }),
                    axios.get('http://localhost:8080/api/tasks', {
                        headers: { Authorization: `Bearer ${token}` },
                    }),
                ])

                setProjects(projectsRes.data.data || [])
                setTasks(tasksRes.data.data || [])
            } catch (error) {
                console.error('Ошибка при загрузке данных:', error)
                router.push('/login') // возможно токен недействителен
            } finally {
                setLoading(false)
            }
        }

        fetchData()
    }, [token, router])

    if (loading) return <div className="p-4">Загрузка...</div>

    return (
        <div className="p-6 space-y-6">
            <h1 className="text-2xl font-bold">Дашборд</h1>

            <section>
                <h2 className="text-xl font-semibold mb-2">Ваши проекты</h2>
                {projects.length === 0 ? (
                    <p className="text-muted-foreground">Проектов нет</p>
                ) : (
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        {projects.map((project) => (
                            <Card key={project.ID}>
                                <CardContent className="p-4">
                                    <h3 className="text-lg font-medium">{project.Title}</h3>
                                    <p className="text-sm text-muted-foreground">{project.Description}</p>
                                </CardContent>
                            </Card>
                        ))}
                    </div>
                )}
            </section>

            <section>
                <h2 className="text-xl font-semibold mb-2">Ваши задачи</h2>
                {tasks.length === 0 ? (
                    <p className="text-muted-foreground">Задач нет</p>
                ) : (
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        {tasks.map((task) => (
                            <Card key={task.ID}>
                                <CardContent className="p-4">
                                    <h3 className="text-lg font-medium">{task.Title}</h3>
                                    <p className="text-sm">{task.Description}</p>
                                    <div className="text-sm mt-2">
                                        <span className="font-semibold">Статус:</span> {task.Status}
                                        <br />
                                        <span className="font-semibold">Приоритет:</span> {task.Priority}
                                    </div>
                                </CardContent>
                            </Card>
                        ))}
                    </div>
                )}
            </section>
        </div>
    )
}
