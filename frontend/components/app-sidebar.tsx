"use client"

import * as React from "react"
import { useEffect, useState } from "react"
import { Plus } from "lucide-react"
import axios from "axios"
import { useRouter } from "next/navigation"
import { DatePicker } from "@/components/date-picker"
import { NavUser } from "@/components/nav-user"
import {
    Sidebar,
    SidebarContent,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarRail,
    SidebarSeparator,
} from "@/components/ui/sidebar"
import {getToken} from "@/lib/auth";


export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
    const [user, setUser] = useState<{ name: string; email: string } | null>(null)

    const router = useRouter()

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const token = getToken()
                if (!token) return router.push('/login')

                const res = await axios.get("http://localhost:8080/api/me", {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                })

                const data = res.data.data
                setUser({ name: data.username, email: data.email })
            } catch (err) {
                console.error("Ошибка получения данных пользователя:", err)
            }
        }

        fetchUser()
    }, [router])

    return (
        <Sidebar {...props}>
            <SidebarHeader className="border-sidebar-border h-16 border-b">
                {user && <NavUser user={user} />}
            </SidebarHeader>
            <SidebarContent>
                <DatePicker />
                <SidebarSeparator className="mx-0" />
                <SidebarMenu>
                    <SidebarMenuItem>
                        <SidebarMenuButton>
                            <Plus />
                            <span>Создать проект</span>
                        </SidebarMenuButton>
                    </SidebarMenuItem>
                </SidebarMenu>
                <SidebarSeparator className="mx-0" />
            </SidebarContent>
            <SidebarRail />
        </Sidebar>
    )
}
