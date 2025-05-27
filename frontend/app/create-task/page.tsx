'use client';

import { useEffect, useState } from "react";
import axios from "axios";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useRouter } from "next/navigation";

export default function CreateTaskPage() {
    const router = useRouter();
    const [projects, setProjects] = useState<any[]>([]);
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [status, setStatus] = useState("open");
    const [priority, setPriority] = useState("medium");
    const [deadline, setDeadline] = useState("");
    const [projectID, setProjectID] = useState("");
    const [userID, setUserID] = useState("");

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            router.push("/login");
            return;
        }

        // Получение текущего пользователя
        axios.get("http://localhost:8080/api/me", {
            headers: { Authorization: `Bearer ${token}` },
        })
            .then(res => {
                setUserID(res.data.data.id);
            })
            .catch(err => {
                console.error("Failed to load user:", err);
                router.push("/login");
            });

        // Получение проектов
        axios.get("http://localhost:8080/api/projects", {
            headers: { Authorization: `Bearer ${token}` },
        })
            .then(res => {
                const data = res.data.data;
                if (Array.isArray(data)) {
                    setProjects(data);
                } else {
                    console.error("Unexpected projects format:", data);
                }
            })
            .catch(err => {
                console.error("Failed to load projects:", err);
            });
    }, []);

    const handleSubmit = async () => {
        const token = localStorage.getItem("token");
        const isoDeadline = new Date(`${deadline}T00:00:00Z`).toISOString();
        if (!token) return;
        try {
            await axios.post("http://localhost:8080/api/tasks", {
                title,
                description,
                status,
                priority,
                isoDeadline,
                project_id: Number(projectID),
                user_id: Number(userID),
            }, {
                headers: { Authorization: `Bearer ${token}` },
            });

            router.push("/");
        } catch (error) {
            console.error("Failed to create task:", error);
        }
    };

    return (
        <div className="max-w-xl mx-auto p-4 space-y-4">
            <h1 className="text-2xl font-bold">Создать задачу</h1>

            <div>
                <Label>Название</Label>
                <Input value={title} onChange={(e) => setTitle(e.target.value)} />
            </div>

            <div>
                <Label>Описание</Label>
                <Textarea value={description} onChange={(e) => setDescription(e.target.value)} />
            </div>

            <div>
                <Label>Статус</Label>
                <Select onValueChange={setStatus} defaultValue={status}>
                    <SelectTrigger><SelectValue placeholder="Выберите статус" /></SelectTrigger>
                    <SelectContent>
                        <SelectItem value="open">Открыта</SelectItem>
                        <SelectItem value="in_progress">В процессе</SelectItem>
                        <SelectItem value="done">Завершена</SelectItem>
                    </SelectContent>
                </Select>
            </div>

            <div>
                <Label>Приоритет</Label>
                <Select onValueChange={setPriority} defaultValue={priority}>
                    <SelectTrigger><SelectValue placeholder="Выберите приоритет" /></SelectTrigger>
                    <SelectContent>
                        <SelectItem value="low">Низкий</SelectItem>
                        <SelectItem value="medium">Средний</SelectItem>
                        <SelectItem value="high">Высокий</SelectItem>
                    </SelectContent>
                </Select>
            </div>

            <div>
                <Label>Срок</Label>
                <Input type="date" value={deadline} onChange={(e) => setDeadline(e.target.value)} />
            </div>

            <div>
                <Label>Проект</Label>
                <Select onValueChange={setProjectID}>
                    <SelectTrigger><SelectValue placeholder="Выберите проект" /></SelectTrigger>
                    <SelectContent>
                        {Array.isArray(projects) && projects.map((project: any) => (
                            <SelectItem key={project.ID} value={String(project.ID)}>
                                {project.Title}
                            </SelectItem>
                        ))}
                    </SelectContent>
                </Select>
            </div>

            <Button onClick={handleSubmit}>Создать</Button>
        </div>
    );
}
