import "@/app/globals.css"
import { ReactNode } from "react"
import {SidebarProvider} from "@/components/ui/sidebar";
import {AppSidebar} from "@/components/app-sidebar";

export default function RootLayout({ children }: { children: ReactNode }) {
    return (
        <html lang="ru">
        <body className="flex min-h-screen">
        <SidebarProvider>
            <AppSidebar />
            <main className="flex-1 p-4 overflow-auto">
                {children}
            </main>
        </SidebarProvider>
        </body>
        </html>
    )
}
